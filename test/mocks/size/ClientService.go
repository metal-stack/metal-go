// Code generated by mockery (devel). DO NOT EDIT.

package size

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	size "github.com/metal-stack/metal-go/api/client/size"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateSize provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateSize(params *size.CreateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.CreateSizeCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.CreateSizeCreated
	if rf, ok := ret.Get(0).(func(*size.CreateSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.CreateSizeCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.CreateSizeCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.CreateSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSize provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteSize(params *size.DeleteSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.DeleteSizeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.DeleteSizeOK
	if rf, ok := ret.Get(0).(func(*size.DeleteSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.DeleteSizeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.DeleteSizeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.DeleteSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSize provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindSize(params *size.FindSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.FindSizeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.FindSizeOK
	if rf, ok := ret.Get(0).(func(*size.FindSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.FindSizeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.FindSizeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.FindSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FromHardware provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FromHardware(params *size.FromHardwareParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.FromHardwareOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.FromHardwareOK
	if rf, ok := ret.Get(0).(func(*size.FromHardwareParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.FromHardwareOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.FromHardwareOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.FromHardwareParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSizes provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListSizes(params *size.ListSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.ListSizesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.ListSizesOK
	if rf, ok := ret.Get(0).(func(*size.ListSizesParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.ListSizesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.ListSizesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.ListSizesParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
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

// UpdateSize provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateSize(params *size.UpdateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...size.ClientOption) (*size.UpdateSizeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *size.UpdateSizeOK
	if rf, ok := ret.Get(0).(func(*size.UpdateSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) *size.UpdateSizeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*size.UpdateSizeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*size.UpdateSizeParams, runtime.ClientAuthInfoWriter, ...size.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
