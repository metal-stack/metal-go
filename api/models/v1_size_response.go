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

// V1SizeResponse v1 size response
//
// swagger:model v1.SizeResponse
type V1SizeResponse struct {

	// the last changed timestamp of this entity
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty"`

	// a list of constraints that defines this size
	// Required: true
	Constraints []*V1SizeConstraint `json:"constraints"`

	// the creation time of this entity
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`
}

// Validate validates this v1 size response
func (m *V1SizeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
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

func (m *V1SizeResponse) validateChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeResponse) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("constraints", "body", m.Constraints); err != nil {
		return err
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SizeResponse) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 size response based on the context it is used
func (m *V1SizeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SizeResponse) contextValidateChanged(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	return nil
}

func (m *V1SizeResponse) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Constraints); i++ {

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1SizeResponse) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SizeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SizeResponse) UnmarshalBinary(b []byte) error {
	var res V1SizeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
