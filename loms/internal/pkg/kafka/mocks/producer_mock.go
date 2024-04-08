// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ProducerMock is an autogenerated mock type for the ProducerMock type
type ProducerMock struct {
	mock.Mock
}

type ProducerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ProducerMock) EXPECT() *ProducerMock_Expecter {
	return &ProducerMock_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *ProducerMock) Close() error {
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

// ProducerMock_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type ProducerMock_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *ProducerMock_Expecter) Close() *ProducerMock_Close_Call {
	return &ProducerMock_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *ProducerMock_Close_Call) Run(run func()) *ProducerMock_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ProducerMock_Close_Call) Return(_a0 error) *ProducerMock_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProducerMock_Close_Call) RunAndReturn(run func() error) *ProducerMock_Close_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessageWithKey provides a mock function with given fields: ctx, topic, key, message
func (_m *ProducerMock) SendMessageWithKey(ctx context.Context, topic string, key string, message interface{}) error {
	ret := _m.Called(ctx, topic, key, message)

	if len(ret) == 0 {
		panic("no return value specified for SendMessageWithKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, topic, key, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProducerMock_SendMessageWithKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessageWithKey'
type ProducerMock_SendMessageWithKey_Call struct {
	*mock.Call
}

// SendMessageWithKey is a helper method to define mock.On call
//   - ctx context.Context
//   - topic string
//   - key string
//   - message interface{}
func (_e *ProducerMock_Expecter) SendMessageWithKey(ctx interface{}, topic interface{}, key interface{}, message interface{}) *ProducerMock_SendMessageWithKey_Call {
	return &ProducerMock_SendMessageWithKey_Call{Call: _e.mock.On("SendMessageWithKey", ctx, topic, key, message)}
}

func (_c *ProducerMock_SendMessageWithKey_Call) Run(run func(ctx context.Context, topic string, key string, message interface{})) *ProducerMock_SendMessageWithKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(interface{}))
	})
	return _c
}

func (_c *ProducerMock_SendMessageWithKey_Call) Return(_a0 error) *ProducerMock_SendMessageWithKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProducerMock_SendMessageWithKey_Call) RunAndReturn(run func(context.Context, string, string, interface{}) error) *ProducerMock_SendMessageWithKey_Call {
	_c.Call.Return(run)
	return _c
}

// NewProducerMock creates a new instance of ProducerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProducerMock {
	mock := &ProducerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}