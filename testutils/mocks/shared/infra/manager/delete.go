// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	manager "github.com/christian-gama/nutrai-api/internal/shared/infra/manager"
	mock "github.com/stretchr/testify/mock"
)

// Delete is an autogenerated mock type for the Delete type
type Delete[Model interface{}] struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Delete[Model]) Delete(ctx context.Context, input manager.DeleteInput[Model]) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.DeleteInput[Model]) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDelete interface {
	mock.TestingT
	Cleanup(func())
}

// NewDelete creates a new instance of Delete. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDelete[Model interface{}](t mockConstructorTestingTNewDelete) *Delete[Model] {
	mock := &Delete[Model]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
