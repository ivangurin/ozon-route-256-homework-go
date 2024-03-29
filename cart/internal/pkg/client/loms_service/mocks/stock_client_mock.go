package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"google.golang.org/grpc"
	"route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
)

// StockClientMockMock implements lomsservice.StockClientMock
type StockClientMockMock struct {
	t minimock.Tester

	funcInfo          func(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption) (sp1 *stock.StockInfoResponse, err error)
	inspectFuncInfo   func(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption)
	afterInfoCounter  uint64
	beforeInfoCounter uint64
	InfoMock          mStockClientMockMockInfo
}

// NewStockClientMockMock returns a mock for lomsservice.StockClientMock
func NewStockClientMockMock(t minimock.Tester) *StockClientMockMock {
	m := &StockClientMockMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.InfoMock = mStockClientMockMockInfo{mock: m}
	m.InfoMock.callArgs = []*StockClientMockMockInfoParams{}

	return m
}

type mStockClientMockMockInfo struct {
	mock               *StockClientMockMock
	defaultExpectation *StockClientMockMockInfoExpectation
	expectations       []*StockClientMockMockInfoExpectation

	callArgs []*StockClientMockMockInfoParams
	mutex    sync.RWMutex
}

// StockClientMockMockInfoExpectation specifies expectation struct of the StockClientMock.Info
type StockClientMockMockInfoExpectation struct {
	mock    *StockClientMockMock
	params  *StockClientMockMockInfoParams
	results *StockClientMockMockInfoResults
	Counter uint64
}

// StockClientMockMockInfoParams contains parameters of the StockClientMock.Info
type StockClientMockMockInfoParams struct {
	ctx  context.Context
	in   *stock.StockInfoRequest
	opts []grpc.CallOption
}

// StockClientMockMockInfoResults contains results of the StockClientMock.Info
type StockClientMockMockInfoResults struct {
	sp1 *stock.StockInfoResponse
	err error
}

// Expect sets up expected params for StockClientMock.Info
func (mmInfo *mStockClientMockMockInfo) Expect(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption) *mStockClientMockMockInfo {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("StockClientMockMock.Info mock is already set by Set")
	}

	if mmInfo.defaultExpectation == nil {
		mmInfo.defaultExpectation = &StockClientMockMockInfoExpectation{}
	}

	mmInfo.defaultExpectation.params = &StockClientMockMockInfoParams{ctx, in, opts}
	for _, e := range mmInfo.expectations {
		if minimock.Equal(e.params, mmInfo.defaultExpectation.params) {
			mmInfo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInfo.defaultExpectation.params)
		}
	}

	return mmInfo
}

// Inspect accepts an inspector function that has same arguments as the StockClientMock.Info
func (mmInfo *mStockClientMockMockInfo) Inspect(f func(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption)) *mStockClientMockMockInfo {
	if mmInfo.mock.inspectFuncInfo != nil {
		mmInfo.mock.t.Fatalf("Inspect function is already set for StockClientMockMock.Info")
	}

	mmInfo.mock.inspectFuncInfo = f

	return mmInfo
}

// Return sets up results that will be returned by StockClientMock.Info
func (mmInfo *mStockClientMockMockInfo) Return(sp1 *stock.StockInfoResponse, err error) *StockClientMockMock {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("StockClientMockMock.Info mock is already set by Set")
	}

	if mmInfo.defaultExpectation == nil {
		mmInfo.defaultExpectation = &StockClientMockMockInfoExpectation{mock: mmInfo.mock}
	}
	mmInfo.defaultExpectation.results = &StockClientMockMockInfoResults{sp1, err}
	return mmInfo.mock
}

// Set uses given function f to mock the StockClientMock.Info method
func (mmInfo *mStockClientMockMockInfo) Set(f func(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption) (sp1 *stock.StockInfoResponse, err error)) *StockClientMockMock {
	if mmInfo.defaultExpectation != nil {
		mmInfo.mock.t.Fatalf("Default expectation is already set for the StockClientMock.Info method")
	}

	if len(mmInfo.expectations) > 0 {
		mmInfo.mock.t.Fatalf("Some expectations are already set for the StockClientMock.Info method")
	}

	mmInfo.mock.funcInfo = f
	return mmInfo.mock
}

// When sets expectation for the StockClientMock.Info which will trigger the result defined by the following
// Then helper
func (mmInfo *mStockClientMockMockInfo) When(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption) *StockClientMockMockInfoExpectation {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("StockClientMockMock.Info mock is already set by Set")
	}

	expectation := &StockClientMockMockInfoExpectation{
		mock:   mmInfo.mock,
		params: &StockClientMockMockInfoParams{ctx, in, opts},
	}
	mmInfo.expectations = append(mmInfo.expectations, expectation)
	return expectation
}

// Then sets up StockClientMock.Info return parameters for the expectation previously defined by the When method
func (e *StockClientMockMockInfoExpectation) Then(sp1 *stock.StockInfoResponse, err error) *StockClientMockMock {
	e.results = &StockClientMockMockInfoResults{sp1, err}
	return e.mock
}

// Info implements lomsservice.StockClientMock
func (mmInfo *StockClientMockMock) Info(ctx context.Context, in *stock.StockInfoRequest, opts ...grpc.CallOption) (sp1 *stock.StockInfoResponse, err error) {
	mm_atomic.AddUint64(&mmInfo.beforeInfoCounter, 1)
	defer mm_atomic.AddUint64(&mmInfo.afterInfoCounter, 1)

	if mmInfo.inspectFuncInfo != nil {
		mmInfo.inspectFuncInfo(ctx, in, opts...)
	}

	mm_params := &StockClientMockMockInfoParams{ctx, in, opts}

	// Record call args
	mmInfo.InfoMock.mutex.Lock()
	mmInfo.InfoMock.callArgs = append(mmInfo.InfoMock.callArgs, mm_params)
	mmInfo.InfoMock.mutex.Unlock()

	for _, e := range mmInfo.InfoMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1, e.results.err
		}
	}

	if mmInfo.InfoMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInfo.InfoMock.defaultExpectation.Counter, 1)
		mm_want := mmInfo.InfoMock.defaultExpectation.params
		mm_got := StockClientMockMockInfoParams{ctx, in, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInfo.t.Errorf("StockClientMockMock.Info got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInfo.InfoMock.defaultExpectation.results
		if mm_results == nil {
			mmInfo.t.Fatal("No results are set for the StockClientMockMock.Info")
		}
		return (*mm_results).sp1, (*mm_results).err
	}
	if mmInfo.funcInfo != nil {
		return mmInfo.funcInfo(ctx, in, opts...)
	}
	mmInfo.t.Fatalf("Unexpected call to StockClientMockMock.Info. %v %v %v", ctx, in, opts)
	return
}

// InfoAfterCounter returns a count of finished StockClientMockMock.Info invocations
func (mmInfo *StockClientMockMock) InfoAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInfo.afterInfoCounter)
}

// InfoBeforeCounter returns a count of StockClientMockMock.Info invocations
func (mmInfo *StockClientMockMock) InfoBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInfo.beforeInfoCounter)
}

// Calls returns a list of arguments used in each call to StockClientMockMock.Info.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInfo *mStockClientMockMockInfo) Calls() []*StockClientMockMockInfoParams {
	mmInfo.mutex.RLock()

	argCopy := make([]*StockClientMockMockInfoParams, len(mmInfo.callArgs))
	copy(argCopy, mmInfo.callArgs)

	mmInfo.mutex.RUnlock()

	return argCopy
}

// MinimockInfoDone returns true if the count of the Info invocations corresponds
// the number of defined expectations
func (m *StockClientMockMock) MinimockInfoDone() bool {
	for _, e := range m.InfoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InfoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInfo != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		return false
	}
	return true
}

// MinimockInfoInspect logs each unmet expectation
func (m *StockClientMockMock) MinimockInfoInspect() {
	for _, e := range m.InfoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StockClientMockMock.Info with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InfoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		if m.InfoMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StockClientMockMock.Info")
		} else {
			m.t.Errorf("Expected call to StockClientMockMock.Info with params: %#v", *m.InfoMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInfo != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		m.t.Error("Expected call to StockClientMockMock.Info")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StockClientMockMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockInfoInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StockClientMockMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *StockClientMockMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockInfoDone()
}
