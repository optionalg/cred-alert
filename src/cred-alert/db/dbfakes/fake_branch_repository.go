// This file was generated by counterfeiter
package dbfakes

import (
	"cred-alert/db"
	"sync"
)

type FakeBranchRepository struct {
	GetBranchesStub        func(repository db.Repository) ([]db.Branch, error)
	getBranchesMutex       sync.RWMutex
	getBranchesArgsForCall []struct {
		repository db.Repository
	}
	getBranchesReturns struct {
		result1 []db.Branch
		result2 error
	}
	UpdateBranchesStub        func(repository db.Repository, branches []db.Branch) error
	updateBranchesMutex       sync.RWMutex
	updateBranchesArgsForCall []struct {
		repository db.Repository
		branches   []db.Branch
	}
	updateBranchesReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBranchRepository) GetBranches(repository db.Repository) ([]db.Branch, error) {
	fake.getBranchesMutex.Lock()
	fake.getBranchesArgsForCall = append(fake.getBranchesArgsForCall, struct {
		repository db.Repository
	}{repository})
	fake.recordInvocation("GetBranches", []interface{}{repository})
	fake.getBranchesMutex.Unlock()
	if fake.GetBranchesStub != nil {
		return fake.GetBranchesStub(repository)
	}
	return fake.getBranchesReturns.result1, fake.getBranchesReturns.result2
}

func (fake *FakeBranchRepository) GetBranchesCallCount() int {
	fake.getBranchesMutex.RLock()
	defer fake.getBranchesMutex.RUnlock()
	return len(fake.getBranchesArgsForCall)
}

func (fake *FakeBranchRepository) GetBranchesArgsForCall(i int) db.Repository {
	fake.getBranchesMutex.RLock()
	defer fake.getBranchesMutex.RUnlock()
	return fake.getBranchesArgsForCall[i].repository
}

func (fake *FakeBranchRepository) GetBranchesReturns(result1 []db.Branch, result2 error) {
	fake.GetBranchesStub = nil
	fake.getBranchesReturns = struct {
		result1 []db.Branch
		result2 error
	}{result1, result2}
}

func (fake *FakeBranchRepository) UpdateBranches(repository db.Repository, branches []db.Branch) error {
	var branchesCopy []db.Branch
	if branches != nil {
		branchesCopy = make([]db.Branch, len(branches))
		copy(branchesCopy, branches)
	}
	fake.updateBranchesMutex.Lock()
	fake.updateBranchesArgsForCall = append(fake.updateBranchesArgsForCall, struct {
		repository db.Repository
		branches   []db.Branch
	}{repository, branchesCopy})
	fake.recordInvocation("UpdateBranches", []interface{}{repository, branchesCopy})
	fake.updateBranchesMutex.Unlock()
	if fake.UpdateBranchesStub != nil {
		return fake.UpdateBranchesStub(repository, branches)
	}
	return fake.updateBranchesReturns.result1
}

func (fake *FakeBranchRepository) UpdateBranchesCallCount() int {
	fake.updateBranchesMutex.RLock()
	defer fake.updateBranchesMutex.RUnlock()
	return len(fake.updateBranchesArgsForCall)
}

func (fake *FakeBranchRepository) UpdateBranchesArgsForCall(i int) (db.Repository, []db.Branch) {
	fake.updateBranchesMutex.RLock()
	defer fake.updateBranchesMutex.RUnlock()
	return fake.updateBranchesArgsForCall[i].repository, fake.updateBranchesArgsForCall[i].branches
}

func (fake *FakeBranchRepository) UpdateBranchesReturns(result1 error) {
	fake.UpdateBranchesStub = nil
	fake.updateBranchesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBranchRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBranchesMutex.RLock()
	defer fake.getBranchesMutex.RUnlock()
	fake.updateBranchesMutex.RLock()
	defer fake.updateBranchesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBranchRepository) recordInvocation(key string, args []interface{}) {
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

var _ db.BranchRepository = new(FakeBranchRepository)