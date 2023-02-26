/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "k8s.io/cloud-provider-gcp/crd/apis/gcpfirewall/v1beta1"
)

// GCPFirewallLister helps list GCPFirewalls.
// All objects returned here must be treated as read-only.
type GCPFirewallLister interface {
	// List lists all GCPFirewalls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.GCPFirewall, err error)
	// GCPFirewalls returns an object that can list and get GCPFirewalls.
	GCPFirewalls(namespace string) GCPFirewallNamespaceLister
	GCPFirewallListerExpansion
}

// gCPFirewallLister implements the GCPFirewallLister interface.
type gCPFirewallLister struct {
	indexer cache.Indexer
}

// NewGCPFirewallLister returns a new GCPFirewallLister.
func NewGCPFirewallLister(indexer cache.Indexer) GCPFirewallLister {
	return &gCPFirewallLister{indexer: indexer}
}

// List lists all GCPFirewalls in the indexer.
func (s *gCPFirewallLister) List(selector labels.Selector) (ret []*v1beta1.GCPFirewall, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.GCPFirewall))
	})
	return ret, err
}

// GCPFirewalls returns an object that can list and get GCPFirewalls.
func (s *gCPFirewallLister) GCPFirewalls(namespace string) GCPFirewallNamespaceLister {
	return gCPFirewallNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPFirewallNamespaceLister helps list and get GCPFirewalls.
// All objects returned here must be treated as read-only.
type GCPFirewallNamespaceLister interface {
	// List lists all GCPFirewalls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.GCPFirewall, err error)
	// Get retrieves the GCPFirewall from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.GCPFirewall, error)
	GCPFirewallNamespaceListerExpansion
}

// gCPFirewallNamespaceLister implements the GCPFirewallNamespaceLister
// interface.
type gCPFirewallNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPFirewalls in the indexer for a given namespace.
func (s gCPFirewallNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.GCPFirewall, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.GCPFirewall))
	})
	return ret, err
}

// Get retrieves the GCPFirewall from the indexer for a given namespace and name.
func (s gCPFirewallNamespaceLister) Get(name string) (*v1beta1.GCPFirewall, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("gcpfirewall"), name)
	}
	return obj.(*v1beta1.GCPFirewall), nil
}
