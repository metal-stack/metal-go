// Code generated by mockery (devel). DO NOT EDIT.

package network

import (
	runtime "github.com/go-openapi/runtime"
	network "github.com/metal-stack/metal-go/api/client/network"
	mock "github.com/stretchr/testify/mock"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AllocateNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateNetwork(params *network.AllocateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.AllocateNetworkCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.AllocateNetworkCreated
	if rf, ok := ret.Get(0).(func(*network.AllocateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.AllocateNetworkCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.AllocateNetworkCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.AllocateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateNetwork(params *network.CreateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.CreateNetworkCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.CreateNetworkCreated
	if rf, ok := ret.Get(0).(func(*network.CreateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.CreateNetworkCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.CreateNetworkCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.CreateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteNetwork(params *network.DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.DeleteNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.DeleteNetworkOK
	if rf, ok := ret.Get(0).(func(*network.DeleteNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.DeleteNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.DeleteNetworkOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.DeleteNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindNetwork(params *network.FindNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.FindNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.FindNetworkOK
	if rf, ok := ret.Get(0).(func(*network.FindNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.FindNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.FindNetworkOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.FindNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNetworks provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindNetworks(params *network.FindNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.FindNetworksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.FindNetworksOK
	if rf, ok := ret.Get(0).(func(*network.FindNetworksParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.FindNetworksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.FindNetworksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.FindNetworksParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeNetwork provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeNetwork(params *network.FreeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.FreeNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.FreeNetworkOK
	if rf, ok := ret.Get(0).(func(*network.FreeNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.FreeNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.FreeNetworkOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.FreeNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListNetworks(params *network.ListNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.ListNetworksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.ListNetworksOK
	if rf, ok := ret.Get(0).(func(*network.ListNetworksParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.ListNetworksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.ListNetworksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.ListNetworksParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
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
func (_m *ClientService) UpdateNetwork(params *network.UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...network.ClientOption) (*network.UpdateNetworkOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *network.UpdateNetworkOK
	if rf, ok := ret.Get(0).(func(*network.UpdateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) *network.UpdateNetworkOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*network.UpdateNetworkOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*network.UpdateNetworkParams, runtime.ClientAuthInfoWriter, ...network.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
