//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvatarInitParameters) DeepCopyInto(out *AvatarInitParameters) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvatarInitParameters.
func (in *AvatarInitParameters) DeepCopy() *AvatarInitParameters {
	if in == nil {
		return nil
	}
	out := new(AvatarInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvatarObservation) DeepCopyInto(out *AvatarObservation) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvatarObservation.
func (in *AvatarObservation) DeepCopy() *AvatarObservation {
	if in == nil {
		return nil
	}
	out := new(AvatarObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvatarParameters) DeepCopyInto(out *AvatarParameters) {
	*out = *in
	if in.Href != nil {
		in, out := &in.Href, &out.Href
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvatarParameters.
func (in *AvatarParameters) DeepCopy() *AvatarParameters {
	if in == nil {
		return nil
	}
	out := new(AvatarParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkInitParameters) DeepCopyInto(out *LinkInitParameters) {
	*out = *in
	if in.Avatar != nil {
		in, out := &in.Avatar, &out.Avatar
		*out = make([]AvatarInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkInitParameters.
func (in *LinkInitParameters) DeepCopy() *LinkInitParameters {
	if in == nil {
		return nil
	}
	out := new(LinkInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkObservation) DeepCopyInto(out *LinkObservation) {
	*out = *in
	if in.Avatar != nil {
		in, out := &in.Avatar, &out.Avatar
		*out = make([]AvatarObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkObservation.
func (in *LinkObservation) DeepCopy() *LinkObservation {
	if in == nil {
		return nil
	}
	out := new(LinkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkParameters) DeepCopyInto(out *LinkParameters) {
	*out = *in
	if in.Avatar != nil {
		in, out := &in.Avatar, &out.Avatar
		*out = make([]AvatarParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkParameters.
func (in *LinkParameters) DeepCopy() *LinkParameters {
	if in == nil {
		return nil
	}
	out := new(LinkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryInitParameters) DeepCopyInto(out *RepositoryInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForkPolicy != nil {
		in, out := &in.ForkPolicy, &out.ForkPolicy
		*out = new(string)
		**out = **in
	}
	if in.HasIssues != nil {
		in, out := &in.HasIssues, &out.HasIssues
		*out = new(bool)
		**out = **in
	}
	if in.HasWiki != nil {
		in, out := &in.HasWiki, &out.HasWiki
		*out = new(bool)
		**out = **in
	}
	if in.InheritBranchingModel != nil {
		in, out := &in.InheritBranchingModel, &out.InheritBranchingModel
		*out = new(bool)
		**out = **in
	}
	if in.InheritDefaultMergeStrategy != nil {
		in, out := &in.InheritDefaultMergeStrategy, &out.InheritDefaultMergeStrategy
		*out = new(bool)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.Link != nil {
		in, out := &in.Link, &out.Link
		*out = make([]LinkInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PipelinesEnabled != nil {
		in, out := &in.PipelinesEnabled, &out.PipelinesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.Scm != nil {
		in, out := &in.Scm, &out.Scm
		*out = new(string)
		**out = **in
	}
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryInitParameters.
func (in *RepositoryInitParameters) DeepCopy() *RepositoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryObservation) DeepCopyInto(out *RepositoryObservation) {
	*out = *in
	if in.CloneHTTPS != nil {
		in, out := &in.CloneHTTPS, &out.CloneHTTPS
		*out = new(string)
		**out = **in
	}
	if in.CloneSSH != nil {
		in, out := &in.CloneSSH, &out.CloneSSH
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForkPolicy != nil {
		in, out := &in.ForkPolicy, &out.ForkPolicy
		*out = new(string)
		**out = **in
	}
	if in.HasIssues != nil {
		in, out := &in.HasIssues, &out.HasIssues
		*out = new(bool)
		**out = **in
	}
	if in.HasWiki != nil {
		in, out := &in.HasWiki, &out.HasWiki
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InheritBranchingModel != nil {
		in, out := &in.InheritBranchingModel, &out.InheritBranchingModel
		*out = new(bool)
		**out = **in
	}
	if in.InheritDefaultMergeStrategy != nil {
		in, out := &in.InheritDefaultMergeStrategy, &out.InheritDefaultMergeStrategy
		*out = new(bool)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.Link != nil {
		in, out := &in.Link, &out.Link
		*out = make([]LinkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PipelinesEnabled != nil {
		in, out := &in.PipelinesEnabled, &out.PipelinesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.Scm != nil {
		in, out := &in.Scm, &out.Scm
		*out = new(string)
		**out = **in
	}
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
	if in.UUID != nil {
		in, out := &in.UUID, &out.UUID
		*out = new(string)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryObservation.
func (in *RepositoryObservation) DeepCopy() *RepositoryObservation {
	if in == nil {
		return nil
	}
	out := new(RepositoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryParameters) DeepCopyInto(out *RepositoryParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForkPolicy != nil {
		in, out := &in.ForkPolicy, &out.ForkPolicy
		*out = new(string)
		**out = **in
	}
	if in.HasIssues != nil {
		in, out := &in.HasIssues, &out.HasIssues
		*out = new(bool)
		**out = **in
	}
	if in.HasWiki != nil {
		in, out := &in.HasWiki, &out.HasWiki
		*out = new(bool)
		**out = **in
	}
	if in.InheritBranchingModel != nil {
		in, out := &in.InheritBranchingModel, &out.InheritBranchingModel
		*out = new(bool)
		**out = **in
	}
	if in.InheritDefaultMergeStrategy != nil {
		in, out := &in.InheritDefaultMergeStrategy, &out.InheritDefaultMergeStrategy
		*out = new(bool)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.Link != nil {
		in, out := &in.Link, &out.Link
		*out = make([]LinkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PipelinesEnabled != nil {
		in, out := &in.PipelinesEnabled, &out.PipelinesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(string)
		**out = **in
	}
	if in.Scm != nil {
		in, out := &in.Scm, &out.Scm
		*out = new(string)
		**out = **in
	}
	if in.Slug != nil {
		in, out := &in.Slug, &out.Slug
		*out = new(string)
		**out = **in
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryParameters.
func (in *RepositoryParameters) DeepCopy() *RepositoryParameters {
	if in == nil {
		return nil
	}
	out := new(RepositoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}