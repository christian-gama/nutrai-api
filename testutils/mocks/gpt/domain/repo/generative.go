// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gpt "github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	mock "github.com/stretchr/testify/mock"

	repo "github.com/christian-gama/nutrai-api/internal/gpt/domain/repo"
)

// Generative is an autogenerated mock type for the Generative type
type Generative struct {
	mock.Mock
}

// ChatCompletion provides a mock function with given fields: ctx, input
func (_m *Generative) ChatCompletion(ctx context.Context, input *repo.ChatCompletionInput) (*gpt.Message, error) {
	ret := _m.Called(ctx, input)

	var r0 *gpt.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *repo.ChatCompletionInput) (*gpt.Message, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *repo.ChatCompletionInput) *gpt.Message); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gpt.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *repo.ChatCompletionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGenerative interface {
	mock.TestingT
	Cleanup(func())
}

// NewGenerative creates a new instance of Generative. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGenerative(t mockConstructorTestingTNewGenerative) *Generative {
	mock := &Generative{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}