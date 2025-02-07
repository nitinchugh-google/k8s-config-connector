//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformDefaultSupportedIDPConfig) DeepCopyInto(out *IdentityPlatformDefaultSupportedIDPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformDefaultSupportedIDPConfig.
func (in *IdentityPlatformDefaultSupportedIDPConfig) DeepCopy() *IdentityPlatformDefaultSupportedIDPConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformDefaultSupportedIDPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformDefaultSupportedIDPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformDefaultSupportedIDPConfigList) DeepCopyInto(out *IdentityPlatformDefaultSupportedIDPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPlatformDefaultSupportedIDPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformDefaultSupportedIDPConfigList.
func (in *IdentityPlatformDefaultSupportedIDPConfigList) DeepCopy() *IdentityPlatformDefaultSupportedIDPConfigList {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformDefaultSupportedIDPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformDefaultSupportedIDPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformDefaultSupportedIDPConfigSpec) DeepCopyInto(out *IdentityPlatformDefaultSupportedIDPConfigSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformDefaultSupportedIDPConfigSpec.
func (in *IdentityPlatformDefaultSupportedIDPConfigSpec) DeepCopy() *IdentityPlatformDefaultSupportedIDPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformDefaultSupportedIDPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformDefaultSupportedIDPConfigStatus) DeepCopyInto(out *IdentityPlatformDefaultSupportedIDPConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformDefaultSupportedIDPConfigStatus.
func (in *IdentityPlatformDefaultSupportedIDPConfigStatus) DeepCopy() *IdentityPlatformDefaultSupportedIDPConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformDefaultSupportedIDPConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformInboundSAMLConfig) DeepCopyInto(out *IdentityPlatformInboundSAMLConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformInboundSAMLConfig.
func (in *IdentityPlatformInboundSAMLConfig) DeepCopy() *IdentityPlatformInboundSAMLConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformInboundSAMLConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformInboundSAMLConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformInboundSAMLConfigList) DeepCopyInto(out *IdentityPlatformInboundSAMLConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPlatformInboundSAMLConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformInboundSAMLConfigList.
func (in *IdentityPlatformInboundSAMLConfigList) DeepCopy() *IdentityPlatformInboundSAMLConfigList {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformInboundSAMLConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformInboundSAMLConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformInboundSAMLConfigSpec) DeepCopyInto(out *IdentityPlatformInboundSAMLConfigSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.IdpConfig.DeepCopyInto(&out.IdpConfig)
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	in.SpConfig.DeepCopyInto(&out.SpConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformInboundSAMLConfigSpec.
func (in *IdentityPlatformInboundSAMLConfigSpec) DeepCopy() *IdentityPlatformInboundSAMLConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformInboundSAMLConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformInboundSAMLConfigStatus) DeepCopyInto(out *IdentityPlatformInboundSAMLConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformInboundSAMLConfigStatus.
func (in *IdentityPlatformInboundSAMLConfigStatus) DeepCopy() *IdentityPlatformInboundSAMLConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformInboundSAMLConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformProjectDefaultConfig) DeepCopyInto(out *IdentityPlatformProjectDefaultConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformProjectDefaultConfig.
func (in *IdentityPlatformProjectDefaultConfig) DeepCopy() *IdentityPlatformProjectDefaultConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformProjectDefaultConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformProjectDefaultConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformProjectDefaultConfigList) DeepCopyInto(out *IdentityPlatformProjectDefaultConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPlatformProjectDefaultConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformProjectDefaultConfigList.
func (in *IdentityPlatformProjectDefaultConfigList) DeepCopy() *IdentityPlatformProjectDefaultConfigList {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformProjectDefaultConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformProjectDefaultConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformProjectDefaultConfigSpec) DeepCopyInto(out *IdentityPlatformProjectDefaultConfigSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SignIn != nil {
		in, out := &in.SignIn, &out.SignIn
		*out = new(ProjectdefaultconfigSignIn)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformProjectDefaultConfigSpec.
func (in *IdentityPlatformProjectDefaultConfigSpec) DeepCopy() *IdentityPlatformProjectDefaultConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformProjectDefaultConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformProjectDefaultConfigStatus) DeepCopyInto(out *IdentityPlatformProjectDefaultConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformProjectDefaultConfigStatus.
func (in *IdentityPlatformProjectDefaultConfigStatus) DeepCopy() *IdentityPlatformProjectDefaultConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformProjectDefaultConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfig) DeepCopyInto(out *IdentityPlatformTenantDefaultSupportedIDPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantDefaultSupportedIDPConfig.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfig) DeepCopy() *IdentityPlatformTenantDefaultSupportedIDPConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantDefaultSupportedIDPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigList) DeepCopyInto(out *IdentityPlatformTenantDefaultSupportedIDPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPlatformTenantDefaultSupportedIDPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantDefaultSupportedIDPConfigList.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigList) DeepCopy() *IdentityPlatformTenantDefaultSupportedIDPConfigList {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantDefaultSupportedIDPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigSpec) DeepCopyInto(out *IdentityPlatformTenantDefaultSupportedIDPConfigSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantDefaultSupportedIDPConfigSpec.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigSpec) DeepCopy() *IdentityPlatformTenantDefaultSupportedIDPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantDefaultSupportedIDPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigStatus) DeepCopyInto(out *IdentityPlatformTenantDefaultSupportedIDPConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantDefaultSupportedIDPConfigStatus.
func (in *IdentityPlatformTenantDefaultSupportedIDPConfigStatus) DeepCopy() *IdentityPlatformTenantDefaultSupportedIDPConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantDefaultSupportedIDPConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantInboundSAMLConfig) DeepCopyInto(out *IdentityPlatformTenantInboundSAMLConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantInboundSAMLConfig.
func (in *IdentityPlatformTenantInboundSAMLConfig) DeepCopy() *IdentityPlatformTenantInboundSAMLConfig {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantInboundSAMLConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformTenantInboundSAMLConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantInboundSAMLConfigList) DeepCopyInto(out *IdentityPlatformTenantInboundSAMLConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPlatformTenantInboundSAMLConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantInboundSAMLConfigList.
func (in *IdentityPlatformTenantInboundSAMLConfigList) DeepCopy() *IdentityPlatformTenantInboundSAMLConfigList {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantInboundSAMLConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPlatformTenantInboundSAMLConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantInboundSAMLConfigSpec) DeepCopyInto(out *IdentityPlatformTenantInboundSAMLConfigSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.IdpConfig.DeepCopyInto(&out.IdpConfig)
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	in.SpConfig.DeepCopyInto(&out.SpConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantInboundSAMLConfigSpec.
func (in *IdentityPlatformTenantInboundSAMLConfigSpec) DeepCopy() *IdentityPlatformTenantInboundSAMLConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantInboundSAMLConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPlatformTenantInboundSAMLConfigStatus) DeepCopyInto(out *IdentityPlatformTenantInboundSAMLConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPlatformTenantInboundSAMLConfigStatus.
func (in *IdentityPlatformTenantInboundSAMLConfigStatus) DeepCopy() *IdentityPlatformTenantInboundSAMLConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPlatformTenantInboundSAMLConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundsamlconfigIdpCertificates) DeepCopyInto(out *InboundsamlconfigIdpCertificates) {
	*out = *in
	if in.X509Certificate != nil {
		in, out := &in.X509Certificate, &out.X509Certificate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundsamlconfigIdpCertificates.
func (in *InboundsamlconfigIdpCertificates) DeepCopy() *InboundsamlconfigIdpCertificates {
	if in == nil {
		return nil
	}
	out := new(InboundsamlconfigIdpCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundsamlconfigIdpConfig) DeepCopyInto(out *InboundsamlconfigIdpConfig) {
	*out = *in
	if in.IdpCertificates != nil {
		in, out := &in.IdpCertificates, &out.IdpCertificates
		*out = make([]InboundsamlconfigIdpCertificates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignRequest != nil {
		in, out := &in.SignRequest, &out.SignRequest
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundsamlconfigIdpConfig.
func (in *InboundsamlconfigIdpConfig) DeepCopy() *InboundsamlconfigIdpConfig {
	if in == nil {
		return nil
	}
	out := new(InboundsamlconfigIdpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundsamlconfigSpCertificates) DeepCopyInto(out *InboundsamlconfigSpCertificates) {
	*out = *in
	if in.X509Certificate != nil {
		in, out := &in.X509Certificate, &out.X509Certificate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundsamlconfigSpCertificates.
func (in *InboundsamlconfigSpCertificates) DeepCopy() *InboundsamlconfigSpCertificates {
	if in == nil {
		return nil
	}
	out := new(InboundsamlconfigSpCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundsamlconfigSpConfig) DeepCopyInto(out *InboundsamlconfigSpConfig) {
	*out = *in
	if in.CallbackUri != nil {
		in, out := &in.CallbackUri, &out.CallbackUri
		*out = new(string)
		**out = **in
	}
	if in.SpCertificates != nil {
		in, out := &in.SpCertificates, &out.SpCertificates
		*out = make([]InboundsamlconfigSpCertificates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpEntityId != nil {
		in, out := &in.SpEntityId, &out.SpEntityId
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundsamlconfigSpConfig.
func (in *InboundsamlconfigSpConfig) DeepCopy() *InboundsamlconfigSpConfig {
	if in == nil {
		return nil
	}
	out := new(InboundsamlconfigSpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectdefaultconfigAnonymous) DeepCopyInto(out *ProjectdefaultconfigAnonymous) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectdefaultconfigAnonymous.
func (in *ProjectdefaultconfigAnonymous) DeepCopy() *ProjectdefaultconfigAnonymous {
	if in == nil {
		return nil
	}
	out := new(ProjectdefaultconfigAnonymous)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectdefaultconfigEmail) DeepCopyInto(out *ProjectdefaultconfigEmail) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PasswordRequired != nil {
		in, out := &in.PasswordRequired, &out.PasswordRequired
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectdefaultconfigEmail.
func (in *ProjectdefaultconfigEmail) DeepCopy() *ProjectdefaultconfigEmail {
	if in == nil {
		return nil
	}
	out := new(ProjectdefaultconfigEmail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectdefaultconfigHashConfig) DeepCopyInto(out *ProjectdefaultconfigHashConfig) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.MemoryCost != nil {
		in, out := &in.MemoryCost, &out.MemoryCost
		*out = new(int)
		**out = **in
	}
	if in.Rounds != nil {
		in, out := &in.Rounds, &out.Rounds
		*out = new(int)
		**out = **in
	}
	if in.SaltSeparator != nil {
		in, out := &in.SaltSeparator, &out.SaltSeparator
		*out = new(string)
		**out = **in
	}
	if in.SignerKey != nil {
		in, out := &in.SignerKey, &out.SignerKey
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectdefaultconfigHashConfig.
func (in *ProjectdefaultconfigHashConfig) DeepCopy() *ProjectdefaultconfigHashConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectdefaultconfigHashConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectdefaultconfigPhoneNumber) DeepCopyInto(out *ProjectdefaultconfigPhoneNumber) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.TestPhoneNumbers != nil {
		in, out := &in.TestPhoneNumbers, &out.TestPhoneNumbers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectdefaultconfigPhoneNumber.
func (in *ProjectdefaultconfigPhoneNumber) DeepCopy() *ProjectdefaultconfigPhoneNumber {
	if in == nil {
		return nil
	}
	out := new(ProjectdefaultconfigPhoneNumber)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectdefaultconfigSignIn) DeepCopyInto(out *ProjectdefaultconfigSignIn) {
	*out = *in
	if in.AllowDuplicateEmails != nil {
		in, out := &in.AllowDuplicateEmails, &out.AllowDuplicateEmails
		*out = new(bool)
		**out = **in
	}
	if in.Anonymous != nil {
		in, out := &in.Anonymous, &out.Anonymous
		*out = new(ProjectdefaultconfigAnonymous)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(ProjectdefaultconfigEmail)
		(*in).DeepCopyInto(*out)
	}
	if in.HashConfig != nil {
		in, out := &in.HashConfig, &out.HashConfig
		*out = make([]ProjectdefaultconfigHashConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
		*out = new(ProjectdefaultconfigPhoneNumber)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectdefaultconfigSignIn.
func (in *ProjectdefaultconfigSignIn) DeepCopy() *ProjectdefaultconfigSignIn {
	if in == nil {
		return nil
	}
	out := new(ProjectdefaultconfigSignIn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantinboundsamlconfigIdpCertificates) DeepCopyInto(out *TenantinboundsamlconfigIdpCertificates) {
	*out = *in
	if in.X509Certificate != nil {
		in, out := &in.X509Certificate, &out.X509Certificate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantinboundsamlconfigIdpCertificates.
func (in *TenantinboundsamlconfigIdpCertificates) DeepCopy() *TenantinboundsamlconfigIdpCertificates {
	if in == nil {
		return nil
	}
	out := new(TenantinboundsamlconfigIdpCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantinboundsamlconfigIdpConfig) DeepCopyInto(out *TenantinboundsamlconfigIdpConfig) {
	*out = *in
	if in.IdpCertificates != nil {
		in, out := &in.IdpCertificates, &out.IdpCertificates
		*out = make([]TenantinboundsamlconfigIdpCertificates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignRequest != nil {
		in, out := &in.SignRequest, &out.SignRequest
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantinboundsamlconfigIdpConfig.
func (in *TenantinboundsamlconfigIdpConfig) DeepCopy() *TenantinboundsamlconfigIdpConfig {
	if in == nil {
		return nil
	}
	out := new(TenantinboundsamlconfigIdpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantinboundsamlconfigSpCertificates) DeepCopyInto(out *TenantinboundsamlconfigSpCertificates) {
	*out = *in
	if in.X509Certificate != nil {
		in, out := &in.X509Certificate, &out.X509Certificate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantinboundsamlconfigSpCertificates.
func (in *TenantinboundsamlconfigSpCertificates) DeepCopy() *TenantinboundsamlconfigSpCertificates {
	if in == nil {
		return nil
	}
	out := new(TenantinboundsamlconfigSpCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantinboundsamlconfigSpConfig) DeepCopyInto(out *TenantinboundsamlconfigSpConfig) {
	*out = *in
	if in.SpCertificates != nil {
		in, out := &in.SpCertificates, &out.SpCertificates
		*out = make([]TenantinboundsamlconfigSpCertificates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantinboundsamlconfigSpConfig.
func (in *TenantinboundsamlconfigSpConfig) DeepCopy() *TenantinboundsamlconfigSpConfig {
	if in == nil {
		return nil
	}
	out := new(TenantinboundsamlconfigSpConfig)
	in.DeepCopyInto(out)
	return out
}
