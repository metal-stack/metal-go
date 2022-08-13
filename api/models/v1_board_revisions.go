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

// V1BoardRevisions v1 board revisions
//
// swagger:model v1.BoardRevisions
type V1BoardRevisions struct {

	// board revisions
	// Required: true
	BoardRevisions map[string][]string `json:"BoardRevisions" yaml:"BoardRevisions"`
}

// Validate validates this v1 board revisions
func (m *V1BoardRevisions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoardRevisions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BoardRevisions) validateBoardRevisions(formats strfmt.Registry) error {

	if err := validate.Required("BoardRevisions", "body", m.BoardRevisions); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 board revisions based on context it is used
func (m *V1BoardRevisions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1BoardRevisions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BoardRevisions) UnmarshalBinary(b []byte) error {
	var res V1BoardRevisions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
