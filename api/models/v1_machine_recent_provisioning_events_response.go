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

// V1MachineRecentProvisioningEventsResponse v1 machine recent provisioning events response
//
// swagger:model v1.MachineRecentProvisioningEventsResponse
type V1MachineRecentProvisioningEventsResponse struct {

	// number of events stored
	// Required: true
	Events *int64 `json:"events" yaml:"events"`

	// slice of machineIDs for which event was not published
	// Required: true
	Failed []string `json:"failed" yaml:"failed"`
}

// Validate validates this v1 machine recent provisioning events response
func (m *V1MachineRecentProvisioningEventsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineRecentProvisioningEventsResponse) validateEvents(formats strfmt.Registry) error {

	if err := validate.Required("events", "body", m.Events); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineRecentProvisioningEventsResponse) validateFailed(formats strfmt.Registry) error {

	if err := validate.Required("failed", "body", m.Failed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine recent provisioning events response based on context it is used
func (m *V1MachineRecentProvisioningEventsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineRecentProvisioningEventsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineRecentProvisioningEventsResponse) UnmarshalBinary(b []byte) error {
	var res V1MachineRecentProvisioningEventsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
