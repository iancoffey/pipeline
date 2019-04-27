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
	scheme "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventBindingsGetter has a method to return a EventBindingInterface.
// A group's client should implement this interface.
type EventBindingsGetter interface {
	EventBindings(namespace string) EventBindingInterface
}

// EventBindingInterface has methods to work with EventBinding resources.
type EventBindingInterface interface {
	Create(*v1alpha1.EventBinding) (*v1alpha1.EventBinding, error)
	Update(*v1alpha1.EventBinding) (*v1alpha1.EventBinding, error)
	UpdateStatus(*v1alpha1.EventBinding) (*v1alpha1.EventBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.EventBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventBinding, err error)
	EventBindingExpansion
}

// eventBindings implements EventBindingInterface
type eventBindings struct {
	client rest.Interface
	ns     string
}

// newEventBindings returns a EventBindings
func newEventBindings(c *TektonV1alpha1Client, namespace string) *eventBindings {
	return &eventBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventBinding, and returns the corresponding eventBinding object, and an error if there is any.
func (c *eventBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.EventBinding, err error) {
	result = &v1alpha1.EventBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventBindings that match those selectors.
func (c *eventBindings) List(opts v1.ListOptions) (result *v1alpha1.EventBindingList, err error) {
	result = &v1alpha1.EventBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventBindings.
func (c *eventBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eventBinding and creates it.  Returns the server's representation of the eventBinding, and an error, if there is any.
func (c *eventBindings) Create(eventBinding *v1alpha1.EventBinding) (result *v1alpha1.EventBinding, err error) {
	result = &v1alpha1.EventBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventbindings").
		Body(eventBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventBinding and updates it. Returns the server's representation of the eventBinding, and an error, if there is any.
func (c *eventBindings) Update(eventBinding *v1alpha1.EventBinding) (result *v1alpha1.EventBinding, err error) {
	result = &v1alpha1.EventBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventbindings").
		Name(eventBinding.Name).
		Body(eventBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eventBindings) UpdateStatus(eventBinding *v1alpha1.EventBinding) (result *v1alpha1.EventBinding, err error) {
	result = &v1alpha1.EventBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventbindings").
		Name(eventBinding.Name).
		SubResource("status").
		Body(eventBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventBinding and deletes it. Returns an error if one occurs.
func (c *eventBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventBinding.
func (c *eventBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventBinding, err error) {
	result = &v1alpha1.EventBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
