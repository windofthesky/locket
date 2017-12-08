// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/locket/metrics"
)

type FakeRequestMetrics struct {
	IncrementRequestsStartedCounterStub        func(requestType string, delta int)
	incrementRequestsStartedCounterMutex       sync.RWMutex
	incrementRequestsStartedCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsSucceededCounterStub        func(requestType string, delta int)
	incrementRequestsSucceededCounterMutex       sync.RWMutex
	incrementRequestsSucceededCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsFailedCounterStub        func(requestType string, delta int)
	incrementRequestsFailedCounterMutex       sync.RWMutex
	incrementRequestsFailedCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	IncrementRequestsInFlightCounterStub        func(requestType string, delta int)
	incrementRequestsInFlightCounterMutex       sync.RWMutex
	incrementRequestsInFlightCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	DecrementRequestsInFlightCounterStub        func(requestType string, delta int)
	decrementRequestsInFlightCounterMutex       sync.RWMutex
	decrementRequestsInFlightCounterArgsForCall []struct {
		requestType string
		delta       int
	}
	UpdateLatencyStub        func(requestType string, dur time.Duration)
	updateLatencyMutex       sync.RWMutex
	updateLatencyArgsForCall []struct {
		requestType string
		dur         time.Duration
	}
	RequestsStartedStub        func(requestType string) uint64
	requestsStartedMutex       sync.RWMutex
	requestsStartedArgsForCall []struct {
		requestType string
	}
	requestsStartedReturns struct {
		result1 uint64
	}
	requestsStartedReturnsOnCall map[int]struct {
		result1 uint64
	}
	RequestsSucceededStub        func(requestType string) uint64
	requestsSucceededMutex       sync.RWMutex
	requestsSucceededArgsForCall []struct {
		requestType string
	}
	requestsSucceededReturns struct {
		result1 uint64
	}
	requestsSucceededReturnsOnCall map[int]struct {
		result1 uint64
	}
	RequestsFailedStub        func(requestType string) uint64
	requestsFailedMutex       sync.RWMutex
	requestsFailedArgsForCall []struct {
		requestType string
	}
	requestsFailedReturns struct {
		result1 uint64
	}
	requestsFailedReturnsOnCall map[int]struct {
		result1 uint64
	}
	RequestsInFlightStub        func(requestType string) uint64
	requestsInFlightMutex       sync.RWMutex
	requestsInFlightArgsForCall []struct {
		requestType string
	}
	requestsInFlightReturns struct {
		result1 uint64
	}
	requestsInFlightReturnsOnCall map[int]struct {
		result1 uint64
	}
	MaxLatencyStub        func(requestType string) time.Duration
	maxLatencyMutex       sync.RWMutex
	maxLatencyArgsForCall []struct {
		requestType string
	}
	maxLatencyReturns struct {
		result1 time.Duration
	}
	maxLatencyReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounter(requestType string, delta int) {
	fake.incrementRequestsStartedCounterMutex.Lock()
	fake.incrementRequestsStartedCounterArgsForCall = append(fake.incrementRequestsStartedCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsStartedCounter", []interface{}{requestType, delta})
	fake.incrementRequestsStartedCounterMutex.Unlock()
	if fake.IncrementRequestsStartedCounterStub != nil {
		fake.IncrementRequestsStartedCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounterCallCount() int {
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	return len(fake.incrementRequestsStartedCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsStartedCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	return fake.incrementRequestsStartedCounterArgsForCall[i].requestType, fake.incrementRequestsStartedCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounter(requestType string, delta int) {
	fake.incrementRequestsSucceededCounterMutex.Lock()
	fake.incrementRequestsSucceededCounterArgsForCall = append(fake.incrementRequestsSucceededCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsSucceededCounter", []interface{}{requestType, delta})
	fake.incrementRequestsSucceededCounterMutex.Unlock()
	if fake.IncrementRequestsSucceededCounterStub != nil {
		fake.IncrementRequestsSucceededCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounterCallCount() int {
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	return len(fake.incrementRequestsSucceededCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsSucceededCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	return fake.incrementRequestsSucceededCounterArgsForCall[i].requestType, fake.incrementRequestsSucceededCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounter(requestType string, delta int) {
	fake.incrementRequestsFailedCounterMutex.Lock()
	fake.incrementRequestsFailedCounterArgsForCall = append(fake.incrementRequestsFailedCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsFailedCounter", []interface{}{requestType, delta})
	fake.incrementRequestsFailedCounterMutex.Unlock()
	if fake.IncrementRequestsFailedCounterStub != nil {
		fake.IncrementRequestsFailedCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounterCallCount() int {
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	return len(fake.incrementRequestsFailedCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsFailedCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	return fake.incrementRequestsFailedCounterArgsForCall[i].requestType, fake.incrementRequestsFailedCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounter(requestType string, delta int) {
	fake.incrementRequestsInFlightCounterMutex.Lock()
	fake.incrementRequestsInFlightCounterArgsForCall = append(fake.incrementRequestsInFlightCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("IncrementRequestsInFlightCounter", []interface{}{requestType, delta})
	fake.incrementRequestsInFlightCounterMutex.Unlock()
	if fake.IncrementRequestsInFlightCounterStub != nil {
		fake.IncrementRequestsInFlightCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounterCallCount() int {
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	return len(fake.incrementRequestsInFlightCounterArgsForCall)
}

func (fake *FakeRequestMetrics) IncrementRequestsInFlightCounterArgsForCall(i int) (string, int) {
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	return fake.incrementRequestsInFlightCounterArgsForCall[i].requestType, fake.incrementRequestsInFlightCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounter(requestType string, delta int) {
	fake.decrementRequestsInFlightCounterMutex.Lock()
	fake.decrementRequestsInFlightCounterArgsForCall = append(fake.decrementRequestsInFlightCounterArgsForCall, struct {
		requestType string
		delta       int
	}{requestType, delta})
	fake.recordInvocation("DecrementRequestsInFlightCounter", []interface{}{requestType, delta})
	fake.decrementRequestsInFlightCounterMutex.Unlock()
	if fake.DecrementRequestsInFlightCounterStub != nil {
		fake.DecrementRequestsInFlightCounterStub(requestType, delta)
	}
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounterCallCount() int {
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	return len(fake.decrementRequestsInFlightCounterArgsForCall)
}

func (fake *FakeRequestMetrics) DecrementRequestsInFlightCounterArgsForCall(i int) (string, int) {
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	return fake.decrementRequestsInFlightCounterArgsForCall[i].requestType, fake.decrementRequestsInFlightCounterArgsForCall[i].delta
}

func (fake *FakeRequestMetrics) UpdateLatency(requestType string, dur time.Duration) {
	fake.updateLatencyMutex.Lock()
	fake.updateLatencyArgsForCall = append(fake.updateLatencyArgsForCall, struct {
		requestType string
		dur         time.Duration
	}{requestType, dur})
	fake.recordInvocation("UpdateLatency", []interface{}{requestType, dur})
	fake.updateLatencyMutex.Unlock()
	if fake.UpdateLatencyStub != nil {
		fake.UpdateLatencyStub(requestType, dur)
	}
}

func (fake *FakeRequestMetrics) UpdateLatencyCallCount() int {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return len(fake.updateLatencyArgsForCall)
}

func (fake *FakeRequestMetrics) UpdateLatencyArgsForCall(i int) (string, time.Duration) {
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	return fake.updateLatencyArgsForCall[i].requestType, fake.updateLatencyArgsForCall[i].dur
}

func (fake *FakeRequestMetrics) RequestsStarted(requestType string) uint64 {
	fake.requestsStartedMutex.Lock()
	ret, specificReturn := fake.requestsStartedReturnsOnCall[len(fake.requestsStartedArgsForCall)]
	fake.requestsStartedArgsForCall = append(fake.requestsStartedArgsForCall, struct {
		requestType string
	}{requestType})
	fake.recordInvocation("RequestsStarted", []interface{}{requestType})
	fake.requestsStartedMutex.Unlock()
	if fake.RequestsStartedStub != nil {
		return fake.RequestsStartedStub(requestType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.requestsStartedReturns.result1
}

func (fake *FakeRequestMetrics) RequestsStartedCallCount() int {
	fake.requestsStartedMutex.RLock()
	defer fake.requestsStartedMutex.RUnlock()
	return len(fake.requestsStartedArgsForCall)
}

func (fake *FakeRequestMetrics) RequestsStartedArgsForCall(i int) string {
	fake.requestsStartedMutex.RLock()
	defer fake.requestsStartedMutex.RUnlock()
	return fake.requestsStartedArgsForCall[i].requestType
}

func (fake *FakeRequestMetrics) RequestsStartedReturns(result1 uint64) {
	fake.RequestsStartedStub = nil
	fake.requestsStartedReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsStartedReturnsOnCall(i int, result1 uint64) {
	fake.RequestsStartedStub = nil
	if fake.requestsStartedReturnsOnCall == nil {
		fake.requestsStartedReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.requestsStartedReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsSucceeded(requestType string) uint64 {
	fake.requestsSucceededMutex.Lock()
	ret, specificReturn := fake.requestsSucceededReturnsOnCall[len(fake.requestsSucceededArgsForCall)]
	fake.requestsSucceededArgsForCall = append(fake.requestsSucceededArgsForCall, struct {
		requestType string
	}{requestType})
	fake.recordInvocation("RequestsSucceeded", []interface{}{requestType})
	fake.requestsSucceededMutex.Unlock()
	if fake.RequestsSucceededStub != nil {
		return fake.RequestsSucceededStub(requestType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.requestsSucceededReturns.result1
}

func (fake *FakeRequestMetrics) RequestsSucceededCallCount() int {
	fake.requestsSucceededMutex.RLock()
	defer fake.requestsSucceededMutex.RUnlock()
	return len(fake.requestsSucceededArgsForCall)
}

func (fake *FakeRequestMetrics) RequestsSucceededArgsForCall(i int) string {
	fake.requestsSucceededMutex.RLock()
	defer fake.requestsSucceededMutex.RUnlock()
	return fake.requestsSucceededArgsForCall[i].requestType
}

func (fake *FakeRequestMetrics) RequestsSucceededReturns(result1 uint64) {
	fake.RequestsSucceededStub = nil
	fake.requestsSucceededReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsSucceededReturnsOnCall(i int, result1 uint64) {
	fake.RequestsSucceededStub = nil
	if fake.requestsSucceededReturnsOnCall == nil {
		fake.requestsSucceededReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.requestsSucceededReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsFailed(requestType string) uint64 {
	fake.requestsFailedMutex.Lock()
	ret, specificReturn := fake.requestsFailedReturnsOnCall[len(fake.requestsFailedArgsForCall)]
	fake.requestsFailedArgsForCall = append(fake.requestsFailedArgsForCall, struct {
		requestType string
	}{requestType})
	fake.recordInvocation("RequestsFailed", []interface{}{requestType})
	fake.requestsFailedMutex.Unlock()
	if fake.RequestsFailedStub != nil {
		return fake.RequestsFailedStub(requestType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.requestsFailedReturns.result1
}

func (fake *FakeRequestMetrics) RequestsFailedCallCount() int {
	fake.requestsFailedMutex.RLock()
	defer fake.requestsFailedMutex.RUnlock()
	return len(fake.requestsFailedArgsForCall)
}

func (fake *FakeRequestMetrics) RequestsFailedArgsForCall(i int) string {
	fake.requestsFailedMutex.RLock()
	defer fake.requestsFailedMutex.RUnlock()
	return fake.requestsFailedArgsForCall[i].requestType
}

func (fake *FakeRequestMetrics) RequestsFailedReturns(result1 uint64) {
	fake.RequestsFailedStub = nil
	fake.requestsFailedReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsFailedReturnsOnCall(i int, result1 uint64) {
	fake.RequestsFailedStub = nil
	if fake.requestsFailedReturnsOnCall == nil {
		fake.requestsFailedReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.requestsFailedReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsInFlight(requestType string) uint64 {
	fake.requestsInFlightMutex.Lock()
	ret, specificReturn := fake.requestsInFlightReturnsOnCall[len(fake.requestsInFlightArgsForCall)]
	fake.requestsInFlightArgsForCall = append(fake.requestsInFlightArgsForCall, struct {
		requestType string
	}{requestType})
	fake.recordInvocation("RequestsInFlight", []interface{}{requestType})
	fake.requestsInFlightMutex.Unlock()
	if fake.RequestsInFlightStub != nil {
		return fake.RequestsInFlightStub(requestType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.requestsInFlightReturns.result1
}

func (fake *FakeRequestMetrics) RequestsInFlightCallCount() int {
	fake.requestsInFlightMutex.RLock()
	defer fake.requestsInFlightMutex.RUnlock()
	return len(fake.requestsInFlightArgsForCall)
}

func (fake *FakeRequestMetrics) RequestsInFlightArgsForCall(i int) string {
	fake.requestsInFlightMutex.RLock()
	defer fake.requestsInFlightMutex.RUnlock()
	return fake.requestsInFlightArgsForCall[i].requestType
}

func (fake *FakeRequestMetrics) RequestsInFlightReturns(result1 uint64) {
	fake.RequestsInFlightStub = nil
	fake.requestsInFlightReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) RequestsInFlightReturnsOnCall(i int, result1 uint64) {
	fake.RequestsInFlightStub = nil
	if fake.requestsInFlightReturnsOnCall == nil {
		fake.requestsInFlightReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.requestsInFlightReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeRequestMetrics) MaxLatency(requestType string) time.Duration {
	fake.maxLatencyMutex.Lock()
	ret, specificReturn := fake.maxLatencyReturnsOnCall[len(fake.maxLatencyArgsForCall)]
	fake.maxLatencyArgsForCall = append(fake.maxLatencyArgsForCall, struct {
		requestType string
	}{requestType})
	fake.recordInvocation("MaxLatency", []interface{}{requestType})
	fake.maxLatencyMutex.Unlock()
	if fake.MaxLatencyStub != nil {
		return fake.MaxLatencyStub(requestType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.maxLatencyReturns.result1
}

func (fake *FakeRequestMetrics) MaxLatencyCallCount() int {
	fake.maxLatencyMutex.RLock()
	defer fake.maxLatencyMutex.RUnlock()
	return len(fake.maxLatencyArgsForCall)
}

func (fake *FakeRequestMetrics) MaxLatencyArgsForCall(i int) string {
	fake.maxLatencyMutex.RLock()
	defer fake.maxLatencyMutex.RUnlock()
	return fake.maxLatencyArgsForCall[i].requestType
}

func (fake *FakeRequestMetrics) MaxLatencyReturns(result1 time.Duration) {
	fake.MaxLatencyStub = nil
	fake.maxLatencyReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeRequestMetrics) MaxLatencyReturnsOnCall(i int, result1 time.Duration) {
	fake.MaxLatencyStub = nil
	if fake.maxLatencyReturnsOnCall == nil {
		fake.maxLatencyReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.maxLatencyReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeRequestMetrics) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementRequestsStartedCounterMutex.RLock()
	defer fake.incrementRequestsStartedCounterMutex.RUnlock()
	fake.incrementRequestsSucceededCounterMutex.RLock()
	defer fake.incrementRequestsSucceededCounterMutex.RUnlock()
	fake.incrementRequestsFailedCounterMutex.RLock()
	defer fake.incrementRequestsFailedCounterMutex.RUnlock()
	fake.incrementRequestsInFlightCounterMutex.RLock()
	defer fake.incrementRequestsInFlightCounterMutex.RUnlock()
	fake.decrementRequestsInFlightCounterMutex.RLock()
	defer fake.decrementRequestsInFlightCounterMutex.RUnlock()
	fake.updateLatencyMutex.RLock()
	defer fake.updateLatencyMutex.RUnlock()
	fake.requestsStartedMutex.RLock()
	defer fake.requestsStartedMutex.RUnlock()
	fake.requestsSucceededMutex.RLock()
	defer fake.requestsSucceededMutex.RUnlock()
	fake.requestsFailedMutex.RLock()
	defer fake.requestsFailedMutex.RUnlock()
	fake.requestsInFlightMutex.RLock()
	defer fake.requestsInFlightMutex.RUnlock()
	fake.maxLatencyMutex.RLock()
	defer fake.maxLatencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestMetrics) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.RequestMetrics = new(FakeRequestMetrics)
