// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineHardwareBase v1 machine hardware base
//
// swagger:model v1.MachineHardwareBase
type V1MachineHardwareBase struct {

	// the number of cpu cores
	// Required: true
	CPUCores *int32 `json:"cpu_cores" yaml:"cpu_cores"`

	// the list of block devices of this machine
	// Required: true
	Disks []*V1MachineBlockDevice `json:"disks" yaml:"disks"`

	// the total memory of the machine
	// Required: true
	Memory *int64 `json:"memory" yaml:"memory"`
}

// Validate validates this v1 machine hardware base
func (m *V1MachineHardwareBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineHardwareBase) validateCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("cpu_cores", "body", m.CPUCores); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineHardwareBase) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MachineHardwareBase) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 machine hardware base based on the context it is used
func (m *V1MachineHardwareBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineHardwareBase) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {

			if swag.IsZero(m.Disks[i]) { // not required
				return nil
			}

			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineHardwareBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineHardwareBase) UnmarshalBinary(b []byte) error {
	var res V1MachineHardwareBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
