// Code generated by mockery v2.12.2. DO NOT EDIT.

package operations

import (
	clientoperations "github.com/metal-stack/metal-go/api/client/operations"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"

	testing "testing"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetVPNAuthKey provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetVPNAuthKey(params *clientoperations.GetVPNAuthKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientoperations.ClientOption) (*clientoperations.GetVPNAuthKeyOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientoperations.GetVPNAuthKeyOK
	if rf, ok := ret.Get(0).(func(*clientoperations.GetVPNAuthKeyParams, runtime.ClientAuthInfoWriter, ...clientoperations.ClientOption) *clientoperations.GetVPNAuthKeyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientoperations.GetVPNAuthKeyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientoperations.GetVPNAuthKeyParams, runtime.ClientAuthInfoWriter, ...clientoperations.ClientOption) error); ok {
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

// NewClientService creates a new instance of ClientService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t testing.TB) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
