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

// V1MachineNetwork prefixes that are reachable within this network
//
// swagger:model v1.MachineNetwork
type V1MachineNetwork struct {

	// ASN number for this network in the bgp configuration
	// Required: true
	Asn *int64 `json:"asn"`

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes"`

	// the ip addresses of the allocated machine in this vrf
	// Required: true
	Ips []string `json:"ips"`

	// if set to true, packets leaving this network get masqueraded behind interface ip
	// Required: true
	Nat *bool `json:"nat"`

	// the networkID of the allocated machine in this vrf
	// Required: true
	Networkid *string `json:"networkid"`

	// the network type, types can be looked up in the network package of metal-lib
	// Required: true
	Networktype *string `json:"networktype"`

	// the prefixes of this network
	// Required: true
	Prefixes []string `json:"prefixes"`

	// indicates whether this network is the private network of this machine
	// Required: true
	Private *bool `json:"private"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay *bool `json:"underlay"`

	// the vrf of the allocated machine
	// Required: true
	Vrf *int64 `json:"vrf"`
}

// Validate validates this v1 machine network
func (m *V1MachineNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationprefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworktype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderlay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineNetwork) validateAsn(formats strfmt.Registry) error {

	if err := validate.Required("asn", "body", m.Asn); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateIps(formats strfmt.Registry) error {

	if err := validate.Required("ips", "body", m.Ips); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateNetworkid(formats strfmt.Registry) error {

	if err := validate.Required("networkid", "body", m.Networkid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateNetworktype(formats strfmt.Registry) error {

	if err := validate.Required("networktype", "body", m.Networktype); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validatePrivate(formats strfmt.Registry) error {

	if err := validate.Required("private", "body", m.Private); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineNetwork) validateVrf(formats strfmt.Registry) error {

	if err := validate.Required("vrf", "body", m.Vrf); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine network based on context it is used
func (m *V1MachineNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineNetwork) UnmarshalBinary(b []byte) error {
	var res V1MachineNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
