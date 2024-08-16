// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1NetworkCreateRequest v1 network create request
//
// swagger:model v1.NetworkCreateRequest
type V1NetworkCreateRequest struct {

	// list of cidrs which are added to the route maps per tenant private network, these are typically pod- and service cidrs, can only be set for private super networks
	AdditionalAnnouncableCIDRs []string `json:"additionalAnnouncableCIDRs" yaml:"additionalAnnouncableCIDRs"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes" yaml:"destinationprefixes"`

	// the unique ID of this entity, auto-generated if left empty
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// free labels that you associate with this network.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// if set to true, packets leaving this network get masqueraded behind interface ip
	// Required: true
	Nat *bool `json:"nat" yaml:"nat"`

	// the id of the parent network
	Parentnetworkid string `json:"parentnetworkid,omitempty" yaml:"parentnetworkid,omitempty"`

	// the partition this network belongs to
	Partitionid string `json:"partitionid,omitempty" yaml:"partitionid,omitempty"`

	// the prefixes of this network
	// Required: true
	Prefixes []string `json:"prefixes" yaml:"prefixes"`

	// if set to true, this network will serve as a partition's super network for the internal machine networks,there can only be one privatesuper network per partition
	// Required: true
	Privatesuper *bool `json:"privatesuper" yaml:"privatesuper"`

	// the project id this network belongs to, can be empty if globally available
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// marks a network as shareable.
	Shared bool `json:"shared,omitempty" yaml:"shared,omitempty"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay *bool `json:"underlay" yaml:"underlay"`

	// the vrf this network is associated with
	Vrf int64 `json:"vrf,omitempty" yaml:"vrf,omitempty"`

	// if set to true, given vrf can be used by multiple networks, which is sometimes useful for network partitioning (default: false)
	Vrfshared bool `json:"vrfshared,omitempty" yaml:"vrfshared,omitempty"`
}

// Validate validates this v1 network create request
func (m *V1NetworkCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationprefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivatesuper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderlay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkCreateRequest) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkCreateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkCreateRequest) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkCreateRequest) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkCreateRequest) validatePrivatesuper(formats strfmt.Registry) error {

	if err := validate.Required("privatesuper", "body", m.Privatesuper); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkCreateRequest) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network create request based on context it is used
func (m *V1NetworkCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkCreateRequest) UnmarshalBinary(b []byte) error {
	var res V1NetworkCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
