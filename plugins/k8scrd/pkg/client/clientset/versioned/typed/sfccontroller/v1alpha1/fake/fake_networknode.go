// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/ligato/sfc-controller/plugins/k8scrd/pkg/apis/sfccontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkNodes implements NetworkNodeInterface
type FakeNetworkNodes struct {
	Fake *FakeSfccontrollerV1alpha1
	ns   string
}

var networknodesResource = schema.GroupVersionResource{Group: "sfccontroller.ligato.github.com", Version: "v1alpha1", Resource: "networknodes"}

var networknodesKind = schema.GroupVersionKind{Group: "sfccontroller.ligato.github.com", Version: "v1alpha1", Kind: "NetworkNode"}

// Get takes name of the networkNode, and returns the corresponding networkNode object, and an error if there is any.
func (c *FakeNetworkNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networknodesResource, c.ns, name), &v1alpha1.NetworkNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkNode), err
}

// List takes label and field selectors, and returns the list of NetworkNodes that match those selectors.
func (c *FakeNetworkNodes) List(opts v1.ListOptions) (result *v1alpha1.NetworkNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networknodesResource, networknodesKind, c.ns, opts), &v1alpha1.NetworkNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkNodeList{}
	for _, item := range obj.(*v1alpha1.NetworkNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkNodes.
func (c *FakeNetworkNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networknodesResource, c.ns, opts))

}

// Create takes the representation of a networkNode and creates it.  Returns the server's representation of the networkNode, and an error, if there is any.
func (c *FakeNetworkNodes) Create(networkNode *v1alpha1.NetworkNode) (result *v1alpha1.NetworkNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networknodesResource, c.ns, networkNode), &v1alpha1.NetworkNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkNode), err
}

// Update takes the representation of a networkNode and updates it. Returns the server's representation of the networkNode, and an error, if there is any.
func (c *FakeNetworkNodes) Update(networkNode *v1alpha1.NetworkNode) (result *v1alpha1.NetworkNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networknodesResource, c.ns, networkNode), &v1alpha1.NetworkNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkNode), err
}

// Delete takes name of the networkNode and deletes it. Returns an error if one occurs.
func (c *FakeNetworkNodes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networknodesResource, c.ns, name), &v1alpha1.NetworkNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networknodesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkNodeList{})
	return err
}

// Patch applies the patch and returns the patched networkNode.
func (c *FakeNetworkNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networknodesResource, c.ns, name, data, subresources...), &v1alpha1.NetworkNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkNode), err
}
