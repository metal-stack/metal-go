// Code generated by mockery v2.7.4. DO NOT EDIT.

package machine

import (
	machine "github.com/metal-stack/metal-go/api/client/machine"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AbortReinstallMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AbortReinstallMachine(params *machine.AbortReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.AbortReinstallMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.AbortReinstallMachineOK
	if rf, ok := ret.Get(0).(func(*machine.AbortReinstallMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.AbortReinstallMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.AbortReinstallMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.AbortReinstallMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddProvisioningEvent provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AddProvisioningEvent(params *machine.AddProvisioningEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.AddProvisioningEventOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.AddProvisioningEventOK
	if rf, ok := ret.Get(0).(func(*machine.AddProvisioningEventParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.AddProvisioningEventOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.AddProvisioningEventOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.AddProvisioningEventParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateMachine(params *machine.AllocateMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.AllocateMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.AllocateMachineOK
	if rf, ok := ret.Get(0).(func(*machine.AllocateMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.AllocateMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.AllocateMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.AllocateMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChassisIdentifyLEDOff provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ChassisIdentifyLEDOff(params *machine.ChassisIdentifyLEDOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.ChassisIdentifyLEDOffOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.ChassisIdentifyLEDOffOK
	if rf, ok := ret.Get(0).(func(*machine.ChassisIdentifyLEDOffParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.ChassisIdentifyLEDOffOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.ChassisIdentifyLEDOffOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.ChassisIdentifyLEDOffParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChassisIdentifyLEDOn provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ChassisIdentifyLEDOn(params *machine.ChassisIdentifyLEDOnParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.ChassisIdentifyLEDOnOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.ChassisIdentifyLEDOnOK
	if rf, ok := ret.Get(0).(func(*machine.ChassisIdentifyLEDOnParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.ChassisIdentifyLEDOnOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.ChassisIdentifyLEDOnOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.ChassisIdentifyLEDOnParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteMachine(params *machine.DeleteMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.DeleteMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.DeleteMachineOK
	if rf, ok := ret.Get(0).(func(*machine.DeleteMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.DeleteMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.DeleteMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.DeleteMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeAllocation provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FinalizeAllocation(params *machine.FinalizeAllocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FinalizeAllocationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FinalizeAllocationOK
	if rf, ok := ret.Get(0).(func(*machine.FinalizeAllocationParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FinalizeAllocationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FinalizeAllocationOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FinalizeAllocationParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIPMIMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindIPMIMachine(params *machine.FindIPMIMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FindIPMIMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FindIPMIMachineOK
	if rf, ok := ret.Get(0).(func(*machine.FindIPMIMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FindIPMIMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FindIPMIMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FindIPMIMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIPMIMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindIPMIMachines(params *machine.FindIPMIMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FindIPMIMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FindIPMIMachinesOK
	if rf, ok := ret.Get(0).(func(*machine.FindIPMIMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FindIPMIMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FindIPMIMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FindIPMIMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindMachine(params *machine.FindMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FindMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FindMachineOK
	if rf, ok := ret.Get(0).(func(*machine.FindMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FindMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FindMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FindMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindMachines(params *machine.FindMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FindMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FindMachinesOK
	if rf, ok := ret.Get(0).(func(*machine.FindMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FindMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FindMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FindMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeMachine(params *machine.FreeMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.FreeMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.FreeMachineOK
	if rf, ok := ret.Get(0).(func(*machine.FreeMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.FreeMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.FreeMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.FreeMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMachineConsolePassword provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetMachineConsolePassword(params *machine.GetMachineConsolePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.GetMachineConsolePasswordOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.GetMachineConsolePasswordOK
	if rf, ok := ret.Get(0).(func(*machine.GetMachineConsolePasswordParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.GetMachineConsolePasswordOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.GetMachineConsolePasswordOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.GetMachineConsolePasswordParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisioningEventContainer provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetProvisioningEventContainer(params *machine.GetProvisioningEventContainerParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.GetProvisioningEventContainerOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.GetProvisioningEventContainerOK
	if rf, ok := ret.Get(0).(func(*machine.GetProvisioningEventContainerParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.GetProvisioningEventContainerOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.GetProvisioningEventContainerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.GetProvisioningEventContainerParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IpmiReport provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) IpmiReport(params *machine.IpmiReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.IpmiReportOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.IpmiReportOK
	if rf, ok := ret.Get(0).(func(*machine.IpmiReportParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.IpmiReportOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.IpmiReportOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.IpmiReportParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListMachines(params *machine.ListMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.ListMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.ListMachinesOK
	if rf, ok := ret.Get(0).(func(*machine.ListMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.ListMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.ListMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.ListMachinesParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineBios provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineBios(params *machine.MachineBiosParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineBiosOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineBiosOK
	if rf, ok := ret.Get(0).(func(*machine.MachineBiosParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineBiosOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineBiosOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineBiosParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineCycle provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineCycle(params *machine.MachineCycleParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineCycleOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineCycleOK
	if rf, ok := ret.Get(0).(func(*machine.MachineCycleParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineCycleOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineCycleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineCycleParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineDisk provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineDisk(params *machine.MachineDiskParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineDiskOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineDiskOK
	if rf, ok := ret.Get(0).(func(*machine.MachineDiskParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineDiskOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineDiskOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineDiskParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineOff provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineOff(params *machine.MachineOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineOffOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineOffOK
	if rf, ok := ret.Get(0).(func(*machine.MachineOffParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineOffOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineOffOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineOffParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineOn provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineOn(params *machine.MachineOnParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineOnOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineOnOK
	if rf, ok := ret.Get(0).(func(*machine.MachineOnParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineOnOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineOnOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineOnParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachinePxe provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachinePxe(params *machine.MachinePxeParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachinePxeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachinePxeOK
	if rf, ok := ret.Get(0).(func(*machine.MachinePxeParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachinePxeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachinePxeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachinePxeParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineReset provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineReset(params *machine.MachineResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.MachineResetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.MachineResetOK
	if rf, ok := ret.Get(0).(func(*machine.MachineResetParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.MachineResetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.MachineResetOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.MachineResetParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) RegisterMachine(params *machine.RegisterMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.RegisterMachineOK, *machine.RegisterMachineCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.RegisterMachineOK
	if rf, ok := ret.Get(0).(func(*machine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.RegisterMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.RegisterMachineOK)
		}
	}

	var r1 *machine.RegisterMachineCreated
	if rf, ok := ret.Get(1).(func(*machine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.RegisterMachineCreated); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*machine.RegisterMachineCreated)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*machine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReinstallMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ReinstallMachine(params *machine.ReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.ReinstallMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.ReinstallMachineOK
	if rf, ok := ret.Get(0).(func(*machine.ReinstallMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.ReinstallMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.ReinstallMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.ReinstallMachineParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetChassisIdentifyLEDState provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) SetChassisIdentifyLEDState(params *machine.SetChassisIdentifyLEDStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.SetChassisIdentifyLEDStateOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.SetChassisIdentifyLEDStateOK
	if rf, ok := ret.Get(0).(func(*machine.SetChassisIdentifyLEDStateParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.SetChassisIdentifyLEDStateOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.SetChassisIdentifyLEDStateOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.SetChassisIdentifyLEDStateParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetMachineState provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) SetMachineState(params *machine.SetMachineStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.SetMachineStateOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.SetMachineStateOK
	if rf, ok := ret.Get(0).(func(*machine.SetMachineStateParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.SetMachineStateOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.SetMachineStateOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.SetMachineStateParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateFirmware provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateFirmware(params *machine.UpdateFirmwareParams, authInfo runtime.ClientAuthInfoWriter, opts ...machine.ClientOption) (*machine.UpdateFirmwareOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *machine.UpdateFirmwareOK
	if rf, ok := ret.Get(0).(func(*machine.UpdateFirmwareParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) *machine.UpdateFirmwareOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machine.UpdateFirmwareOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*machine.UpdateFirmwareParams, runtime.ClientAuthInfoWriter, ...machine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
