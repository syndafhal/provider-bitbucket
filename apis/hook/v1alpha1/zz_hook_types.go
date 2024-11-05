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

type HookInitParameters struct {

	// Whether the webhook configuration is active or not (Default: true).
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The name / description to show in the UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The events this webhook is subscribed to. Valid values can be found at Bitbucket Event Payloads Docs.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Whether a webhook history is enabled.
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	// A Webhook secret value. Passing a null or empty secret or not passing a secret will leave the webhook's secret unset. This value is not returned on read and cannot resolve diffs or be imported as its not returned back from bitbucket API.
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Whether to skip certificate verification or not (Default: true).
	SkipCertVerification *bool `json:"skipCertVerification,omitempty" tf:"skip_cert_verification,omitempty"`

	// Where to POST to.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type HookObservation struct {

	// Whether the webhook configuration is active or not (Default: true).
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The name / description to show in the UI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The events this webhook is subscribed to. Valid values can be found at Bitbucket Event Payloads Docs.
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Whether a webhook history is enabled.
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The owner of this repository. Can be you or any team you
	// have write access to.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The name of the repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Whether a webhook secret is set.
	SecretSet *bool `json:"secretSet,omitempty" tf:"secret_set,omitempty"`

	// Whether to skip certificate verification or not (Default: true).
	SkipCertVerification *bool `json:"skipCertVerification,omitempty" tf:"skip_cert_verification,omitempty"`

	// Where to POST to.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The UUID of the workspace webhook.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type HookParameters struct {

	// Whether the webhook configuration is active or not (Default: true).
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The name / description to show in the UI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The events this webhook is subscribed to. Valid values can be found at Bitbucket Event Payloads Docs.
	// +kubebuilder:validation:Optional
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// Whether a webhook history is enabled.
	// +kubebuilder:validation:Optional
	HistoryEnabled *bool `json:"historyEnabled,omitempty" tf:"history_enabled,omitempty"`

	// The owner of this repository. Can be you or any team you
	// have write access to.
	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`

	// The name of the repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-bitbucket/apis/repository/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// A Webhook secret value. Passing a null or empty secret or not passing a secret will leave the webhook's secret unset. This value is not returned on read and cannot resolve diffs or be imported as its not returned back from bitbucket API.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// Whether to skip certificate verification or not (Default: true).
	// +kubebuilder:validation:Optional
	SkipCertVerification *bool `json:"skipCertVerification,omitempty" tf:"skip_cert_verification,omitempty"`

	// Where to POST to.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// HookSpec defines the desired state of Hook
type HookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HookParameters `json:"forProvider"`
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
	InitProvider HookInitParameters `json:"initProvider,omitempty"`
}

// HookStatus defines the observed state of Hook.
type HookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Hook is the Schema for the Hooks API. Provides a Bitbucket Webhook
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,bitbucket}
type Hook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.events) || (has(self.initProvider) && has(self.initProvider.events))",message="spec.forProvider.events is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   HookSpec   `json:"spec"`
	Status HookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HookList contains a list of Hooks
type HookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hook `json:"items"`
}

// Repository type metadata.
var (
	Hook_Kind             = "Hook"
	Hook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Hook_Kind}.String()
	Hook_KindAPIVersion   = Hook_Kind + "." + CRDGroupVersion.String()
	Hook_GroupVersionKind = CRDGroupVersion.WithKind(Hook_Kind)
)

func init() {
	SchemeBuilder.Register(&Hook{}, &HookList{})
}
