// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"

	order "route256.ozon.ru/project/cart/internal/pb/api/order/v1"
)

// OrderClientMock is an autogenerated mock type for the OrderClientMock type
type OrderClientMock struct {
	mock.Mock
}

type OrderClientMock_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderClientMock) EXPECT() *OrderClientMock_Expecter {
	return &OrderClientMock_Expecter{mock: &_m.Mock}
}

// Cancel provides a mock function with given fields: ctx, in, opts
func (_m *OrderClientMock) Cancel(ctx context.Context, in *order.OrderCancelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Cancel")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderCancelRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderCancelRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderCancelRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderClientMock_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type OrderClientMock_Cancel_Call struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
//   - ctx context.Context
//   - in *order.OrderCancelRequest
//   - opts ...grpc.CallOption
func (_e *OrderClientMock_Expecter) Cancel(ctx interface{}, in interface{}, opts ...interface{}) *OrderClientMock_Cancel_Call {
	return &OrderClientMock_Cancel_Call{Call: _e.mock.On("Cancel",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderClientMock_Cancel_Call) Run(run func(ctx context.Context, in *order.OrderCancelRequest, opts ...grpc.CallOption)) *OrderClientMock_Cancel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*order.OrderCancelRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderClientMock_Cancel_Call) Return(_a0 *emptypb.Empty, _a1 error) *OrderClientMock_Cancel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderClientMock_Cancel_Call) RunAndReturn(run func(context.Context, *order.OrderCancelRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *OrderClientMock_Cancel_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *OrderClientMock) Create(ctx context.Context, in *order.OrderCreateRequest, opts ...grpc.CallOption) (*order.OrderCreateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *order.OrderCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderCreateRequest, ...grpc.CallOption) (*order.OrderCreateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderCreateRequest, ...grpc.CallOption) *order.OrderCreateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.OrderCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderCreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderClientMock_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type OrderClientMock_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - in *order.OrderCreateRequest
//   - opts ...grpc.CallOption
func (_e *OrderClientMock_Expecter) Create(ctx interface{}, in interface{}, opts ...interface{}) *OrderClientMock_Create_Call {
	return &OrderClientMock_Create_Call{Call: _e.mock.On("Create",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderClientMock_Create_Call) Run(run func(ctx context.Context, in *order.OrderCreateRequest, opts ...grpc.CallOption)) *OrderClientMock_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*order.OrderCreateRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderClientMock_Create_Call) Return(_a0 *order.OrderCreateResponse, _a1 error) *OrderClientMock_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderClientMock_Create_Call) RunAndReturn(run func(context.Context, *order.OrderCreateRequest, ...grpc.CallOption) (*order.OrderCreateResponse, error)) *OrderClientMock_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, in, opts
func (_m *OrderClientMock) GetByIDs(ctx context.Context, in *order.GetOrdersByIDsRequest, opts ...grpc.CallOption) (*order.GetOrdersByIDsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 *order.GetOrdersByIDsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *order.GetOrdersByIDsRequest, ...grpc.CallOption) (*order.GetOrdersByIDsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *order.GetOrdersByIDsRequest, ...grpc.CallOption) *order.GetOrdersByIDsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.GetOrdersByIDsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *order.GetOrdersByIDsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderClientMock_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type OrderClientMock_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - in *order.GetOrdersByIDsRequest
//   - opts ...grpc.CallOption
func (_e *OrderClientMock_Expecter) GetByIDs(ctx interface{}, in interface{}, opts ...interface{}) *OrderClientMock_GetByIDs_Call {
	return &OrderClientMock_GetByIDs_Call{Call: _e.mock.On("GetByIDs",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderClientMock_GetByIDs_Call) Run(run func(ctx context.Context, in *order.GetOrdersByIDsRequest, opts ...grpc.CallOption)) *OrderClientMock_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*order.GetOrdersByIDsRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderClientMock_GetByIDs_Call) Return(_a0 *order.GetOrdersByIDsResponse, _a1 error) *OrderClientMock_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderClientMock_GetByIDs_Call) RunAndReturn(run func(context.Context, *order.GetOrdersByIDsRequest, ...grpc.CallOption) (*order.GetOrdersByIDsResponse, error)) *OrderClientMock_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: ctx, in, opts
func (_m *OrderClientMock) Info(ctx context.Context, in *order.OrderInfoRequest, opts ...grpc.CallOption) (*order.OrderInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Info")
	}

	var r0 *order.OrderInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderInfoRequest, ...grpc.CallOption) (*order.OrderInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderInfoRequest, ...grpc.CallOption) *order.OrderInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.OrderInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderClientMock_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type OrderClientMock_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - ctx context.Context
//   - in *order.OrderInfoRequest
//   - opts ...grpc.CallOption
func (_e *OrderClientMock_Expecter) Info(ctx interface{}, in interface{}, opts ...interface{}) *OrderClientMock_Info_Call {
	return &OrderClientMock_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderClientMock_Info_Call) Run(run func(ctx context.Context, in *order.OrderInfoRequest, opts ...grpc.CallOption)) *OrderClientMock_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*order.OrderInfoRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderClientMock_Info_Call) Return(_a0 *order.OrderInfoResponse, _a1 error) *OrderClientMock_Info_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderClientMock_Info_Call) RunAndReturn(run func(context.Context, *order.OrderInfoRequest, ...grpc.CallOption) (*order.OrderInfoResponse, error)) *OrderClientMock_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Pay provides a mock function with given fields: ctx, in, opts
func (_m *OrderClientMock) Pay(ctx context.Context, in *order.OrderPayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Pay")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderPayRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *order.OrderPayRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *order.OrderPayRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderClientMock_Pay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pay'
type OrderClientMock_Pay_Call struct {
	*mock.Call
}

// Pay is a helper method to define mock.On call
//   - ctx context.Context
//   - in *order.OrderPayRequest
//   - opts ...grpc.CallOption
func (_e *OrderClientMock_Expecter) Pay(ctx interface{}, in interface{}, opts ...interface{}) *OrderClientMock_Pay_Call {
	return &OrderClientMock_Pay_Call{Call: _e.mock.On("Pay",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OrderClientMock_Pay_Call) Run(run func(ctx context.Context, in *order.OrderPayRequest, opts ...grpc.CallOption)) *OrderClientMock_Pay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*order.OrderPayRequest), variadicArgs...)
	})
	return _c
}

func (_c *OrderClientMock_Pay_Call) Return(_a0 *emptypb.Empty, _a1 error) *OrderClientMock_Pay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderClientMock_Pay_Call) RunAndReturn(run func(context.Context, *order.OrderPayRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *OrderClientMock_Pay_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderClientMock creates a new instance of OrderClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderClientMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderClientMock {
	mock := &OrderClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
