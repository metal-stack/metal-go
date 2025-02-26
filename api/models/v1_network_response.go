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

// V1NetworkResponse v1 network response
//
// swagger:model v1.NetworkResponse
type V1NetworkResponse struct {

	// list of cidrs which are added to the route maps per tenant private network, these are typically pod- and service cidrs, can only be set for private super networks
	AdditionalAnnouncableCIDRs []string `json:"additionalAnnouncableCIDRs" yaml:"additionalAnnouncableCIDRs"`

	// the last changed timestamp of this entity
	// Read Only: true
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty" yaml:"changed,omitempty"`

	// consumption of ips and prefixes in this network
	// Required: true
	Consumption *V1NetworkConsumption `json:"consumption" yaml:"consumption"`

	// the creation time of this entity
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty" yaml:"created,omitempty"`

	// if privatesuper, this defines the bitlen of child prefixes per addressfamily if not nil
	Defaultchildprefixlength map[string]int64 `json:"defaultchildprefixlength,omitempty" yaml:"defaultchildprefixlength,omitempty"`

	// a description for this entity
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// the destination prefixes of this network
	// Required: true
	Destinationprefixes []string `json:"destinationprefixes" yaml:"destinationprefixes"`

	// the unique ID of this entity
	// Required: true
	ID *string `json:"id" yaml:"id"`

	// free labels that you associate with this network.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// a readable name for this entity
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// if set to true, packets leaving this ipv4 network get masqueraded behind interface ip
	// Required: true
	Nat *bool `json:"nat" yaml:"nat"`

	// the id of the parent network
	Parentnetworkid string `json:"parentnetworkid,omitempty" yaml:"parentnetworkid,omitempty"`

	// the partition this network belongs to
	Partitionid string `json:"partitionid,omitempty" yaml:"partitionid,omitempty"`

	// the prefixes of this network
	// Required: true
	Prefixes []string `json:"prefixes" yaml:"prefixes"`

	// if set to true, this network will serve as a partition's super network for the internal machine networks,there can only be one privatesuper network per partition
	// Required: true
	Privatesuper *bool `json:"privatesuper" yaml:"privatesuper"`

	// the project id this network belongs to, can be empty if globally available
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// marks a network as shareable.
	Shared bool `json:"shared,omitempty" yaml:"shared,omitempty"`

	// if set to true, this network can be used for underlay communication
	// Required: true
	Underlay *bool `json:"underlay" yaml:"underlay"`

	// usage of IPv4 ips and prefixes in this network
	// Required: true
	Usage *V1NetworkUsage `json:"usage" yaml:"usage"`

	// the vrf this network is associated with
	Vrf int64 `json:"vrf,omitempty" yaml:"vrf,omitempty"`

	// if set to true, given vrf can be used by multiple networks, which is sometimes useful for network partitioning (default: false)
	Vrfshared bool `json:"vrfshared,omitempty" yaml:"vrfshared,omitempty"`
}

// Validate validates this v1 network response
func (m *V1NetworkResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationprefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivatesuper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnderlay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkResponse) validateChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateConsumption(formats strfmt.Registry) error {

	if err := validate.Required("consumption", "body", m.Consumption); err != nil {
		return err
	}

	if m.Consumption != nil {
		if err := m.Consumption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption")
			}
			return err
		}
	}

	return nil
}

func (m *V1NetworkResponse) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateDestinationprefixes(formats strfmt.Registry) error {

	if err := validate.Required("destinationprefixes", "body", m.Destinationprefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateNat(formats strfmt.Registry) error {

	if err := validate.Required("nat", "body", m.Nat); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validatePrefixes(formats strfmt.Registry) error {

	if err := validate.Required("prefixes", "body", m.Prefixes); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validatePrivatesuper(formats strfmt.Registry) error {

	if err := validate.Required("privatesuper", "body", m.Privatesuper); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateUnderlay(formats strfmt.Registry) error {

	if err := validate.Required("underlay", "body", m.Underlay); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 network response based on the context it is used
func (m *V1NetworkResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsumption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkResponse) contextValidateChanged(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "changed", "body", strfmt.DateTime(m.Changed)); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) contextValidateConsumption(ctx context.Context, formats strfmt.Registry) error {

	if m.Consumption != nil {

		if err := m.Consumption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumption")
			}
			return err
		}
	}

	return nil
}

func (m *V1NetworkResponse) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkResponse) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {

		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkResponse) UnmarshalBinary(b []byte) error {
	var res V1NetworkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
