package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

// RepositoryMockMock implements stockstorage.RepositoryMock
type RepositoryMockMock struct {
	t minimock.Tester

	funcCancelReserve          func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)
	inspectFuncCancelReserve   func(ctx context.Context, items mm_stockstorage.ReserveItems)
	afterCancelReserveCounter  uint64
	beforeCancelReserveCounter uint64
	CancelReserveMock          mRepositoryMockMockCancelReserve

	funcGetBySku          func(ctx context.Context, sku int64) (u1 uint16, err error)
	inspectFuncGetBySku   func(ctx context.Context, sku int64)
	afterGetBySkuCounter  uint64
	beforeGetBySkuCounter uint64
	GetBySkuMock          mRepositoryMockMockGetBySku

	funcRemoveReserve          func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)
	inspectFuncRemoveReserve   func(ctx context.Context, items mm_stockstorage.ReserveItems)
	afterRemoveReserveCounter  uint64
	beforeRemoveReserveCounter uint64
	RemoveReserveMock          mRepositoryMockMockRemoveReserve

	funcReserve          func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)
	inspectFuncReserve   func(ctx context.Context, items mm_stockstorage.ReserveItems)
	afterReserveCounter  uint64
	beforeReserveCounter uint64
	ReserveMock          mRepositoryMockMockReserve
}

// NewRepositoryMockMock returns a mock for stockstorage.RepositoryMock
func NewRepositoryMockMock(t minimock.Tester) *RepositoryMockMock {
	m := &RepositoryMockMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CancelReserveMock = mRepositoryMockMockCancelReserve{mock: m}
	m.CancelReserveMock.callArgs = []*RepositoryMockMockCancelReserveParams{}

	m.GetBySkuMock = mRepositoryMockMockGetBySku{mock: m}
	m.GetBySkuMock.callArgs = []*RepositoryMockMockGetBySkuParams{}

	m.RemoveReserveMock = mRepositoryMockMockRemoveReserve{mock: m}
	m.RemoveReserveMock.callArgs = []*RepositoryMockMockRemoveReserveParams{}

	m.ReserveMock = mRepositoryMockMockReserve{mock: m}
	m.ReserveMock.callArgs = []*RepositoryMockMockReserveParams{}

	return m
}

type mRepositoryMockMockCancelReserve struct {
	mock               *RepositoryMockMock
	defaultExpectation *RepositoryMockMockCancelReserveExpectation
	expectations       []*RepositoryMockMockCancelReserveExpectation

	callArgs []*RepositoryMockMockCancelReserveParams
	mutex    sync.RWMutex
}

// RepositoryMockMockCancelReserveExpectation specifies expectation struct of the RepositoryMock.CancelReserve
type RepositoryMockMockCancelReserveExpectation struct {
	mock    *RepositoryMockMock
	params  *RepositoryMockMockCancelReserveParams
	results *RepositoryMockMockCancelReserveResults
	Counter uint64
}

// RepositoryMockMockCancelReserveParams contains parameters of the RepositoryMock.CancelReserve
type RepositoryMockMockCancelReserveParams struct {
	ctx   context.Context
	items mm_stockstorage.ReserveItems
}

// RepositoryMockMockCancelReserveResults contains results of the RepositoryMock.CancelReserve
type RepositoryMockMockCancelReserveResults struct {
	err error
}

// Expect sets up expected params for RepositoryMock.CancelReserve
func (mmCancelReserve *mRepositoryMockMockCancelReserve) Expect(ctx context.Context, items mm_stockstorage.ReserveItems) *mRepositoryMockMockCancelReserve {
	if mmCancelReserve.mock.funcCancelReserve != nil {
		mmCancelReserve.mock.t.Fatalf("RepositoryMockMock.CancelReserve mock is already set by Set")
	}

	if mmCancelReserve.defaultExpectation == nil {
		mmCancelReserve.defaultExpectation = &RepositoryMockMockCancelReserveExpectation{}
	}

	mmCancelReserve.defaultExpectation.params = &RepositoryMockMockCancelReserveParams{ctx, items}
	for _, e := range mmCancelReserve.expectations {
		if minimock.Equal(e.params, mmCancelReserve.defaultExpectation.params) {
			mmCancelReserve.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCancelReserve.defaultExpectation.params)
		}
	}

	return mmCancelReserve
}

// Inspect accepts an inspector function that has same arguments as the RepositoryMock.CancelReserve
func (mmCancelReserve *mRepositoryMockMockCancelReserve) Inspect(f func(ctx context.Context, items mm_stockstorage.ReserveItems)) *mRepositoryMockMockCancelReserve {
	if mmCancelReserve.mock.inspectFuncCancelReserve != nil {
		mmCancelReserve.mock.t.Fatalf("Inspect function is already set for RepositoryMockMock.CancelReserve")
	}

	mmCancelReserve.mock.inspectFuncCancelReserve = f

	return mmCancelReserve
}

// Return sets up results that will be returned by RepositoryMock.CancelReserve
func (mmCancelReserve *mRepositoryMockMockCancelReserve) Return(err error) *RepositoryMockMock {
	if mmCancelReserve.mock.funcCancelReserve != nil {
		mmCancelReserve.mock.t.Fatalf("RepositoryMockMock.CancelReserve mock is already set by Set")
	}

	if mmCancelReserve.defaultExpectation == nil {
		mmCancelReserve.defaultExpectation = &RepositoryMockMockCancelReserveExpectation{mock: mmCancelReserve.mock}
	}
	mmCancelReserve.defaultExpectation.results = &RepositoryMockMockCancelReserveResults{err}
	return mmCancelReserve.mock
}

// Set uses given function f to mock the RepositoryMock.CancelReserve method
func (mmCancelReserve *mRepositoryMockMockCancelReserve) Set(f func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)) *RepositoryMockMock {
	if mmCancelReserve.defaultExpectation != nil {
		mmCancelReserve.mock.t.Fatalf("Default expectation is already set for the RepositoryMock.CancelReserve method")
	}

	if len(mmCancelReserve.expectations) > 0 {
		mmCancelReserve.mock.t.Fatalf("Some expectations are already set for the RepositoryMock.CancelReserve method")
	}

	mmCancelReserve.mock.funcCancelReserve = f
	return mmCancelReserve.mock
}

// When sets expectation for the RepositoryMock.CancelReserve which will trigger the result defined by the following
// Then helper
func (mmCancelReserve *mRepositoryMockMockCancelReserve) When(ctx context.Context, items mm_stockstorage.ReserveItems) *RepositoryMockMockCancelReserveExpectation {
	if mmCancelReserve.mock.funcCancelReserve != nil {
		mmCancelReserve.mock.t.Fatalf("RepositoryMockMock.CancelReserve mock is already set by Set")
	}

	expectation := &RepositoryMockMockCancelReserveExpectation{
		mock:   mmCancelReserve.mock,
		params: &RepositoryMockMockCancelReserveParams{ctx, items},
	}
	mmCancelReserve.expectations = append(mmCancelReserve.expectations, expectation)
	return expectation
}

// Then sets up RepositoryMock.CancelReserve return parameters for the expectation previously defined by the When method
func (e *RepositoryMockMockCancelReserveExpectation) Then(err error) *RepositoryMockMock {
	e.results = &RepositoryMockMockCancelReserveResults{err}
	return e.mock
}

// CancelReserve implements stockstorage.RepositoryMock
func (mmCancelReserve *RepositoryMockMock) CancelReserve(ctx context.Context, items mm_stockstorage.ReserveItems) (err error) {
	mm_atomic.AddUint64(&mmCancelReserve.beforeCancelReserveCounter, 1)
	defer mm_atomic.AddUint64(&mmCancelReserve.afterCancelReserveCounter, 1)

	if mmCancelReserve.inspectFuncCancelReserve != nil {
		mmCancelReserve.inspectFuncCancelReserve(ctx, items)
	}

	mm_params := &RepositoryMockMockCancelReserveParams{ctx, items}

	// Record call args
	mmCancelReserve.CancelReserveMock.mutex.Lock()
	mmCancelReserve.CancelReserveMock.callArgs = append(mmCancelReserve.CancelReserveMock.callArgs, mm_params)
	mmCancelReserve.CancelReserveMock.mutex.Unlock()

	for _, e := range mmCancelReserve.CancelReserveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCancelReserve.CancelReserveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCancelReserve.CancelReserveMock.defaultExpectation.Counter, 1)
		mm_want := mmCancelReserve.CancelReserveMock.defaultExpectation.params
		mm_got := RepositoryMockMockCancelReserveParams{ctx, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCancelReserve.t.Errorf("RepositoryMockMock.CancelReserve got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCancelReserve.CancelReserveMock.defaultExpectation.results
		if mm_results == nil {
			mmCancelReserve.t.Fatal("No results are set for the RepositoryMockMock.CancelReserve")
		}
		return (*mm_results).err
	}
	if mmCancelReserve.funcCancelReserve != nil {
		return mmCancelReserve.funcCancelReserve(ctx, items)
	}
	mmCancelReserve.t.Fatalf("Unexpected call to RepositoryMockMock.CancelReserve. %v %v", ctx, items)
	return
}

// CancelReserveAfterCounter returns a count of finished RepositoryMockMock.CancelReserve invocations
func (mmCancelReserve *RepositoryMockMock) CancelReserveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCancelReserve.afterCancelReserveCounter)
}

// CancelReserveBeforeCounter returns a count of RepositoryMockMock.CancelReserve invocations
func (mmCancelReserve *RepositoryMockMock) CancelReserveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCancelReserve.beforeCancelReserveCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMockMock.CancelReserve.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCancelReserve *mRepositoryMockMockCancelReserve) Calls() []*RepositoryMockMockCancelReserveParams {
	mmCancelReserve.mutex.RLock()

	argCopy := make([]*RepositoryMockMockCancelReserveParams, len(mmCancelReserve.callArgs))
	copy(argCopy, mmCancelReserve.callArgs)

	mmCancelReserve.mutex.RUnlock()

	return argCopy
}

// MinimockCancelReserveDone returns true if the count of the CancelReserve invocations corresponds
// the number of defined expectations
func (m *RepositoryMockMock) MinimockCancelReserveDone() bool {
	for _, e := range m.CancelReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CancelReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCancelReserveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCancelReserve != nil && mm_atomic.LoadUint64(&m.afterCancelReserveCounter) < 1 {
		return false
	}
	return true
}

// MinimockCancelReserveInspect logs each unmet expectation
func (m *RepositoryMockMock) MinimockCancelReserveInspect() {
	for _, e := range m.CancelReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMockMock.CancelReserve with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CancelReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCancelReserveCounter) < 1 {
		if m.CancelReserveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMockMock.CancelReserve")
		} else {
			m.t.Errorf("Expected call to RepositoryMockMock.CancelReserve with params: %#v", *m.CancelReserveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCancelReserve != nil && mm_atomic.LoadUint64(&m.afterCancelReserveCounter) < 1 {
		m.t.Error("Expected call to RepositoryMockMock.CancelReserve")
	}
}

type mRepositoryMockMockGetBySku struct {
	mock               *RepositoryMockMock
	defaultExpectation *RepositoryMockMockGetBySkuExpectation
	expectations       []*RepositoryMockMockGetBySkuExpectation

	callArgs []*RepositoryMockMockGetBySkuParams
	mutex    sync.RWMutex
}

// RepositoryMockMockGetBySkuExpectation specifies expectation struct of the RepositoryMock.GetBySku
type RepositoryMockMockGetBySkuExpectation struct {
	mock    *RepositoryMockMock
	params  *RepositoryMockMockGetBySkuParams
	results *RepositoryMockMockGetBySkuResults
	Counter uint64
}

// RepositoryMockMockGetBySkuParams contains parameters of the RepositoryMock.GetBySku
type RepositoryMockMockGetBySkuParams struct {
	ctx context.Context
	sku int64
}

// RepositoryMockMockGetBySkuResults contains results of the RepositoryMock.GetBySku
type RepositoryMockMockGetBySkuResults struct {
	u1  uint16
	err error
}

// Expect sets up expected params for RepositoryMock.GetBySku
func (mmGetBySku *mRepositoryMockMockGetBySku) Expect(ctx context.Context, sku int64) *mRepositoryMockMockGetBySku {
	if mmGetBySku.mock.funcGetBySku != nil {
		mmGetBySku.mock.t.Fatalf("RepositoryMockMock.GetBySku mock is already set by Set")
	}

	if mmGetBySku.defaultExpectation == nil {
		mmGetBySku.defaultExpectation = &RepositoryMockMockGetBySkuExpectation{}
	}

	mmGetBySku.defaultExpectation.params = &RepositoryMockMockGetBySkuParams{ctx, sku}
	for _, e := range mmGetBySku.expectations {
		if minimock.Equal(e.params, mmGetBySku.defaultExpectation.params) {
			mmGetBySku.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetBySku.defaultExpectation.params)
		}
	}

	return mmGetBySku
}

// Inspect accepts an inspector function that has same arguments as the RepositoryMock.GetBySku
func (mmGetBySku *mRepositoryMockMockGetBySku) Inspect(f func(ctx context.Context, sku int64)) *mRepositoryMockMockGetBySku {
	if mmGetBySku.mock.inspectFuncGetBySku != nil {
		mmGetBySku.mock.t.Fatalf("Inspect function is already set for RepositoryMockMock.GetBySku")
	}

	mmGetBySku.mock.inspectFuncGetBySku = f

	return mmGetBySku
}

// Return sets up results that will be returned by RepositoryMock.GetBySku
func (mmGetBySku *mRepositoryMockMockGetBySku) Return(u1 uint16, err error) *RepositoryMockMock {
	if mmGetBySku.mock.funcGetBySku != nil {
		mmGetBySku.mock.t.Fatalf("RepositoryMockMock.GetBySku mock is already set by Set")
	}

	if mmGetBySku.defaultExpectation == nil {
		mmGetBySku.defaultExpectation = &RepositoryMockMockGetBySkuExpectation{mock: mmGetBySku.mock}
	}
	mmGetBySku.defaultExpectation.results = &RepositoryMockMockGetBySkuResults{u1, err}
	return mmGetBySku.mock
}

// Set uses given function f to mock the RepositoryMock.GetBySku method
func (mmGetBySku *mRepositoryMockMockGetBySku) Set(f func(ctx context.Context, sku int64) (u1 uint16, err error)) *RepositoryMockMock {
	if mmGetBySku.defaultExpectation != nil {
		mmGetBySku.mock.t.Fatalf("Default expectation is already set for the RepositoryMock.GetBySku method")
	}

	if len(mmGetBySku.expectations) > 0 {
		mmGetBySku.mock.t.Fatalf("Some expectations are already set for the RepositoryMock.GetBySku method")
	}

	mmGetBySku.mock.funcGetBySku = f
	return mmGetBySku.mock
}

// When sets expectation for the RepositoryMock.GetBySku which will trigger the result defined by the following
// Then helper
func (mmGetBySku *mRepositoryMockMockGetBySku) When(ctx context.Context, sku int64) *RepositoryMockMockGetBySkuExpectation {
	if mmGetBySku.mock.funcGetBySku != nil {
		mmGetBySku.mock.t.Fatalf("RepositoryMockMock.GetBySku mock is already set by Set")
	}

	expectation := &RepositoryMockMockGetBySkuExpectation{
		mock:   mmGetBySku.mock,
		params: &RepositoryMockMockGetBySkuParams{ctx, sku},
	}
	mmGetBySku.expectations = append(mmGetBySku.expectations, expectation)
	return expectation
}

// Then sets up RepositoryMock.GetBySku return parameters for the expectation previously defined by the When method
func (e *RepositoryMockMockGetBySkuExpectation) Then(u1 uint16, err error) *RepositoryMockMock {
	e.results = &RepositoryMockMockGetBySkuResults{u1, err}
	return e.mock
}

// GetBySku implements stockstorage.RepositoryMock
func (mmGetBySku *RepositoryMockMock) GetBySku(ctx context.Context, sku int64) (u1 uint16, err error) {
	mm_atomic.AddUint64(&mmGetBySku.beforeGetBySkuCounter, 1)
	defer mm_atomic.AddUint64(&mmGetBySku.afterGetBySkuCounter, 1)

	if mmGetBySku.inspectFuncGetBySku != nil {
		mmGetBySku.inspectFuncGetBySku(ctx, sku)
	}

	mm_params := &RepositoryMockMockGetBySkuParams{ctx, sku}

	// Record call args
	mmGetBySku.GetBySkuMock.mutex.Lock()
	mmGetBySku.GetBySkuMock.callArgs = append(mmGetBySku.GetBySkuMock.callArgs, mm_params)
	mmGetBySku.GetBySkuMock.mutex.Unlock()

	for _, e := range mmGetBySku.GetBySkuMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetBySku.GetBySkuMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetBySku.GetBySkuMock.defaultExpectation.Counter, 1)
		mm_want := mmGetBySku.GetBySkuMock.defaultExpectation.params
		mm_got := RepositoryMockMockGetBySkuParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetBySku.t.Errorf("RepositoryMockMock.GetBySku got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetBySku.GetBySkuMock.defaultExpectation.results
		if mm_results == nil {
			mmGetBySku.t.Fatal("No results are set for the RepositoryMockMock.GetBySku")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetBySku.funcGetBySku != nil {
		return mmGetBySku.funcGetBySku(ctx, sku)
	}
	mmGetBySku.t.Fatalf("Unexpected call to RepositoryMockMock.GetBySku. %v %v", ctx, sku)
	return
}

// GetBySkuAfterCounter returns a count of finished RepositoryMockMock.GetBySku invocations
func (mmGetBySku *RepositoryMockMock) GetBySkuAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetBySku.afterGetBySkuCounter)
}

// GetBySkuBeforeCounter returns a count of RepositoryMockMock.GetBySku invocations
func (mmGetBySku *RepositoryMockMock) GetBySkuBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetBySku.beforeGetBySkuCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMockMock.GetBySku.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetBySku *mRepositoryMockMockGetBySku) Calls() []*RepositoryMockMockGetBySkuParams {
	mmGetBySku.mutex.RLock()

	argCopy := make([]*RepositoryMockMockGetBySkuParams, len(mmGetBySku.callArgs))
	copy(argCopy, mmGetBySku.callArgs)

	mmGetBySku.mutex.RUnlock()

	return argCopy
}

// MinimockGetBySkuDone returns true if the count of the GetBySku invocations corresponds
// the number of defined expectations
func (m *RepositoryMockMock) MinimockGetBySkuDone() bool {
	for _, e := range m.GetBySkuMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetBySkuMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetBySkuCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetBySku != nil && mm_atomic.LoadUint64(&m.afterGetBySkuCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetBySkuInspect logs each unmet expectation
func (m *RepositoryMockMock) MinimockGetBySkuInspect() {
	for _, e := range m.GetBySkuMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMockMock.GetBySku with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetBySkuMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetBySkuCounter) < 1 {
		if m.GetBySkuMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMockMock.GetBySku")
		} else {
			m.t.Errorf("Expected call to RepositoryMockMock.GetBySku with params: %#v", *m.GetBySkuMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetBySku != nil && mm_atomic.LoadUint64(&m.afterGetBySkuCounter) < 1 {
		m.t.Error("Expected call to RepositoryMockMock.GetBySku")
	}
}

type mRepositoryMockMockRemoveReserve struct {
	mock               *RepositoryMockMock
	defaultExpectation *RepositoryMockMockRemoveReserveExpectation
	expectations       []*RepositoryMockMockRemoveReserveExpectation

	callArgs []*RepositoryMockMockRemoveReserveParams
	mutex    sync.RWMutex
}

// RepositoryMockMockRemoveReserveExpectation specifies expectation struct of the RepositoryMock.RemoveReserve
type RepositoryMockMockRemoveReserveExpectation struct {
	mock    *RepositoryMockMock
	params  *RepositoryMockMockRemoveReserveParams
	results *RepositoryMockMockRemoveReserveResults
	Counter uint64
}

// RepositoryMockMockRemoveReserveParams contains parameters of the RepositoryMock.RemoveReserve
type RepositoryMockMockRemoveReserveParams struct {
	ctx   context.Context
	items mm_stockstorage.ReserveItems
}

// RepositoryMockMockRemoveReserveResults contains results of the RepositoryMock.RemoveReserve
type RepositoryMockMockRemoveReserveResults struct {
	err error
}

// Expect sets up expected params for RepositoryMock.RemoveReserve
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) Expect(ctx context.Context, items mm_stockstorage.ReserveItems) *mRepositoryMockMockRemoveReserve {
	if mmRemoveReserve.mock.funcRemoveReserve != nil {
		mmRemoveReserve.mock.t.Fatalf("RepositoryMockMock.RemoveReserve mock is already set by Set")
	}

	if mmRemoveReserve.defaultExpectation == nil {
		mmRemoveReserve.defaultExpectation = &RepositoryMockMockRemoveReserveExpectation{}
	}

	mmRemoveReserve.defaultExpectation.params = &RepositoryMockMockRemoveReserveParams{ctx, items}
	for _, e := range mmRemoveReserve.expectations {
		if minimock.Equal(e.params, mmRemoveReserve.defaultExpectation.params) {
			mmRemoveReserve.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRemoveReserve.defaultExpectation.params)
		}
	}

	return mmRemoveReserve
}

// Inspect accepts an inspector function that has same arguments as the RepositoryMock.RemoveReserve
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) Inspect(f func(ctx context.Context, items mm_stockstorage.ReserveItems)) *mRepositoryMockMockRemoveReserve {
	if mmRemoveReserve.mock.inspectFuncRemoveReserve != nil {
		mmRemoveReserve.mock.t.Fatalf("Inspect function is already set for RepositoryMockMock.RemoveReserve")
	}

	mmRemoveReserve.mock.inspectFuncRemoveReserve = f

	return mmRemoveReserve
}

// Return sets up results that will be returned by RepositoryMock.RemoveReserve
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) Return(err error) *RepositoryMockMock {
	if mmRemoveReserve.mock.funcRemoveReserve != nil {
		mmRemoveReserve.mock.t.Fatalf("RepositoryMockMock.RemoveReserve mock is already set by Set")
	}

	if mmRemoveReserve.defaultExpectation == nil {
		mmRemoveReserve.defaultExpectation = &RepositoryMockMockRemoveReserveExpectation{mock: mmRemoveReserve.mock}
	}
	mmRemoveReserve.defaultExpectation.results = &RepositoryMockMockRemoveReserveResults{err}
	return mmRemoveReserve.mock
}

// Set uses given function f to mock the RepositoryMock.RemoveReserve method
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) Set(f func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)) *RepositoryMockMock {
	if mmRemoveReserve.defaultExpectation != nil {
		mmRemoveReserve.mock.t.Fatalf("Default expectation is already set for the RepositoryMock.RemoveReserve method")
	}

	if len(mmRemoveReserve.expectations) > 0 {
		mmRemoveReserve.mock.t.Fatalf("Some expectations are already set for the RepositoryMock.RemoveReserve method")
	}

	mmRemoveReserve.mock.funcRemoveReserve = f
	return mmRemoveReserve.mock
}

// When sets expectation for the RepositoryMock.RemoveReserve which will trigger the result defined by the following
// Then helper
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) When(ctx context.Context, items mm_stockstorage.ReserveItems) *RepositoryMockMockRemoveReserveExpectation {
	if mmRemoveReserve.mock.funcRemoveReserve != nil {
		mmRemoveReserve.mock.t.Fatalf("RepositoryMockMock.RemoveReserve mock is already set by Set")
	}

	expectation := &RepositoryMockMockRemoveReserveExpectation{
		mock:   mmRemoveReserve.mock,
		params: &RepositoryMockMockRemoveReserveParams{ctx, items},
	}
	mmRemoveReserve.expectations = append(mmRemoveReserve.expectations, expectation)
	return expectation
}

// Then sets up RepositoryMock.RemoveReserve return parameters for the expectation previously defined by the When method
func (e *RepositoryMockMockRemoveReserveExpectation) Then(err error) *RepositoryMockMock {
	e.results = &RepositoryMockMockRemoveReserveResults{err}
	return e.mock
}

// RemoveReserve implements stockstorage.RepositoryMock
func (mmRemoveReserve *RepositoryMockMock) RemoveReserve(ctx context.Context, items mm_stockstorage.ReserveItems) (err error) {
	mm_atomic.AddUint64(&mmRemoveReserve.beforeRemoveReserveCounter, 1)
	defer mm_atomic.AddUint64(&mmRemoveReserve.afterRemoveReserveCounter, 1)

	if mmRemoveReserve.inspectFuncRemoveReserve != nil {
		mmRemoveReserve.inspectFuncRemoveReserve(ctx, items)
	}

	mm_params := &RepositoryMockMockRemoveReserveParams{ctx, items}

	// Record call args
	mmRemoveReserve.RemoveReserveMock.mutex.Lock()
	mmRemoveReserve.RemoveReserveMock.callArgs = append(mmRemoveReserve.RemoveReserveMock.callArgs, mm_params)
	mmRemoveReserve.RemoveReserveMock.mutex.Unlock()

	for _, e := range mmRemoveReserve.RemoveReserveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmRemoveReserve.RemoveReserveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRemoveReserve.RemoveReserveMock.defaultExpectation.Counter, 1)
		mm_want := mmRemoveReserve.RemoveReserveMock.defaultExpectation.params
		mm_got := RepositoryMockMockRemoveReserveParams{ctx, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRemoveReserve.t.Errorf("RepositoryMockMock.RemoveReserve got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRemoveReserve.RemoveReserveMock.defaultExpectation.results
		if mm_results == nil {
			mmRemoveReserve.t.Fatal("No results are set for the RepositoryMockMock.RemoveReserve")
		}
		return (*mm_results).err
	}
	if mmRemoveReserve.funcRemoveReserve != nil {
		return mmRemoveReserve.funcRemoveReserve(ctx, items)
	}
	mmRemoveReserve.t.Fatalf("Unexpected call to RepositoryMockMock.RemoveReserve. %v %v", ctx, items)
	return
}

// RemoveReserveAfterCounter returns a count of finished RepositoryMockMock.RemoveReserve invocations
func (mmRemoveReserve *RepositoryMockMock) RemoveReserveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRemoveReserve.afterRemoveReserveCounter)
}

// RemoveReserveBeforeCounter returns a count of RepositoryMockMock.RemoveReserve invocations
func (mmRemoveReserve *RepositoryMockMock) RemoveReserveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRemoveReserve.beforeRemoveReserveCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMockMock.RemoveReserve.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRemoveReserve *mRepositoryMockMockRemoveReserve) Calls() []*RepositoryMockMockRemoveReserveParams {
	mmRemoveReserve.mutex.RLock()

	argCopy := make([]*RepositoryMockMockRemoveReserveParams, len(mmRemoveReserve.callArgs))
	copy(argCopy, mmRemoveReserve.callArgs)

	mmRemoveReserve.mutex.RUnlock()

	return argCopy
}

// MinimockRemoveReserveDone returns true if the count of the RemoveReserve invocations corresponds
// the number of defined expectations
func (m *RepositoryMockMock) MinimockRemoveReserveDone() bool {
	for _, e := range m.RemoveReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RemoveReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRemoveReserveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRemoveReserve != nil && mm_atomic.LoadUint64(&m.afterRemoveReserveCounter) < 1 {
		return false
	}
	return true
}

// MinimockRemoveReserveInspect logs each unmet expectation
func (m *RepositoryMockMock) MinimockRemoveReserveInspect() {
	for _, e := range m.RemoveReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMockMock.RemoveReserve with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RemoveReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRemoveReserveCounter) < 1 {
		if m.RemoveReserveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMockMock.RemoveReserve")
		} else {
			m.t.Errorf("Expected call to RepositoryMockMock.RemoveReserve with params: %#v", *m.RemoveReserveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRemoveReserve != nil && mm_atomic.LoadUint64(&m.afterRemoveReserveCounter) < 1 {
		m.t.Error("Expected call to RepositoryMockMock.RemoveReserve")
	}
}

type mRepositoryMockMockReserve struct {
	mock               *RepositoryMockMock
	defaultExpectation *RepositoryMockMockReserveExpectation
	expectations       []*RepositoryMockMockReserveExpectation

	callArgs []*RepositoryMockMockReserveParams
	mutex    sync.RWMutex
}

// RepositoryMockMockReserveExpectation specifies expectation struct of the RepositoryMock.Reserve
type RepositoryMockMockReserveExpectation struct {
	mock    *RepositoryMockMock
	params  *RepositoryMockMockReserveParams
	results *RepositoryMockMockReserveResults
	Counter uint64
}

// RepositoryMockMockReserveParams contains parameters of the RepositoryMock.Reserve
type RepositoryMockMockReserveParams struct {
	ctx   context.Context
	items mm_stockstorage.ReserveItems
}

// RepositoryMockMockReserveResults contains results of the RepositoryMock.Reserve
type RepositoryMockMockReserveResults struct {
	err error
}

// Expect sets up expected params for RepositoryMock.Reserve
func (mmReserve *mRepositoryMockMockReserve) Expect(ctx context.Context, items mm_stockstorage.ReserveItems) *mRepositoryMockMockReserve {
	if mmReserve.mock.funcReserve != nil {
		mmReserve.mock.t.Fatalf("RepositoryMockMock.Reserve mock is already set by Set")
	}

	if mmReserve.defaultExpectation == nil {
		mmReserve.defaultExpectation = &RepositoryMockMockReserveExpectation{}
	}

	mmReserve.defaultExpectation.params = &RepositoryMockMockReserveParams{ctx, items}
	for _, e := range mmReserve.expectations {
		if minimock.Equal(e.params, mmReserve.defaultExpectation.params) {
			mmReserve.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReserve.defaultExpectation.params)
		}
	}

	return mmReserve
}

// Inspect accepts an inspector function that has same arguments as the RepositoryMock.Reserve
func (mmReserve *mRepositoryMockMockReserve) Inspect(f func(ctx context.Context, items mm_stockstorage.ReserveItems)) *mRepositoryMockMockReserve {
	if mmReserve.mock.inspectFuncReserve != nil {
		mmReserve.mock.t.Fatalf("Inspect function is already set for RepositoryMockMock.Reserve")
	}

	mmReserve.mock.inspectFuncReserve = f

	return mmReserve
}

// Return sets up results that will be returned by RepositoryMock.Reserve
func (mmReserve *mRepositoryMockMockReserve) Return(err error) *RepositoryMockMock {
	if mmReserve.mock.funcReserve != nil {
		mmReserve.mock.t.Fatalf("RepositoryMockMock.Reserve mock is already set by Set")
	}

	if mmReserve.defaultExpectation == nil {
		mmReserve.defaultExpectation = &RepositoryMockMockReserveExpectation{mock: mmReserve.mock}
	}
	mmReserve.defaultExpectation.results = &RepositoryMockMockReserveResults{err}
	return mmReserve.mock
}

// Set uses given function f to mock the RepositoryMock.Reserve method
func (mmReserve *mRepositoryMockMockReserve) Set(f func(ctx context.Context, items mm_stockstorage.ReserveItems) (err error)) *RepositoryMockMock {
	if mmReserve.defaultExpectation != nil {
		mmReserve.mock.t.Fatalf("Default expectation is already set for the RepositoryMock.Reserve method")
	}

	if len(mmReserve.expectations) > 0 {
		mmReserve.mock.t.Fatalf("Some expectations are already set for the RepositoryMock.Reserve method")
	}

	mmReserve.mock.funcReserve = f
	return mmReserve.mock
}

// When sets expectation for the RepositoryMock.Reserve which will trigger the result defined by the following
// Then helper
func (mmReserve *mRepositoryMockMockReserve) When(ctx context.Context, items mm_stockstorage.ReserveItems) *RepositoryMockMockReserveExpectation {
	if mmReserve.mock.funcReserve != nil {
		mmReserve.mock.t.Fatalf("RepositoryMockMock.Reserve mock is already set by Set")
	}

	expectation := &RepositoryMockMockReserveExpectation{
		mock:   mmReserve.mock,
		params: &RepositoryMockMockReserveParams{ctx, items},
	}
	mmReserve.expectations = append(mmReserve.expectations, expectation)
	return expectation
}

// Then sets up RepositoryMock.Reserve return parameters for the expectation previously defined by the When method
func (e *RepositoryMockMockReserveExpectation) Then(err error) *RepositoryMockMock {
	e.results = &RepositoryMockMockReserveResults{err}
	return e.mock
}

// Reserve implements stockstorage.RepositoryMock
func (mmReserve *RepositoryMockMock) Reserve(ctx context.Context, items mm_stockstorage.ReserveItems) (err error) {
	mm_atomic.AddUint64(&mmReserve.beforeReserveCounter, 1)
	defer mm_atomic.AddUint64(&mmReserve.afterReserveCounter, 1)

	if mmReserve.inspectFuncReserve != nil {
		mmReserve.inspectFuncReserve(ctx, items)
	}

	mm_params := &RepositoryMockMockReserveParams{ctx, items}

	// Record call args
	mmReserve.ReserveMock.mutex.Lock()
	mmReserve.ReserveMock.callArgs = append(mmReserve.ReserveMock.callArgs, mm_params)
	mmReserve.ReserveMock.mutex.Unlock()

	for _, e := range mmReserve.ReserveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmReserve.ReserveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReserve.ReserveMock.defaultExpectation.Counter, 1)
		mm_want := mmReserve.ReserveMock.defaultExpectation.params
		mm_got := RepositoryMockMockReserveParams{ctx, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReserve.t.Errorf("RepositoryMockMock.Reserve got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReserve.ReserveMock.defaultExpectation.results
		if mm_results == nil {
			mmReserve.t.Fatal("No results are set for the RepositoryMockMock.Reserve")
		}
		return (*mm_results).err
	}
	if mmReserve.funcReserve != nil {
		return mmReserve.funcReserve(ctx, items)
	}
	mmReserve.t.Fatalf("Unexpected call to RepositoryMockMock.Reserve. %v %v", ctx, items)
	return
}

// ReserveAfterCounter returns a count of finished RepositoryMockMock.Reserve invocations
func (mmReserve *RepositoryMockMock) ReserveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserve.afterReserveCounter)
}

// ReserveBeforeCounter returns a count of RepositoryMockMock.Reserve invocations
func (mmReserve *RepositoryMockMock) ReserveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReserve.beforeReserveCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMockMock.Reserve.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReserve *mRepositoryMockMockReserve) Calls() []*RepositoryMockMockReserveParams {
	mmReserve.mutex.RLock()

	argCopy := make([]*RepositoryMockMockReserveParams, len(mmReserve.callArgs))
	copy(argCopy, mmReserve.callArgs)

	mmReserve.mutex.RUnlock()

	return argCopy
}

// MinimockReserveDone returns true if the count of the Reserve invocations corresponds
// the number of defined expectations
func (m *RepositoryMockMock) MinimockReserveDone() bool {
	for _, e := range m.ReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserve != nil && mm_atomic.LoadUint64(&m.afterReserveCounter) < 1 {
		return false
	}
	return true
}

// MinimockReserveInspect logs each unmet expectation
func (m *RepositoryMockMock) MinimockReserveInspect() {
	for _, e := range m.ReserveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMockMock.Reserve with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReserveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReserveCounter) < 1 {
		if m.ReserveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMockMock.Reserve")
		} else {
			m.t.Errorf("Expected call to RepositoryMockMock.Reserve with params: %#v", *m.ReserveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReserve != nil && mm_atomic.LoadUint64(&m.afterReserveCounter) < 1 {
		m.t.Error("Expected call to RepositoryMockMock.Reserve")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMockMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCancelReserveInspect()

		m.MinimockGetBySkuInspect()

		m.MinimockRemoveReserveInspect()

		m.MinimockReserveInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMockMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RepositoryMockMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCancelReserveDone() &&
		m.MinimockGetBySkuDone() &&
		m.MinimockRemoveReserveDone() &&
		m.MinimockReserveDone()
}