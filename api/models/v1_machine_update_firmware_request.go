// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineUpdateFirmwareRequest v1 machine update firmware request
//
// swagger:model v1.MachineUpdateFirmwareRequest
type V1MachineUpdateFirmwareRequest struct {

	// a description why the machine has been updated
	// Required: true
	Description *string `json:"description" yaml:"description"`

	// the firmware kind, i.e. [bios|bmc]
	// Required: true
	// Enum: ["bios","bmc"]
	Kind *string `json:"kind" yaml:"kind"`

	// the update revision
	// Required: true
	Revision *string `json:"revision" yaml:"revision"`
}

// Validate validates this v1 machine update firmware request
func (m *V1MachineUpdateFirmwareRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineUpdateFirmwareRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var v1MachineUpdateFirmwareRequestTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bios","bmc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MachineUpdateFirmwareRequestTypeKindPropEnum = append(v1MachineUpdateFirmwareRequestTypeKindPropEnum, v)
	}
}

const (

	// V1MachineUpdateFirmwareRequestKindBios captures enum value "bios"
	V1MachineUpdateFirmwareRequestKindBios string = "bios"

	// V1MachineUpdateFirmwareRequestKindBmc captures enum value "bmc"
	V1MachineUpdateFirmwareRequestKindBmc string = "bmc"
)

// prop value enum
func (m *V1MachineUpdateFirmwareRequest) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1MachineUpdateFirmwareRequestTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1MachineUpdateFirmwareRequest) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineUpdateFirmwareRequest) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine update firmware request based on context it is used
func (m *V1MachineUpdateFirmwareRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineUpdateFirmwareRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineUpdateFirmwareRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineUpdateFirmwareRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
