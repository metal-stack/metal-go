// Code generated by mockery (devel). DO NOT EDIT.

package firewall

import (
	firewall "github.com/metal-stack/metal-go/api/client/firewall"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AllocateFirewall provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateFirewall(params *firewall.AllocateFirewallParams, authInfo runtime.ClientAuthInfoWriter, opts ...firewall.ClientOption) (*firewall.AllocateFirewallOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firewall.AllocateFirewallOK
	if rf, ok := ret.Get(0).(func(*firewall.AllocateFirewallParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) *firewall.AllocateFirewallOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firewall.AllocateFirewallOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firewall.AllocateFirewallParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFirewall provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindFirewall(params *firewall.FindFirewallParams, authInfo runtime.ClientAuthInfoWriter, opts ...firewall.ClientOption) (*firewall.FindFirewallOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firewall.FindFirewallOK
	if rf, ok := ret.Get(0).(func(*firewall.FindFirewallParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) *firewall.FindFirewallOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firewall.FindFirewallOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firewall.FindFirewallParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFirewalls provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindFirewalls(params *firewall.FindFirewallsParams, authInfo runtime.ClientAuthInfoWriter, opts ...firewall.ClientOption) (*firewall.FindFirewallsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firewall.FindFirewallsOK
	if rf, ok := ret.Get(0).(func(*firewall.FindFirewallsParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) *firewall.FindFirewallsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firewall.FindFirewallsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firewall.FindFirewallsParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFirewalls provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListFirewalls(params *firewall.ListFirewallsParams, authInfo runtime.ClientAuthInfoWriter, opts ...firewall.ClientOption) (*firewall.ListFirewallsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *firewall.ListFirewallsOK
	if rf, ok := ret.Get(0).(func(*firewall.ListFirewallsParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) *firewall.ListFirewallsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firewall.ListFirewallsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*firewall.ListFirewallsParams, runtime.ClientAuthInfoWriter, ...firewall.ClientOption) error); ok {
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
