// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networksecurity/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkSecurityClientTLSPolicies implements NetworkSecurityClientTLSPolicyInterface
type FakeNetworkSecurityClientTLSPolicies struct {
	Fake *FakeNetworksecurityV1beta1
	ns   string
}

var networksecurityclienttlspoliciesResource = schema.GroupVersionResource{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Resource: "networksecurityclienttlspolicies"}

var networksecurityclienttlspoliciesKind = schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityClientTLSPolicy"}

// Get takes name of the networkSecurityClientTLSPolicy, and returns the corresponding networkSecurityClientTLSPolicy object, and an error if there is any.
func (c *FakeNetworkSecurityClientTLSPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkSecurityClientTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksecurityclienttlspoliciesResource, c.ns, name), &v1beta1.NetworkSecurityClientTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityClientTLSPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkSecurityClientTLSPolicies that match those selectors.
func (c *FakeNetworkSecurityClientTLSPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkSecurityClientTLSPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksecurityclienttlspoliciesResource, networksecurityclienttlspoliciesKind, c.ns, opts), &v1beta1.NetworkSecurityClientTLSPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkSecurityClientTLSPolicyList{ListMeta: obj.(*v1beta1.NetworkSecurityClientTLSPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkSecurityClientTLSPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSecurityClientTLSPolicies.
func (c *FakeNetworkSecurityClientTLSPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksecurityclienttlspoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkSecurityClientTLSPolicy and creates it.  Returns the server's representation of the networkSecurityClientTLSPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityClientTLSPolicies) Create(ctx context.Context, networkSecurityClientTLSPolicy *v1beta1.NetworkSecurityClientTLSPolicy, opts v1.CreateOptions) (result *v1beta1.NetworkSecurityClientTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksecurityclienttlspoliciesResource, c.ns, networkSecurityClientTLSPolicy), &v1beta1.NetworkSecurityClientTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityClientTLSPolicy), err
}

// Update takes the representation of a networkSecurityClientTLSPolicy and updates it. Returns the server's representation of the networkSecurityClientTLSPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityClientTLSPolicies) Update(ctx context.Context, networkSecurityClientTLSPolicy *v1beta1.NetworkSecurityClientTLSPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkSecurityClientTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksecurityclienttlspoliciesResource, c.ns, networkSecurityClientTLSPolicy), &v1beta1.NetworkSecurityClientTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityClientTLSPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkSecurityClientTLSPolicies) UpdateStatus(ctx context.Context, networkSecurityClientTLSPolicy *v1beta1.NetworkSecurityClientTLSPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkSecurityClientTLSPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networksecurityclienttlspoliciesResource, "status", c.ns, networkSecurityClientTLSPolicy), &v1beta1.NetworkSecurityClientTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityClientTLSPolicy), err
}

// Delete takes name of the networkSecurityClientTLSPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSecurityClientTLSPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networksecurityclienttlspoliciesResource, c.ns, name, opts), &v1beta1.NetworkSecurityClientTLSPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSecurityClientTLSPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksecurityclienttlspoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkSecurityClientTLSPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkSecurityClientTLSPolicy.
func (c *FakeNetworkSecurityClientTLSPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkSecurityClientTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksecurityclienttlspoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkSecurityClientTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityClientTLSPolicy), err
}
