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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	"github.com/cellery-io/mesh-controller/pkg/controller"
)

func CreateService(composite *v1alpha1.Composite, serviceTemplate v1alpha1.ServiceTemplateSpec) *v1alpha1.Service {
	serviceSpec := serviceTemplate.Spec.DeepCopy()
	serviceSpec.Container.Name = serviceTemplate.Name

	// Default to Deployment if not specified
	if len(serviceSpec.Type) == 0 {
		serviceSpec.Type = v1alpha1.ServiceTypeDeployment
	}

	return &v1alpha1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        ServiceName(composite, serviceTemplate),
			Namespace:   composite.Namespace,
			Labels:      addTokenServiceLabels(createLabels(composite)),
			Annotations: createServiceAnnotations(composite),
			OwnerReferences: []metav1.OwnerReference{
				*controller.CreateCompositeOwnerRef(composite),
			},
		},
		Spec: *serviceSpec,
	}
}
