// This file was generated by counterfeiter
package dbfakes

import (
	"cred-alert/db"
	"sync"
)

type FakeCredentialRepository struct {
	ForScanWithIDStub        func(int) ([]db.Credential, error)
	forScanWithIDMutex       sync.RWMutex
	forScanWithIDArgsForCall []struct {
		arg1 int
	}
	forScanWithIDReturns struct {
		result1 []db.Credential
		result2 error
	}
	UniqueSHAsForRepoAndRulesVersionStub        func(db.Repository, int) ([]string, error)
	uniqueSHAsForRepoAndRulesVersionMutex       sync.RWMutex
	uniqueSHAsForRepoAndRulesVersionArgsForCall []struct {
		arg1 db.Repository
		arg2 int
	}
	uniqueSHAsForRepoAndRulesVersionReturns struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredentialRepository) ForScanWithID(arg1 int) ([]db.Credential, error) {
	fake.forScanWithIDMutex.Lock()
	fake.forScanWithIDArgsForCall = append(fake.forScanWithIDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("ForScanWithID", []interface{}{arg1})
	fake.forScanWithIDMutex.Unlock()
	if fake.ForScanWithIDStub != nil {
		return fake.ForScanWithIDStub(arg1)
	}
	return fake.forScanWithIDReturns.result1, fake.forScanWithIDReturns.result2
}

func (fake *FakeCredentialRepository) ForScanWithIDCallCount() int {
	fake.forScanWithIDMutex.RLock()
	defer fake.forScanWithIDMutex.RUnlock()
	return len(fake.forScanWithIDArgsForCall)
}

func (fake *FakeCredentialRepository) ForScanWithIDArgsForCall(i int) int {
	fake.forScanWithIDMutex.RLock()
	defer fake.forScanWithIDMutex.RUnlock()
	return fake.forScanWithIDArgsForCall[i].arg1
}

func (fake *FakeCredentialRepository) ForScanWithIDReturns(result1 []db.Credential, result2 error) {
	fake.ForScanWithIDStub = nil
	fake.forScanWithIDReturns = struct {
		result1 []db.Credential
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialRepository) UniqueSHAsForRepoAndRulesVersion(arg1 db.Repository, arg2 int) ([]string, error) {
	fake.uniqueSHAsForRepoAndRulesVersionMutex.Lock()
	fake.uniqueSHAsForRepoAndRulesVersionArgsForCall = append(fake.uniqueSHAsForRepoAndRulesVersionArgsForCall, struct {
		arg1 db.Repository
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("UniqueSHAsForRepoAndRulesVersion", []interface{}{arg1, arg2})
	fake.uniqueSHAsForRepoAndRulesVersionMutex.Unlock()
	if fake.UniqueSHAsForRepoAndRulesVersionStub != nil {
		return fake.UniqueSHAsForRepoAndRulesVersionStub(arg1, arg2)
	}
	return fake.uniqueSHAsForRepoAndRulesVersionReturns.result1, fake.uniqueSHAsForRepoAndRulesVersionReturns.result2
}

func (fake *FakeCredentialRepository) UniqueSHAsForRepoAndRulesVersionCallCount() int {
	fake.uniqueSHAsForRepoAndRulesVersionMutex.RLock()
	defer fake.uniqueSHAsForRepoAndRulesVersionMutex.RUnlock()
	return len(fake.uniqueSHAsForRepoAndRulesVersionArgsForCall)
}

func (fake *FakeCredentialRepository) UniqueSHAsForRepoAndRulesVersionArgsForCall(i int) (db.Repository, int) {
	fake.uniqueSHAsForRepoAndRulesVersionMutex.RLock()
	defer fake.uniqueSHAsForRepoAndRulesVersionMutex.RUnlock()
	return fake.uniqueSHAsForRepoAndRulesVersionArgsForCall[i].arg1, fake.uniqueSHAsForRepoAndRulesVersionArgsForCall[i].arg2
}

func (fake *FakeCredentialRepository) UniqueSHAsForRepoAndRulesVersionReturns(result1 []string, result2 error) {
	fake.UniqueSHAsForRepoAndRulesVersionStub = nil
	fake.uniqueSHAsForRepoAndRulesVersionReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.forScanWithIDMutex.RLock()
	defer fake.forScanWithIDMutex.RUnlock()
	fake.uniqueSHAsForRepoAndRulesVersionMutex.RLock()
	defer fake.uniqueSHAsForRepoAndRulesVersionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCredentialRepository) recordInvocation(key string, args []interface{}) {
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

var _ db.CredentialRepository = new(FakeCredentialRepository)
