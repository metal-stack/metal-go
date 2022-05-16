// Code generated by mockery v2.12.2. DO NOT EDIT.

package machine

import (
	clientmachine "github.com/metal-stack/metal-go/client/machine"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"

	testing "testing"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AbortReinstallMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AbortReinstallMachine(params *clientmachine.AbortReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.AbortReinstallMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.AbortReinstallMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.AbortReinstallMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.AbortReinstallMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.AbortReinstallMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.AbortReinstallMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddProvisioningEvent provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AddProvisioningEvent(params *clientmachine.AddProvisioningEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.AddProvisioningEventOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.AddProvisioningEventOK
	if rf, ok := ret.Get(0).(func(*clientmachine.AddProvisioningEventParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.AddProvisioningEventOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.AddProvisioningEventOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.AddProvisioningEventParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddProvisioningEvents provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AddProvisioningEvents(params *clientmachine.AddProvisioningEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.AddProvisioningEventsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.AddProvisioningEventsOK
	if rf, ok := ret.Get(0).(func(*clientmachine.AddProvisioningEventsParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.AddProvisioningEventsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.AddProvisioningEventsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.AddProvisioningEventsParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateMachine(params *clientmachine.AllocateMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.AllocateMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.AllocateMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.AllocateMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.AllocateMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.AllocateMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.AllocateMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChassisIdentifyLEDOff provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ChassisIdentifyLEDOff(params *clientmachine.ChassisIdentifyLEDOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.ChassisIdentifyLEDOffOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.ChassisIdentifyLEDOffOK
	if rf, ok := ret.Get(0).(func(*clientmachine.ChassisIdentifyLEDOffParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.ChassisIdentifyLEDOffOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.ChassisIdentifyLEDOffOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.ChassisIdentifyLEDOffParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChassisIdentifyLEDOn provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ChassisIdentifyLEDOn(params *clientmachine.ChassisIdentifyLEDOnParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.ChassisIdentifyLEDOnOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.ChassisIdentifyLEDOnOK
	if rf, ok := ret.Get(0).(func(*clientmachine.ChassisIdentifyLEDOnParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.ChassisIdentifyLEDOnOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.ChassisIdentifyLEDOnOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.ChassisIdentifyLEDOnParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteMachine(params *clientmachine.DeleteMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.DeleteMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.DeleteMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.DeleteMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.DeleteMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.DeleteMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.DeleteMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeAllocation provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FinalizeAllocation(params *clientmachine.FinalizeAllocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FinalizeAllocationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FinalizeAllocationOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FinalizeAllocationParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FinalizeAllocationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FinalizeAllocationOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FinalizeAllocationParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIPMIMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindIPMIMachine(params *clientmachine.FindIPMIMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FindIPMIMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FindIPMIMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FindIPMIMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FindIPMIMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FindIPMIMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FindIPMIMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIPMIMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindIPMIMachines(params *clientmachine.FindIPMIMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FindIPMIMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FindIPMIMachinesOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FindIPMIMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FindIPMIMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FindIPMIMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FindIPMIMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindMachine(params *clientmachine.FindMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FindMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FindMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FindMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FindMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FindMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FindMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindMachines(params *clientmachine.FindMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FindMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FindMachinesOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FindMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FindMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FindMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FindMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeMachine(params *clientmachine.FreeMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.FreeMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.FreeMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.FreeMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.FreeMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.FreeMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.FreeMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMachineConsolePassword provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetMachineConsolePassword(params *clientmachine.GetMachineConsolePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.GetMachineConsolePasswordOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.GetMachineConsolePasswordOK
	if rf, ok := ret.Get(0).(func(*clientmachine.GetMachineConsolePasswordParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.GetMachineConsolePasswordOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.GetMachineConsolePasswordOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.GetMachineConsolePasswordParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisioningEventContainer provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetProvisioningEventContainer(params *clientmachine.GetProvisioningEventContainerParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.GetProvisioningEventContainerOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.GetProvisioningEventContainerOK
	if rf, ok := ret.Get(0).(func(*clientmachine.GetProvisioningEventContainerParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.GetProvisioningEventContainerOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.GetProvisioningEventContainerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.GetProvisioningEventContainerParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IpmiReport provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) IpmiReport(params *clientmachine.IpmiReportParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.IpmiReportOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.IpmiReportOK
	if rf, ok := ret.Get(0).(func(*clientmachine.IpmiReportParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.IpmiReportOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.IpmiReportOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.IpmiReportParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMachines provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListMachines(params *clientmachine.ListMachinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.ListMachinesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.ListMachinesOK
	if rf, ok := ret.Get(0).(func(*clientmachine.ListMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.ListMachinesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.ListMachinesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.ListMachinesParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineBios provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineBios(params *clientmachine.MachineBiosParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineBiosOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineBiosOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineBiosParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineBiosOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineBiosOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineBiosParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineCycle provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineCycle(params *clientmachine.MachineCycleParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineCycleOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineCycleOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineCycleParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineCycleOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineCycleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineCycleParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineDisk provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineDisk(params *clientmachine.MachineDiskParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineDiskOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineDiskOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineDiskParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineDiskOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineDiskOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineDiskParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineOff provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineOff(params *clientmachine.MachineOffParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineOffOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineOffOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineOffParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineOffOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineOffOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineOffParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineOn provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineOn(params *clientmachine.MachineOnParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineOnOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineOnOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineOnParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineOnOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineOnOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineOnParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachinePxe provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachinePxe(params *clientmachine.MachinePxeParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachinePxeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachinePxeOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachinePxeParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachinePxeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachinePxeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachinePxeParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MachineReset provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MachineReset(params *clientmachine.MachineResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.MachineResetOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.MachineResetOK
	if rf, ok := ret.Get(0).(func(*clientmachine.MachineResetParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.MachineResetOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.MachineResetOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.MachineResetParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) RegisterMachine(params *clientmachine.RegisterMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.RegisterMachineOK, *clientmachine.RegisterMachineCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.RegisterMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.RegisterMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.RegisterMachineOK)
		}
	}

	var r1 *clientmachine.RegisterMachineCreated
	if rf, ok := ret.Get(1).(func(*clientmachine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.RegisterMachineCreated); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*clientmachine.RegisterMachineCreated)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*clientmachine.RegisterMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReinstallMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ReinstallMachine(params *clientmachine.ReinstallMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.ReinstallMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.ReinstallMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.ReinstallMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.ReinstallMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.ReinstallMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.ReinstallMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetChassisIdentifyLEDState provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) SetChassisIdentifyLEDState(params *clientmachine.SetChassisIdentifyLEDStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.SetChassisIdentifyLEDStateOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.SetChassisIdentifyLEDStateOK
	if rf, ok := ret.Get(0).(func(*clientmachine.SetChassisIdentifyLEDStateParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.SetChassisIdentifyLEDStateOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.SetChassisIdentifyLEDStateOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.SetChassisIdentifyLEDStateParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetMachineState provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) SetMachineState(params *clientmachine.SetMachineStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.SetMachineStateOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.SetMachineStateOK
	if rf, ok := ret.Get(0).(func(*clientmachine.SetMachineStateParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.SetMachineStateOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.SetMachineStateOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.SetMachineStateParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
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
func (_m *ClientService) UpdateFirmware(params *clientmachine.UpdateFirmwareParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.UpdateFirmwareOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.UpdateFirmwareOK
	if rf, ok := ret.Get(0).(func(*clientmachine.UpdateFirmwareParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.UpdateFirmwareOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.UpdateFirmwareOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.UpdateFirmwareParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMachine provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateMachine(params *clientmachine.UpdateMachineParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmachine.ClientOption) (*clientmachine.UpdateMachineOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmachine.UpdateMachineOK
	if rf, ok := ret.Get(0).(func(*clientmachine.UpdateMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) *clientmachine.UpdateMachineOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmachine.UpdateMachineOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientmachine.UpdateMachineParams, runtime.ClientAuthInfoWriter, ...clientmachine.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientService creates a new instance of ClientService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t testing.TB) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
