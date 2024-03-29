// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockFunction is an autogenerated mock type for the Function type
type MockFunction struct {
	mock.Mock
}

type MockFunction_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFunction) EXPECT() *MockFunction_Expecter {
	return &MockFunction_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockFunction) Execute(_a0 context.Context) (interface{}, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (interface{}, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunction_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockFunction_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockFunction_Expecter) Execute(_a0 interface{}) *MockFunction_Execute_Call {
	return &MockFunction_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockFunction_Execute_Call) Run(run func(_a0 context.Context)) *MockFunction_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockFunction_Execute_Call) Return(_a0 interface{}, _a1 error) *MockFunction_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunction_Execute_Call) RunAndReturn(run func(context.Context) (interface{}, error)) *MockFunction_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFunction creates a new instance of MockFunction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFunction(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFunction {
	mock := &MockFunction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
