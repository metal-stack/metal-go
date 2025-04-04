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

// V1SwitchNotifyRequest v1 switch notify request
//
// swagger:model v1.SwitchNotifyRequest
type V1SwitchNotifyRequest struct {

	// the current bgp port states
	BgpPortStates map[string]V1SwitchBGPPortState `json:"bgp_port_states,omitempty" yaml:"bgp_port_states,omitempty"`

	// error
	// Required: true
	Error *string `json:"error" yaml:"error"`

	// the current switch port states
	// Required: true
	PortStates map[string]string `json:"port_states" yaml:"port_states"`

	// the duration of the switch synchronization
	// Required: true
	SyncDuration *int64 `json:"sync_duration" yaml:"sync_duration"`
}

// Validate validates this v1 switch notify request
func (m *V1SwitchNotifyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBgpPortStates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortStates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchNotifyRequest) validateBgpPortStates(formats strfmt.Registry) error {
	if swag.IsZero(m.BgpPortStates) { // not required
		return nil
	}

	for k := range m.BgpPortStates {

		if err := validate.Required("bgp_port_states"+"."+k, "body", m.BgpPortStates[k]); err != nil {
			return err
		}
		if val, ok := m.BgpPortStates[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bgp_port_states" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bgp_port_states" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SwitchNotifyRequest) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchNotifyRequest) validatePortStates(formats strfmt.Registry) error {

	if err := validate.Required("port_states", "body", m.PortStates); err != nil {
		return err
	}

	return nil
}

func (m *V1SwitchNotifyRequest) validateSyncDuration(formats strfmt.Registry) error {

	if err := validate.Required("sync_duration", "body", m.SyncDuration); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 switch notify request based on the context it is used
func (m *V1SwitchNotifyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBgpPortStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SwitchNotifyRequest) contextValidateBgpPortStates(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.BgpPortStates {

		if val, ok := m.BgpPortStates[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SwitchNotifyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SwitchNotifyRequest) UnmarshalBinary(b []byte) error {
	var res V1SwitchNotifyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
