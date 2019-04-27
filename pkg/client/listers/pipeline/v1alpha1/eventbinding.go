/*
Copyright 2018 The Knative Authors

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
package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventBindingLister helps list EventBindings.
type EventBindingLister interface {
	// List lists all EventBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EventBinding, err error)
	// EventBindings returns an object that can list and get EventBindings.
	EventBindings(namespace string) EventBindingNamespaceLister
	EventBindingListerExpansion
}

// eventBindingLister implements the EventBindingLister interface.
type eventBindingLister struct {
	indexer cache.Indexer
}

// NewEventBindingLister returns a new EventBindingLister.
func NewEventBindingLister(indexer cache.Indexer) EventBindingLister {
	return &eventBindingLister{indexer: indexer}
}

// List lists all EventBindings in the indexer.
func (s *eventBindingLister) List(selector labels.Selector) (ret []*v1alpha1.EventBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventBinding))
	})
	return ret, err
}

// EventBindings returns an object that can list and get EventBindings.
func (s *eventBindingLister) EventBindings(namespace string) EventBindingNamespaceLister {
	return eventBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventBindingNamespaceLister helps list and get EventBindings.
type EventBindingNamespaceLister interface {
	// List lists all EventBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EventBinding, err error)
	// Get retrieves the EventBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EventBinding, error)
	EventBindingNamespaceListerExpansion
}

// eventBindingNamespaceLister implements the EventBindingNamespaceLister
// interface.
type eventBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventBindings in the indexer for a given namespace.
func (s eventBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventBinding))
	})
	return ret, err
}

// Get retrieves the EventBinding from the indexer for a given namespace and name.
func (s eventBindingNamespaceLister) Get(name string) (*v1alpha1.EventBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventbinding"), name)
	}
	return obj.(*v1alpha1.EventBinding), nil
}
