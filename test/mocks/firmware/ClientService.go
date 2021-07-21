// Code generated by mockery v2.7.4. DO NOT EDIT.

package firmware

import (
	firmware "github.com/metal-stack/metal-go/api/client/firmware"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// ListFirmwares provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListFirmwares(params *firmware.ListFirmwaresParams, authInfo runtime.ClientAuthInfoWriter, opts ...firmware.ClientOption) (*firmware.ListFirmwaresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firmware.ListFirmwaresOK
	if rf, ok := ret.Get(0).(func(*firmware.ListFirmwaresParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) *firmware.ListFirmwaresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firmware.ListFirmwaresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firmware.ListFirmwaresParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFirmware provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) RemoveFirmware(params *firmware.RemoveFirmwareParams, authInfo runtime.ClientAuthInfoWriter, opts ...firmware.ClientOption) (*firmware.RemoveFirmwareOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firmware.RemoveFirmwareOK
	if rf, ok := ret.Get(0).(func(*firmware.RemoveFirmwareParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) *firmware.RemoveFirmwareOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firmware.RemoveFirmwareOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firmware.RemoveFirmwareParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) error); ok {
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

// UploadFirmware provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UploadFirmware(params *firmware.UploadFirmwareParams, authInfo runtime.ClientAuthInfoWriter, opts ...firmware.ClientOption) (*firmware.UploadFirmwareOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firmware.UploadFirmwareOK
	if rf, ok := ret.Get(0).(func(*firmware.UploadFirmwareParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) *firmware.UploadFirmwareOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firmware.UploadFirmwareOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firmware.UploadFirmwareParams, runtime.ClientAuthInfoWriter, ...firmware.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
