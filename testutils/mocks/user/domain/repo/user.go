// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	querying "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	mock "github.com/stretchr/testify/mock"

	repo "github.com/christian-gama/nutrai-api/internal/user/domain/repo"

	user "github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx, input, preload
func (_m *User) All(ctx context.Context, input repo.AllUsersInput, preload ...string) (*querying.PaginationOutput[*user.User], error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *querying.PaginationOutput[*user.User]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.AllUsersInput, ...string) (*querying.PaginationOutput[*user.User], error)); ok {
		return rf(ctx, input, preload...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repo.AllUsersInput, ...string) *querying.PaginationOutput[*user.User]); ok {
		r0 = rf(ctx, input, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*querying.PaginationOutput[*user.User])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repo.AllUsersInput, ...string) error); ok {
		r1 = rf(ctx, input, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, input
func (_m *User) Delete(ctx context.Context, input repo.DeleteUserInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.DeleteUserInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, input, preload
func (_m *User) Find(ctx context.Context, input repo.FindUserInput, preload ...string) (*user.User, error) {
	_va := make([]interface{}, len(preload))
	for _i := range preload {
		_va[_i] = preload[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindUserInput, ...string) (*user.User, error)); ok {
		return rf(ctx, input, preload...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repo.FindUserInput, ...string) *user.User); ok {
		r0 = rf(ctx, input, preload...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repo.FindUserInput, ...string) error); ok {
		r1 = rf(ctx, input, preload...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, input
func (_m *User) Save(ctx context.Context, input repo.SaveUserInput) (*user.User, error) {
	ret := _m.Called(ctx, input)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.SaveUserInput) (*user.User, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repo.SaveUserInput) *user.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repo.SaveUserInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, input
func (_m *User) Update(ctx context.Context, input repo.UpdateUserInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.UpdateUserInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUser(t mockConstructorTestingTNewUser) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
