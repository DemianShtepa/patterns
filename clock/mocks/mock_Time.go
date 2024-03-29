// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	clock "github.com/demianshtepa/patterns/clock"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockTime is an autogenerated mock type for the Time type
type MockTime struct {
	mock.Mock
}

type MockTime_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTime) EXPECT() *MockTime_Expecter {
	return &MockTime_Expecter{mock: &_m.Mock}
}

// After provides a mock function with given fields: _a0
func (_m *MockTime) After(_a0 time.Duration) <-chan time.Time {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for After")
	}

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// MockTime_After_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'After'
type MockTime_After_Call struct {
	*mock.Call
}

// After is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockTime_Expecter) After(_a0 interface{}) *MockTime_After_Call {
	return &MockTime_After_Call{Call: _e.mock.On("After", _a0)}
}

func (_c *MockTime_After_Call) Run(run func(_a0 time.Duration)) *MockTime_After_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockTime_After_Call) Return(_a0 <-chan time.Time) *MockTime_After_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTime_After_Call) RunAndReturn(run func(time.Duration) <-chan time.Time) *MockTime_After_Call {
	_c.Call.Return(run)
	return _c
}

// NewTicker provides a mock function with given fields: _a0
func (_m *MockTime) NewTicker(_a0 time.Duration) clock.Ticker {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for NewTicker")
	}

	var r0 clock.Ticker
	if rf, ok := ret.Get(0).(func(time.Duration) clock.Ticker); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clock.Ticker)
		}
	}

	return r0
}

// MockTime_NewTicker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewTicker'
type MockTime_NewTicker_Call struct {
	*mock.Call
}

// NewTicker is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockTime_Expecter) NewTicker(_a0 interface{}) *MockTime_NewTicker_Call {
	return &MockTime_NewTicker_Call{Call: _e.mock.On("NewTicker", _a0)}
}

func (_c *MockTime_NewTicker_Call) Run(run func(_a0 time.Duration)) *MockTime_NewTicker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockTime_NewTicker_Call) Return(_a0 clock.Ticker) *MockTime_NewTicker_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTime_NewTicker_Call) RunAndReturn(run func(time.Duration) clock.Ticker) *MockTime_NewTicker_Call {
	_c.Call.Return(run)
	return _c
}

// Now provides a mock function with given fields:
func (_m *MockTime) Now() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Now")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// MockTime_Now_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Now'
type MockTime_Now_Call struct {
	*mock.Call
}

// Now is a helper method to define mock.On call
func (_e *MockTime_Expecter) Now() *MockTime_Now_Call {
	return &MockTime_Now_Call{Call: _e.mock.On("Now")}
}

func (_c *MockTime_Now_Call) Run(run func()) *MockTime_Now_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTime_Now_Call) Return(_a0 time.Time) *MockTime_Now_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTime_Now_Call) RunAndReturn(run func() time.Time) *MockTime_Now_Call {
	_c.Call.Return(run)
	return _c
}

// Sleep provides a mock function with given fields: _a0
func (_m *MockTime) Sleep(_a0 time.Duration) {
	_m.Called(_a0)
}

// MockTime_Sleep_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sleep'
type MockTime_Sleep_Call struct {
	*mock.Call
}

// Sleep is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockTime_Expecter) Sleep(_a0 interface{}) *MockTime_Sleep_Call {
	return &MockTime_Sleep_Call{Call: _e.mock.On("Sleep", _a0)}
}

func (_c *MockTime_Sleep_Call) Run(run func(_a0 time.Duration)) *MockTime_Sleep_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockTime_Sleep_Call) Return() *MockTime_Sleep_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTime_Sleep_Call) RunAndReturn(run func(time.Duration)) *MockTime_Sleep_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTime creates a new instance of MockTime. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTime(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTime {
	mock := &MockTime{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
