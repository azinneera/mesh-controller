/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
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

package v1alpha1

import (
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	scheme "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TokenServicesGetter has a method to return a TokenServiceInterface.
// A group's client should implement this interface.
type TokenServicesGetter interface {
	TokenServices(namespace string) TokenServiceInterface
}

// TokenServiceInterface has methods to work with TokenService resources.
type TokenServiceInterface interface {
	Create(*v1alpha1.TokenService) (*v1alpha1.TokenService, error)
	Update(*v1alpha1.TokenService) (*v1alpha1.TokenService, error)
	UpdateStatus(*v1alpha1.TokenService) (*v1alpha1.TokenService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TokenService, error)
	List(opts v1.ListOptions) (*v1alpha1.TokenServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TokenService, err error)
	TokenServiceExpansion
}

// tokenServices implements TokenServiceInterface
type tokenServices struct {
	client rest.Interface
	ns     string
}

// newTokenServices returns a TokenServices
func newTokenServices(c *MeshV1alpha1Client, namespace string) *tokenServices {
	return &tokenServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tokenService, and returns the corresponding tokenService object, and an error if there is any.
func (c *tokenServices) Get(name string, options v1.GetOptions) (result *v1alpha1.TokenService, err error) {
	result = &v1alpha1.TokenService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tokenservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TokenServices that match those selectors.
func (c *tokenServices) List(opts v1.ListOptions) (result *v1alpha1.TokenServiceList, err error) {
	result = &v1alpha1.TokenServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tokenservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tokenServices.
func (c *tokenServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tokenservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a tokenService and creates it.  Returns the server's representation of the tokenService, and an error, if there is any.
func (c *tokenServices) Create(tokenService *v1alpha1.TokenService) (result *v1alpha1.TokenService, err error) {
	result = &v1alpha1.TokenService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tokenservices").
		Body(tokenService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tokenService and updates it. Returns the server's representation of the tokenService, and an error, if there is any.
func (c *tokenServices) Update(tokenService *v1alpha1.TokenService) (result *v1alpha1.TokenService, err error) {
	result = &v1alpha1.TokenService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tokenservices").
		Name(tokenService.Name).
		Body(tokenService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *tokenServices) UpdateStatus(tokenService *v1alpha1.TokenService) (result *v1alpha1.TokenService, err error) {
	result = &v1alpha1.TokenService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tokenservices").
		Name(tokenService.Name).
		SubResource("status").
		Body(tokenService).
		Do().
		Into(result)
	return
}

// Delete takes name of the tokenService and deletes it. Returns an error if one occurs.
func (c *tokenServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokenservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tokenServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokenservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tokenService.
func (c *tokenServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TokenService, err error) {
	result = &v1alpha1.TokenService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tokenservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
