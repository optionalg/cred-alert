// This file was generated by counterfeiter
package notificationsfakes

import (
	"cred-alert/notifications"
	"cred-alert/scanners"
	"sync"

	"github.com/pivotal-golang/lager"
)

type FakeNotifier struct {
	SendNotificationStub        func(logger lager.Logger, repository string, sha string, line scanners.Line, private bool) error
	sendNotificationMutex       sync.RWMutex
	sendNotificationArgsForCall []struct {
		logger     lager.Logger
		repository string
		sha        string
		line       scanners.Line
		private    bool
	}
	sendNotificationReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNotifier) SendNotification(logger lager.Logger, repository string, sha string, line scanners.Line, private bool) error {
	fake.sendNotificationMutex.Lock()
	fake.sendNotificationArgsForCall = append(fake.sendNotificationArgsForCall, struct {
		logger     lager.Logger
		repository string
		sha        string
		line       scanners.Line
		private    bool
	}{logger, repository, sha, line, private})
	fake.recordInvocation("SendNotification", []interface{}{logger, repository, sha, line, private})
	fake.sendNotificationMutex.Unlock()
	if fake.SendNotificationStub != nil {
		return fake.SendNotificationStub(logger, repository, sha, line, private)
	} else {
		return fake.sendNotificationReturns.result1
	}
}

func (fake *FakeNotifier) SendNotificationCallCount() int {
	fake.sendNotificationMutex.RLock()
	defer fake.sendNotificationMutex.RUnlock()
	return len(fake.sendNotificationArgsForCall)
}

func (fake *FakeNotifier) SendNotificationArgsForCall(i int) (lager.Logger, string, string, scanners.Line, bool) {
	fake.sendNotificationMutex.RLock()
	defer fake.sendNotificationMutex.RUnlock()
	return fake.sendNotificationArgsForCall[i].logger, fake.sendNotificationArgsForCall[i].repository, fake.sendNotificationArgsForCall[i].sha, fake.sendNotificationArgsForCall[i].line, fake.sendNotificationArgsForCall[i].private
}

func (fake *FakeNotifier) SendNotificationReturns(result1 error) {
	fake.SendNotificationStub = nil
	fake.sendNotificationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendNotificationMutex.RLock()
	defer fake.sendNotificationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNotifier) recordInvocation(key string, args []interface{}) {
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

var _ notifications.Notifier = new(FakeNotifier)
