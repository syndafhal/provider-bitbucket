// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AvatarInitParameters struct {

	// href of the avatar.
	Href *string `json:"href,omitempty" tf:"href,omitempty"`
}

type AvatarObservation struct {

	// href of the avatar.
	Href *string `json:"href,omitempty" tf:"href,omitempty"`
}

type AvatarParameters struct {

	// href of the avatar.
	// +kubebuilder:validation:Optional
	Href *string `json:"href,omitempty" tf:"href,omitempty"`
}

type LinkInitParameters struct {

	// An avatar link to a resource related to this object. See Avatar Below.
	Avatar []AvatarInitParameters `json:"avatar,omitempty" tf:"avatar,omitempty"`
}

type LinkObservation struct {

	// An avatar link to a resource related to this object. See Avatar Below.
	Avatar []AvatarObservation `json:"avatar,omitempty" tf:"avatar,omitempty"`
}

type LinkParameters struct {

	// An avatar link to a resource related to this object. See Avatar Below.
	// +kubebuilder:validation:Optional
	Avatar []AvatarParameters `json:"avatar,omitempty" tf:"avatar,omitempty"`
}

type RepositoryInitParameters struct {

	// What the description of the repo is.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// What the fork policy should be. Defaults to
	// allow_forks. Valid values are allow_forks, no_public_forks, no_forks.
	ForkPolicy *string `json:"forkPolicy,omitempty" tf:"fork_policy,omitempty"`

	// If this should have issues turned on or not.
	HasIssues *bool `json:"hasIssues,omitempty" tf:"has_issues,omitempty"`

	// If this should have wiki turned on or not.
	HasWiki *bool `json:"hasWiki,omitempty" tf:"has_wiki,omitempty"`

	// Whether to inherit branching model from project.
	InheritBranchingModel *bool `json:"inheritBranchingModel,omitempty" tf:"inherit_branching_model,omitempty"`

	// Whether to inherit default merge strategy from project.
	InheritDefaultMergeStrategy *bool `json:"inheritDefaultMergeStrategy,omitempty" tf:"inherit_default_merge_strategy,omitempty"`

	// If this should be private or not. Defaults to true.
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// What the language of this repository should be.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// A set of links to a resource related to this object. See Link Below.
	Link []LinkInitParameters `json:"link,omitempty" tf:"link,omitempty"`

	// The owner of this repository. Can be you or any team you
	// have write access to.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Turn on to enable pipelines support.
	PipelinesEnabled *bool `json:"pipelinesEnabled,omitempty" tf:"pipelines_enabled,omitempty"`

	// If you want to have this repo associated with a
	// project.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// What SCM you want to use. Valid options are hg or git.
	// Defaults to git.
	Scm *string `json:"scm,omitempty" tf:"scm,omitempty"`

	// The slug of the repository.
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// URL of website associated with this repository.
	Website *string `json:"website,omitempty" tf:"website,omitempty"`
}

type RepositoryObservation struct {

	// The HTTPS clone URL.
	CloneHTTPS *string `json:"cloneHttps,omitempty" tf:"clone_https,omitempty"`

	// The SSH clone URL.
	CloneSSH *string `json:"cloneSsh,omitempty" tf:"clone_ssh,omitempty"`

	// What the description of the repo is.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// What the fork policy should be. Defaults to
	// allow_forks. Valid values are allow_forks, no_public_forks, no_forks.
	ForkPolicy *string `json:"forkPolicy,omitempty" tf:"fork_policy,omitempty"`

	// If this should have issues turned on or not.
	HasIssues *bool `json:"hasIssues,omitempty" tf:"has_issues,omitempty"`

	// If this should have wiki turned on or not.
	HasWiki *bool `json:"hasWiki,omitempty" tf:"has_wiki,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to inherit branching model from project.
	InheritBranchingModel *bool `json:"inheritBranchingModel,omitempty" tf:"inherit_branching_model,omitempty"`

	// Whether to inherit default merge strategy from project.
	InheritDefaultMergeStrategy *bool `json:"inheritDefaultMergeStrategy,omitempty" tf:"inherit_default_merge_strategy,omitempty"`

	// If this should be private or not. Defaults to true.
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// What the language of this repository should be.
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// A set of links to a resource related to this object. See Link Below.
	Link []LinkObservation `json:"link,omitempty" tf:"link,omitempty"`

	// The owner of this repository. Can be you or any team you
	// have write access to.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Turn on to enable pipelines support.
	PipelinesEnabled *bool `json:"pipelinesEnabled,omitempty" tf:"pipelines_enabled,omitempty"`

	// If you want to have this repo associated with a
	// project.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// What SCM you want to use. Valid options are hg or git.
	// Defaults to git.
	Scm *string `json:"scm,omitempty" tf:"scm,omitempty"`

	// The slug of the repository.
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// the uuid of the repository resource.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// URL of website associated with this repository.
	Website *string `json:"website,omitempty" tf:"website,omitempty"`
}

type RepositoryParameters struct {

	// What the description of the repo is.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// What the fork policy should be. Defaults to
	// allow_forks. Valid values are allow_forks, no_public_forks, no_forks.
	// +kubebuilder:validation:Optional
	ForkPolicy *string `json:"forkPolicy,omitempty" tf:"fork_policy,omitempty"`

	// If this should have issues turned on or not.
	// +kubebuilder:validation:Optional
	HasIssues *bool `json:"hasIssues,omitempty" tf:"has_issues,omitempty"`

	// If this should have wiki turned on or not.
	// +kubebuilder:validation:Optional
	HasWiki *bool `json:"hasWiki,omitempty" tf:"has_wiki,omitempty"`

	// Whether to inherit branching model from project.
	// +kubebuilder:validation:Optional
	InheritBranchingModel *bool `json:"inheritBranchingModel,omitempty" tf:"inherit_branching_model,omitempty"`

	// Whether to inherit default merge strategy from project.
	// +kubebuilder:validation:Optional
	InheritDefaultMergeStrategy *bool `json:"inheritDefaultMergeStrategy,omitempty" tf:"inherit_default_merge_strategy,omitempty"`

	// If this should be private or not. Defaults to true.
	// +kubebuilder:validation:Optional
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// What the language of this repository should be.
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// A set of links to a resource related to this object. See Link Below.
	// +kubebuilder:validation:Optional
	Link []LinkParameters `json:"link,omitempty" tf:"link,omitempty"`

	// The owner of this repository. Can be you or any team you
	// have write access to.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Turn on to enable pipelines support.
	// +kubebuilder:validation:Optional
	PipelinesEnabled *bool `json:"pipelinesEnabled,omitempty" tf:"pipelines_enabled,omitempty"`

	// If you want to have this repo associated with a
	// project.
	// +kubebuilder:validation:Optional
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// What SCM you want to use. Valid options are hg or git.
	// Defaults to git.
	// +kubebuilder:validation:Optional
	Scm *string `json:"scm,omitempty" tf:"scm,omitempty"`

	// The slug of the repository.
	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// URL of website associated with this repository.
	// +kubebuilder:validation:Optional
	Website *string `json:"website,omitempty" tf:"website,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Repository is the Schema for the Repositorys API. Provides a Bitbucket Repository
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,bitbucket}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.owner) || (has(self.initProvider) && has(self.initProvider.owner))",message="spec.forProvider.owner is a required parameter"
	Spec   RepositorySpec   `json:"spec"`
	Status RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}