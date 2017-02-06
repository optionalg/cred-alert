// This file was generated by counterfeiter
package notificationsfakes

import (
	"cred-alert/notifications"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeRouter struct {
	DeliverStub        func(logger lager.Logger, batch []notifications.Notification) error
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		logger lager.Logger
		batch  []notifications.Notification
	}
	deliverReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouter) Deliver(logger lager.Logger, batch []notifications.Notification) error {
	var batchCopy []notifications.Notification
	if batch != nil {
		batchCopy = make([]notifications.Notification, len(batch))
		copy(batchCopy, batch)
	}
	fake.deliverMutex.Lock()
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		logger lager.Logger
		batch  []notifications.Notification
	}{logger, batchCopy})
	fake.recordInvocation("Deliver", []interface{}{logger, batchCopy})
	fake.deliverMutex.Unlock()
	if fake.DeliverStub != nil {
		return fake.DeliverStub(logger, batch)
	}
	return fake.deliverReturns.result1
}

func (fake *FakeRouter) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *FakeRouter) DeliverArgsForCall(i int) (lager.Logger, []notifications.Notification) {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return fake.deliverArgsForCall[i].logger, fake.deliverArgsForCall[i].batch
}

func (fake *FakeRouter) DeliverReturns(result1 error) {
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRouter) recordInvocation(key string, args []interface{}) {
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

var _ notifications.Router = new(FakeRouter)
