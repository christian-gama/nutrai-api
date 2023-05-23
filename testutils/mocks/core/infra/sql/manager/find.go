// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	manager "github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	mock "github.com/stretchr/testify/mock"
)

// Find is an autogenerated mock type for the Find type
type Find[Model interface{}] struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, input
func (_m *Find[Model]) Find(ctx context.Context, input manager.FindInput[Model]) (*Model, error) {
	ret := _m.Called(ctx, input)

	var r0 *Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.FindInput[Model]) (*Model, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, manager.FindInput[Model]) *Model); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, manager.FindInput[Model]) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFind interface {
	mock.TestingT
	Cleanup(func())
}

// NewFind creates a new instance of Find. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFind[Model interface{}](t mockConstructorTestingTNewFind) *Find[Model] {
	mock := &Find[Model]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
