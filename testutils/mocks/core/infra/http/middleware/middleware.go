// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx
func (_m *Middleware) Handle(ctx *gin.Context) {
	_m.Called(ctx)
}

type mockConstructorTestingTNewMiddleware interface {
	mock.TestingT
	Cleanup(func())
}

// NewMiddleware creates a new instance of Middleware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMiddleware(t mockConstructorTestingTNewMiddleware) *Middleware {
	mock := &Middleware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
