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

// V1MachineIPMI The IPMI connection data
//
// swagger:model v1.MachineIPMI
type V1MachineIPMI struct {

	// address
	// Required: true
	Address *string `json:"address" yaml:"address"`

	// bmcversion
	// Required: true
	Bmcversion *string `json:"bmcversion" yaml:"bmcversion"`

	// fru
	// Required: true
	Fru *V1MachineFru `json:"fru" yaml:"fru"`

	// interface
	// Required: true
	Interface *string `json:"interface" yaml:"interface"`

	// last updated
	// Required: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated" yaml:"last_updated"`

	// mac
	// Required: true
	Mac *string `json:"mac" yaml:"mac"`

	// password
	// Required: true
	Password *string `json:"password" yaml:"password"`

	// powermetric
	// Required: true
	Powermetric *V1PowerMetric `json:"powermetric" yaml:"powermetric"`

	// powerstate
	// Required: true
	Powerstate *string `json:"powerstate" yaml:"powerstate"`

	// powersupplies
	// Required: true
	Powersupplies []*V1PowerSupply `json:"powersupplies" yaml:"powersupplies"`

	// user
	// Required: true
	User *string `json:"user" yaml:"user"`
}

// Validate validates this v1 machine IP m i
func (m *V1MachineIPMI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBmcversion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFru(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowermetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerstate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowersupplies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineIPMI) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validateBmcversion(formats strfmt.Registry) error {

	if err := validate.Required("bmcversion", "body", m.Bmcversion); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validateFru(formats strfmt.Registry) error {

	if err := validate.Required("fru", "body", m.Fru); err != nil {
		return err
	}

	if m.Fru != nil {
		if err := m.Fru.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fru")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fru")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineIPMI) validateInterface(formats strfmt.Registry) error {

	if err := validate.Required("interface", "body", m.Interface); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", m.Mac); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validatePowermetric(formats strfmt.Registry) error {

	if err := validate.Required("powermetric", "body", m.Powermetric); err != nil {
		return err
	}

	if m.Powermetric != nil {
		if err := m.Powermetric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powermetric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powermetric")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineIPMI) validatePowerstate(formats strfmt.Registry) error {

	if err := validate.Required("powerstate", "body", m.Powerstate); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineIPMI) validatePowersupplies(formats strfmt.Registry) error {

	if err := validate.Required("powersupplies", "body", m.Powersupplies); err != nil {
		return err
	}

	for i := 0; i < len(m.Powersupplies); i++ {
		if swag.IsZero(m.Powersupplies[i]) { // not required
			continue
		}

		if m.Powersupplies[i] != nil {
			if err := m.Powersupplies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powersupplies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powersupplies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MachineIPMI) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 machine IP m i based on the context it is used
func (m *V1MachineIPMI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFru(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowermetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowersupplies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineIPMI) contextValidateFru(ctx context.Context, formats strfmt.Registry) error {

	if m.Fru != nil {

		if err := m.Fru.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fru")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fru")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineIPMI) contextValidatePowermetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Powermetric != nil {

		if err := m.Powermetric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powermetric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powermetric")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineIPMI) contextValidatePowersupplies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Powersupplies); i++ {

		if m.Powersupplies[i] != nil {

			if swag.IsZero(m.Powersupplies[i]) { // not required
				return nil
			}

			if err := m.Powersupplies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powersupplies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powersupplies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineIPMI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineIPMI) UnmarshalBinary(b []byte) error {
	var res V1MachineIPMI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
