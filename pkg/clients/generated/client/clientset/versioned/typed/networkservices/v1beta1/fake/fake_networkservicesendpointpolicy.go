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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesEndpointPolicies implements NetworkServicesEndpointPolicyInterface
type FakeNetworkServicesEndpointPolicies struct {
	Fake *FakeNetworkservicesV1beta1
	ns   string
}

var networkservicesendpointpoliciesResource = schema.GroupVersionResource{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Resource: "networkservicesendpointpolicies"}

var networkservicesendpointpoliciesKind = schema.GroupVersionKind{Group: "networkservices.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkServicesEndpointPolicy"}

// Get takes name of the networkServicesEndpointPolicy, and returns the corresponding networkServicesEndpointPolicy object, and an error if there is any.
func (c *FakeNetworkServicesEndpointPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkServicesEndpointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesendpointpoliciesResource, c.ns, name), &v1beta1.NetworkServicesEndpointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesEndpointPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkServicesEndpointPolicies that match those selectors.
func (c *FakeNetworkServicesEndpointPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkServicesEndpointPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesendpointpoliciesResource, networkservicesendpointpoliciesKind, c.ns, opts), &v1beta1.NetworkServicesEndpointPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkServicesEndpointPolicyList{ListMeta: obj.(*v1beta1.NetworkServicesEndpointPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkServicesEndpointPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesEndpointPolicies.
func (c *FakeNetworkServicesEndpointPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesendpointpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesEndpointPolicy and creates it.  Returns the server's representation of the networkServicesEndpointPolicy, and an error, if there is any.
func (c *FakeNetworkServicesEndpointPolicies) Create(ctx context.Context, networkServicesEndpointPolicy *v1beta1.NetworkServicesEndpointPolicy, opts v1.CreateOptions) (result *v1beta1.NetworkServicesEndpointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesendpointpoliciesResource, c.ns, networkServicesEndpointPolicy), &v1beta1.NetworkServicesEndpointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesEndpointPolicy), err
}

// Update takes the representation of a networkServicesEndpointPolicy and updates it. Returns the server's representation of the networkServicesEndpointPolicy, and an error, if there is any.
func (c *FakeNetworkServicesEndpointPolicies) Update(ctx context.Context, networkServicesEndpointPolicy *v1beta1.NetworkServicesEndpointPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkServicesEndpointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesendpointpoliciesResource, c.ns, networkServicesEndpointPolicy), &v1beta1.NetworkServicesEndpointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesEndpointPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesEndpointPolicies) UpdateStatus(ctx context.Context, networkServicesEndpointPolicy *v1beta1.NetworkServicesEndpointPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkServicesEndpointPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicesendpointpoliciesResource, "status", c.ns, networkServicesEndpointPolicy), &v1beta1.NetworkServicesEndpointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesEndpointPolicy), err
}

// Delete takes name of the networkServicesEndpointPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesEndpointPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesendpointpoliciesResource, c.ns, name, opts), &v1beta1.NetworkServicesEndpointPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesEndpointPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesendpointpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkServicesEndpointPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesEndpointPolicy.
func (c *FakeNetworkServicesEndpointPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkServicesEndpointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesendpointpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkServicesEndpointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesEndpointPolicy), err
}
