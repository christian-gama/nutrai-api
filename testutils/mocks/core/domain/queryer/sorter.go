// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	queryer "github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	mock "github.com/stretchr/testify/mock"
)

// Sorter is an autogenerated mock type for the Sorter type
type Sorter struct {
	mock.Mock
}

// Add provides a mock function with given fields: field, isDesc
func (_m *Sorter) Add(field string, isDesc bool) queryer.Sorter {
	ret := _m.Called(field, isDesc)

	var r0 queryer.Sorter
	if rf, ok := ret.Get(0).(func(string, bool) queryer.Sorter); ok {
		r0 = rf(field, isDesc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queryer.Sorter)
		}
	}

	return r0
}

// Field provides a mock function with given fields: idx
func (_m *Sorter) Field(idx int) string {
	ret := _m.Called(idx)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsDesc provides a mock function with given fields: idx
func (_m *Sorter) IsDesc(idx int) bool {
	ret := _m.Called(idx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Slice provides a mock function with given fields:
func (_m *Sorter) Slice() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

type mockConstructorTestingTNewSorter interface {
	mock.TestingT
	Cleanup(func())
}

// NewSorter creates a new instance of Sorter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSorter(t mockConstructorTestingTNewSorter) *Sorter {
	mock := &Sorter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
