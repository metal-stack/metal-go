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

// V1MachineFinalizeAllocationRequest v1 machine finalize allocation request
//
// swagger:model v1.MachineFinalizeAllocationRequest
type V1MachineFinalizeAllocationRequest struct {

	// the bootloader ID
	// Required: true
	Bootloaderid *string `json:"bootloaderid" yaml:"bootloaderid"`

	// the cmdline
	// Required: true
	Cmdline *string `json:"cmdline" yaml:"cmdline"`

	// the console password which was generated while provisioning
	// Required: true
	ConsolePassword *string `json:"console_password" yaml:"console_password"`

	// the initrd image
	// Required: true
	Initrd *string `json:"initrd" yaml:"initrd"`

	// the kernel
	// Required: true
	Kernel *string `json:"kernel" yaml:"kernel"`

	// the partition that has the OS installed
	// Required: true
	Ospartition *string `json:"ospartition" yaml:"ospartition"`

	// the device name of the primary disk
	// Required: true
	Primarydisk *string `json:"primarydisk" yaml:"primarydisk"`
}

// Validate validates this v1 machine finalize allocation request
func (m *V1MachineFinalizeAllocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootloaderid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCmdline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsolePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitrd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOspartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimarydisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateBootloaderid(formats strfmt.Registry) error {

	if err := validate.Required("bootloaderid", "body", m.Bootloaderid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateCmdline(formats strfmt.Registry) error {

	if err := validate.Required("cmdline", "body", m.Cmdline); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateConsolePassword(formats strfmt.Registry) error {

	if err := validate.Required("console_password", "body", m.ConsolePassword); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateInitrd(formats strfmt.Registry) error {

	if err := validate.Required("initrd", "body", m.Initrd); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateKernel(formats strfmt.Registry) error {

	if err := validate.Required("kernel", "body", m.Kernel); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validateOspartition(formats strfmt.Registry) error {

	if err := validate.Required("ospartition", "body", m.Ospartition); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineFinalizeAllocationRequest) validatePrimarydisk(formats strfmt.Registry) error {

	if err := validate.Required("primarydisk", "body", m.Primarydisk); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 machine finalize allocation request based on context it is used
func (m *V1MachineFinalizeAllocationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineFinalizeAllocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineFinalizeAllocationRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineFinalizeAllocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
