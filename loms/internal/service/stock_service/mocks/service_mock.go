package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ServiceMockMock implements stockservice.ServiceMock
type ServiceMockMock struct {
	t minimock.Tester

	funcInfo          func(sku int64) (u1 uint16, err error)
	inspectFuncInfo   func(sku int64)
	afterInfoCounter  uint64
	beforeInfoCounter uint64
	InfoMock          mServiceMockMockInfo
}

// NewServiceMockMock returns a mock for stockservice.ServiceMock
func NewServiceMockMock(t minimock.Tester) *ServiceMockMock {
	m := &ServiceMockMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.InfoMock = mServiceMockMockInfo{mock: m}
	m.InfoMock.callArgs = []*ServiceMockMockInfoParams{}

	return m
}

type mServiceMockMockInfo struct {
	mock               *ServiceMockMock
	defaultExpectation *ServiceMockMockInfoExpectation
	expectations       []*ServiceMockMockInfoExpectation

	callArgs []*ServiceMockMockInfoParams
	mutex    sync.RWMutex
}

// ServiceMockMockInfoExpectation specifies expectation struct of the ServiceMock.Info
type ServiceMockMockInfoExpectation struct {
	mock    *ServiceMockMock
	params  *ServiceMockMockInfoParams
	results *ServiceMockMockInfoResults
	Counter uint64
}

// ServiceMockMockInfoParams contains parameters of the ServiceMock.Info
type ServiceMockMockInfoParams struct {
	sku int64
}

// ServiceMockMockInfoResults contains results of the ServiceMock.Info
type ServiceMockMockInfoResults struct {
	u1  uint16
	err error
}

// Expect sets up expected params for ServiceMock.Info
func (mmInfo *mServiceMockMockInfo) Expect(sku int64) *mServiceMockMockInfo {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("ServiceMockMock.Info mock is already set by Set")
	}

	if mmInfo.defaultExpectation == nil {
		mmInfo.defaultExpectation = &ServiceMockMockInfoExpectation{}
	}

	mmInfo.defaultExpectation.params = &ServiceMockMockInfoParams{sku}
	for _, e := range mmInfo.expectations {
		if minimock.Equal(e.params, mmInfo.defaultExpectation.params) {
			mmInfo.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmInfo.defaultExpectation.params)
		}
	}

	return mmInfo
}

// Inspect accepts an inspector function that has same arguments as the ServiceMock.Info
func (mmInfo *mServiceMockMockInfo) Inspect(f func(sku int64)) *mServiceMockMockInfo {
	if mmInfo.mock.inspectFuncInfo != nil {
		mmInfo.mock.t.Fatalf("Inspect function is already set for ServiceMockMock.Info")
	}

	mmInfo.mock.inspectFuncInfo = f

	return mmInfo
}

// Return sets up results that will be returned by ServiceMock.Info
func (mmInfo *mServiceMockMockInfo) Return(u1 uint16, err error) *ServiceMockMock {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("ServiceMockMock.Info mock is already set by Set")
	}

	if mmInfo.defaultExpectation == nil {
		mmInfo.defaultExpectation = &ServiceMockMockInfoExpectation{mock: mmInfo.mock}
	}
	mmInfo.defaultExpectation.results = &ServiceMockMockInfoResults{u1, err}
	return mmInfo.mock
}

// Set uses given function f to mock the ServiceMock.Info method
func (mmInfo *mServiceMockMockInfo) Set(f func(sku int64) (u1 uint16, err error)) *ServiceMockMock {
	if mmInfo.defaultExpectation != nil {
		mmInfo.mock.t.Fatalf("Default expectation is already set for the ServiceMock.Info method")
	}

	if len(mmInfo.expectations) > 0 {
		mmInfo.mock.t.Fatalf("Some expectations are already set for the ServiceMock.Info method")
	}

	mmInfo.mock.funcInfo = f
	return mmInfo.mock
}

// When sets expectation for the ServiceMock.Info which will trigger the result defined by the following
// Then helper
func (mmInfo *mServiceMockMockInfo) When(sku int64) *ServiceMockMockInfoExpectation {
	if mmInfo.mock.funcInfo != nil {
		mmInfo.mock.t.Fatalf("ServiceMockMock.Info mock is already set by Set")
	}

	expectation := &ServiceMockMockInfoExpectation{
		mock:   mmInfo.mock,
		params: &ServiceMockMockInfoParams{sku},
	}
	mmInfo.expectations = append(mmInfo.expectations, expectation)
	return expectation
}

// Then sets up ServiceMock.Info return parameters for the expectation previously defined by the When method
func (e *ServiceMockMockInfoExpectation) Then(u1 uint16, err error) *ServiceMockMock {
	e.results = &ServiceMockMockInfoResults{u1, err}
	return e.mock
}

// Info implements stockservice.ServiceMock
func (mmInfo *ServiceMockMock) Info(sku int64) (u1 uint16, err error) {
	mm_atomic.AddUint64(&mmInfo.beforeInfoCounter, 1)
	defer mm_atomic.AddUint64(&mmInfo.afterInfoCounter, 1)

	if mmInfo.inspectFuncInfo != nil {
		mmInfo.inspectFuncInfo(sku)
	}

	mm_params := &ServiceMockMockInfoParams{sku}

	// Record call args
	mmInfo.InfoMock.mutex.Lock()
	mmInfo.InfoMock.callArgs = append(mmInfo.InfoMock.callArgs, mm_params)
	mmInfo.InfoMock.mutex.Unlock()

	for _, e := range mmInfo.InfoMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmInfo.InfoMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmInfo.InfoMock.defaultExpectation.Counter, 1)
		mm_want := mmInfo.InfoMock.defaultExpectation.params
		mm_got := ServiceMockMockInfoParams{sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmInfo.t.Errorf("ServiceMockMock.Info got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmInfo.InfoMock.defaultExpectation.results
		if mm_results == nil {
			mmInfo.t.Fatal("No results are set for the ServiceMockMock.Info")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmInfo.funcInfo != nil {
		return mmInfo.funcInfo(sku)
	}
	mmInfo.t.Fatalf("Unexpected call to ServiceMockMock.Info. %v", sku)
	return
}

// InfoAfterCounter returns a count of finished ServiceMockMock.Info invocations
func (mmInfo *ServiceMockMock) InfoAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInfo.afterInfoCounter)
}

// InfoBeforeCounter returns a count of ServiceMockMock.Info invocations
func (mmInfo *ServiceMockMock) InfoBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmInfo.beforeInfoCounter)
}

// Calls returns a list of arguments used in each call to ServiceMockMock.Info.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmInfo *mServiceMockMockInfo) Calls() []*ServiceMockMockInfoParams {
	mmInfo.mutex.RLock()

	argCopy := make([]*ServiceMockMockInfoParams, len(mmInfo.callArgs))
	copy(argCopy, mmInfo.callArgs)

	mmInfo.mutex.RUnlock()

	return argCopy
}

// MinimockInfoDone returns true if the count of the Info invocations corresponds
// the number of defined expectations
func (m *ServiceMockMock) MinimockInfoDone() bool {
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
func (m *ServiceMockMock) MinimockInfoInspect() {
	for _, e := range m.InfoMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ServiceMockMock.Info with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.InfoMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		if m.InfoMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ServiceMockMock.Info")
		} else {
			m.t.Errorf("Expected call to ServiceMockMock.Info with params: %#v", *m.InfoMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcInfo != nil && mm_atomic.LoadUint64(&m.afterInfoCounter) < 1 {
		m.t.Error("Expected call to ServiceMockMock.Info")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ServiceMockMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockInfoInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ServiceMockMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ServiceMockMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockInfoDone()
}
