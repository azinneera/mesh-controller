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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned/typed/mesh/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMeshV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMeshV1alpha1) AutoscalePolicies(namespace string) v1alpha1.AutoscalePolicyInterface {
	return &FakeAutoscalePolicies{c, namespace}
}

func (c *FakeMeshV1alpha1) Cells(namespace string) v1alpha1.CellInterface {
	return &FakeCells{c, namespace}
}

func (c *FakeMeshV1alpha1) Composites(namespace string) v1alpha1.CompositeInterface {
	return &FakeComposites{c, namespace}
}

func (c *FakeMeshV1alpha1) Gateways(namespace string) v1alpha1.GatewayInterface {
	return &FakeGateways{c, namespace}
}

func (c *FakeMeshV1alpha1) Services(namespace string) v1alpha1.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeMeshV1alpha1) TokenServices(namespace string) v1alpha1.TokenServiceInterface {
	return &FakeTokenServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMeshV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
