/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// This file was automatically generated by informer-gen

package internalversion

import (
	security "github.com/openshift/origin/pkg/security/apis/security"
	internalinterfaces "github.com/openshift/origin/pkg/security/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/security/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/security/generated/listers/security/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// SecurityContextConstraintsInformer provides access to a shared informer and lister for
// SecurityContextConstraints.
type SecurityContextConstraintsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.SecurityContextConstraintsLister
}

type securityContextConstraintsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSecurityContextConstraintsInformer constructs a new informer for SecurityContextConstraints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecurityContextConstraintsInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecurityContextConstraintsInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSecurityContextConstraintsInformer constructs a new informer for SecurityContextConstraints type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecurityContextConstraintsInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Security().SecurityContextConstraints().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Security().SecurityContextConstraints().Watch(options)
			},
		},
		&security.SecurityContextConstraints{},
		resyncPeriod,
		indexers,
	)
}

func (f *securityContextConstraintsInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecurityContextConstraintsInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *securityContextConstraintsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&security.SecurityContextConstraints{}, f.defaultInformer)
}

func (f *securityContextConstraintsInformer) Lister() internalversion.SecurityContextConstraintsLister {
	return internalversion.NewSecurityContextConstraintsLister(f.Informer().GetIndexer())
}