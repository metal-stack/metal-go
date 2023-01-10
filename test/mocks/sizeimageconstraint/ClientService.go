// Code generated by mockery v2.16.0. DO NOT EDIT.

package sizeimageconstraint

import (
	clientsizeimageconstraint "github.com/metal-stack/metal-go/api/client/sizeimageconstraint"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateSizeImageConstraint(params *clientsizeimageconstraint.CreateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.CreateSizeImageConstraintCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.CreateSizeImageConstraintCreated
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.CreateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.CreateSizeImageConstraintCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.CreateSizeImageConstraintCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.CreateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteSizeImageConstraint(params *clientsizeimageconstraint.DeleteSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.DeleteSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.DeleteSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.DeleteSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.DeleteSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.DeleteSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.DeleteSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindSizeImageConstraint(params *clientsizeimageconstraint.FindSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.FindSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.FindSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.FindSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.FindSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.FindSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.FindSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSizeImageConstraints provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListSizeImageConstraints(params *clientsizeimageconstraint.ListSizeImageConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.ListSizeImageConstraintsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.ListSizeImageConstraintsOK
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.ListSizeImageConstraintsParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.ListSizeImageConstraintsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.ListSizeImageConstraintsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.ListSizeImageConstraintsParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
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

// TrySizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) TrySizeImageConstraint(params *clientsizeimageconstraint.TrySizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.TrySizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.TrySizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.TrySizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.TrySizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.TrySizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.TrySizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateSizeImageConstraint(params *clientsizeimageconstraint.UpdateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientsizeimageconstraint.ClientOption) (*clientsizeimageconstraint.UpdateSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientsizeimageconstraint.UpdateSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*clientsizeimageconstraint.UpdateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) *clientsizeimageconstraint.UpdateSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientsizeimageconstraint.UpdateSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientsizeimageconstraint.UpdateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...clientsizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t mockConstructorTestingTNewClientService) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
