/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	catalogcattleiov1 "github.com/rancher/rancher/pkg/apis/catalog.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperations implements OperationInterface
type FakeOperations struct {
	Fake *FakeCatalogV1
	ns   string
}

var operationsResource = schema.GroupVersionResource{Group: "catalog.cattle.io", Version: "v1", Resource: "operations"}

var operationsKind = schema.GroupVersionKind{Group: "catalog.cattle.io", Version: "v1", Kind: "Operation"}

// Get takes name of the operation, and returns the corresponding operation object, and an error if there is any.
func (c *FakeOperations) Get(ctx context.Context, name string, options v1.GetOptions) (result *catalogcattleiov1.Operation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operationsResource, c.ns, name), &catalogcattleiov1.Operation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.Operation), err
}

// List takes label and field selectors, and returns the list of Operations that match those selectors.
func (c *FakeOperations) List(ctx context.Context, opts v1.ListOptions) (result *catalogcattleiov1.OperationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operationsResource, operationsKind, c.ns, opts), &catalogcattleiov1.OperationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &catalogcattleiov1.OperationList{ListMeta: obj.(*catalogcattleiov1.OperationList).ListMeta}
	for _, item := range obj.(*catalogcattleiov1.OperationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operations.
func (c *FakeOperations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operationsResource, c.ns, opts))

}

// Create takes the representation of a operation and creates it.  Returns the server's representation of the operation, and an error, if there is any.
func (c *FakeOperations) Create(ctx context.Context, operation *catalogcattleiov1.Operation, opts v1.CreateOptions) (result *catalogcattleiov1.Operation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operationsResource, c.ns, operation), &catalogcattleiov1.Operation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.Operation), err
}

// Update takes the representation of a operation and updates it. Returns the server's representation of the operation, and an error, if there is any.
func (c *FakeOperations) Update(ctx context.Context, operation *catalogcattleiov1.Operation, opts v1.UpdateOptions) (result *catalogcattleiov1.Operation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operationsResource, c.ns, operation), &catalogcattleiov1.Operation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.Operation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperations) UpdateStatus(ctx context.Context, operation *catalogcattleiov1.Operation, opts v1.UpdateOptions) (*catalogcattleiov1.Operation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(operationsResource, "status", c.ns, operation), &catalogcattleiov1.Operation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.Operation), err
}

// Delete takes name of the operation and deletes it. Returns an error if one occurs.
func (c *FakeOperations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(operationsResource, c.ns, name, opts), &catalogcattleiov1.Operation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &catalogcattleiov1.OperationList{})
	return err
}

// Patch applies the patch and returns the patched operation.
func (c *FakeOperations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *catalogcattleiov1.Operation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operationsResource, c.ns, name, pt, data, subresources...), &catalogcattleiov1.Operation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.Operation), err
}
