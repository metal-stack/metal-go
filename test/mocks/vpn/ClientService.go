// Code generated by mockery v2.39.1. DO NOT EDIT.

package vpn

import (
	clientvpn "github.com/metal-stack/metal-go/api/client/vpn"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetVPNAuthKey provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetVPNAuthKey(params *clientvpn.GetVPNAuthKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientvpn.ClientOption) (*clientvpn.GetVPNAuthKeyOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetVPNAuthKey")
	}

	var r0 *clientvpn.GetVPNAuthKeyOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientvpn.GetVPNAuthKeyParams, runtime.ClientAuthInfoWriter, ...clientvpn.ClientOption) (*clientvpn.GetVPNAuthKeyOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientvpn.GetVPNAuthKeyParams, runtime.ClientAuthInfoWriter, ...clientvpn.ClientOption) *clientvpn.GetVPNAuthKeyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientvpn.GetVPNAuthKeyOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientvpn.GetVPNAuthKeyParams, runtime.ClientAuthInfoWriter, ...clientvpn.ClientOption) error); ok {
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

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
