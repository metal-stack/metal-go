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

// V1FilesystemLayoutResponse v1 filesystem layout response
//
// swagger:model v1.FilesystemLayoutResponse
type V1FilesystemLayoutResponse struct {

	// constraints
	// Required: true
	Constraints *V1FilesystemLayoutConstraints `json:"Constraints"`

	// disks
	// Required: true
	Disks []*V1Disk `json:"Disks"`

	// filesystems
	// Required: true
	Filesystems []*V1Filesystem `json:"Filesystems"`

	// raid
	// Required: true
	Raid []*V1Raid `json:"Raid"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 filesystem layout response
func (m *V1FilesystemLayoutResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaid(formats); err != nil {
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

func (m *V1FilesystemLayoutResponse) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("Constraints", "body", m.Constraints); err != nil {
		return err
	}

	if m.Constraints != nil {
		if err := m.Constraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constraints")
			}
			return err
		}
	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("Disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateFilesystems(formats strfmt.Registry) error {

	if err := validate.Required("Filesystems", "body", m.Filesystems); err != nil {
		return err
	}

	for i := 0; i < len(m.Filesystems); i++ {
		if swag.IsZero(m.Filesystems[i]) { // not required
			continue
		}

		if m.Filesystems[i] != nil {
			if err := m.Filesystems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Filesystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateRaid(formats strfmt.Registry) error {

	if err := validate.Required("Raid", "body", m.Raid); err != nil {
		return err
	}

	for i := 0; i < len(m.Raid); i++ {
		if swag.IsZero(m.Raid[i]) { // not required
			continue
		}

		if m.Raid[i] != nil {
			if err := m.Raid[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Raid" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 filesystem layout response based on the context it is used
func (m *V1FilesystemLayoutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilesystems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.Constraints != nil {
		if err := m.Constraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Constraints")
			}
			return err
		}
	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateFilesystems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filesystems); i++ {

		if m.Filesystems[i] != nil {
			if err := m.Filesystems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Filesystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateRaid(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Raid); i++ {

		if m.Raid[i] != nil {
			if err := m.Raid[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Raid" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FilesystemLayoutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FilesystemLayoutResponse) UnmarshalBinary(b []byte) error {
	var res V1FilesystemLayoutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
