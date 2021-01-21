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

// V1MachineBlockDevice v1 machine block device
//
// swagger:model v1.MachineBlockDevice
type V1MachineBlockDevice struct {

	// the name of this block device
	// Required: true
	Name *string `json:"name"`

	// the partitions of this disk
	Partitions []*V1MachineDiskPartition `json:"partitions"`

	// whether this disk has the OS installed
	// Required: true
	Primary *bool `json:"primary"`

	// the size of this block device
	// Required: true
	Size *int64 `json:"size"`
}

// Validate validates this v1 machine block device
func (m *V1MachineBlockDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineBlockDevice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineBlockDevice) validatePartitions(formats strfmt.Registry) error {
	if swag.IsZero(m.Partitions) { // not required
		return nil
	}

	for i := 0; i < len(m.Partitions); i++ {
		if swag.IsZero(m.Partitions[i]) { // not required
			continue
		}

		if m.Partitions[i] != nil {
			if err := m.Partitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MachineBlockDevice) validatePrimary(formats strfmt.Registry) error {

	if err := validate.Required("primary", "body", m.Primary); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineBlockDevice) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 machine block device based on the context it is used
func (m *V1MachineBlockDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineBlockDevice) contextValidatePartitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Partitions); i++ {

		if m.Partitions[i] != nil {
			if err := m.Partitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineBlockDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineBlockDevice) UnmarshalBinary(b []byte) error {
	var res V1MachineBlockDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
