// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NetworkAllocateRequest v1 network allocate request
//
// swagger:model v1.NetworkAllocateRequest
type V1NetworkAllocateRequest struct {

	// can be ipv4 or ipv6, defaults to ipv4
	AddressFamily string `json:"address_family,omitempty" yaml:"address_family,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the destination prefixes of this network
	Destinationprefixes []string `json:"destinationprefixes" yaml:"destinationprefixes"`

	// free labels that you associate with this network.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// the bitlen of the prefix to allocate, defaults to childprefixlength of super prefix
	Length int64 `json:"length,omitempty" yaml:"length,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// if set to true, packets leaving this network get masqueraded behind interface ip
	Nat bool `json:"nat,omitempty" yaml:"nat,omitempty"`

	// the partition this network belongs to
	Partitionid string `json:"partitionid,omitempty" yaml:"partitionid,omitempty"`

	// the project id this network belongs to, can be empty if globally available
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// marks a network as shareable.
	Shared bool `json:"shared,omitempty" yaml:"shared,omitempty"`
}

// Validate validates this v1 network allocate request
func (m *V1NetworkAllocateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 network allocate request based on context it is used
func (m *V1NetworkAllocateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
