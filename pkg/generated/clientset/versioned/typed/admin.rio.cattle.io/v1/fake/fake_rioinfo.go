/*
Copyright 2020 Rancher Labs.

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
	adminriocattleiov1 "github.com/rancher/rio/pkg/apis/admin.rio.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRioInfos implements RioInfoInterface
type FakeRioInfos struct {
	Fake *FakeAdminV1
}

var rioinfosResource = schema.GroupVersionResource{Group: "admin.rio.cattle.io", Version: "v1", Resource: "rioinfos"}

var rioinfosKind = schema.GroupVersionKind{Group: "admin.rio.cattle.io", Version: "v1", Kind: "RioInfo"}

// Get takes name of the rioInfo, and returns the corresponding rioInfo object, and an error if there is any.
func (c *FakeRioInfos) Get(name string, options v1.GetOptions) (result *adminriocattleiov1.RioInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rioinfosResource, name), &adminriocattleiov1.RioInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*adminriocattleiov1.RioInfo), err
}

// List takes label and field selectors, and returns the list of RioInfos that match those selectors.
func (c *FakeRioInfos) List(opts v1.ListOptions) (result *adminriocattleiov1.RioInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rioinfosResource, rioinfosKind, opts), &adminriocattleiov1.RioInfoList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &adminriocattleiov1.RioInfoList{ListMeta: obj.(*adminriocattleiov1.RioInfoList).ListMeta}
	for _, item := range obj.(*adminriocattleiov1.RioInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rioInfos.
func (c *FakeRioInfos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rioinfosResource, opts))
}

// Create takes the representation of a rioInfo and creates it.  Returns the server's representation of the rioInfo, and an error, if there is any.
func (c *FakeRioInfos) Create(rioInfo *adminriocattleiov1.RioInfo) (result *adminriocattleiov1.RioInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rioinfosResource, rioInfo), &adminriocattleiov1.RioInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*adminriocattleiov1.RioInfo), err
}

// Update takes the representation of a rioInfo and updates it. Returns the server's representation of the rioInfo, and an error, if there is any.
func (c *FakeRioInfos) Update(rioInfo *adminriocattleiov1.RioInfo) (result *adminriocattleiov1.RioInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rioinfosResource, rioInfo), &adminriocattleiov1.RioInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*adminriocattleiov1.RioInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRioInfos) UpdateStatus(rioInfo *adminriocattleiov1.RioInfo) (*adminriocattleiov1.RioInfo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(rioinfosResource, "status", rioInfo), &adminriocattleiov1.RioInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*adminriocattleiov1.RioInfo), err
}

// Delete takes name of the rioInfo and deletes it. Returns an error if one occurs.
func (c *FakeRioInfos) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(rioinfosResource, name), &adminriocattleiov1.RioInfo{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRioInfos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rioinfosResource, listOptions)

	_, err := c.Fake.Invokes(action, &adminriocattleiov1.RioInfoList{})
	return err
}

// Patch applies the patch and returns the patched rioInfo.
func (c *FakeRioInfos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *adminriocattleiov1.RioInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rioinfosResource, name, pt, data, subresources...), &adminriocattleiov1.RioInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*adminriocattleiov1.RioInfo), err
}
