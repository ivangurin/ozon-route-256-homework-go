// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
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

// CancelReserve provides a mock function with given fields: ctx, items
func (_m *RepositoryMock) CancelReserve(ctx context.Context, items stockstorage.ReserveItems) error {
	ret := _m.Called(ctx, items)

	if len(ret) == 0 {
		panic("no return value specified for CancelReserve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stockstorage.ReserveItems) error); ok {
		r0 = rf(ctx, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryMock_CancelReserve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelReserve'
type RepositoryMock_CancelReserve_Call struct {
	*mock.Call
}

// CancelReserve is a helper method to define mock.On call
//   - ctx context.Context
//   - items stockstorage.ReserveItems
func (_e *RepositoryMock_Expecter) CancelReserve(ctx interface{}, items interface{}) *RepositoryMock_CancelReserve_Call {
	return &RepositoryMock_CancelReserve_Call{Call: _e.mock.On("CancelReserve", ctx, items)}
}

func (_c *RepositoryMock_CancelReserve_Call) Run(run func(ctx context.Context, items stockstorage.ReserveItems)) *RepositoryMock_CancelReserve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(stockstorage.ReserveItems))
	})
	return _c
}

func (_c *RepositoryMock_CancelReserve_Call) Return(_a0 error) *RepositoryMock_CancelReserve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryMock_CancelReserve_Call) RunAndReturn(run func(context.Context, stockstorage.ReserveItems) error) *RepositoryMock_CancelReserve_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySku provides a mock function with given fields: ctx, sku
func (_m *RepositoryMock) GetBySku(ctx context.Context, sku int64) (uint16, error) {
	ret := _m.Called(ctx, sku)

	if len(ret) == 0 {
		panic("no return value specified for GetBySku")
	}

	var r0 uint16
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (uint16, error)); ok {
		return rf(ctx, sku)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) uint16); ok {
		r0 = rf(ctx, sku)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, sku)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryMock_GetBySku_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySku'
type RepositoryMock_GetBySku_Call struct {
	*mock.Call
}

// GetBySku is a helper method to define mock.On call
//   - ctx context.Context
//   - sku int64
func (_e *RepositoryMock_Expecter) GetBySku(ctx interface{}, sku interface{}) *RepositoryMock_GetBySku_Call {
	return &RepositoryMock_GetBySku_Call{Call: _e.mock.On("GetBySku", ctx, sku)}
}

func (_c *RepositoryMock_GetBySku_Call) Run(run func(ctx context.Context, sku int64)) *RepositoryMock_GetBySku_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *RepositoryMock_GetBySku_Call) Return(_a0 uint16, _a1 error) *RepositoryMock_GetBySku_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryMock_GetBySku_Call) RunAndReturn(run func(context.Context, int64) (uint16, error)) *RepositoryMock_GetBySku_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveReserve provides a mock function with given fields: ctx, items
func (_m *RepositoryMock) RemoveReserve(ctx context.Context, items stockstorage.ReserveItems) error {
	ret := _m.Called(ctx, items)

	if len(ret) == 0 {
		panic("no return value specified for RemoveReserve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stockstorage.ReserveItems) error); ok {
		r0 = rf(ctx, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryMock_RemoveReserve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveReserve'
type RepositoryMock_RemoveReserve_Call struct {
	*mock.Call
}

// RemoveReserve is a helper method to define mock.On call
//   - ctx context.Context
//   - items stockstorage.ReserveItems
func (_e *RepositoryMock_Expecter) RemoveReserve(ctx interface{}, items interface{}) *RepositoryMock_RemoveReserve_Call {
	return &RepositoryMock_RemoveReserve_Call{Call: _e.mock.On("RemoveReserve", ctx, items)}
}

func (_c *RepositoryMock_RemoveReserve_Call) Run(run func(ctx context.Context, items stockstorage.ReserveItems)) *RepositoryMock_RemoveReserve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(stockstorage.ReserveItems))
	})
	return _c
}

func (_c *RepositoryMock_RemoveReserve_Call) Return(_a0 error) *RepositoryMock_RemoveReserve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryMock_RemoveReserve_Call) RunAndReturn(run func(context.Context, stockstorage.ReserveItems) error) *RepositoryMock_RemoveReserve_Call {
	_c.Call.Return(run)
	return _c
}

// Reserve provides a mock function with given fields: ctx, items
func (_m *RepositoryMock) Reserve(ctx context.Context, items stockstorage.ReserveItems) error {
	ret := _m.Called(ctx, items)

	if len(ret) == 0 {
		panic("no return value specified for Reserve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stockstorage.ReserveItems) error); ok {
		r0 = rf(ctx, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryMock_Reserve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reserve'
type RepositoryMock_Reserve_Call struct {
	*mock.Call
}

// Reserve is a helper method to define mock.On call
//   - ctx context.Context
//   - items stockstorage.ReserveItems
func (_e *RepositoryMock_Expecter) Reserve(ctx interface{}, items interface{}) *RepositoryMock_Reserve_Call {
	return &RepositoryMock_Reserve_Call{Call: _e.mock.On("Reserve", ctx, items)}
}

func (_c *RepositoryMock_Reserve_Call) Run(run func(ctx context.Context, items stockstorage.ReserveItems)) *RepositoryMock_Reserve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(stockstorage.ReserveItems))
	})
	return _c
}

func (_c *RepositoryMock_Reserve_Call) Return(_a0 error) *RepositoryMock_Reserve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryMock_Reserve_Call) RunAndReturn(run func(context.Context, stockstorage.ReserveItems) error) *RepositoryMock_Reserve_Call {
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
