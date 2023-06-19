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

	// constraints which must match that this layout is taken, if sizes and images are empty these are develop layouts
	// Required: true
	Constraints *V1FilesystemLayoutConstraints `json:"constraints" yaml:"constraints"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// list of disks that belong to this layout
	Disks []*V1Disk `json:"disks" yaml:"disks"`

	// list of filesystems to create
	Filesystems []*V1Filesystem `json:"filesystems" yaml:"filesystems"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// list of logicalvolumes to create
	Logicalvolumes []*V1LogicalVolume `json:"logicalvolumes" yaml:"logicalvolumes"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// list of raid arrays to create
	Raid []*V1Raid `json:"raid" yaml:"raid"`

	// list of volumegroups to create
	Volumegroups []*V1VolumeGroup `json:"volumegroups" yaml:"volumegroups"`
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

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalvolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumegroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FilesystemLayoutResponse) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("constraints", "body", m.Constraints); err != nil {
		return err
	}

	if m.Constraints != nil {
		if err := m.Constraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraints")
			}
			return err
		}
	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
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

func (m *V1FilesystemLayoutResponse) validateFilesystems(formats strfmt.Registry) error {
	if swag.IsZero(m.Filesystems) { // not required
		return nil
	}

	for i := 0; i < len(m.Filesystems); i++ {
		if swag.IsZero(m.Filesystems[i]) { // not required
			continue
		}

		if m.Filesystems[i] != nil {
			if err := m.Filesystems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesystems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesystems" + "." + strconv.Itoa(i))
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

func (m *V1FilesystemLayoutResponse) validateLogicalvolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Logicalvolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Logicalvolumes); i++ {
		if swag.IsZero(m.Logicalvolumes[i]) { // not required
			continue
		}

		if m.Logicalvolumes[i] != nil {
			if err := m.Logicalvolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logicalvolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logicalvolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateRaid(formats strfmt.Registry) error {
	if swag.IsZero(m.Raid) { // not required
		return nil
	}

	for i := 0; i < len(m.Raid); i++ {
		if swag.IsZero(m.Raid[i]) { // not required
			continue
		}

		if m.Raid[i] != nil {
			if err := m.Raid[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raid" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raid" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) validateVolumegroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumegroups) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumegroups); i++ {
		if swag.IsZero(m.Volumegroups[i]) { // not required
			continue
		}

		if m.Volumegroups[i] != nil {
			if err := m.Volumegroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumegroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumegroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

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

	if err := m.contextValidateLogicalvolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumegroups(ctx, formats); err != nil {
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
				return ve.ValidateName("constraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constraints")
			}
			return err
		}
	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1FilesystemLayoutResponse) contextValidateFilesystems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filesystems); i++ {

		if m.Filesystems[i] != nil {

			if swag.IsZero(m.Filesystems[i]) { // not required
				return nil
			}

			if err := m.Filesystems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filesystems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filesystems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateLogicalvolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logicalvolumes); i++ {

		if m.Logicalvolumes[i] != nil {

			if swag.IsZero(m.Logicalvolumes[i]) { // not required
				return nil
			}

			if err := m.Logicalvolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logicalvolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logicalvolumes" + "." + strconv.Itoa(i))
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

			if swag.IsZero(m.Raid[i]) { // not required
				return nil
			}

			if err := m.Raid[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("raid" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("raid" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1FilesystemLayoutResponse) contextValidateVolumegroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumegroups); i++ {

		if m.Volumegroups[i] != nil {

			if swag.IsZero(m.Volumegroups[i]) { // not required
				return nil
			}

			if err := m.Volumegroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumegroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumegroups" + "." + strconv.Itoa(i))
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
