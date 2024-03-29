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

// V1MachineConsolePasswordRequest v1 machine console password request
//
// swagger:model v1.MachineConsolePasswordRequest
type V1MachineConsolePasswordRequest struct {

	// id of the machine to get the consolepassword for
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// reason why the consolepassword is requested, typically a incident number with short description
	// Required: true
	Reason *string `json:"reason" yaml:"reason"`
}

// Validate validates this v1 machine console password request
func (m *V1MachineConsolePasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineConsolePasswordRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineConsolePasswordRequest) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine console password request based on context it is used
func (m *V1MachineConsolePasswordRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineConsolePasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineConsolePasswordRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineConsolePasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
