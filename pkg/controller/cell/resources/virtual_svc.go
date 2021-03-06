/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package resources

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cellery-io/mesh-controller/pkg/apis/istio/networking/v1alpha3"
	"github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	listers "github.com/cellery-io/mesh-controller/pkg/client/listers/mesh/v1alpha1"
	"github.com/cellery-io/mesh-controller/pkg/controller"
	controllercommons "github.com/cellery-io/mesh-controller/pkg/controller/commons"
)

func CreateCellVirtualService(cell *v1alpha1.Cell, cellLister listers.CellLister) (*v1alpha3.VirtualService, error) {
	hostNames, httpRoutes, tcpRoutes, err := buildInterCellRoutingInfo(cell, cellLister)
	if err != nil {
		return nil, err
	}
	if len(hostNames) == 0 || (len(httpRoutes) == 0 && len(tcpRoutes) == 0) {
		// No virtual service needed
		return nil, nil
	}
	return &v1alpha3.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:      CellVirtualServiceName(cell),
			Namespace: cell.Namespace,
			Labels:    createLabels(cell),
			OwnerReferences: []metav1.OwnerReference{
				*controller.CreateCellOwnerRef(cell),
			},
		},
		Spec: v1alpha3.VirtualServiceSpec{
			Hosts:    hostNames,
			Gateways: []string{"mesh"},
			Http:     httpRoutes,
			// TCP is not supported atm
			// TODO: support TCP
			//Tcp:      tcpRoutes,
		},
	}, nil
}

func buildInterCellRoutingInfo(cell *v1alpha1.Cell, cellLister listers.CellLister) ([]string, []*v1alpha3.HTTPRoute, []*v1alpha3.TCPRoute, error) {
	var intercellHttpRoutes []*v1alpha3.HTTPRoute
	var intercellTcpRoutes []*v1alpha3.TCPRoute
	var hostNames []string
	// get dependencies from cell annotations,
	dependencies, err := controllercommons.ExtractDependencies(cell.Annotations)
	if err != nil {
		return nil, nil, nil, err
	}
	// if the source cell is a web cell, we need to create a few additional routing rules
	isWebCell := cell.Spec.GatewayTemplate.Spec.Host != ""
	// for each dependency, create a route
	for _, dependency := range dependencies {
		dependencyInst := dependency[controllercommons.Instance]
		if dependencyInst == "" {
			return nil, nil, nil, fmt.Errorf("unable to extract dependency instance from annotations")
		}
		dependencyKind := dependency[controllercommons.Kind]
		if dependencyKind == "" {
			return nil, nil, nil, fmt.Errorf("unable to extract dependency kind from annotations")
		}
		// retrieve the cell using the cell instance name
		depCell, err := cellLister.Cells(cell.Namespace).Get(dependencyInst)
		if err != nil {
			return nil, nil, nil, err
		}
		if dependencyKind == controllercommons.CellKind {
			depCell, err := cellLister.Cells(cell.Namespace).Get(dependencyInst)
			if err != nil {
				return nil, nil, nil, err
			}
			if len(depCell.Spec.GatewayTemplate.Spec.HTTPRoutes) > 0 {
				hostNames = append(hostNames, controllercommons.BuildHostNameForCellDependency(dependencyInst))
				// build http routes
				intercellHttpRoutes = append(intercellHttpRoutes, controllercommons.BuildHttpRoutesForCellDependency(cell.Name, dependencyInst, isWebCell)...)
			}
		} else if dependencyKind == controllercommons.CompositeKind {
			// retrieve the cell using the cell instance name
			depComposite, err := cellLister.Cells(cell.Namespace).Get(dependencyInst)
			if err != nil {
				return nil, nil, nil, err
			}
			if len(depComposite.Spec.ServiceTemplates) > 0 {
				hostNames = append(hostNames, controllercommons.BuildHostNamesForCompositeDependency(dependencyInst, depComposite.Spec.ServiceTemplates)...)
				intercellHttpRoutes = append(intercellHttpRoutes, controllercommons.BuildHttpRoutesForCompositeDependency(cell.Name, dependencyInst, depComposite.Spec.ServiceTemplates, isWebCell)...)
			}
		} else {
			// unknown dependency kind
			return nil, nil, nil, fmt.Errorf("unknown dependency kind '%s'", dependencyKind)
		}
		if len(depCell.Spec.GatewayTemplate.Spec.TCPRoutes) > 0 {
			// TCP is not supported atm
			// TODO: support TCP
			// hostNames = append(hostNames, buildHostName(dependencyInst))
			// build tcp routes
			// intercellTcpRoutes = append(intercellTcpRoutes, buildTcpRoutes(depCell, dependencyInst)...)
		}
	}

	return hostNames, intercellHttpRoutes, intercellTcpRoutes, nil
}

//func buildHostName(dependencyInst string) string {
//	return GatewayK8sServiceName(GatewayNameFromInstanceName(dependencyInst))
//}
//
//func buildHttpRoutes(cell *v1alpha1.Cell, dependencyInst string) []*v1alpha3.HTTPRoute {
//	instanceIdMatch1Rule := &v1alpha3.HTTPRoute{
//		Match: []*v1alpha3.HTTPMatchRequest{
//			{
//				Authority: &v1alpha3.StringMatch{
//					Regex: fmt.Sprintf("^(%s)(--gateway-service)(\\S*)$", dependencyInst),
//				},
//				Headers: map[string]*v1alpha3.StringMatch{
//					instanceId: {
//						Exact: "1",
//					},
//				},
//				SourceLabels: map[string]string{
//					mesh.CellLabelKeySource:      cell.Name,
//					mesh.ComponentLabelKeySource: "true",
//				},
//			},
//		},
//		Route: []*v1alpha3.DestinationWeight{
//			{
//				Destination: &v1alpha3.Destination{
//					Host: GatewayK8sServiceName(GatewayNameFromInstanceName(dependencyInst)),
//				},
//			},
//		},
//	}
//	instanceIdMatch2Rule := &v1alpha3.HTTPRoute{
//		Match: []*v1alpha3.HTTPMatchRequest{
//			{
//				Authority: &v1alpha3.StringMatch{
//					Regex: fmt.Sprintf("^(%s)(--gateway-service)(\\S*)$", dependencyInst),
//				},
//				Headers: map[string]*v1alpha3.StringMatch{
//					instanceId: {
//						Exact: "2",
//					},
//				},
//				SourceLabels: map[string]string{
//					mesh.CellLabelKeySource:      cell.Name,
//					mesh.ComponentLabelKeySource: "true",
//				},
//			},
//		},
//		Route: []*v1alpha3.DestinationWeight{
//			{
//				Destination: &v1alpha3.Destination{
//					Host: GatewayK8sServiceName(GatewayNameFromInstanceName(dependencyInst)),
//				},
//			},
//		},
//	}
//	percentageBasedRule := &v1alpha3.HTTPRoute{
//		Match: []*v1alpha3.HTTPMatchRequest{
//			{
//				Authority: &v1alpha3.StringMatch{
//					Regex: fmt.Sprintf("^(%s)(--gateway-service)(\\S*)$", dependencyInst),
//				},
//				SourceLabels: map[string]string{
//					mesh.CellLabelKeySource:      cell.Name,
//					mesh.ComponentLabelKeySource: "true",
//				},
//			},
//		},
//		Route: []*v1alpha3.DestinationWeight{
//			{
//				Destination: &v1alpha3.Destination{
//					Host: GatewayK8sServiceName(GatewayNameFromInstanceName(dependencyInst)),
//				},
//			},
//		},
//	}
//	return []*v1alpha3.HTTPRoute{instanceIdMatch1Rule, instanceIdMatch2Rule, percentageBasedRule}
//}
//
//func buildTcpRoutes(cell *v1alpha1.Cell, dependencyInst string) []*v1alpha3.TCPRoute {
//	var intercellRoutes []*v1alpha3.TCPRoute
//	for _, cellTcpRoute := range cell.Spec.GatewayTemplate.Spec.TCPRoutes {
//		route := v1alpha3.TCPRoute{
//			Match: []*v1alpha3.L4MatchAttributes{
//				{
//					Port: cellTcpRoute.Port,
//					SourceLabels: map[string]string{
//						mesh.CellLabelKey:      cell.Name,
//						mesh.ComponentLabelKey: "true",
//					},
//				},
//			},
//			Route: []*v1alpha3.DestinationWeight{
//				{
//					Destination: &v1alpha3.Destination{
//						Host: GatewayK8sServiceName(GatewayNameFromInstanceName(dependencyInst)),
//						Port: &v1alpha3.PortSelector{
//							Number: cellTcpRoute.Port,
//						},
//					},
//				},
//			},
//		}
//		intercellRoutes = append(intercellRoutes, &route)
//	}
//	return intercellRoutes
//}
//
//func extractDependencies(cell *v1alpha1.Cell) ([]map[string]string, error) {
//	cellDependencies := cell.Annotations[mesh.CellDependenciesAnnotationKey]
//	var dependencies []map[string]string
//	if cellDependencies == "" {
//		// no dependencies
//		return dependencies, nil
//	}
//	err := json.Unmarshal([]byte(cellDependencies), &dependencies)
//	if err != nil {
//		return dependencies, err
//	}
//	return dependencies, nil
//}
//
//func BuildVirtualServiceiedConfig(vs *v1alpha3.VirtualService) *v1alpha3.VirtualService {
//	return &v1alpha3.VirtualService{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "VirtualService",
//			APIVersion: v1alpha3.SchemeGroupVersion.String(),
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      vs.Name,
//			Namespace: vs.Namespace,
//		},
//		Spec: vs.Spec,
//	}
//}
//
//func Annotate(vs *v1alpha3.VirtualService, name string, value string) {
//	annotations := make(map[string]string, len(vs.ObjectMeta.Annotations)+1)
//	annotations[name] = value
//	for k, v := range vs.ObjectMeta.Annotations {
//		annotations[k] = v
//	}
//	vs.Annotations = annotations
//}
