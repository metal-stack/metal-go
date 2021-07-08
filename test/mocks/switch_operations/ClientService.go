// Code generated by mockery (devel). DO NOT EDIT.

package switch_operations

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	switch_operations "github.com/metal-stack/metal-go/api/client/switch_operations"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// DeleteSwitch provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteSwitch(params *switch_operations.DeleteSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.DeleteSwitchOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.DeleteSwitchOK
	if rf, ok := ret.Get(0).(func(*switch_operations.DeleteSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.DeleteSwitchOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.DeleteSwitchOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*switch_operations.DeleteSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSwitch provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindSwitch(params *switch_operations.FindSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.FindSwitchOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.FindSwitchOK
	if rf, ok := ret.Get(0).(func(*switch_operations.FindSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.FindSwitchOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.FindSwitchOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*switch_operations.FindSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSwitches provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListSwitches(params *switch_operations.ListSwitchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.ListSwitchesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.ListSwitchesOK
	if rf, ok := ret.Get(0).(func(*switch_operations.ListSwitchesParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.ListSwitchesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.ListSwitchesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*switch_operations.ListSwitchesParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifySwitch provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) NotifySwitch(params *switch_operations.NotifySwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.NotifySwitchOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.NotifySwitchOK
	if rf, ok := ret.Get(0).(func(*switch_operations.NotifySwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.NotifySwitchOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.NotifySwitchOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*switch_operations.NotifySwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterSwitch provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) RegisterSwitch(params *switch_operations.RegisterSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.RegisterSwitchOK, *switch_operations.RegisterSwitchCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.RegisterSwitchOK
	if rf, ok := ret.Get(0).(func(*switch_operations.RegisterSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.RegisterSwitchOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.RegisterSwitchOK)
		}
	}

	var r1 *switch_operations.RegisterSwitchCreated
	if rf, ok := ret.Get(1).(func(*switch_operations.RegisterSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.RegisterSwitchCreated); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*switch_operations.RegisterSwitchCreated)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*switch_operations.RegisterSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r2 = rf(params, authInfo, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateSwitch provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateSwitch(params *switch_operations.UpdateSwitchParams, authInfo runtime.ClientAuthInfoWriter, opts ...switch_operations.ClientOption) (*switch_operations.UpdateSwitchOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *switch_operations.UpdateSwitchOK
	if rf, ok := ret.Get(0).(func(*switch_operations.UpdateSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) *switch_operations.UpdateSwitchOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*switch_operations.UpdateSwitchOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*switch_operations.UpdateSwitchParams, runtime.ClientAuthInfoWriter, ...switch_operations.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
