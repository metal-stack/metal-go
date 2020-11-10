// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PartitionUpdateRequest v1 partition update request
//
// swagger:model v1.PartitionUpdateRequest
type V1PartitionUpdateRequest struct {

	// the boot configuration of this partition
	Bootconfig *V1PartitionBootConfiguration `json:"bootconfig,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	// Unique: true
	ID *string `json:"id"`

	// the address to the management service of this partition
	Mgmtserviceaddress string `json:"mgmtserviceaddress,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 partition update request
func (m *V1PartitionUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PartitionUpdateRequest) validateBootconfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Bootconfig) { // not required
		return nil
	}

	if m.Bootconfig != nil {
		if err := m.Bootconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootconfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1PartitionUpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PartitionUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PartitionUpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1PartitionUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
