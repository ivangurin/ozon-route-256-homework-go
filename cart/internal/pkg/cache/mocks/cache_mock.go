// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// CacheMock is an autogenerated mock type for the CacheMock type
type CacheMock struct {
	mock.Mock
}

type CacheMock_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheMock) EXPECT() *CacheMock_Expecter {
	return &CacheMock_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *CacheMock) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheMock_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type CacheMock_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *CacheMock_Expecter) Close() *CacheMock_Close_Call {
	return &CacheMock_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *CacheMock_Close_Call) Run(run func()) *CacheMock_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CacheMock_Close_Call) Return(_a0 error) *CacheMock_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheMock_Close_Call) RunAndReturn(run func() error) *CacheMock_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *CacheMock) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (bool, error)); ok {
		return rf(ctx, key, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheMock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CacheMock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
func (_e *CacheMock_Expecter) Get(ctx interface{}, key interface{}, value interface{}) *CacheMock_Get_Call {
	return &CacheMock_Get_Call{Call: _e.mock.On("Get", ctx, key, value)}
}

func (_c *CacheMock_Get_Call) Run(run func(ctx context.Context, key string, value interface{})) *CacheMock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *CacheMock_Get_Call) Return(_a0 bool, _a1 error) *CacheMock_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheMock_Get_Call) RunAndReturn(run func(context.Context, string, interface{}) (bool, error)) *CacheMock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, ttl
func (_m *CacheMock) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheMock_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type CacheMock_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - ttl time.Duration
func (_e *CacheMock_Expecter) Set(ctx interface{}, key interface{}, value interface{}, ttl interface{}) *CacheMock_Set_Call {
	return &CacheMock_Set_Call{Call: _e.mock.On("Set", ctx, key, value, ttl)}
}

func (_c *CacheMock_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, ttl time.Duration)) *CacheMock_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *CacheMock_Set_Call) Return(_a0 error) *CacheMock_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheMock_Set_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) error) *CacheMock_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewCacheMock creates a new instance of CacheMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCacheMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *CacheMock {
	mock := &CacheMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
