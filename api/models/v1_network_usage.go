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

// V1NetworkUsage v1 network usage
//
// swagger:model v1.NetworkUsage
type V1NetworkUsage struct {

	// the total available IPs
	// Required: true
	AvailableIps *int64 `json:"available_ips"`

	// the total available Prefixes
	// Required: true
	AvailablePrefixes *int64 `json:"available_prefixes"`

	// the total used IPs
	// Required: true
	UsedIps *int64 `json:"used_ips"`

	// the total used Prefixes
	// Required: true
	UsedPrefixes *int64 `json:"used_prefixes"`
}

// Validate validates this v1 network usage
func (m *V1NetworkUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailablePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedPrefixes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkUsage) validateAvailableIps(formats strfmt.Registry) error {

	if err := validate.Required("available_ips", "body", m.AvailableIps); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateAvailablePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("available_prefixes", "body", m.AvailablePrefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateUsedIps(formats strfmt.Registry) error {

	if err := validate.Required("used_ips", "body", m.UsedIps); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkUsage) validateUsedPrefixes(formats strfmt.Registry) error {

	if err := validate.Required("used_prefixes", "body", m.UsedPrefixes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network usage based on context it is used
func (m *V1NetworkUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkUsage) UnmarshalBinary(b []byte) error {
	var res V1NetworkUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
