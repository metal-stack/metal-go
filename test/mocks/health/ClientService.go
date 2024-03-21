<<<<<<< HEAD
// Code generated by mockery v2.42.0. DO NOT EDIT.
=======
// Code generated by mockery v2.42.1. DO NOT EDIT.
>>>>>>> e7d1a0151b0e198a4a6d168e11a9cf836e214d88

package health

import (
	clienthealth "github.com/metal-stack/metal-go/api/client/health"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// Health provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Health(params *clienthealth.HealthParams, authInfo runtime.ClientAuthInfoWriter, opts ...clienthealth.ClientOption) (*clienthealth.HealthOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Health")
	}

	var r0 *clienthealth.HealthOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clienthealth.HealthParams, runtime.ClientAuthInfoWriter, ...clienthealth.ClientOption) (*clienthealth.HealthOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clienthealth.HealthParams, runtime.ClientAuthInfoWriter, ...clienthealth.ClientOption) *clienthealth.HealthOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clienthealth.HealthOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clienthealth.HealthParams, runtime.ClientAuthInfoWriter, ...clienthealth.ClientOption) error); ok {
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
