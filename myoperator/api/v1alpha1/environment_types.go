/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	FinalizerKubernetes FinalizerName = "kubernetes"
)

// FinalizerName is the name identifying a finalizer during namespace lifecycle.
type FinalizerName string

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Finalizers is an opaque list of values that must be empty to permanently remove object from storage
	Finalizers []FinalizerName `json:"finalizers,omitempty"`
}

// EnvironmentPhase defines the phase in which the Environment is
type EnvironmentPhase string

// These are the valid phases of a Environment.
const (
	// EnvironmentActive means the Environment is available for use in the system
	EnvironmentActive EnvironmentPhase = "Active"
	// EnvironmentTerminating means the Environment is undergoing graceful termination
	EnvironmentTerminating EnvironmentPhase = "Terminating"
)

// ConditionStatus defines conditions of resources
type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition;
// "ConditionFalse" means a resource is not in the condition; "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// EnvironmentStatus defines the observed state of Environment
type EnvironmentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Phase is the current lifecycle phase of the Environment.
	// +optional
	Phase EnvironmentPhase `json:"phase,omitempty"`
	// +optional
	Conditions []EnvironmentCondition `json:"conditions,omitempty"`
}

// EnvironmentConditionType defines constants reporting on status during namespace lifetime and deletion progress
type EnvironmentConditionType string

// These are valid conditions of a Environment.
const (
	EnvironmentDeletionDiscoveryFailure EnvironmentConditionType = "NamespaceDeletionDiscoveryFailure"
	EnvironmentDeletionContentFailure   EnvironmentConditionType = "NamespaceDeletionContentFailure"
	EnvironmentDeletionGVParsingFailure EnvironmentConditionType = "NamespaceDeletionGroupVersionParsingFailure"
)

// EnvironmentCondition contains details about state of Environment.
type EnvironmentCondition struct {
	// Type of namespace controller condition.
	Type EnvironmentConditionType `json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status,omitempty"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Environment is the Schema for the environments API
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentSpec   `json:"spec,omitempty"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EnvironmentList contains a list of Environment
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
