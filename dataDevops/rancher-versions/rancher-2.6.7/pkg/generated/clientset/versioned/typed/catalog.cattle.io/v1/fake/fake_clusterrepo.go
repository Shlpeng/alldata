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

// FakeClusterRepos implements ClusterRepoInterface
type FakeClusterRepos struct {
	Fake *FakeCatalogV1
}

var clusterreposResource = schema.GroupVersionResource{Group: "catalog.cattle.io", Version: "v1", Resource: "clusterrepos"}

var clusterreposKind = schema.GroupVersionKind{Group: "catalog.cattle.io", Version: "v1", Kind: "ClusterRepo"}

// Get takes name of the clusterRepo, and returns the corresponding clusterRepo object, and an error if there is any.
func (c *FakeClusterRepos) Get(ctx context.Context, name string, options v1.GetOptions) (result *catalogcattleiov1.ClusterRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterreposResource, name), &catalogcattleiov1.ClusterRepo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.ClusterRepo), err
}

// List takes label and field selectors, and returns the list of ClusterRepos that match those selectors.
func (c *FakeClusterRepos) List(ctx context.Context, opts v1.ListOptions) (result *catalogcattleiov1.ClusterRepoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterreposResource, clusterreposKind, opts), &catalogcattleiov1.ClusterRepoList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &catalogcattleiov1.ClusterRepoList{ListMeta: obj.(*catalogcattleiov1.ClusterRepoList).ListMeta}
	for _, item := range obj.(*catalogcattleiov1.ClusterRepoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRepos.
func (c *FakeClusterRepos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterreposResource, opts))
}

// Create takes the representation of a clusterRepo and creates it.  Returns the server's representation of the clusterRepo, and an error, if there is any.
func (c *FakeClusterRepos) Create(ctx context.Context, clusterRepo *catalogcattleiov1.ClusterRepo, opts v1.CreateOptions) (result *catalogcattleiov1.ClusterRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterreposResource, clusterRepo), &catalogcattleiov1.ClusterRepo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.ClusterRepo), err
}

// Update takes the representation of a clusterRepo and updates it. Returns the server's representation of the clusterRepo, and an error, if there is any.
func (c *FakeClusterRepos) Update(ctx context.Context, clusterRepo *catalogcattleiov1.ClusterRepo, opts v1.UpdateOptions) (result *catalogcattleiov1.ClusterRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterreposResource, clusterRepo), &catalogcattleiov1.ClusterRepo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.ClusterRepo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRepos) UpdateStatus(ctx context.Context, clusterRepo *catalogcattleiov1.ClusterRepo, opts v1.UpdateOptions) (*catalogcattleiov1.ClusterRepo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterreposResource, "status", clusterRepo), &catalogcattleiov1.ClusterRepo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.ClusterRepo), err
}

// Delete takes name of the clusterRepo and deletes it. Returns an error if one occurs.
func (c *FakeClusterRepos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterreposResource, name, opts), &catalogcattleiov1.ClusterRepo{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRepos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterreposResource, listOpts)

	_, err := c.Fake.Invokes(action, &catalogcattleiov1.ClusterRepoList{})
	return err
}

// Patch applies the patch and returns the patched clusterRepo.
func (c *FakeClusterRepos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *catalogcattleiov1.ClusterRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterreposResource, name, pt, data, subresources...), &catalogcattleiov1.ClusterRepo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*catalogcattleiov1.ClusterRepo), err
}
