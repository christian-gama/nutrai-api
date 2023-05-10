// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/christian-gama/nutrai-api/internal/user/app/service"
	mock "github.com/stretchr/testify/mock"
)

// HashPasswordHandler is an autogenerated mock type for the HashPasswordHandler type
type HashPasswordHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, input
func (_m *HashPasswordHandler) Handle(ctx context.Context, input *service.HashPasswordInput) (*service.HashPasswordOutput, error) {
	ret := _m.Called(ctx, input)

	var r0 *service.HashPasswordOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *service.HashPasswordInput) (*service.HashPasswordOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *service.HashPasswordInput) *service.HashPasswordOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.HashPasswordOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *service.HashPasswordInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHashPasswordHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHashPasswordHandler creates a new instance of HashPasswordHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHashPasswordHandler(t mockConstructorTestingTNewHashPasswordHandler) *HashPasswordHandler {
	mock := &HashPasswordHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
