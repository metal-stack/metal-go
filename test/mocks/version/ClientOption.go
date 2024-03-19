// Code generated by mockery v2.42.0. DO NOT EDIT.

package version

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"
)

// ClientOption is an autogenerated mock type for the ClientOption type
type ClientOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ClientOption) Execute(_a0 *runtime.ClientOperation) {
	_m.Called(_a0)
}

// NewClientOption creates a new instance of ClientOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientOption {
	mock := &ClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
