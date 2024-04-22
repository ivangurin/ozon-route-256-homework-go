// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ServiceMock is an autogenerated mock type for the ServiceMock type
type ServiceMock struct {
	mock.Mock
}

type ServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceMock) EXPECT() *ServiceMock_Expecter {
	return &ServiceMock_Expecter{mock: &_m.Mock}
}

// SendMessages provides a mock function with given fields: ctx
func (_m *ServiceMock) SendMessages(ctx context.Context) {
	_m.Called(ctx)
}

// ServiceMock_SendMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessages'
type ServiceMock_SendMessages_Call struct {
	*mock.Call
}

// SendMessages is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ServiceMock_Expecter) SendMessages(ctx interface{}) *ServiceMock_SendMessages_Call {
	return &ServiceMock_SendMessages_Call{Call: _e.mock.On("SendMessages", ctx)}
}

func (_c *ServiceMock_SendMessages_Call) Run(run func(ctx context.Context)) *ServiceMock_SendMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ServiceMock_SendMessages_Call) Return() *ServiceMock_SendMessages_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServiceMock_SendMessages_Call) RunAndReturn(run func(context.Context)) *ServiceMock_SendMessages_Call {
	_c.Call.Return(run)
	return _c
}

// StopSendMessages provides a mock function with given fields:
func (_m *ServiceMock) StopSendMessages() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StopSendMessages")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceMock_StopSendMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopSendMessages'
type ServiceMock_StopSendMessages_Call struct {
	*mock.Call
}

// StopSendMessages is a helper method to define mock.On call
func (_e *ServiceMock_Expecter) StopSendMessages() *ServiceMock_StopSendMessages_Call {
	return &ServiceMock_StopSendMessages_Call{Call: _e.mock.On("StopSendMessages")}
}

func (_c *ServiceMock_StopSendMessages_Call) Run(run func()) *ServiceMock_StopSendMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceMock_StopSendMessages_Call) Return(_a0 error) *ServiceMock_StopSendMessages_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceMock_StopSendMessages_Call) RunAndReturn(run func() error) *ServiceMock_StopSendMessages_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceMock creates a new instance of ServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceMock {
	mock := &ServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}