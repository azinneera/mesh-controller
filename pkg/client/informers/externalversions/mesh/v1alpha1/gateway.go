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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	meshv1alpha1 "github.com/cellery-io/mesh-controller/pkg/apis/mesh/v1alpha1"
	versioned "github.com/cellery-io/mesh-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cellery-io/mesh-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cellery-io/mesh-controller/pkg/client/listers/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GatewayInformer provides access to a shared informer and lister for
// Gateways.
type GatewayInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GatewayLister
}

type gatewayInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGatewayInformer constructs a new informer for Gateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGatewayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGatewayInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGatewayInformer constructs a new informer for Gateway type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGatewayInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().Gateways(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().Gateways(namespace).Watch(options)
			},
		},
		&meshv1alpha1.Gateway{},
		resyncPeriod,
		indexers,
	)
}

func (f *gatewayInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGatewayInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gatewayInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meshv1alpha1.Gateway{}, f.defaultInformer)
}

func (f *gatewayInformer) Lister() v1alpha1.GatewayLister {
	return v1alpha1.NewGatewayLister(f.Informer().GetIndexer())
}
