// Code generated by mockery v2.53.3. DO NOT EDIT.

package network

import (
	clientnetwork "github.com/metal-stack/metal-go/api/client/network"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AllocateNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateNetwork(params *clientnetwork.AllocateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.AllocateNetworkCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AllocateNetwork")
	}

	var r0 *clientnetwork.AllocateNetworkCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.AllocateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.AllocateNetworkCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.AllocateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.AllocateNetworkCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.AllocateNetworkCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.AllocateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateNetwork(params *clientnetwork.CreateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.CreateNetworkCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateNetwork")
	}

	var r0 *clientnetwork.CreateNetworkCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.CreateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.CreateNetworkCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.CreateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.CreateNetworkCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.CreateNetworkCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.CreateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteNetwork(params *clientnetwork.DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.DeleteNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNetwork")
	}

	var r0 *clientnetwork.DeleteNetworkOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.DeleteNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.DeleteNetworkOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.DeleteNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.DeleteNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.DeleteNetworkOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.DeleteNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindNetwork(params *clientnetwork.FindNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.FindNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindNetwork")
	}

	var r0 *clientnetwork.FindNetworkOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.FindNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.FindNetworkOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.FindNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.FindNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.FindNetworkOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.FindNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNetworks provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindNetworks(params *clientnetwork.FindNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.FindNetworksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindNetworks")
	}

	var r0 *clientnetwork.FindNetworksOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.FindNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.FindNetworksOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.FindNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.FindNetworksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.FindNetworksOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.FindNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeNetwork(params *clientnetwork.FreeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.FreeNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FreeNetwork")
	}

	var r0 *clientnetwork.FreeNetworkOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.FreeNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.FreeNetworkOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.FreeNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.FreeNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.FreeNetworkOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.FreeNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeNetworkDeprecated provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeNetworkDeprecated(params *clientnetwork.FreeNetworkDeprecatedParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.FreeNetworkDeprecatedOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FreeNetworkDeprecated")
	}

	var r0 *clientnetwork.FreeNetworkDeprecatedOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.FreeNetworkDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.FreeNetworkDeprecatedOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.FreeNetworkDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.FreeNetworkDeprecatedOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.FreeNetworkDeprecatedOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.FreeNetworkDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListNetworks(params *clientnetwork.ListNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.ListNetworksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListNetworks")
	}

	var r0 *clientnetwork.ListNetworksOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.ListNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.ListNetworksOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.ListNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.ListNetworksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.ListNetworksOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.ListNetworksParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
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

// UpdateNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateNetwork(params *clientnetwork.UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientnetwork.ClientOption) (*clientnetwork.UpdateNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNetwork")
	}

	var r0 *clientnetwork.UpdateNetworkOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientnetwork.UpdateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) (*clientnetwork.UpdateNetworkOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientnetwork.UpdateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) *clientnetwork.UpdateNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientnetwork.UpdateNetworkOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientnetwork.UpdateNetworkParams, runtime.ClientAuthInfoWriter, ...clientnetwork.ClientOption) error); ok {
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
