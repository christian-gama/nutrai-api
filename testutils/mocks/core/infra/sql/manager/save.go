// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	manager "github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	mock "github.com/stretchr/testify/mock"
)

// Save is an autogenerated mock type for the Save type
type Save[Model interface{}] struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, input
func (_m *Save[Model]) Save(ctx context.Context, input manager.SaveInput[Model]) (*Model, error) {
	ret := _m.Called(ctx, input)

	var r0 *Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.SaveInput[Model]) (*Model, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, manager.SaveInput[Model]) *Model); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, manager.SaveInput[Model]) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSave interface {
	mock.TestingT
	Cleanup(func())
}

// NewSave creates a new instance of Save. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSave[Model interface{}](t mockConstructorTestingTNewSave) *Save[Model] {
	mock := &Save[Model]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
