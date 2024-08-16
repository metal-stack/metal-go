// Code generated by mockery v2.45.0. DO NOT EDIT.

package filesystemlayout

import (
	clientfilesystemlayout "github.com/metal-stack/metal-go/api/client/filesystemlayout"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateFilesystemLayout(params *clientfilesystemlayout.CreateFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.CreateFilesystemLayoutCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.CreateFilesystemLayoutCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.CreateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.CreateFilesystemLayoutCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.CreateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.CreateFilesystemLayoutCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.CreateFilesystemLayoutCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.CreateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteFilesystemLayout(params *clientfilesystemlayout.DeleteFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.DeleteFilesystemLayoutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.DeleteFilesystemLayoutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.DeleteFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.DeleteFilesystemLayoutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.DeleteFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.DeleteFilesystemLayoutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.DeleteFilesystemLayoutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.DeleteFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetFilesystemLayout(params *clientfilesystemlayout.GetFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.GetFilesystemLayoutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.GetFilesystemLayoutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.GetFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.GetFilesystemLayoutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.GetFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.GetFilesystemLayoutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.GetFilesystemLayoutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.GetFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFilesystemLayouts provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListFilesystemLayouts(params *clientfilesystemlayout.ListFilesystemLayoutsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.ListFilesystemLayoutsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFilesystemLayouts")
	}

	var r0 *clientfilesystemlayout.ListFilesystemLayoutsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.ListFilesystemLayoutsParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.ListFilesystemLayoutsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.ListFilesystemLayoutsParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.ListFilesystemLayoutsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.ListFilesystemLayoutsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.ListFilesystemLayoutsParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MatchFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) MatchFilesystemLayout(params *clientfilesystemlayout.MatchFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.MatchFilesystemLayoutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MatchFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.MatchFilesystemLayoutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.MatchFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.MatchFilesystemLayoutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.MatchFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.MatchFilesystemLayoutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.MatchFilesystemLayoutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.MatchFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
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

// TryFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) TryFilesystemLayout(params *clientfilesystemlayout.TryFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.TryFilesystemLayoutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TryFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.TryFilesystemLayoutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.TryFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.TryFilesystemLayoutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.TryFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.TryFilesystemLayoutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.TryFilesystemLayoutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.TryFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFilesystemLayout provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateFilesystemLayout(params *clientfilesystemlayout.UpdateFilesystemLayoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.UpdateFilesystemLayoutOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFilesystemLayout")
	}

	var r0 *clientfilesystemlayout.UpdateFilesystemLayoutOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.UpdateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) (*clientfilesystemlayout.UpdateFilesystemLayoutOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfilesystemlayout.UpdateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) *clientfilesystemlayout.UpdateFilesystemLayoutOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfilesystemlayout.UpdateFilesystemLayoutOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfilesystemlayout.UpdateFilesystemLayoutParams, runtime.ClientAuthInfoWriter, ...clientfilesystemlayout.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
