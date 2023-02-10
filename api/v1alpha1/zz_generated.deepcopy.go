//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


Licensed under the Mozilla Public License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.mozilla.org/en-US/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/TykTechnologies/tyk-operator/api/model"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDefinitionSpec) DeepCopyInto(out *APIDefinitionSpec) {
	*out = *in
	in.APIDefinitionSpec.DeepCopyInto(&out.APIDefinitionSpec)
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDefinitionSpec.
func (in *APIDefinitionSpec) DeepCopy() *APIDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(APIDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDescription) DeepCopyInto(out *APIDescription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDescription.
func (in *APIDescription) DeepCopy() *APIDescription {
	if in == nil {
		return nil
	}
	out := new(APIDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIDescription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDescriptionBase) DeepCopyInto(out *APIDescriptionBase) {
	*out = *in
	in.APIDescription.DeepCopyInto(&out.APIDescription)
	if in.APIDocumentation != nil {
		in, out := &in.APIDocumentation, &out.APIDocumentation
		*out = new(APIDocumentation)
		**out = **in
	}
	if in.PolicyRef != nil {
		in, out := &in.PolicyRef, &out.PolicyRef
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDescriptionBase.
func (in *APIDescriptionBase) DeepCopy() *APIDescriptionBase {
	if in == nil {
		return nil
	}
	out := new(APIDescriptionBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDescriptionList) DeepCopyInto(out *APIDescriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIDescription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDescriptionList.
func (in *APIDescriptionList) DeepCopy() *APIDescriptionList {
	if in == nil {
		return nil
	}
	out := new(APIDescriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIDescriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDescriptionSpec) DeepCopyInto(out *APIDescriptionSpec) {
	*out = *in
	in.APIDescriptionBase.DeepCopyInto(&out.APIDescriptionBase)
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDescriptionSpec.
func (in *APIDescriptionSpec) DeepCopy() *APIDescriptionSpec {
	if in == nil {
		return nil
	}
	out := new(APIDescriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDescriptionStatus) DeepCopyInto(out *APIDescriptionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDescriptionStatus.
func (in *APIDescriptionStatus) DeepCopy() *APIDescriptionStatus {
	if in == nil {
		return nil
	}
	out := new(APIDescriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIDocumentation) DeepCopyInto(out *APIDocumentation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIDocumentation.
func (in *APIDocumentation) DeepCopy() *APIDocumentation {
	if in == nil {
		return nil
	}
	out := new(APIDocumentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APILimit) DeepCopyInto(out *APILimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APILimit.
func (in *APILimit) DeepCopy() *APILimit {
	if in == nil {
		return nil
	}
	out := new(APILimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessDefinition) DeepCopyInto(out *AccessDefinition) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedTypes != nil {
		in, out := &in.AllowedTypes, &out.AllowedTypes
		*out = make(GraphQLTypeList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestrictedTypes != nil {
		in, out := &in.RestrictedTypes, &out.RestrictedTypes
		*out = make(GraphQLTypeList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FieldAccessRights != nil {
		in, out := &in.FieldAccessRights, &out.FieldAccessRights
		*out = make([]FieldAccessDefinition, len(*in))
		copy(*out, *in)
	}
	if in.AllowedURLs != nil {
		in, out := &in.AllowedURLs, &out.AllowedURLs
		*out = make([]AccessSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessDefinition.
func (in *AccessDefinition) DeepCopy() *AccessDefinition {
	if in == nil {
		return nil
	}
	out := new(AccessDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessSpec) DeepCopyInto(out *AccessSpec) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessSpec.
func (in *AccessSpec) DeepCopy() *AccessSpec {
	if in == nil {
		return nil
	}
	out := new(AccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiDefinition) DeepCopyInto(out *ApiDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiDefinition.
func (in *ApiDefinition) DeepCopy() *ApiDefinition {
	if in == nil {
		return nil
	}
	out := new(ApiDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiDefinitionList) DeepCopyInto(out *ApiDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApiDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiDefinitionList.
func (in *ApiDefinitionList) DeepCopy() *ApiDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ApiDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApiDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiDefinitionStatus) DeepCopyInto(out *ApiDefinitionStatus) {
	*out = *in
	if in.LinkedByPolicies != nil {
		in, out := &in.LinkedByPolicies, &out.LinkedByPolicies
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedByAPIs != nil {
		in, out := &in.LinkedByAPIs, &out.LinkedByAPIs
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedToAPIs != nil {
		in, out := &in.LinkedToAPIs, &out.LinkedToAPIs
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiDefinitionStatus.
func (in *ApiDefinitionStatus) DeepCopy() *ApiDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ApiDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.Ingress = in.Ingress
	if in.UserOwners != nil {
		in, out := &in.UserOwners, &out.UserOwners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserGroupOwners != nil {
		in, out := &in.UserGroupOwners, &out.UserGroupOwners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldAccessDefinition) DeepCopyInto(out *FieldAccessDefinition) {
	*out = *in
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldAccessDefinition.
func (in *FieldAccessDefinition) DeepCopy() *FieldAccessDefinition {
	if in == nil {
		return nil
	}
	out := new(FieldAccessDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldLimits) DeepCopyInto(out *FieldLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldLimits.
func (in *FieldLimits) DeepCopy() *FieldLimits {
	if in == nil {
		return nil
	}
	out := new(FieldLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphQLType) DeepCopyInto(out *GraphQLType) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLType.
func (in *GraphQLType) DeepCopy() *GraphQLType {
	if in == nil {
		return nil
	}
	out := new(GraphQLType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in GraphQLTypeList) DeepCopyInto(out *GraphQLTypeList) {
	{
		in := &in
		*out = make(GraphQLTypeList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphQLTypeList.
func (in GraphQLTypeList) DeepCopy() GraphQLTypeList {
	if in == nil {
		return nil
	}
	out := new(GraphQLTypeList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorContext) DeepCopyInto(out *OperatorContext) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorContext.
func (in *OperatorContext) DeepCopy() *OperatorContext {
	if in == nil {
		return nil
	}
	out := new(OperatorContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorContext) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorContextList) DeepCopyInto(out *OperatorContextList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorContextList.
func (in *OperatorContextList) DeepCopy() *OperatorContextList {
	if in == nil {
		return nil
	}
	out := new(OperatorContextList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorContextList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorContextSpec) DeepCopyInto(out *OperatorContextSpec) {
	*out = *in
	if in.FromSecret != nil {
		in, out := &in.FromSecret, &out.FromSecret
		*out = new(model.Target)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = new(Environment)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorContextSpec.
func (in *OperatorContextSpec) DeepCopy() *OperatorContextSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorContextSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorContextStatus) DeepCopyInto(out *OperatorContextStatus) {
	*out = *in
	if in.LinkedApiDefinitions != nil {
		in, out := &in.LinkedApiDefinitions, &out.LinkedApiDefinitions
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedApiDescriptions != nil {
		in, out := &in.LinkedApiDescriptions, &out.LinkedApiDescriptions
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedPortalAPICatalogues != nil {
		in, out := &in.LinkedPortalAPICatalogues, &out.LinkedPortalAPICatalogues
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedSecurityPolicies != nil {
		in, out := &in.LinkedSecurityPolicies, &out.LinkedSecurityPolicies
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
	if in.LinkedPortalConfigs != nil {
		in, out := &in.LinkedPortalConfigs, &out.LinkedPortalConfigs
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorContextStatus.
func (in *OperatorContextStatus) DeepCopy() *OperatorContextStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorContextStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyPartitions) DeepCopyInto(out *PolicyPartitions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyPartitions.
func (in *PolicyPartitions) DeepCopy() *PolicyPartitions {
	if in == nil {
		return nil
	}
	out := new(PolicyPartitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalAPICatalogue) DeepCopyInto(out *PortalAPICatalogue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalAPICatalogue.
func (in *PortalAPICatalogue) DeepCopy() *PortalAPICatalogue {
	if in == nil {
		return nil
	}
	out := new(PortalAPICatalogue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortalAPICatalogue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalAPICatalogueList) DeepCopyInto(out *PortalAPICatalogueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PortalAPICatalogue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalAPICatalogueList.
func (in *PortalAPICatalogueList) DeepCopy() *PortalAPICatalogueList {
	if in == nil {
		return nil
	}
	out := new(PortalAPICatalogueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortalAPICatalogueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalAPICatalogueSpec) DeepCopyInto(out *PortalAPICatalogueSpec) {
	*out = *in
	if in.APIDescriptionList != nil {
		in, out := &in.APIDescriptionList, &out.APIDescriptionList
		*out = make([]*PortalCatalogueDescription, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PortalCatalogueDescription)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalAPICatalogueSpec.
func (in *PortalAPICatalogueSpec) DeepCopy() *PortalAPICatalogueSpec {
	if in == nil {
		return nil
	}
	out := new(PortalAPICatalogueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalAPICatalogueStatus) DeepCopyInto(out *PortalAPICatalogueStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalAPICatalogueStatus.
func (in *PortalAPICatalogueStatus) DeepCopy() *PortalAPICatalogueStatus {
	if in == nil {
		return nil
	}
	out := new(PortalAPICatalogueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalCatalogueDescription) DeepCopyInto(out *PortalCatalogueDescription) {
	*out = *in
	in.APIDescriptionBase.DeepCopyInto(&out.APIDescriptionBase)
	if in.APIDescriptionRef != nil {
		in, out := &in.APIDescriptionRef, &out.APIDescriptionRef
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalCatalogueDescription.
func (in *PortalCatalogueDescription) DeepCopy() *PortalCatalogueDescription {
	if in == nil {
		return nil
	}
	out := new(PortalCatalogueDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalConfig) DeepCopyInto(out *PortalConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalConfig.
func (in *PortalConfig) DeepCopy() *PortalConfig {
	if in == nil {
		return nil
	}
	out := new(PortalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortalConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalConfigList) DeepCopyInto(out *PortalConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PortalConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalConfigList.
func (in *PortalConfigList) DeepCopy() *PortalConfigList {
	if in == nil {
		return nil
	}
	out := new(PortalConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortalConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalConfigSpec) DeepCopyInto(out *PortalConfigSpec) {
	*out = *in
	in.PortalModelPortalConfig.DeepCopyInto(&out.PortalModelPortalConfig)
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(model.Target)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalConfigSpec.
func (in *PortalConfigSpec) DeepCopy() *PortalConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PortalConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortalConfigStatus) DeepCopyInto(out *PortalConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortalConfigStatus.
func (in *PortalConfigStatus) DeepCopy() *PortalConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PortalConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicy) DeepCopyInto(out *SecurityPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicy.
func (in *SecurityPolicy) DeepCopy() *SecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyList) DeepCopyInto(out *SecurityPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyList.
func (in *SecurityPolicyList) DeepCopy() *SecurityPolicyList {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicySpec) DeepCopyInto(out *SecurityPolicySpec) {
	*out = *in
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(model.Target)
		**out = **in
	}
	if in.AccessRightsArray != nil {
		in, out := &in.AccessRightsArray, &out.AccessRightsArray
		*out = make([]*AccessDefinition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AccessDefinition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AccessRights != nil {
		in, out := &in.AccessRights, &out.AccessRights
		*out = make(map[string]AccessDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Partitions != nil {
		in, out := &in.Partitions, &out.Partitions
		*out = new(PolicyPartitions)
		**out = **in
	}
	if in.MetaData != nil {
		in, out := &in.MetaData, &out.MetaData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicySpec.
func (in *SecurityPolicySpec) DeepCopy() *SecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityPolicyStatus) DeepCopyInto(out *SecurityPolicyStatus) {
	*out = *in
	if in.LinkedAPIs != nil {
		in, out := &in.LinkedAPIs, &out.LinkedAPIs
		*out = make([]model.Target, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityPolicyStatus.
func (in *SecurityPolicyStatus) DeepCopy() *SecurityPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubGraph) DeepCopyInto(out *SubGraph) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubGraph.
func (in *SubGraph) DeepCopy() *SubGraph {
	if in == nil {
		return nil
	}
	out := new(SubGraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubGraph) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubGraphList) DeepCopyInto(out *SubGraphList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubGraph, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubGraphList.
func (in *SubGraphList) DeepCopy() *SubGraphList {
	if in == nil {
		return nil
	}
	out := new(SubGraphList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubGraphList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubGraphSpec) DeepCopyInto(out *SubGraphSpec) {
	*out = *in
	out.SubGraphSpec = in.SubGraphSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubGraphSpec.
func (in *SubGraphSpec) DeepCopy() *SubGraphSpec {
	if in == nil {
		return nil
	}
	out := new(SubGraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubGraphStatus) DeepCopyInto(out *SubGraphStatus) {
	*out = *in
	out.SubGraphStatus = in.SubGraphStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubGraphStatus.
func (in *SubGraphStatus) DeepCopy() *SubGraphStatus {
	if in == nil {
		return nil
	}
	out := new(SubGraphStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuperGraph) DeepCopyInto(out *SuperGraph) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuperGraph.
func (in *SuperGraph) DeepCopy() *SuperGraph {
	if in == nil {
		return nil
	}
	out := new(SuperGraph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SuperGraph) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuperGraphList) DeepCopyInto(out *SuperGraphList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SuperGraph, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuperGraphList.
func (in *SuperGraphList) DeepCopy() *SuperGraphList {
	if in == nil {
		return nil
	}
	out := new(SuperGraphList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SuperGraphList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuperGraphSpec) DeepCopyInto(out *SuperGraphSpec) {
	*out = *in
	in.SuperGraphSpec.DeepCopyInto(&out.SuperGraphSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuperGraphSpec.
func (in *SuperGraphSpec) DeepCopy() *SuperGraphSpec {
	if in == nil {
		return nil
	}
	out := new(SuperGraphSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuperGraphStatus) DeepCopyInto(out *SuperGraphStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuperGraphStatus.
func (in *SuperGraphStatus) DeepCopy() *SuperGraphStatus {
	if in == nil {
		return nil
	}
	out := new(SuperGraphStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TykOASApiDefinition) DeepCopyInto(out *TykOASApiDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TykOASApiDefinition.
func (in *TykOASApiDefinition) DeepCopy() *TykOASApiDefinition {
	if in == nil {
		return nil
	}
	out := new(TykOASApiDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TykOASApiDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TykOASApiDefinitionList) DeepCopyInto(out *TykOASApiDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TykOASApiDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TykOASApiDefinitionList.
func (in *TykOASApiDefinitionList) DeepCopy() *TykOASApiDefinitionList {
	if in == nil {
		return nil
	}
	out := new(TykOASApiDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TykOASApiDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TykOASApiDefinitionSpec) DeepCopyInto(out *TykOASApiDefinitionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TykOASApiDefinitionSpec.
func (in *TykOASApiDefinitionSpec) DeepCopy() *TykOASApiDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(TykOASApiDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TykOASApiDefinitionStatus) DeepCopyInto(out *TykOASApiDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TykOASApiDefinitionStatus.
func (in *TykOASApiDefinitionStatus) DeepCopy() *TykOASApiDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(TykOASApiDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}
