// Code generated by counterfeiter. DO NOT EDIT.
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/kawasaki"
	lager "code.cloudfoundry.org/lager/v3"
)

type FakeFirewallOpener struct {
	BulkOpenStub        func(lager.Logger, string, string, []garden.NetOutRule) error
	bulkOpenMutex       sync.RWMutex
	bulkOpenArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 []garden.NetOutRule
	}
	bulkOpenReturns struct {
		result1 error
	}
	bulkOpenReturnsOnCall map[int]struct {
		result1 error
	}
	OpenStub        func(lager.Logger, string, string, garden.NetOutRule) error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.NetOutRule
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFirewallOpener) BulkOpen(arg1 lager.Logger, arg2 string, arg3 string, arg4 []garden.NetOutRule) error {
	var arg4Copy []garden.NetOutRule
	if arg4 != nil {
		arg4Copy = make([]garden.NetOutRule, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.bulkOpenMutex.Lock()
	ret, specificReturn := fake.bulkOpenReturnsOnCall[len(fake.bulkOpenArgsForCall)]
	fake.bulkOpenArgsForCall = append(fake.bulkOpenArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 []garden.NetOutRule
	}{arg1, arg2, arg3, arg4Copy})
	stub := fake.BulkOpenStub
	fakeReturns := fake.bulkOpenReturns
	fake.recordInvocation("BulkOpen", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.bulkOpenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFirewallOpener) BulkOpenCallCount() int {
	fake.bulkOpenMutex.RLock()
	defer fake.bulkOpenMutex.RUnlock()
	return len(fake.bulkOpenArgsForCall)
}

func (fake *FakeFirewallOpener) BulkOpenCalls(stub func(lager.Logger, string, string, []garden.NetOutRule) error) {
	fake.bulkOpenMutex.Lock()
	defer fake.bulkOpenMutex.Unlock()
	fake.BulkOpenStub = stub
}

func (fake *FakeFirewallOpener) BulkOpenArgsForCall(i int) (lager.Logger, string, string, []garden.NetOutRule) {
	fake.bulkOpenMutex.RLock()
	defer fake.bulkOpenMutex.RUnlock()
	argsForCall := fake.bulkOpenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeFirewallOpener) BulkOpenReturns(result1 error) {
	fake.bulkOpenMutex.Lock()
	defer fake.bulkOpenMutex.Unlock()
	fake.BulkOpenStub = nil
	fake.bulkOpenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirewallOpener) BulkOpenReturnsOnCall(i int, result1 error) {
	fake.bulkOpenMutex.Lock()
	defer fake.bulkOpenMutex.Unlock()
	fake.BulkOpenStub = nil
	if fake.bulkOpenReturnsOnCall == nil {
		fake.bulkOpenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bulkOpenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirewallOpener) Open(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.NetOutRule) error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.NetOutRule
	}{arg1, arg2, arg3, arg4})
	stub := fake.OpenStub
	fakeReturns := fake.openReturns
	fake.recordInvocation("Open", []interface{}{arg1, arg2, arg3, arg4})
	fake.openMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFirewallOpener) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeFirewallOpener) OpenCalls(stub func(lager.Logger, string, string, garden.NetOutRule) error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = stub
}

func (fake *FakeFirewallOpener) OpenArgsForCall(i int) (lager.Logger, string, string, garden.NetOutRule) {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	argsForCall := fake.openArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeFirewallOpener) OpenReturns(result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirewallOpener) OpenReturnsOnCall(i int, result1 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFirewallOpener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bulkOpenMutex.RLock()
	defer fake.bulkOpenMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFirewallOpener) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.FirewallOpener = new(FakeFirewallOpener)
