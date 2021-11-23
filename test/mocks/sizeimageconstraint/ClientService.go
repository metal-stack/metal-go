// Code generated by mockery v2.7.4. DO NOT EDIT.

package sizeimageconstraint

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	sizeimageconstraint "github.com/metal-stack/metal-go/api/client/sizeimageconstraint"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateSizeImageConstraint(params *sizeimageconstraint.CreateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...sizeimageconstraint.ClientOption) (*sizeimageconstraint.CreateSizeImageConstraintCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sizeimageconstraint.CreateSizeImageConstraintCreated
	if rf, ok := ret.Get(0).(func(*sizeimageconstraint.CreateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) *sizeimageconstraint.CreateSizeImageConstraintCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sizeimageconstraint.CreateSizeImageConstraintCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sizeimageconstraint.CreateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteSizeImageConstraint(params *sizeimageconstraint.DeleteSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...sizeimageconstraint.ClientOption) (*sizeimageconstraint.DeleteSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sizeimageconstraint.DeleteSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*sizeimageconstraint.DeleteSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) *sizeimageconstraint.DeleteSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sizeimageconstraint.DeleteSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sizeimageconstraint.DeleteSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindSizeImageConstraint(params *sizeimageconstraint.FindSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...sizeimageconstraint.ClientOption) (*sizeimageconstraint.FindSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sizeimageconstraint.FindSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*sizeimageconstraint.FindSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) *sizeimageconstraint.FindSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sizeimageconstraint.FindSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sizeimageconstraint.FindSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSizeImageConstraints provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListSizeImageConstraints(params *sizeimageconstraint.ListSizeImageConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...sizeimageconstraint.ClientOption) (*sizeimageconstraint.ListSizeImageConstraintsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sizeimageconstraint.ListSizeImageConstraintsOK
	if rf, ok := ret.Get(0).(func(*sizeimageconstraint.ListSizeImageConstraintsParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) *sizeimageconstraint.ListSizeImageConstraintsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sizeimageconstraint.ListSizeImageConstraintsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sizeimageconstraint.ListSizeImageConstraintsParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) error); ok {
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

// UpdateSizeImageConstraint provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateSizeImageConstraint(params *sizeimageconstraint.UpdateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...sizeimageconstraint.ClientOption) (*sizeimageconstraint.UpdateSizeImageConstraintOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sizeimageconstraint.UpdateSizeImageConstraintOK
	if rf, ok := ret.Get(0).(func(*sizeimageconstraint.UpdateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) *sizeimageconstraint.UpdateSizeImageConstraintOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sizeimageconstraint.UpdateSizeImageConstraintOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sizeimageconstraint.UpdateSizeImageConstraintParams, runtime.ClientAuthInfoWriter, ...sizeimageconstraint.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
