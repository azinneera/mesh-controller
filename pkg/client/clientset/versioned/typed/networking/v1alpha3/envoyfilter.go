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

package v1alpha3

import (
	v1alpha3 "github.com/cellery-io/mesh-controller/pkg/apis/istio/networking/v1alpha3"
	scheme "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnvoyFiltersGetter has a method to return a EnvoyFilterInterface.
// A group's client should implement this interface.
type EnvoyFiltersGetter interface {
	EnvoyFilters(namespace string) EnvoyFilterInterface
}

// EnvoyFilterInterface has methods to work with EnvoyFilter resources.
type EnvoyFilterInterface interface {
	Create(*v1alpha3.EnvoyFilter) (*v1alpha3.EnvoyFilter, error)
	Update(*v1alpha3.EnvoyFilter) (*v1alpha3.EnvoyFilter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha3.EnvoyFilter, error)
	List(opts v1.ListOptions) (*v1alpha3.EnvoyFilterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.EnvoyFilter, err error)
	EnvoyFilterExpansion
}

// envoyFilters implements EnvoyFilterInterface
type envoyFilters struct {
	client rest.Interface
	ns     string
}

// newEnvoyFilters returns a EnvoyFilters
func newEnvoyFilters(c *NetworkingV1alpha3Client, namespace string) *envoyFilters {
	return &envoyFilters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the envoyFilter, and returns the corresponding envoyFilter object, and an error if there is any.
func (c *envoyFilters) Get(name string, options v1.GetOptions) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EnvoyFilters that match those selectors.
func (c *envoyFilters) List(opts v1.ListOptions) (result *v1alpha3.EnvoyFilterList, err error) {
	result = &v1alpha3.EnvoyFilterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested envoyFilters.
func (c *envoyFilters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a envoyFilter and creates it.  Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *envoyFilters) Create(envoyFilter *v1alpha3.EnvoyFilter) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("envoyfilters").
		Body(envoyFilter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a envoyFilter and updates it. Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *envoyFilters) Update(envoyFilter *v1alpha3.EnvoyFilter) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(envoyFilter.Name).
		Body(envoyFilter).
		Do().
		Into(result)
	return
}

// Delete takes name of the envoyFilter and deletes it. Returns an error if one occurs.
func (c *envoyFilters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("envoyfilters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *envoyFilters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("envoyfilters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched envoyFilter.
func (c *envoyFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.EnvoyFilter, err error) {
	result = &v1alpha3.EnvoyFilter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("envoyfilters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
