/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	dnsv1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	versioned "github.com/gardener/external-dns-management/pkg/client/dns/clientset/versioned"
	internalinterfaces "github.com/gardener/external-dns-management/pkg/client/dns/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/external-dns-management/pkg/client/dns/listers/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSResourceWatchInformer provides access to a shared informer and lister for
// DNSResourceWatches.
type DNSResourceWatchInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DNSResourceWatchLister
}

type dNSResourceWatchInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSResourceWatchInformer constructs a new informer for DNSAnnotation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSResourceWatchInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSResourceWatchInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSResourceWatchInformer constructs a new informer for DNSAnnotation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSResourceWatchInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSResourceWatches(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSResourceWatches(namespace).Watch(options)
			},
		},
		&dnsv1alpha1.DNSAnnotation{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSResourceWatchInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSResourceWatchInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSResourceWatchInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dnsv1alpha1.DNSAnnotation{}, f.defaultInformer)
}

func (f *dNSResourceWatchInformer) Lister() v1alpha1.DNSResourceWatchLister {
	return v1alpha1.NewDNSResourceWatchLister(f.Informer().GetIndexer())
}