// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlc "route256.ozon.ru/project/loms/internal/repository/kafka_storage/sqlc"
)

// RepositoryMock is an autogenerated mock type for the RepositoryMock type
type RepositoryMock struct {
	mock.Mock
}

type RepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *RepositoryMock) EXPECT() *RepositoryMock_Expecter {
	return &RepositoryMock_Expecter{mock: &_m.Mock}
}

// SendMessages provides a mock function with given fields: ctx, callback
func (_m *RepositoryMock) SendMessages(ctx context.Context, callback func(context.Context, *sqlc.KafkaOutbox) error) error {
	ret := _m.Called(ctx, callback)

	if len(ret) == 0 {
		panic("no return value specified for SendMessages")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, *sqlc.KafkaOutbox) error) error); ok {
		r0 = rf(ctx, callback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryMock_SendMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessages'
type RepositoryMock_SendMessages_Call struct {
	*mock.Call
}

// SendMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - callback func(context.Context , *sqlc.KafkaOutbox) error
func (_e *RepositoryMock_Expecter) SendMessages(ctx interface{}, callback interface{}) *RepositoryMock_SendMessages_Call {
	return &RepositoryMock_SendMessages_Call{Call: _e.mock.On("SendMessages", ctx, callback)}
}

func (_c *RepositoryMock_SendMessages_Call) Run(run func(ctx context.Context, callback func(context.Context, *sqlc.KafkaOutbox) error)) *RepositoryMock_SendMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(context.Context, *sqlc.KafkaOutbox) error))
	})
	return _c
}

func (_c *RepositoryMock_SendMessages_Call) Return(_a0 error) *RepositoryMock_SendMessages_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryMock_SendMessages_Call) RunAndReturn(run func(context.Context, func(context.Context, *sqlc.KafkaOutbox) error) error) *RepositoryMock_SendMessages_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepositoryMock creates a new instance of RepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryMock {
	mock := &RepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
