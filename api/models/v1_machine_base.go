// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MachineBase A machine representing a bare metal machine.
//
// swagger:model v1.MachineBase
type V1MachineBase struct {

	// the allocation data of an allocated machine
	Allocation *V1MachineAllocation `json:"allocation,omitempty" yaml:"allocation,omitempty"`

	// bios information of this machine
	// Required: true
	Bios *V1MachineBIOS `json:"bios" yaml:"bios"`

	// recent events of this machine during provisioning
	// Required: true
	Events *V1MachineRecentProvisioningEvents `json:"events" yaml:"events"`

	// the hardware of this machine
	// Required: true
	Hardware *V1MachineHardware `json:"hardware" yaml:"hardware"`

	// the state of this chassis identify LED
	// Required: true
	Ledstate *V1ChassisIdentifyLEDState `json:"ledstate" yaml:"ledstate"`

	// the liveliness of this machine
	// Required: true
	// Enum: [Alive Dead Hibernated Unknown]
	Liveliness *string `json:"liveliness" yaml:"liveliness"`

	// the partition assigned to this machine
	// Read Only: true
	Partition *V1PartitionResponse `json:"partition,omitempty" yaml:"partition,omitempty"`

	// the rack assigned to this machine
	// Read Only: true
	Rackid string `json:"rackid,omitempty" yaml:"rackid,omitempty"`

	// the size of this machine
	// Read Only: true
	Size *V1SizeResponse `json:"size,omitempty" yaml:"size,omitempty"`

	// the state of this machine
	// Required: true
	State *V1MachineState `json:"state" yaml:"state"`

	// tags for this machine
	// Required: true
	Tags []string `json:"tags" yaml:"tags"`
}

// Validate validates this v1 machine base
func (m *V1MachineBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBios(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLedstate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLiveliness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineBase) validateAllocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocation) { // not required
		return nil
	}

	if m.Allocation != nil {
		if err := m.Allocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateBios(formats strfmt.Registry) error {

	if err := validate.Required("bios", "body", m.Bios); err != nil {
		return err
	}

	if m.Bios != nil {
		if err := m.Bios.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bios")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bios")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateEvents(formats strfmt.Registry) error {

	if err := validate.Required("events", "body", m.Events); err != nil {
		return err
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateHardware(formats strfmt.Registry) error {

	if err := validate.Required("hardware", "body", m.Hardware); err != nil {
		return err
	}

	if m.Hardware != nil {
		if err := m.Hardware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateLedstate(formats strfmt.Registry) error {

	if err := validate.Required("ledstate", "body", m.Ledstate); err != nil {
		return err
	}

	if m.Ledstate != nil {
		if err := m.Ledstate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ledstate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ledstate")
			}
			return err
		}
	}

	return nil
}

var v1MachineBaseTypeLivelinessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Alive","Dead","Hibernated","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MachineBaseTypeLivelinessPropEnum = append(v1MachineBaseTypeLivelinessPropEnum, v)
	}
}

const (

	// V1MachineBaseLivelinessAlive captures enum value "Alive"
	V1MachineBaseLivelinessAlive string = "Alive"

	// V1MachineBaseLivelinessDead captures enum value "Dead"
	V1MachineBaseLivelinessDead string = "Dead"

	// V1MachineBaseLivelinessHibernated captures enum value "Hibernated"
	V1MachineBaseLivelinessHibernated string = "Hibernated"

	// V1MachineBaseLivelinessUnknown captures enum value "Unknown"
	V1MachineBaseLivelinessUnknown string = "Unknown"
)

// prop value enum
func (m *V1MachineBase) validateLivelinessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1MachineBaseTypeLivelinessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1MachineBase) validateLiveliness(formats strfmt.Registry) error {

	if err := validate.Required("liveliness", "body", m.Liveliness); err != nil {
		return err
	}

	// value enum
	if err := m.validateLivelinessEnum("liveliness", "body", *m.Liveliness); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineBase) validatePartition(formats strfmt.Registry) error {
	if swag.IsZero(m.Partition) { // not required
		return nil
	}

	if m.Partition != nil {
		if err := m.Partition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partition")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 machine base based on the context it is used
func (m *V1MachineBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBios(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHardware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLedstate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRackid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MachineBase) contextValidateAllocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocation != nil {
		if err := m.Allocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateBios(ctx context.Context, formats strfmt.Registry) error {

	if m.Bios != nil {
		if err := m.Bios.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bios")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bios")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.Events != nil {
		if err := m.Events.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("events")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateHardware(ctx context.Context, formats strfmt.Registry) error {

	if m.Hardware != nil {
		if err := m.Hardware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateLedstate(ctx context.Context, formats strfmt.Registry) error {

	if m.Ledstate != nil {
		if err := m.Ledstate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ledstate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ledstate")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidatePartition(ctx context.Context, formats strfmt.Registry) error {

	if m.Partition != nil {
		if err := m.Partition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partition")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateRackid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rackid", "body", string(m.Rackid)); err != nil {
		return err
	}

	return nil
}

func (m *V1MachineBase) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *V1MachineBase) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MachineBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MachineBase) UnmarshalBinary(b []byte) error {
	var res V1MachineBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
