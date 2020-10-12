// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineAllocateRequest v1 machine allocate request
//
// swagger:model v1.MachineAllocateRequest
type V1MachineAllocateRequest struct {

	// a description for this entity
	Description string `json:"description,omitempty"`

	// the hostname for the allocated machine (defaults to metal)
	Hostname string `json:"hostname,omitempty"`

	// the image id to assign this machine to
	// Required: true
	Imageid *string `json:"imageid"`

	// the ips to attach to this machine additionally
	Ips []string `json:"ips"`

	// a readable name for this entity
	Name string `json:"name,omitempty"`

	// the networks that this machine will be placed in.
	Networks []*V1MachineAllocationNetwork `json:"networks"`

	// the partition id to assign this machine to
	// Required: true
	Partitionid *string `json:"partitionid"`

	// the project id to assign this machine to
	// Required: true
	Projectid *string `json:"projectid"`

	// the size id to assign this machine to
	// Required: true
	Sizeid *string `json:"sizeid"`

	// the public ssh keys to access the machine with
	// Required: true
	SSHPubKeys []string `json:"ssh_pub_keys"`

	// tags for this machine
	Tags []string `json:"tags"`

	// cloud-init.io compatible userdata must be base64 encoded
	UserData string `json:"user_data,omitempty"`

	// if this field is set, this specific machine will be allocated if it is not in available state and not currently allocated. this field overrules size and partition
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 machine allocate request
func (m *V1MachineAllocateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPubKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineAllocateRequest) validateImageid(formats strfmt.Registry) error {

	if err := validate.Required("imageid", "body", m.Imageid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocateRequest) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1MachineAllocateRequest) validatePartitionid(formats strfmt.Registry) error {

	if err := validate.Required("partitionid", "body", m.Partitionid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocateRequest) validateProjectid(formats strfmt.Registry) error {

	if err := validate.Required("projectid", "body", m.Projectid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocateRequest) validateSizeid(formats strfmt.Registry) error {

	if err := validate.Required("sizeid", "body", m.Sizeid); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineAllocateRequest) validateSSHPubKeys(formats strfmt.Registry) error {

	if err := validate.Required("ssh_pub_keys", "body", m.SSHPubKeys); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineAllocateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineAllocateRequest) UnmarshalBinary(b []byte) error {
	var res V1MachineAllocateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
