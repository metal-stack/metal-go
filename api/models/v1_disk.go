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

// V1Disk v1 disk
//
// swagger:model v1.Disk
type V1Disk struct {

	// the device to create the partitions
	// Required: true
	Device *string `json:"device" yaml:"device"`

	// list of partitions to create on this disk
	Partitions []*V1DiskPartition `json:"partitions" yaml:"partitions"`

	// if set to true, this disk will be wiped before reinstallation
	// Required: true
	Wipeonreinstall *bool `json:"wipeonreinstall" yaml:"wipeonreinstall"`
}

// Validate validates this v1 disk
func (m *V1Disk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWipeonreinstall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Disk) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *V1Disk) validatePartitions(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Disk) validateWipeonreinstall(formats strfmt.Registry) error {

	if err := validate.Required("wipeonreinstall", "body", m.Wipeonreinstall); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 disk based on the context it is used
func (m *V1Disk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Disk) contextValidatePartitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Partitions); i++ {

		if m.Partitions[i] != nil {

			if swag.IsZero(m.Partitions[i]) { // not required
				return nil
			}

			if err := m.Partitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Disk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Disk) UnmarshalBinary(b []byte) error {
	var res V1Disk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
