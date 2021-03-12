/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	scheme "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClickHouseOperatorConfigurationsGetter has a method to return a ClickHouseOperatorConfigurationInterface.
// A group's client should implement this interface.
type ClickHouseOperatorConfigurationsGetter interface {
	ClickHouseOperatorConfigurations(namespace string) ClickHouseOperatorConfigurationInterface
}

// ClickHouseOperatorConfigurationInterface has methods to work with ClickHouseOperatorConfiguration resources.
type ClickHouseOperatorConfigurationInterface interface {
	Create(*v1.ClickHouseOperatorConfiguration) (*v1.ClickHouseOperatorConfiguration, error)
	Update(*v1.ClickHouseOperatorConfiguration) (*v1.ClickHouseOperatorConfiguration, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClickHouseOperatorConfiguration, error)
	List(opts metav1.ListOptions) (*v1.ClickHouseOperatorConfigurationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClickHouseOperatorConfiguration, err error)
	ClickHouseOperatorConfigurationExpansion
}

// clickHouseOperatorConfigurations implements ClickHouseOperatorConfigurationInterface
type clickHouseOperatorConfigurations struct {
	client rest.Interface
	ns     string
}

// newClickHouseOperatorConfigurations returns a ClickHouseOperatorConfigurations
func newClickHouseOperatorConfigurations(c *ClickhouseV1Client, namespace string) *clickHouseOperatorConfigurations {
	return &clickHouseOperatorConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clickHouseOperatorConfiguration, and returns the corresponding clickHouseOperatorConfiguration object, and an error if there is any.
func (c *clickHouseOperatorConfigurations) Get(name string, options metav1.GetOptions) (result *v1.ClickHouseOperatorConfiguration, err error) {
	result = &v1.ClickHouseOperatorConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClickHouseOperatorConfigurations that match those selectors.
func (c *clickHouseOperatorConfigurations) List(opts metav1.ListOptions) (result *v1.ClickHouseOperatorConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClickHouseOperatorConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clickHouseOperatorConfigurations.
func (c *clickHouseOperatorConfigurations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clickHouseOperatorConfiguration and creates it.  Returns the server's representation of the clickHouseOperatorConfiguration, and an error, if there is any.
func (c *clickHouseOperatorConfigurations) Create(clickHouseOperatorConfiguration *v1.ClickHouseOperatorConfiguration) (result *v1.ClickHouseOperatorConfiguration, err error) {
	result = &v1.ClickHouseOperatorConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		Body(clickHouseOperatorConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clickHouseOperatorConfiguration and updates it. Returns the server's representation of the clickHouseOperatorConfiguration, and an error, if there is any.
func (c *clickHouseOperatorConfigurations) Update(clickHouseOperatorConfiguration *v1.ClickHouseOperatorConfiguration) (result *v1.ClickHouseOperatorConfiguration, err error) {
	result = &v1.ClickHouseOperatorConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		Name(clickHouseOperatorConfiguration.Name).
		Body(clickHouseOperatorConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the clickHouseOperatorConfiguration and deletes it. Returns an error if one occurs.
func (c *clickHouseOperatorConfigurations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clickHouseOperatorConfigurations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clickHouseOperatorConfiguration.
func (c *clickHouseOperatorConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClickHouseOperatorConfiguration, err error) {
	result = &v1.ClickHouseOperatorConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clickhouseoperatorconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}