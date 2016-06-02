// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/dadoo"
)

type FakePidGetter struct {
	PidStub        func(pidFilePath string) (int, error)
	pidMutex       sync.RWMutex
	pidArgsForCall []struct {
		pidFilePath string
	}
	pidReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakePidGetter) Pid(pidFilePath string) (int, error) {
	fake.pidMutex.Lock()
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct {
		pidFilePath string
	}{pidFilePath})
	fake.pidMutex.Unlock()
	if fake.PidStub != nil {
		return fake.PidStub(pidFilePath)
	} else {
		return fake.pidReturns.result1, fake.pidReturns.result2
	}
}

func (fake *FakePidGetter) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

func (fake *FakePidGetter) PidArgsForCall(i int) string {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return fake.pidArgsForCall[i].pidFilePath
}

func (fake *FakePidGetter) PidReturns(result1 int, result2 error) {
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ dadoo.PidGetter = new(FakePidGetter)
