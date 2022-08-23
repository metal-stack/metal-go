// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ProjectFindRequest v1 project find request
//
// swagger:model v1.ProjectFindRequest
type V1ProjectFindRequest struct {

	// description
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// id
	ID string `json:"id,omitempty" yaml:"id,omitempty"`

	// name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty" yaml:"tenant_id,omitempty"`
}

// Validate validates this v1 project find request
func (m *V1ProjectFindRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 project find request based on context it is used
func (m *V1ProjectFindRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ProjectFindRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProjectFindRequest) UnmarshalBinary(b []byte) error {
	var res V1ProjectFindRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
