// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1NetworkAllocateRequest v1 network allocate request
// swagger:model v1.NetworkAllocateRequest
type V1NetworkAllocateRequest struct {

	// a description for this entity
	Description string `json:"description,omitempty"`

	// free labels that you associate with this network.
	Labels map[string]string `json:"labels,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// if set to true, packets leaving this network get masqueraded behind interface ip.
	Nat bool `json:"nat,omitempty"`

	// the partition this network belongs to
	Partitionid string `json:"partitionid,omitempty"`

	// the project id this network belongs to, can be empty if globally available
	Projectid string `json:"projectid,omitempty"`

	// marks a network as shareable.
	Shared bool `json:"shared,omitempty"`
}

// Validate validates this v1 network allocate request
func (m *V1NetworkAllocateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkAllocateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkAllocateRequest) UnmarshalBinary(b []byte) error {
	var res V1NetworkAllocateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
