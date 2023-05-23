// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	manager "github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	mock "github.com/stretchr/testify/mock"

	queryer "github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
)

// Repository is an autogenerated mock type for the Repository type
type Repository[Model interface{}] struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx, input
func (_m *Repository[Model]) All(ctx context.Context, input manager.AllInput[Model]) (*queryer.PaginationOutput[*Model], error) {
	ret := _m.Called(ctx, input)

	var r0 *queryer.PaginationOutput[*Model]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.AllInput[Model]) (*queryer.PaginationOutput[*Model], error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, manager.AllInput[Model]) *queryer.PaginationOutput[*Model]); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*queryer.PaginationOutput[*Model])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, manager.AllInput[Model]) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Repository[Model]) Delete(ctx context.Context, input manager.DeleteInput[Model]) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.DeleteInput[Model]) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, input
func (_m *Repository[Model]) Find(ctx context.Context, input manager.FindInput[Model]) (*Model, error) {
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

// Save provides a mock function with given fields: ctx, input
func (_m *Repository[Model]) Save(ctx context.Context, input manager.SaveInput[Model]) (*Model, error) {
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

// Update provides a mock function with given fields: ctx, input
func (_m *Repository[Model]) Update(ctx context.Context, input manager.UpdateInput[Model]) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, manager.UpdateInput[Model]) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository[Model interface{}](t mockConstructorTestingTNewRepository) *Repository[Model] {
	mock := &Repository[Model]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
