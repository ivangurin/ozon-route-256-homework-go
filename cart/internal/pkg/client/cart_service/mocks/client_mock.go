package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_cartservice "route256.ozon.ru/project/cart/internal/pkg/client/cart_service"
)

// ClientMockMock implements cartservice.ClientMock
type ClientMockMock struct {
	t minimock.Tester

	funcAddItem          func(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) (err error)
	inspectFuncAddItem   func(ctx context.Context, UserID int64, SkuID int64, Quantity uint16)
	afterAddItemCounter  uint64
	beforeAddItemCounter uint64
	AddItemMock          mClientMockMockAddItem

	funcDeleteItem          func(ctx context.Context, UserID int64, SkuID int64) (err error)
	inspectFuncDeleteItem   func(ctx context.Context, UserID int64, SkuID int64)
	afterDeleteItemCounter  uint64
	beforeDeleteItemCounter uint64
	DeleteItemMock          mClientMockMockDeleteItem

	funcDeleteItemsByUserID          func(ctx context.Context, UserID int64) (err error)
	inspectFuncDeleteItemsByUserID   func(ctx context.Context, UserID int64)
	afterDeleteItemsByUserIDCounter  uint64
	beforeDeleteItemsByUserIDCounter uint64
	DeleteItemsByUserIDMock          mClientMockMockDeleteItemsByUserID

	funcGetItemsByUserID          func(ctx context.Context, UserID int64) (gp1 *mm_cartservice.GetItmesByUserIDResponse, err error)
	inspectFuncGetItemsByUserID   func(ctx context.Context, UserID int64)
	afterGetItemsByUserIDCounter  uint64
	beforeGetItemsByUserIDCounter uint64
	GetItemsByUserIDMock          mClientMockMockGetItemsByUserID
}

// NewClientMockMock returns a mock for cartservice.ClientMock
func NewClientMockMock(t minimock.Tester) *ClientMockMock {
	m := &ClientMockMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddItemMock = mClientMockMockAddItem{mock: m}
	m.AddItemMock.callArgs = []*ClientMockMockAddItemParams{}

	m.DeleteItemMock = mClientMockMockDeleteItem{mock: m}
	m.DeleteItemMock.callArgs = []*ClientMockMockDeleteItemParams{}

	m.DeleteItemsByUserIDMock = mClientMockMockDeleteItemsByUserID{mock: m}
	m.DeleteItemsByUserIDMock.callArgs = []*ClientMockMockDeleteItemsByUserIDParams{}

	m.GetItemsByUserIDMock = mClientMockMockGetItemsByUserID{mock: m}
	m.GetItemsByUserIDMock.callArgs = []*ClientMockMockGetItemsByUserIDParams{}

	return m
}

type mClientMockMockAddItem struct {
	mock               *ClientMockMock
	defaultExpectation *ClientMockMockAddItemExpectation
	expectations       []*ClientMockMockAddItemExpectation

	callArgs []*ClientMockMockAddItemParams
	mutex    sync.RWMutex
}

// ClientMockMockAddItemExpectation specifies expectation struct of the ClientMock.AddItem
type ClientMockMockAddItemExpectation struct {
	mock    *ClientMockMock
	params  *ClientMockMockAddItemParams
	results *ClientMockMockAddItemResults
	Counter uint64
}

// ClientMockMockAddItemParams contains parameters of the ClientMock.AddItem
type ClientMockMockAddItemParams struct {
	ctx      context.Context
	UserID   int64
	SkuID    int64
	Quantity uint16
}

// ClientMockMockAddItemResults contains results of the ClientMock.AddItem
type ClientMockMockAddItemResults struct {
	err error
}

// Expect sets up expected params for ClientMock.AddItem
func (mmAddItem *mClientMockMockAddItem) Expect(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) *mClientMockMockAddItem {
	if mmAddItem.mock.funcAddItem != nil {
		mmAddItem.mock.t.Fatalf("ClientMockMock.AddItem mock is already set by Set")
	}

	if mmAddItem.defaultExpectation == nil {
		mmAddItem.defaultExpectation = &ClientMockMockAddItemExpectation{}
	}

	mmAddItem.defaultExpectation.params = &ClientMockMockAddItemParams{ctx, UserID, SkuID, Quantity}
	for _, e := range mmAddItem.expectations {
		if minimock.Equal(e.params, mmAddItem.defaultExpectation.params) {
			mmAddItem.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddItem.defaultExpectation.params)
		}
	}

	return mmAddItem
}

// Inspect accepts an inspector function that has same arguments as the ClientMock.AddItem
func (mmAddItem *mClientMockMockAddItem) Inspect(f func(ctx context.Context, UserID int64, SkuID int64, Quantity uint16)) *mClientMockMockAddItem {
	if mmAddItem.mock.inspectFuncAddItem != nil {
		mmAddItem.mock.t.Fatalf("Inspect function is already set for ClientMockMock.AddItem")
	}

	mmAddItem.mock.inspectFuncAddItem = f

	return mmAddItem
}

// Return sets up results that will be returned by ClientMock.AddItem
func (mmAddItem *mClientMockMockAddItem) Return(err error) *ClientMockMock {
	if mmAddItem.mock.funcAddItem != nil {
		mmAddItem.mock.t.Fatalf("ClientMockMock.AddItem mock is already set by Set")
	}

	if mmAddItem.defaultExpectation == nil {
		mmAddItem.defaultExpectation = &ClientMockMockAddItemExpectation{mock: mmAddItem.mock}
	}
	mmAddItem.defaultExpectation.results = &ClientMockMockAddItemResults{err}
	return mmAddItem.mock
}

// Set uses given function f to mock the ClientMock.AddItem method
func (mmAddItem *mClientMockMockAddItem) Set(f func(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) (err error)) *ClientMockMock {
	if mmAddItem.defaultExpectation != nil {
		mmAddItem.mock.t.Fatalf("Default expectation is already set for the ClientMock.AddItem method")
	}

	if len(mmAddItem.expectations) > 0 {
		mmAddItem.mock.t.Fatalf("Some expectations are already set for the ClientMock.AddItem method")
	}

	mmAddItem.mock.funcAddItem = f
	return mmAddItem.mock
}

// When sets expectation for the ClientMock.AddItem which will trigger the result defined by the following
// Then helper
func (mmAddItem *mClientMockMockAddItem) When(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) *ClientMockMockAddItemExpectation {
	if mmAddItem.mock.funcAddItem != nil {
		mmAddItem.mock.t.Fatalf("ClientMockMock.AddItem mock is already set by Set")
	}

	expectation := &ClientMockMockAddItemExpectation{
		mock:   mmAddItem.mock,
		params: &ClientMockMockAddItemParams{ctx, UserID, SkuID, Quantity},
	}
	mmAddItem.expectations = append(mmAddItem.expectations, expectation)
	return expectation
}

// Then sets up ClientMock.AddItem return parameters for the expectation previously defined by the When method
func (e *ClientMockMockAddItemExpectation) Then(err error) *ClientMockMock {
	e.results = &ClientMockMockAddItemResults{err}
	return e.mock
}

// AddItem implements cartservice.ClientMock
func (mmAddItem *ClientMockMock) AddItem(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) (err error) {
	mm_atomic.AddUint64(&mmAddItem.beforeAddItemCounter, 1)
	defer mm_atomic.AddUint64(&mmAddItem.afterAddItemCounter, 1)

	if mmAddItem.inspectFuncAddItem != nil {
		mmAddItem.inspectFuncAddItem(ctx, UserID, SkuID, Quantity)
	}

	mm_params := &ClientMockMockAddItemParams{ctx, UserID, SkuID, Quantity}

	// Record call args
	mmAddItem.AddItemMock.mutex.Lock()
	mmAddItem.AddItemMock.callArgs = append(mmAddItem.AddItemMock.callArgs, mm_params)
	mmAddItem.AddItemMock.mutex.Unlock()

	for _, e := range mmAddItem.AddItemMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddItem.AddItemMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddItem.AddItemMock.defaultExpectation.Counter, 1)
		mm_want := mmAddItem.AddItemMock.defaultExpectation.params
		mm_got := ClientMockMockAddItemParams{ctx, UserID, SkuID, Quantity}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddItem.t.Errorf("ClientMockMock.AddItem got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddItem.AddItemMock.defaultExpectation.results
		if mm_results == nil {
			mmAddItem.t.Fatal("No results are set for the ClientMockMock.AddItem")
		}
		return (*mm_results).err
	}
	if mmAddItem.funcAddItem != nil {
		return mmAddItem.funcAddItem(ctx, UserID, SkuID, Quantity)
	}
	mmAddItem.t.Fatalf("Unexpected call to ClientMockMock.AddItem. %v %v %v %v", ctx, UserID, SkuID, Quantity)
	return
}

// AddItemAfterCounter returns a count of finished ClientMockMock.AddItem invocations
func (mmAddItem *ClientMockMock) AddItemAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddItem.afterAddItemCounter)
}

// AddItemBeforeCounter returns a count of ClientMockMock.AddItem invocations
func (mmAddItem *ClientMockMock) AddItemBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddItem.beforeAddItemCounter)
}

// Calls returns a list of arguments used in each call to ClientMockMock.AddItem.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddItem *mClientMockMockAddItem) Calls() []*ClientMockMockAddItemParams {
	mmAddItem.mutex.RLock()

	argCopy := make([]*ClientMockMockAddItemParams, len(mmAddItem.callArgs))
	copy(argCopy, mmAddItem.callArgs)

	mmAddItem.mutex.RUnlock()

	return argCopy
}

// MinimockAddItemDone returns true if the count of the AddItem invocations corresponds
// the number of defined expectations
func (m *ClientMockMock) MinimockAddItemDone() bool {
	for _, e := range m.AddItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddItemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddItemCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddItem != nil && mm_atomic.LoadUint64(&m.afterAddItemCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddItemInspect logs each unmet expectation
func (m *ClientMockMock) MinimockAddItemInspect() {
	for _, e := range m.AddItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMockMock.AddItem with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddItemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddItemCounter) < 1 {
		if m.AddItemMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMockMock.AddItem")
		} else {
			m.t.Errorf("Expected call to ClientMockMock.AddItem with params: %#v", *m.AddItemMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddItem != nil && mm_atomic.LoadUint64(&m.afterAddItemCounter) < 1 {
		m.t.Error("Expected call to ClientMockMock.AddItem")
	}
}

type mClientMockMockDeleteItem struct {
	mock               *ClientMockMock
	defaultExpectation *ClientMockMockDeleteItemExpectation
	expectations       []*ClientMockMockDeleteItemExpectation

	callArgs []*ClientMockMockDeleteItemParams
	mutex    sync.RWMutex
}

// ClientMockMockDeleteItemExpectation specifies expectation struct of the ClientMock.DeleteItem
type ClientMockMockDeleteItemExpectation struct {
	mock    *ClientMockMock
	params  *ClientMockMockDeleteItemParams
	results *ClientMockMockDeleteItemResults
	Counter uint64
}

// ClientMockMockDeleteItemParams contains parameters of the ClientMock.DeleteItem
type ClientMockMockDeleteItemParams struct {
	ctx    context.Context
	UserID int64
	SkuID  int64
}

// ClientMockMockDeleteItemResults contains results of the ClientMock.DeleteItem
type ClientMockMockDeleteItemResults struct {
	err error
}

// Expect sets up expected params for ClientMock.DeleteItem
func (mmDeleteItem *mClientMockMockDeleteItem) Expect(ctx context.Context, UserID int64, SkuID int64) *mClientMockMockDeleteItem {
	if mmDeleteItem.mock.funcDeleteItem != nil {
		mmDeleteItem.mock.t.Fatalf("ClientMockMock.DeleteItem mock is already set by Set")
	}

	if mmDeleteItem.defaultExpectation == nil {
		mmDeleteItem.defaultExpectation = &ClientMockMockDeleteItemExpectation{}
	}

	mmDeleteItem.defaultExpectation.params = &ClientMockMockDeleteItemParams{ctx, UserID, SkuID}
	for _, e := range mmDeleteItem.expectations {
		if minimock.Equal(e.params, mmDeleteItem.defaultExpectation.params) {
			mmDeleteItem.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteItem.defaultExpectation.params)
		}
	}

	return mmDeleteItem
}

// Inspect accepts an inspector function that has same arguments as the ClientMock.DeleteItem
func (mmDeleteItem *mClientMockMockDeleteItem) Inspect(f func(ctx context.Context, UserID int64, SkuID int64)) *mClientMockMockDeleteItem {
	if mmDeleteItem.mock.inspectFuncDeleteItem != nil {
		mmDeleteItem.mock.t.Fatalf("Inspect function is already set for ClientMockMock.DeleteItem")
	}

	mmDeleteItem.mock.inspectFuncDeleteItem = f

	return mmDeleteItem
}

// Return sets up results that will be returned by ClientMock.DeleteItem
func (mmDeleteItem *mClientMockMockDeleteItem) Return(err error) *ClientMockMock {
	if mmDeleteItem.mock.funcDeleteItem != nil {
		mmDeleteItem.mock.t.Fatalf("ClientMockMock.DeleteItem mock is already set by Set")
	}

	if mmDeleteItem.defaultExpectation == nil {
		mmDeleteItem.defaultExpectation = &ClientMockMockDeleteItemExpectation{mock: mmDeleteItem.mock}
	}
	mmDeleteItem.defaultExpectation.results = &ClientMockMockDeleteItemResults{err}
	return mmDeleteItem.mock
}

// Set uses given function f to mock the ClientMock.DeleteItem method
func (mmDeleteItem *mClientMockMockDeleteItem) Set(f func(ctx context.Context, UserID int64, SkuID int64) (err error)) *ClientMockMock {
	if mmDeleteItem.defaultExpectation != nil {
		mmDeleteItem.mock.t.Fatalf("Default expectation is already set for the ClientMock.DeleteItem method")
	}

	if len(mmDeleteItem.expectations) > 0 {
		mmDeleteItem.mock.t.Fatalf("Some expectations are already set for the ClientMock.DeleteItem method")
	}

	mmDeleteItem.mock.funcDeleteItem = f
	return mmDeleteItem.mock
}

// When sets expectation for the ClientMock.DeleteItem which will trigger the result defined by the following
// Then helper
func (mmDeleteItem *mClientMockMockDeleteItem) When(ctx context.Context, UserID int64, SkuID int64) *ClientMockMockDeleteItemExpectation {
	if mmDeleteItem.mock.funcDeleteItem != nil {
		mmDeleteItem.mock.t.Fatalf("ClientMockMock.DeleteItem mock is already set by Set")
	}

	expectation := &ClientMockMockDeleteItemExpectation{
		mock:   mmDeleteItem.mock,
		params: &ClientMockMockDeleteItemParams{ctx, UserID, SkuID},
	}
	mmDeleteItem.expectations = append(mmDeleteItem.expectations, expectation)
	return expectation
}

// Then sets up ClientMock.DeleteItem return parameters for the expectation previously defined by the When method
func (e *ClientMockMockDeleteItemExpectation) Then(err error) *ClientMockMock {
	e.results = &ClientMockMockDeleteItemResults{err}
	return e.mock
}

// DeleteItem implements cartservice.ClientMock
func (mmDeleteItem *ClientMockMock) DeleteItem(ctx context.Context, UserID int64, SkuID int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteItem.beforeDeleteItemCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteItem.afterDeleteItemCounter, 1)

	if mmDeleteItem.inspectFuncDeleteItem != nil {
		mmDeleteItem.inspectFuncDeleteItem(ctx, UserID, SkuID)
	}

	mm_params := &ClientMockMockDeleteItemParams{ctx, UserID, SkuID}

	// Record call args
	mmDeleteItem.DeleteItemMock.mutex.Lock()
	mmDeleteItem.DeleteItemMock.callArgs = append(mmDeleteItem.DeleteItemMock.callArgs, mm_params)
	mmDeleteItem.DeleteItemMock.mutex.Unlock()

	for _, e := range mmDeleteItem.DeleteItemMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteItem.DeleteItemMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteItem.DeleteItemMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteItem.DeleteItemMock.defaultExpectation.params
		mm_got := ClientMockMockDeleteItemParams{ctx, UserID, SkuID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteItem.t.Errorf("ClientMockMock.DeleteItem got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteItem.DeleteItemMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteItem.t.Fatal("No results are set for the ClientMockMock.DeleteItem")
		}
		return (*mm_results).err
	}
	if mmDeleteItem.funcDeleteItem != nil {
		return mmDeleteItem.funcDeleteItem(ctx, UserID, SkuID)
	}
	mmDeleteItem.t.Fatalf("Unexpected call to ClientMockMock.DeleteItem. %v %v %v", ctx, UserID, SkuID)
	return
}

// DeleteItemAfterCounter returns a count of finished ClientMockMock.DeleteItem invocations
func (mmDeleteItem *ClientMockMock) DeleteItemAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteItem.afterDeleteItemCounter)
}

// DeleteItemBeforeCounter returns a count of ClientMockMock.DeleteItem invocations
func (mmDeleteItem *ClientMockMock) DeleteItemBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteItem.beforeDeleteItemCounter)
}

// Calls returns a list of arguments used in each call to ClientMockMock.DeleteItem.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteItem *mClientMockMockDeleteItem) Calls() []*ClientMockMockDeleteItemParams {
	mmDeleteItem.mutex.RLock()

	argCopy := make([]*ClientMockMockDeleteItemParams, len(mmDeleteItem.callArgs))
	copy(argCopy, mmDeleteItem.callArgs)

	mmDeleteItem.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteItemDone returns true if the count of the DeleteItem invocations corresponds
// the number of defined expectations
func (m *ClientMockMock) MinimockDeleteItemDone() bool {
	for _, e := range m.DeleteItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteItemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteItemCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteItem != nil && mm_atomic.LoadUint64(&m.afterDeleteItemCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteItemInspect logs each unmet expectation
func (m *ClientMockMock) MinimockDeleteItemInspect() {
	for _, e := range m.DeleteItemMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMockMock.DeleteItem with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteItemMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteItemCounter) < 1 {
		if m.DeleteItemMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMockMock.DeleteItem")
		} else {
			m.t.Errorf("Expected call to ClientMockMock.DeleteItem with params: %#v", *m.DeleteItemMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteItem != nil && mm_atomic.LoadUint64(&m.afterDeleteItemCounter) < 1 {
		m.t.Error("Expected call to ClientMockMock.DeleteItem")
	}
}

type mClientMockMockDeleteItemsByUserID struct {
	mock               *ClientMockMock
	defaultExpectation *ClientMockMockDeleteItemsByUserIDExpectation
	expectations       []*ClientMockMockDeleteItemsByUserIDExpectation

	callArgs []*ClientMockMockDeleteItemsByUserIDParams
	mutex    sync.RWMutex
}

// ClientMockMockDeleteItemsByUserIDExpectation specifies expectation struct of the ClientMock.DeleteItemsByUserID
type ClientMockMockDeleteItemsByUserIDExpectation struct {
	mock    *ClientMockMock
	params  *ClientMockMockDeleteItemsByUserIDParams
	results *ClientMockMockDeleteItemsByUserIDResults
	Counter uint64
}

// ClientMockMockDeleteItemsByUserIDParams contains parameters of the ClientMock.DeleteItemsByUserID
type ClientMockMockDeleteItemsByUserIDParams struct {
	ctx    context.Context
	UserID int64
}

// ClientMockMockDeleteItemsByUserIDResults contains results of the ClientMock.DeleteItemsByUserID
type ClientMockMockDeleteItemsByUserIDResults struct {
	err error
}

// Expect sets up expected params for ClientMock.DeleteItemsByUserID
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) Expect(ctx context.Context, UserID int64) *mClientMockMockDeleteItemsByUserID {
	if mmDeleteItemsByUserID.mock.funcDeleteItemsByUserID != nil {
		mmDeleteItemsByUserID.mock.t.Fatalf("ClientMockMock.DeleteItemsByUserID mock is already set by Set")
	}

	if mmDeleteItemsByUserID.defaultExpectation == nil {
		mmDeleteItemsByUserID.defaultExpectation = &ClientMockMockDeleteItemsByUserIDExpectation{}
	}

	mmDeleteItemsByUserID.defaultExpectation.params = &ClientMockMockDeleteItemsByUserIDParams{ctx, UserID}
	for _, e := range mmDeleteItemsByUserID.expectations {
		if minimock.Equal(e.params, mmDeleteItemsByUserID.defaultExpectation.params) {
			mmDeleteItemsByUserID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteItemsByUserID.defaultExpectation.params)
		}
	}

	return mmDeleteItemsByUserID
}

// Inspect accepts an inspector function that has same arguments as the ClientMock.DeleteItemsByUserID
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) Inspect(f func(ctx context.Context, UserID int64)) *mClientMockMockDeleteItemsByUserID {
	if mmDeleteItemsByUserID.mock.inspectFuncDeleteItemsByUserID != nil {
		mmDeleteItemsByUserID.mock.t.Fatalf("Inspect function is already set for ClientMockMock.DeleteItemsByUserID")
	}

	mmDeleteItemsByUserID.mock.inspectFuncDeleteItemsByUserID = f

	return mmDeleteItemsByUserID
}

// Return sets up results that will be returned by ClientMock.DeleteItemsByUserID
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) Return(err error) *ClientMockMock {
	if mmDeleteItemsByUserID.mock.funcDeleteItemsByUserID != nil {
		mmDeleteItemsByUserID.mock.t.Fatalf("ClientMockMock.DeleteItemsByUserID mock is already set by Set")
	}

	if mmDeleteItemsByUserID.defaultExpectation == nil {
		mmDeleteItemsByUserID.defaultExpectation = &ClientMockMockDeleteItemsByUserIDExpectation{mock: mmDeleteItemsByUserID.mock}
	}
	mmDeleteItemsByUserID.defaultExpectation.results = &ClientMockMockDeleteItemsByUserIDResults{err}
	return mmDeleteItemsByUserID.mock
}

// Set uses given function f to mock the ClientMock.DeleteItemsByUserID method
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) Set(f func(ctx context.Context, UserID int64) (err error)) *ClientMockMock {
	if mmDeleteItemsByUserID.defaultExpectation != nil {
		mmDeleteItemsByUserID.mock.t.Fatalf("Default expectation is already set for the ClientMock.DeleteItemsByUserID method")
	}

	if len(mmDeleteItemsByUserID.expectations) > 0 {
		mmDeleteItemsByUserID.mock.t.Fatalf("Some expectations are already set for the ClientMock.DeleteItemsByUserID method")
	}

	mmDeleteItemsByUserID.mock.funcDeleteItemsByUserID = f
	return mmDeleteItemsByUserID.mock
}

// When sets expectation for the ClientMock.DeleteItemsByUserID which will trigger the result defined by the following
// Then helper
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) When(ctx context.Context, UserID int64) *ClientMockMockDeleteItemsByUserIDExpectation {
	if mmDeleteItemsByUserID.mock.funcDeleteItemsByUserID != nil {
		mmDeleteItemsByUserID.mock.t.Fatalf("ClientMockMock.DeleteItemsByUserID mock is already set by Set")
	}

	expectation := &ClientMockMockDeleteItemsByUserIDExpectation{
		mock:   mmDeleteItemsByUserID.mock,
		params: &ClientMockMockDeleteItemsByUserIDParams{ctx, UserID},
	}
	mmDeleteItemsByUserID.expectations = append(mmDeleteItemsByUserID.expectations, expectation)
	return expectation
}

// Then sets up ClientMock.DeleteItemsByUserID return parameters for the expectation previously defined by the When method
func (e *ClientMockMockDeleteItemsByUserIDExpectation) Then(err error) *ClientMockMock {
	e.results = &ClientMockMockDeleteItemsByUserIDResults{err}
	return e.mock
}

// DeleteItemsByUserID implements cartservice.ClientMock
func (mmDeleteItemsByUserID *ClientMockMock) DeleteItemsByUserID(ctx context.Context, UserID int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteItemsByUserID.beforeDeleteItemsByUserIDCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteItemsByUserID.afterDeleteItemsByUserIDCounter, 1)

	if mmDeleteItemsByUserID.inspectFuncDeleteItemsByUserID != nil {
		mmDeleteItemsByUserID.inspectFuncDeleteItemsByUserID(ctx, UserID)
	}

	mm_params := &ClientMockMockDeleteItemsByUserIDParams{ctx, UserID}

	// Record call args
	mmDeleteItemsByUserID.DeleteItemsByUserIDMock.mutex.Lock()
	mmDeleteItemsByUserID.DeleteItemsByUserIDMock.callArgs = append(mmDeleteItemsByUserID.DeleteItemsByUserIDMock.callArgs, mm_params)
	mmDeleteItemsByUserID.DeleteItemsByUserIDMock.mutex.Unlock()

	for _, e := range mmDeleteItemsByUserID.DeleteItemsByUserIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteItemsByUserID.DeleteItemsByUserIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteItemsByUserID.DeleteItemsByUserIDMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteItemsByUserID.DeleteItemsByUserIDMock.defaultExpectation.params
		mm_got := ClientMockMockDeleteItemsByUserIDParams{ctx, UserID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteItemsByUserID.t.Errorf("ClientMockMock.DeleteItemsByUserID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteItemsByUserID.DeleteItemsByUserIDMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteItemsByUserID.t.Fatal("No results are set for the ClientMockMock.DeleteItemsByUserID")
		}
		return (*mm_results).err
	}
	if mmDeleteItemsByUserID.funcDeleteItemsByUserID != nil {
		return mmDeleteItemsByUserID.funcDeleteItemsByUserID(ctx, UserID)
	}
	mmDeleteItemsByUserID.t.Fatalf("Unexpected call to ClientMockMock.DeleteItemsByUserID. %v %v", ctx, UserID)
	return
}

// DeleteItemsByUserIDAfterCounter returns a count of finished ClientMockMock.DeleteItemsByUserID invocations
func (mmDeleteItemsByUserID *ClientMockMock) DeleteItemsByUserIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteItemsByUserID.afterDeleteItemsByUserIDCounter)
}

// DeleteItemsByUserIDBeforeCounter returns a count of ClientMockMock.DeleteItemsByUserID invocations
func (mmDeleteItemsByUserID *ClientMockMock) DeleteItemsByUserIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteItemsByUserID.beforeDeleteItemsByUserIDCounter)
}

// Calls returns a list of arguments used in each call to ClientMockMock.DeleteItemsByUserID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteItemsByUserID *mClientMockMockDeleteItemsByUserID) Calls() []*ClientMockMockDeleteItemsByUserIDParams {
	mmDeleteItemsByUserID.mutex.RLock()

	argCopy := make([]*ClientMockMockDeleteItemsByUserIDParams, len(mmDeleteItemsByUserID.callArgs))
	copy(argCopy, mmDeleteItemsByUserID.callArgs)

	mmDeleteItemsByUserID.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteItemsByUserIDDone returns true if the count of the DeleteItemsByUserID invocations corresponds
// the number of defined expectations
func (m *ClientMockMock) MinimockDeleteItemsByUserIDDone() bool {
	for _, e := range m.DeleteItemsByUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteItemsByUserIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteItemsByUserIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteItemsByUserID != nil && mm_atomic.LoadUint64(&m.afterDeleteItemsByUserIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteItemsByUserIDInspect logs each unmet expectation
func (m *ClientMockMock) MinimockDeleteItemsByUserIDInspect() {
	for _, e := range m.DeleteItemsByUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMockMock.DeleteItemsByUserID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteItemsByUserIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteItemsByUserIDCounter) < 1 {
		if m.DeleteItemsByUserIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMockMock.DeleteItemsByUserID")
		} else {
			m.t.Errorf("Expected call to ClientMockMock.DeleteItemsByUserID with params: %#v", *m.DeleteItemsByUserIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteItemsByUserID != nil && mm_atomic.LoadUint64(&m.afterDeleteItemsByUserIDCounter) < 1 {
		m.t.Error("Expected call to ClientMockMock.DeleteItemsByUserID")
	}
}

type mClientMockMockGetItemsByUserID struct {
	mock               *ClientMockMock
	defaultExpectation *ClientMockMockGetItemsByUserIDExpectation
	expectations       []*ClientMockMockGetItemsByUserIDExpectation

	callArgs []*ClientMockMockGetItemsByUserIDParams
	mutex    sync.RWMutex
}

// ClientMockMockGetItemsByUserIDExpectation specifies expectation struct of the ClientMock.GetItemsByUserID
type ClientMockMockGetItemsByUserIDExpectation struct {
	mock    *ClientMockMock
	params  *ClientMockMockGetItemsByUserIDParams
	results *ClientMockMockGetItemsByUserIDResults
	Counter uint64
}

// ClientMockMockGetItemsByUserIDParams contains parameters of the ClientMock.GetItemsByUserID
type ClientMockMockGetItemsByUserIDParams struct {
	ctx    context.Context
	UserID int64
}

// ClientMockMockGetItemsByUserIDResults contains results of the ClientMock.GetItemsByUserID
type ClientMockMockGetItemsByUserIDResults struct {
	gp1 *mm_cartservice.GetItmesByUserIDResponse
	err error
}

// Expect sets up expected params for ClientMock.GetItemsByUserID
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) Expect(ctx context.Context, UserID int64) *mClientMockMockGetItemsByUserID {
	if mmGetItemsByUserID.mock.funcGetItemsByUserID != nil {
		mmGetItemsByUserID.mock.t.Fatalf("ClientMockMock.GetItemsByUserID mock is already set by Set")
	}

	if mmGetItemsByUserID.defaultExpectation == nil {
		mmGetItemsByUserID.defaultExpectation = &ClientMockMockGetItemsByUserIDExpectation{}
	}

	mmGetItemsByUserID.defaultExpectation.params = &ClientMockMockGetItemsByUserIDParams{ctx, UserID}
	for _, e := range mmGetItemsByUserID.expectations {
		if minimock.Equal(e.params, mmGetItemsByUserID.defaultExpectation.params) {
			mmGetItemsByUserID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetItemsByUserID.defaultExpectation.params)
		}
	}

	return mmGetItemsByUserID
}

// Inspect accepts an inspector function that has same arguments as the ClientMock.GetItemsByUserID
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) Inspect(f func(ctx context.Context, UserID int64)) *mClientMockMockGetItemsByUserID {
	if mmGetItemsByUserID.mock.inspectFuncGetItemsByUserID != nil {
		mmGetItemsByUserID.mock.t.Fatalf("Inspect function is already set for ClientMockMock.GetItemsByUserID")
	}

	mmGetItemsByUserID.mock.inspectFuncGetItemsByUserID = f

	return mmGetItemsByUserID
}

// Return sets up results that will be returned by ClientMock.GetItemsByUserID
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) Return(gp1 *mm_cartservice.GetItmesByUserIDResponse, err error) *ClientMockMock {
	if mmGetItemsByUserID.mock.funcGetItemsByUserID != nil {
		mmGetItemsByUserID.mock.t.Fatalf("ClientMockMock.GetItemsByUserID mock is already set by Set")
	}

	if mmGetItemsByUserID.defaultExpectation == nil {
		mmGetItemsByUserID.defaultExpectation = &ClientMockMockGetItemsByUserIDExpectation{mock: mmGetItemsByUserID.mock}
	}
	mmGetItemsByUserID.defaultExpectation.results = &ClientMockMockGetItemsByUserIDResults{gp1, err}
	return mmGetItemsByUserID.mock
}

// Set uses given function f to mock the ClientMock.GetItemsByUserID method
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) Set(f func(ctx context.Context, UserID int64) (gp1 *mm_cartservice.GetItmesByUserIDResponse, err error)) *ClientMockMock {
	if mmGetItemsByUserID.defaultExpectation != nil {
		mmGetItemsByUserID.mock.t.Fatalf("Default expectation is already set for the ClientMock.GetItemsByUserID method")
	}

	if len(mmGetItemsByUserID.expectations) > 0 {
		mmGetItemsByUserID.mock.t.Fatalf("Some expectations are already set for the ClientMock.GetItemsByUserID method")
	}

	mmGetItemsByUserID.mock.funcGetItemsByUserID = f
	return mmGetItemsByUserID.mock
}

// When sets expectation for the ClientMock.GetItemsByUserID which will trigger the result defined by the following
// Then helper
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) When(ctx context.Context, UserID int64) *ClientMockMockGetItemsByUserIDExpectation {
	if mmGetItemsByUserID.mock.funcGetItemsByUserID != nil {
		mmGetItemsByUserID.mock.t.Fatalf("ClientMockMock.GetItemsByUserID mock is already set by Set")
	}

	expectation := &ClientMockMockGetItemsByUserIDExpectation{
		mock:   mmGetItemsByUserID.mock,
		params: &ClientMockMockGetItemsByUserIDParams{ctx, UserID},
	}
	mmGetItemsByUserID.expectations = append(mmGetItemsByUserID.expectations, expectation)
	return expectation
}

// Then sets up ClientMock.GetItemsByUserID return parameters for the expectation previously defined by the When method
func (e *ClientMockMockGetItemsByUserIDExpectation) Then(gp1 *mm_cartservice.GetItmesByUserIDResponse, err error) *ClientMockMock {
	e.results = &ClientMockMockGetItemsByUserIDResults{gp1, err}
	return e.mock
}

// GetItemsByUserID implements cartservice.ClientMock
func (mmGetItemsByUserID *ClientMockMock) GetItemsByUserID(ctx context.Context, UserID int64) (gp1 *mm_cartservice.GetItmesByUserIDResponse, err error) {
	mm_atomic.AddUint64(&mmGetItemsByUserID.beforeGetItemsByUserIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetItemsByUserID.afterGetItemsByUserIDCounter, 1)

	if mmGetItemsByUserID.inspectFuncGetItemsByUserID != nil {
		mmGetItemsByUserID.inspectFuncGetItemsByUserID(ctx, UserID)
	}

	mm_params := &ClientMockMockGetItemsByUserIDParams{ctx, UserID}

	// Record call args
	mmGetItemsByUserID.GetItemsByUserIDMock.mutex.Lock()
	mmGetItemsByUserID.GetItemsByUserIDMock.callArgs = append(mmGetItemsByUserID.GetItemsByUserIDMock.callArgs, mm_params)
	mmGetItemsByUserID.GetItemsByUserIDMock.mutex.Unlock()

	for _, e := range mmGetItemsByUserID.GetItemsByUserIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gp1, e.results.err
		}
	}

	if mmGetItemsByUserID.GetItemsByUserIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetItemsByUserID.GetItemsByUserIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetItemsByUserID.GetItemsByUserIDMock.defaultExpectation.params
		mm_got := ClientMockMockGetItemsByUserIDParams{ctx, UserID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetItemsByUserID.t.Errorf("ClientMockMock.GetItemsByUserID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetItemsByUserID.GetItemsByUserIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetItemsByUserID.t.Fatal("No results are set for the ClientMockMock.GetItemsByUserID")
		}
		return (*mm_results).gp1, (*mm_results).err
	}
	if mmGetItemsByUserID.funcGetItemsByUserID != nil {
		return mmGetItemsByUserID.funcGetItemsByUserID(ctx, UserID)
	}
	mmGetItemsByUserID.t.Fatalf("Unexpected call to ClientMockMock.GetItemsByUserID. %v %v", ctx, UserID)
	return
}

// GetItemsByUserIDAfterCounter returns a count of finished ClientMockMock.GetItemsByUserID invocations
func (mmGetItemsByUserID *ClientMockMock) GetItemsByUserIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetItemsByUserID.afterGetItemsByUserIDCounter)
}

// GetItemsByUserIDBeforeCounter returns a count of ClientMockMock.GetItemsByUserID invocations
func (mmGetItemsByUserID *ClientMockMock) GetItemsByUserIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetItemsByUserID.beforeGetItemsByUserIDCounter)
}

// Calls returns a list of arguments used in each call to ClientMockMock.GetItemsByUserID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetItemsByUserID *mClientMockMockGetItemsByUserID) Calls() []*ClientMockMockGetItemsByUserIDParams {
	mmGetItemsByUserID.mutex.RLock()

	argCopy := make([]*ClientMockMockGetItemsByUserIDParams, len(mmGetItemsByUserID.callArgs))
	copy(argCopy, mmGetItemsByUserID.callArgs)

	mmGetItemsByUserID.mutex.RUnlock()

	return argCopy
}

// MinimockGetItemsByUserIDDone returns true if the count of the GetItemsByUserID invocations corresponds
// the number of defined expectations
func (m *ClientMockMock) MinimockGetItemsByUserIDDone() bool {
	for _, e := range m.GetItemsByUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetItemsByUserIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetItemsByUserIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetItemsByUserID != nil && mm_atomic.LoadUint64(&m.afterGetItemsByUserIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetItemsByUserIDInspect logs each unmet expectation
func (m *ClientMockMock) MinimockGetItemsByUserIDInspect() {
	for _, e := range m.GetItemsByUserIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMockMock.GetItemsByUserID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetItemsByUserIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetItemsByUserIDCounter) < 1 {
		if m.GetItemsByUserIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMockMock.GetItemsByUserID")
		} else {
			m.t.Errorf("Expected call to ClientMockMock.GetItemsByUserID with params: %#v", *m.GetItemsByUserIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetItemsByUserID != nil && mm_atomic.LoadUint64(&m.afterGetItemsByUserIDCounter) < 1 {
		m.t.Error("Expected call to ClientMockMock.GetItemsByUserID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMockMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddItemInspect()

		m.MinimockDeleteItemInspect()

		m.MinimockDeleteItemsByUserIDInspect()

		m.MinimockGetItemsByUserIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMockMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ClientMockMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddItemDone() &&
		m.MinimockDeleteItemDone() &&
		m.MinimockDeleteItemsByUserIDDone() &&
		m.MinimockGetItemsByUserIDDone()
}
