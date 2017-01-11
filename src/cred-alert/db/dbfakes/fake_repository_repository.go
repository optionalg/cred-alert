// This file was generated by counterfeiter
package dbfakes

import (
	"cred-alert/db"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeRepositoryRepository struct {
	CreateStub        func(*db.Repository) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *db.Repository
	}
	createReturns struct {
		result1 error
	}
	FindStub        func(owner string, name string) (db.Repository, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		owner string
		name  string
	}
	findReturns struct {
		result1 db.Repository
		result2 error
	}
	AllStub        func() ([]db.Repository, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []db.Repository
		result2 error
	}
	ActiveStub        func() ([]db.Repository, error)
	activeMutex       sync.RWMutex
	activeArgsForCall []struct{}
	activeReturns     struct {
		result1 []db.Repository
		result2 error
	}
	AllForOrganizationStub        func(string) ([]db.Repository, error)
	allForOrganizationMutex       sync.RWMutex
	allForOrganizationArgsForCall []struct {
		arg1 string
	}
	allForOrganizationReturns struct {
		result1 []db.Repository
		result2 error
	}
	NotScannedWithVersionStub        func(int) ([]db.Repository, error)
	notScannedWithVersionMutex       sync.RWMutex
	notScannedWithVersionArgsForCall []struct {
		arg1 int
	}
	notScannedWithVersionReturns struct {
		result1 []db.Repository
		result2 error
	}
	MarkAsClonedStub        func(string, string, string) error
	markAsClonedMutex       sync.RWMutex
	markAsClonedArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	markAsClonedReturns struct {
		result1 error
	}
	RegisterFailedFetchStub        func(lager.Logger, *db.Repository) error
	registerFailedFetchMutex       sync.RWMutex
	registerFailedFetchArgsForCall []struct {
		arg1 lager.Logger
		arg2 *db.Repository
	}
	registerFailedFetchReturns struct {
		result1 error
	}
	UpdateCredentialCountStub        func(*db.Repository, map[string]uint) error
	updateCredentialCountMutex       sync.RWMutex
	updateCredentialCountArgsForCall []struct {
		arg1 *db.Repository
		arg2 map[string]uint
	}
	updateCredentialCountReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepositoryRepository) Create(arg1 *db.Repository) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *db.Repository
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeRepositoryRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRepositoryRepository) CreateArgsForCall(i int) *db.Repository {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeRepositoryRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepositoryRepository) Find(owner string, name string) (db.Repository, error) {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		owner string
		name  string
	}{owner, name})
	fake.recordInvocation("Find", []interface{}{owner, name})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(owner, name)
	} else {
		return fake.findReturns.result1, fake.findReturns.result2
	}
}

func (fake *FakeRepositoryRepository) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeRepositoryRepository) FindArgsForCall(i int) (string, string) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].owner, fake.findArgsForCall[i].name
}

func (fake *FakeRepositoryRepository) FindReturns(result1 db.Repository, result2 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryRepository) All() ([]db.Repository, error) {
	fake.allMutex.Lock()
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	} else {
		return fake.allReturns.result1, fake.allReturns.result2
	}
}

func (fake *FakeRepositoryRepository) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *FakeRepositoryRepository) AllReturns(result1 []db.Repository, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryRepository) Active() ([]db.Repository, error) {
	fake.activeMutex.Lock()
	fake.activeArgsForCall = append(fake.activeArgsForCall, struct{}{})
	fake.recordInvocation("Active", []interface{}{})
	fake.activeMutex.Unlock()
	if fake.ActiveStub != nil {
		return fake.ActiveStub()
	} else {
		return fake.activeReturns.result1, fake.activeReturns.result2
	}
}

func (fake *FakeRepositoryRepository) ActiveCallCount() int {
	fake.activeMutex.RLock()
	defer fake.activeMutex.RUnlock()
	return len(fake.activeArgsForCall)
}

func (fake *FakeRepositoryRepository) ActiveReturns(result1 []db.Repository, result2 error) {
	fake.ActiveStub = nil
	fake.activeReturns = struct {
		result1 []db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryRepository) AllForOrganization(arg1 string) ([]db.Repository, error) {
	fake.allForOrganizationMutex.Lock()
	fake.allForOrganizationArgsForCall = append(fake.allForOrganizationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("AllForOrganization", []interface{}{arg1})
	fake.allForOrganizationMutex.Unlock()
	if fake.AllForOrganizationStub != nil {
		return fake.AllForOrganizationStub(arg1)
	} else {
		return fake.allForOrganizationReturns.result1, fake.allForOrganizationReturns.result2
	}
}

func (fake *FakeRepositoryRepository) AllForOrganizationCallCount() int {
	fake.allForOrganizationMutex.RLock()
	defer fake.allForOrganizationMutex.RUnlock()
	return len(fake.allForOrganizationArgsForCall)
}

func (fake *FakeRepositoryRepository) AllForOrganizationArgsForCall(i int) string {
	fake.allForOrganizationMutex.RLock()
	defer fake.allForOrganizationMutex.RUnlock()
	return fake.allForOrganizationArgsForCall[i].arg1
}

func (fake *FakeRepositoryRepository) AllForOrganizationReturns(result1 []db.Repository, result2 error) {
	fake.AllForOrganizationStub = nil
	fake.allForOrganizationReturns = struct {
		result1 []db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryRepository) NotScannedWithVersion(arg1 int) ([]db.Repository, error) {
	fake.notScannedWithVersionMutex.Lock()
	fake.notScannedWithVersionArgsForCall = append(fake.notScannedWithVersionArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("NotScannedWithVersion", []interface{}{arg1})
	fake.notScannedWithVersionMutex.Unlock()
	if fake.NotScannedWithVersionStub != nil {
		return fake.NotScannedWithVersionStub(arg1)
	} else {
		return fake.notScannedWithVersionReturns.result1, fake.notScannedWithVersionReturns.result2
	}
}

func (fake *FakeRepositoryRepository) NotScannedWithVersionCallCount() int {
	fake.notScannedWithVersionMutex.RLock()
	defer fake.notScannedWithVersionMutex.RUnlock()
	return len(fake.notScannedWithVersionArgsForCall)
}

func (fake *FakeRepositoryRepository) NotScannedWithVersionArgsForCall(i int) int {
	fake.notScannedWithVersionMutex.RLock()
	defer fake.notScannedWithVersionMutex.RUnlock()
	return fake.notScannedWithVersionArgsForCall[i].arg1
}

func (fake *FakeRepositoryRepository) NotScannedWithVersionReturns(result1 []db.Repository, result2 error) {
	fake.NotScannedWithVersionStub = nil
	fake.notScannedWithVersionReturns = struct {
		result1 []db.Repository
		result2 error
	}{result1, result2}
}

func (fake *FakeRepositoryRepository) MarkAsCloned(arg1 string, arg2 string, arg3 string) error {
	fake.markAsClonedMutex.Lock()
	fake.markAsClonedArgsForCall = append(fake.markAsClonedArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("MarkAsCloned", []interface{}{arg1, arg2, arg3})
	fake.markAsClonedMutex.Unlock()
	if fake.MarkAsClonedStub != nil {
		return fake.MarkAsClonedStub(arg1, arg2, arg3)
	} else {
		return fake.markAsClonedReturns.result1
	}
}

func (fake *FakeRepositoryRepository) MarkAsClonedCallCount() int {
	fake.markAsClonedMutex.RLock()
	defer fake.markAsClonedMutex.RUnlock()
	return len(fake.markAsClonedArgsForCall)
}

func (fake *FakeRepositoryRepository) MarkAsClonedArgsForCall(i int) (string, string, string) {
	fake.markAsClonedMutex.RLock()
	defer fake.markAsClonedMutex.RUnlock()
	return fake.markAsClonedArgsForCall[i].arg1, fake.markAsClonedArgsForCall[i].arg2, fake.markAsClonedArgsForCall[i].arg3
}

func (fake *FakeRepositoryRepository) MarkAsClonedReturns(result1 error) {
	fake.MarkAsClonedStub = nil
	fake.markAsClonedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepositoryRepository) RegisterFailedFetch(arg1 lager.Logger, arg2 *db.Repository) error {
	fake.registerFailedFetchMutex.Lock()
	fake.registerFailedFetchArgsForCall = append(fake.registerFailedFetchArgsForCall, struct {
		arg1 lager.Logger
		arg2 *db.Repository
	}{arg1, arg2})
	fake.recordInvocation("RegisterFailedFetch", []interface{}{arg1, arg2})
	fake.registerFailedFetchMutex.Unlock()
	if fake.RegisterFailedFetchStub != nil {
		return fake.RegisterFailedFetchStub(arg1, arg2)
	} else {
		return fake.registerFailedFetchReturns.result1
	}
}

func (fake *FakeRepositoryRepository) RegisterFailedFetchCallCount() int {
	fake.registerFailedFetchMutex.RLock()
	defer fake.registerFailedFetchMutex.RUnlock()
	return len(fake.registerFailedFetchArgsForCall)
}

func (fake *FakeRepositoryRepository) RegisterFailedFetchArgsForCall(i int) (lager.Logger, *db.Repository) {
	fake.registerFailedFetchMutex.RLock()
	defer fake.registerFailedFetchMutex.RUnlock()
	return fake.registerFailedFetchArgsForCall[i].arg1, fake.registerFailedFetchArgsForCall[i].arg2
}

func (fake *FakeRepositoryRepository) RegisterFailedFetchReturns(result1 error) {
	fake.RegisterFailedFetchStub = nil
	fake.registerFailedFetchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepositoryRepository) UpdateCredentialCount(arg1 *db.Repository, arg2 map[string]uint) error {
	fake.updateCredentialCountMutex.Lock()
	fake.updateCredentialCountArgsForCall = append(fake.updateCredentialCountArgsForCall, struct {
		arg1 *db.Repository
		arg2 map[string]uint
	}{arg1, arg2})
	fake.recordInvocation("UpdateCredentialCount", []interface{}{arg1, arg2})
	fake.updateCredentialCountMutex.Unlock()
	if fake.UpdateCredentialCountStub != nil {
		return fake.UpdateCredentialCountStub(arg1, arg2)
	} else {
		return fake.updateCredentialCountReturns.result1
	}
}

func (fake *FakeRepositoryRepository) UpdateCredentialCountCallCount() int {
	fake.updateCredentialCountMutex.RLock()
	defer fake.updateCredentialCountMutex.RUnlock()
	return len(fake.updateCredentialCountArgsForCall)
}

func (fake *FakeRepositoryRepository) UpdateCredentialCountArgsForCall(i int) (*db.Repository, map[string]uint) {
	fake.updateCredentialCountMutex.RLock()
	defer fake.updateCredentialCountMutex.RUnlock()
	return fake.updateCredentialCountArgsForCall[i].arg1, fake.updateCredentialCountArgsForCall[i].arg2
}

func (fake *FakeRepositoryRepository) UpdateCredentialCountReturns(result1 error) {
	fake.UpdateCredentialCountStub = nil
	fake.updateCredentialCountReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepositoryRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.activeMutex.RLock()
	defer fake.activeMutex.RUnlock()
	fake.allForOrganizationMutex.RLock()
	defer fake.allForOrganizationMutex.RUnlock()
	fake.notScannedWithVersionMutex.RLock()
	defer fake.notScannedWithVersionMutex.RUnlock()
	fake.markAsClonedMutex.RLock()
	defer fake.markAsClonedMutex.RUnlock()
	fake.registerFailedFetchMutex.RLock()
	defer fake.registerFailedFetchMutex.RUnlock()
	fake.updateCredentialCountMutex.RLock()
	defer fake.updateCredentialCountMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRepositoryRepository) recordInvocation(key string, args []interface{}) {
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

var _ db.RepositoryRepository = new(FakeRepositoryRepository)
