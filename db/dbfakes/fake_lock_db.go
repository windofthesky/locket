// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/locket/db"
	"code.cloudfoundry.org/locket/models"
)

type FakeLockDB struct {
	LockStub        func(logger lager.Logger, resource *models.Resource, ttl int64) error
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		logger   lager.Logger
		resource *models.Resource
		ttl      int64
	}
	lockReturns struct {
		result1 error
	}
	ReleaseStub        func(logger lager.Logger, resource *models.Resource) error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
		logger   lager.Logger
		resource *models.Resource
	}
	releaseReturns struct {
		result1 error
	}
	FetchStub        func(logger lager.Logger, key string) (*models.Resource, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		logger lager.Logger
		key    string
	}
	fetchReturns struct {
		result1 *models.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockDB) Lock(logger lager.Logger, resource *models.Resource, ttl int64) error {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		logger   lager.Logger
		resource *models.Resource
		ttl      int64
	}{logger, resource, ttl})
	fake.recordInvocation("Lock", []interface{}{logger, resource, ttl})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(logger, resource, ttl)
	} else {
		return fake.lockReturns.result1
	}
}

func (fake *FakeLockDB) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLockDB) LockArgsForCall(i int) (lager.Logger, *models.Resource, int64) {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].logger, fake.lockArgsForCall[i].resource, fake.lockArgsForCall[i].ttl
}

func (fake *FakeLockDB) LockReturns(result1 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockDB) Release(logger lager.Logger, resource *models.Resource) error {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
		logger   lager.Logger
		resource *models.Resource
	}{logger, resource})
	fake.recordInvocation("Release", []interface{}{logger, resource})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub(logger, resource)
	} else {
		return fake.releaseReturns.result1
	}
}

func (fake *FakeLockDB) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLockDB) ReleaseArgsForCall(i int) (lager.Logger, *models.Resource) {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.releaseArgsForCall[i].logger, fake.releaseArgsForCall[i].resource
}

func (fake *FakeLockDB) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockDB) Fetch(logger lager.Logger, key string) (*models.Resource, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		logger lager.Logger
		key    string
	}{logger, key})
	fake.recordInvocation("Fetch", []interface{}{logger, key})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(logger, key)
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeLockDB) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeLockDB) FetchArgsForCall(i int) (lager.Logger, string) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.fetchArgsForCall[i].logger, fake.fetchArgsForCall[i].key
}

func (fake *FakeLockDB) FetchReturns(result1 *models.Resource, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 *models.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLockDB) recordInvocation(key string, args []interface{}) {
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

var _ db.LockDB = new(FakeLockDB)
