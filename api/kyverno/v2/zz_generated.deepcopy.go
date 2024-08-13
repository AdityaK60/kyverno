//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	v2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionRequestInfoObject) DeepCopyInto(out *AdmissionRequestInfoObject) {
	*out = *in
	if in.AdmissionRequest != nil {
		in, out := &in.AdmissionRequest, &out.AdmissionRequest
		*out = new(v1.AdmissionRequest)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionRequestInfoObject.
func (in *AdmissionRequestInfoObject) DeepCopy() *AdmissionRequestInfoObject {
	if in == nil {
		return nil
	}
	out := new(AdmissionRequestInfoObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnyAllConditions) DeepCopyInto(out *AnyAllConditions) {
	*out = *in
	if in.AnyConditions != nil {
		in, out := &in.AnyConditions, &out.AnyConditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllConditions != nil {
		in, out := &in.AllConditions, &out.AllConditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnyAllConditions.
func (in *AnyAllConditions) DeepCopy() *AnyAllConditions {
	if in == nil {
		return nil
	}
	out := new(AnyAllConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupPolicy) DeepCopyInto(out *CleanupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupPolicy.
func (in *CleanupPolicy) DeepCopy() *CleanupPolicy {
	if in == nil {
		return nil
	}
	out := new(CleanupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CleanupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupPolicyList) DeepCopyInto(out *CleanupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CleanupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupPolicyList.
func (in *CleanupPolicyList) DeepCopy() *CleanupPolicyList {
	if in == nil {
		return nil
	}
	out := new(CleanupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CleanupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupPolicySpec) DeepCopyInto(out *CleanupPolicySpec) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = make([]kyvernov1.ContextEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.MatchResources.DeepCopyInto(&out.MatchResources)
	if in.ExcludeResources != nil {
		in, out := &in.ExcludeResources, &out.ExcludeResources
		*out = new(v2beta1.MatchResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(AnyAllConditions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupPolicySpec.
func (in *CleanupPolicySpec) DeepCopy() *CleanupPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CleanupPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CleanupPolicyStatus) DeepCopyInto(out *CleanupPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastExecutionTime.DeepCopyInto(&out.LastExecutionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CleanupPolicyStatus.
func (in *CleanupPolicyStatus) DeepCopy() *CleanupPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CleanupPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCleanupPolicy) DeepCopyInto(out *ClusterCleanupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCleanupPolicy.
func (in *ClusterCleanupPolicy) DeepCopy() *ClusterCleanupPolicy {
	if in == nil {
		return nil
	}
	out := new(ClusterCleanupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterCleanupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCleanupPolicyList) DeepCopyInto(out *ClusterCleanupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterCleanupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCleanupPolicyList.
func (in *ClusterCleanupPolicyList) DeepCopy() *ClusterCleanupPolicyList {
	if in == nil {
		return nil
	}
	out := new(ClusterCleanupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterCleanupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.RawKey != nil {
		in, out := &in.RawKey, &out.RawKey
		*out = (*in).DeepCopy()
	}
	if in.RawValue != nil {
		in, out := &in.RawValue, &out.RawValue
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Exception) DeepCopyInto(out *Exception) {
	*out = *in
	if in.RuleNames != nil {
		in, out := &in.RuleNames, &out.RuleNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Exception.
func (in *Exception) DeepCopy() *Exception {
	if in == nil {
		return nil
	}
	out := new(Exception)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyException) DeepCopyInto(out *PolicyException) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyException.
func (in *PolicyException) DeepCopy() *PolicyException {
	if in == nil {
		return nil
	}
	out := new(PolicyException)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyException) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyExceptionList) DeepCopyInto(out *PolicyExceptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyException, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyExceptionList.
func (in *PolicyExceptionList) DeepCopy() *PolicyExceptionList {
	if in == nil {
		return nil
	}
	out := new(PolicyExceptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyExceptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyExceptionSpec) DeepCopyInto(out *PolicyExceptionSpec) {
	*out = *in
	if in.Background != nil {
		in, out := &in.Background, &out.Background
		*out = new(bool)
		**out = **in
	}
	in.Match.DeepCopyInto(&out.Match)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(AnyAllConditions)
		(*in).DeepCopyInto(*out)
	}
	if in.Exceptions != nil {
		in, out := &in.Exceptions, &out.Exceptions
		*out = make([]Exception, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSecurity != nil {
		in, out := &in.PodSecurity, &out.PodSecurity
		*out = make([]kyvernov1.PodSecurityStandard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyExceptionSpec.
func (in *PolicyExceptionSpec) DeepCopy() *PolicyExceptionSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyExceptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestInfo) DeepCopyInto(out *RequestInfo) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRoles != nil {
		in, out := &in.ClusterRoles, &out.ClusterRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AdmissionUserInfo.DeepCopyInto(&out.AdmissionUserInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestInfo.
func (in *RequestInfo) DeepCopy() *RequestInfo {
	if in == nil {
		return nil
	}
	out := new(RequestInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleContext) DeepCopyInto(out *RuleContext) {
	*out = *in
	out.Trigger = in.Trigger
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleContext.
func (in *RuleContext) DeepCopy() *RuleContext {
	if in == nil {
		return nil
	}
	out := new(RuleContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequest) DeepCopyInto(out *UpdateRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequest.
func (in *UpdateRequest) DeepCopy() *UpdateRequest {
	if in == nil {
		return nil
	}
	out := new(UpdateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestList) DeepCopyInto(out *UpdateRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpdateRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestList.
func (in *UpdateRequestList) DeepCopy() *UpdateRequestList {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpdateRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestSpec) DeepCopyInto(out *UpdateRequestSpec) {
	*out = *in
	if in.RuleContext != nil {
		in, out := &in.RuleContext, &out.RuleContext
		*out = make([]RuleContext, len(*in))
		copy(*out, *in)
	}
	out.Resource = in.Resource
	in.Context.DeepCopyInto(&out.Context)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestSpec.
func (in *UpdateRequestSpec) DeepCopy() *UpdateRequestSpec {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestSpecContext) DeepCopyInto(out *UpdateRequestSpecContext) {
	*out = *in
	in.UserRequestInfo.DeepCopyInto(&out.UserRequestInfo)
	in.AdmissionRequestInfo.DeepCopyInto(&out.AdmissionRequestInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestSpecContext.
func (in *UpdateRequestSpecContext) DeepCopy() *UpdateRequestSpecContext {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestSpecContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateRequestStatus) DeepCopyInto(out *UpdateRequestStatus) {
	*out = *in
	if in.GeneratedResources != nil {
		in, out := &in.GeneratedResources, &out.GeneratedResources
		*out = make([]kyvernov1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateRequestStatus.
func (in *UpdateRequestStatus) DeepCopy() *UpdateRequestStatus {
	if in == nil {
		return nil
	}
	out := new(UpdateRequestStatus)
	in.DeepCopyInto(out)
	return out
}
