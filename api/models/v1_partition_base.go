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
)

// V1PartitionBase v1 partition base
//
// swagger:model v1.PartitionBase
type V1PartitionBase struct {

	// the dns servers for this partition
	DNSServers []*V1DNSServer `json:"dns_servers" yaml:"dns_servers"`

	// free labels that you associate with this partition
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// the address to the management service of this partition
	Mgmtserviceaddress string `json:"mgmtserviceaddress,omitempty" yaml:"mgmtserviceaddress,omitempty"`

	// the ntp servers for this partition
	NtpServers []*V1NTPServer `json:"ntp_servers" yaml:"ntp_servers"`
}

// Validate validates this v1 partition base
func (m *V1PartitionBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtpServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PartitionBase) validateDNSServers(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServers) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServers); i++ {
		if swag.IsZero(m.DNSServers[i]) { // not required
			continue
		}

		if m.DNSServers[i] != nil {
			if err := m.DNSServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PartitionBase) validateNtpServers(formats strfmt.Registry) error {
	if swag.IsZero(m.NtpServers) { // not required
		return nil
	}

	for i := 0; i < len(m.NtpServers); i++ {
		if swag.IsZero(m.NtpServers[i]) { // not required
			continue
		}

		if m.NtpServers[i] != nil {
			if err := m.NtpServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 partition base based on the context it is used
func (m *V1PartitionBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNtpServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PartitionBase) contextValidateDNSServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServers); i++ {

		if m.DNSServers[i] != nil {

			if swag.IsZero(m.DNSServers[i]) { // not required
				return nil
			}

			if err := m.DNSServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1PartitionBase) contextValidateNtpServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NtpServers); i++ {

		if m.NtpServers[i] != nil {

			if swag.IsZero(m.NtpServers[i]) { // not required
				return nil
			}

			if err := m.NtpServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntp_servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PartitionBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PartitionBase) UnmarshalBinary(b []byte) error {
	var res V1PartitionBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
