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

// V1NetworkImmutable a network which contains prefixes from which IP addresses can be allocated
// prefixes that are reachable within this network
//
// swagger:model v1.NetworkImmutable
type V1NetworkImmutable struct {

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes" yaml:"destinationprefixes"`

	// if set to true, packets leaving this network get masqueraded behind interface ip
	// Required: true
	Nat *bool `json:"nat" yaml:"nat"`

	// the id of the parent network
	Parentnetworkid string `json:"parentnetworkid,omitempty" yaml:"parentnetworkid,omitempty"`

	// the prefixes of this network
	// Required: true
	Prefixes []string `json:"prefixes" yaml:"prefixes"`

	// if set to true, this network will serve as a partition's super network for the internal machine networks,there can only be one privatesuper network per partition
	// Required: true
	Privatesuper *bool `json:"privatesuper" yaml:"privatesuper"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay *bool `json:"underlay" yaml:"underlay"`

	// the vrf this network is associated with
	Vrf int64 `json:"vrf,omitempty" yaml:"vrf,omitempty"`

	// if set to true, given vrf can be used by multiple networks, which is sometimes useful for network partioning (default: false)
	Vrfshared bool `json:"vrfshared,omitempty" yaml:"vrfshared,omitempty"`
}

// Validate validates this v1 network immutable
func (m *V1NetworkImmutable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationprefixes(formats); err != nil {
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

func (m *V1NetworkImmutable) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkImmutable) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkImmutable) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkImmutable) validatePrivatesuper(formats strfmt.Registry) error {

	if err := validate.Required("privatesuper", "body", m.Privatesuper); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkImmutable) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network immutable based on context it is used
func (m *V1NetworkImmutable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkImmutable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkImmutable) UnmarshalBinary(b []byte) error {
	var res V1NetworkImmutable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
