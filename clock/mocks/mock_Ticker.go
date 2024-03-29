// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockTicker is an autogenerated mock type for the Ticker type
type MockTicker struct {
	mock.Mock
}

type MockTicker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTicker) EXPECT() *MockTicker_Expecter {
	return &MockTicker_Expecter{mock: &_m.Mock}
}

// Chan provides a mock function with given fields:
func (_m *MockTicker) Chan() <-chan time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Chan")
	}

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func() <-chan time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// MockTicker_Chan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chan'
type MockTicker_Chan_Call struct {
	*mock.Call
}

// Chan is a helper method to define mock.On call
func (_e *MockTicker_Expecter) Chan() *MockTicker_Chan_Call {
	return &MockTicker_Chan_Call{Call: _e.mock.On("Chan")}
}

func (_c *MockTicker_Chan_Call) Run(run func()) *MockTicker_Chan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTicker_Chan_Call) Return(_a0 <-chan time.Time) *MockTicker_Chan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTicker_Chan_Call) RunAndReturn(run func() <-chan time.Time) *MockTicker_Chan_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields: _a0
func (_m *MockTicker) Reset(_a0 time.Duration) {
	_m.Called(_a0)
}

// MockTicker_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type MockTicker_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockTicker_Expecter) Reset(_a0 interface{}) *MockTicker_Reset_Call {
	return &MockTicker_Reset_Call{Call: _e.mock.On("Reset", _a0)}
}

func (_c *MockTicker_Reset_Call) Run(run func(_a0 time.Duration)) *MockTicker_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockTicker_Reset_Call) Return() *MockTicker_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTicker_Reset_Call) RunAndReturn(run func(time.Duration)) *MockTicker_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockTicker) Stop() {
	_m.Called()
}

// MockTicker_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockTicker_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockTicker_Expecter) Stop() *MockTicker_Stop_Call {
	return &MockTicker_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockTicker_Stop_Call) Run(run func()) *MockTicker_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTicker_Stop_Call) Return() *MockTicker_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTicker_Stop_Call) RunAndReturn(run func()) *MockTicker_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTicker creates a new instance of MockTicker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTicker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTicker {
	mock := &MockTicker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
