package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlobReplicaSet describes a BlobReplicaSet resource
type BlobReplicaSet struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec BlobReplicaSpec `json:"spec"`
}

// BlobReplicaSpec is the spec for a BlobReplicaSet resource
type BlobReplicaSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
	Message   string `json:"message"`
	SomeValue *int32 `json:"someValue"`
	Replicas *int32 `json:"replicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlobReplicaSetList is a list of BlobReplicaSet resources
type BlobReplicaSetList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata"`

	Items []BlobReplicaSet `json:"items"`
}
