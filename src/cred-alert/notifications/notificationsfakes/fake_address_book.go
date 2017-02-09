// This file was generated by counterfeiter
package notificationsfakes

import (
	"cred-alert/notifications"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeAddressBook struct {
	AddressForRepoStub        func(logger lager.Logger, owner, name string) []notifications.Address
	addressForRepoMutex       sync.RWMutex
	addressForRepoArgsForCall []struct {
		logger lager.Logger
		owner  string
		name   string
	}
	addressForRepoReturns struct {
		result1 []notifications.Address
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAddressBook) AddressForRepo(logger lager.Logger, owner string, name string) []notifications.Address {
	fake.addressForRepoMutex.Lock()
	fake.addressForRepoArgsForCall = append(fake.addressForRepoArgsForCall, struct {
		logger lager.Logger
		owner  string
		name   string
	}{logger, owner, name})
	fake.recordInvocation("AddressForRepo", []interface{}{logger, owner, name})
	fake.addressForRepoMutex.Unlock()
	if fake.AddressForRepoStub != nil {
		return fake.AddressForRepoStub(logger, owner, name)
	}
	return fake.addressForRepoReturns.result1
}

func (fake *FakeAddressBook) AddressForRepoCallCount() int {
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	return len(fake.addressForRepoArgsForCall)
}

func (fake *FakeAddressBook) AddressForRepoArgsForCall(i int) (lager.Logger, string, string) {
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	return fake.addressForRepoArgsForCall[i].logger, fake.addressForRepoArgsForCall[i].owner, fake.addressForRepoArgsForCall[i].name
}

func (fake *FakeAddressBook) AddressForRepoReturns(result1 []notifications.Address) {
	fake.AddressForRepoStub = nil
	fake.addressForRepoReturns = struct {
		result1 []notifications.Address
	}{result1}
}

func (fake *FakeAddressBook) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addressForRepoMutex.RLock()
	defer fake.addressForRepoMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAddressBook) recordInvocation(key string, args []interface{}) {
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

var _ notifications.AddressBook = new(FakeAddressBook)