// Code generated by mockery v2.53.3. DO NOT EDIT.

package image

import (
	clientimage "github.com/metal-stack/metal-go/api/client/image"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateImage provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateImage(params *clientimage.CreateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.CreateImageCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateImage")
	}

	var r0 *clientimage.CreateImageCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.CreateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.CreateImageCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.CreateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.CreateImageCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.CreateImageCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.CreateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImage provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteImage(params *clientimage.DeleteImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.DeleteImageOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteImage")
	}

	var r0 *clientimage.DeleteImageOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.DeleteImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.DeleteImageOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.DeleteImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.DeleteImageOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.DeleteImageOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.DeleteImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindImage provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindImage(params *clientimage.FindImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.FindImageOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindImage")
	}

	var r0 *clientimage.FindImageOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.FindImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.FindImageOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.FindImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.FindImageOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.FindImageOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.FindImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindImages provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindImages(params *clientimage.FindImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.FindImagesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindImages")
	}

	var r0 *clientimage.FindImagesOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.FindImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.FindImagesOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.FindImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.FindImagesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.FindImagesOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.FindImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLatestImage provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindLatestImage(params *clientimage.FindLatestImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.FindLatestImageOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindLatestImage")
	}

	var r0 *clientimage.FindLatestImageOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.FindLatestImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.FindLatestImageOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.FindLatestImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.FindLatestImageOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.FindLatestImageOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.FindLatestImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImages provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListImages(params *clientimage.ListImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.ListImagesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListImages")
	}

	var r0 *clientimage.ListImagesOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.ListImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.ListImagesOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.ListImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.ListImagesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.ListImagesOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.ListImagesParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryImagesByID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) QueryImagesByID(params *clientimage.QueryImagesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.QueryImagesByIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryImagesByID")
	}

	var r0 *clientimage.QueryImagesByIDOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.QueryImagesByIDParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.QueryImagesByIDOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.QueryImagesByIDParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.QueryImagesByIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.QueryImagesByIDOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.QueryImagesByIDParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
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

// UpdateImage provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateImage(params *clientimage.UpdateImageParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientimage.ClientOption) (*clientimage.UpdateImageOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateImage")
	}

	var r0 *clientimage.UpdateImageOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientimage.UpdateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) (*clientimage.UpdateImageOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientimage.UpdateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) *clientimage.UpdateImageOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientimage.UpdateImageOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientimage.UpdateImageParams, runtime.ClientAuthInfoWriter, ...clientimage.ClientOption) error); ok {
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
