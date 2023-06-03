// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	controller "github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	gin "github.com/gin-gonic/gin"

	http "github.com/christian-gama/nutrai-api/internal/core/infra/http"

	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx
func (_m *Controller) Handle(ctx *gin.Context) {
	_m.Called(ctx)
}

// IsPublic provides a mock function with given fields:
func (_m *Controller) IsPublic() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *Controller) Method() http.Method {
	ret := _m.Called()

	var r0 http.Method
	if rf, ok := ret.Get(0).(func() http.Method); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(http.Method)
	}

	return r0
}

// Params provides a mock function with given fields:
func (_m *Controller) Params() controller.Params {
	ret := _m.Called()

	var r0 controller.Params
	if rf, ok := ret.Get(0).(func() controller.Params); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(controller.Params)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *Controller) Path() controller.Path {
	ret := _m.Called()

	var r0 controller.Path
	if rf, ok := ret.Get(0).(func() controller.Path); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(controller.Path)
	}

	return r0
}

// RPM provides a mock function with given fields:
func (_m *Controller) RPM() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewController interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t mockConstructorTestingTNewController) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
