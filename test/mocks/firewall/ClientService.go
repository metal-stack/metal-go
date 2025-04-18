// Code generated by mockery v2.53.3. DO NOT EDIT.

package firewall

import (
	clientfirewall "github.com/metal-stack/metal-go/api/client/firewall"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AllocateFirewall provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateFirewall(params *clientfirewall.AllocateFirewallParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfirewall.ClientOption) (*clientfirewall.AllocateFirewallOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AllocateFirewall")
	}

	var r0 *clientfirewall.AllocateFirewallOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfirewall.AllocateFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) (*clientfirewall.AllocateFirewallOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfirewall.AllocateFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) *clientfirewall.AllocateFirewallOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfirewall.AllocateFirewallOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfirewall.AllocateFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFirewall provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindFirewall(params *clientfirewall.FindFirewallParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfirewall.ClientOption) (*clientfirewall.FindFirewallOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindFirewall")
	}

	var r0 *clientfirewall.FindFirewallOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfirewall.FindFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) (*clientfirewall.FindFirewallOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfirewall.FindFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) *clientfirewall.FindFirewallOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfirewall.FindFirewallOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfirewall.FindFirewallParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFirewalls provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindFirewalls(params *clientfirewall.FindFirewallsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfirewall.ClientOption) (*clientfirewall.FindFirewallsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindFirewalls")
	}

	var r0 *clientfirewall.FindFirewallsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfirewall.FindFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) (*clientfirewall.FindFirewallsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfirewall.FindFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) *clientfirewall.FindFirewallsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfirewall.FindFirewallsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfirewall.FindFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFirewalls provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListFirewalls(params *clientfirewall.ListFirewallsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientfirewall.ClientOption) (*clientfirewall.ListFirewallsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFirewalls")
	}

	var r0 *clientfirewall.ListFirewallsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientfirewall.ListFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) (*clientfirewall.ListFirewallsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientfirewall.ListFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) *clientfirewall.ListFirewallsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientfirewall.ListFirewallsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientfirewall.ListFirewallsParams, runtime.ClientAuthInfoWriter, ...clientfirewall.ClientOption) error); ok {
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
