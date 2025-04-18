// Code generated by mockery v2.53.3. DO NOT EDIT.

package partition

import (
	clientpartition "github.com/metal-stack/metal-go/api/client/partition"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreatePartition provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePartition(params *clientpartition.CreatePartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.CreatePartitionCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePartition")
	}

	var r0 *clientpartition.CreatePartitionCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.CreatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.CreatePartitionCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.CreatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.CreatePartitionCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.CreatePartitionCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.CreatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePartition provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeletePartition(params *clientpartition.DeletePartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.DeletePartitionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePartition")
	}

	var r0 *clientpartition.DeletePartitionOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.DeletePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.DeletePartitionOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.DeletePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.DeletePartitionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.DeletePartitionOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.DeletePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPartition provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindPartition(params *clientpartition.FindPartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.FindPartitionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindPartition")
	}

	var r0 *clientpartition.FindPartitionOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.FindPartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.FindPartitionOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.FindPartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.FindPartitionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.FindPartitionOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.FindPartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPartitions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListPartitions(params *clientpartition.ListPartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.ListPartitionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPartitions")
	}

	var r0 *clientpartition.ListPartitionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.ListPartitionsParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.ListPartitionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.ListPartitionsParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.ListPartitionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.ListPartitionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.ListPartitionsParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PartitionCapacity provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PartitionCapacity(params *clientpartition.PartitionCapacityParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.PartitionCapacityOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PartitionCapacity")
	}

	var r0 *clientpartition.PartitionCapacityOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.PartitionCapacityParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.PartitionCapacityOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.PartitionCapacityParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.PartitionCapacityOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.PartitionCapacityOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.PartitionCapacityParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
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

// UpdatePartition provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdatePartition(params *clientpartition.UpdatePartitionParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientpartition.ClientOption) (*clientpartition.UpdatePartitionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePartition")
	}

	var r0 *clientpartition.UpdatePartitionOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientpartition.UpdatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) (*clientpartition.UpdatePartitionOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientpartition.UpdatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) *clientpartition.UpdatePartitionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientpartition.UpdatePartitionOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientpartition.UpdatePartitionParams, runtime.ClientAuthInfoWriter, ...clientpartition.ClientOption) error); ok {
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
