// Code generated by counterfeiter. DO NOT EDIT.
package kawasakifakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
	lager "code.cloudfoundry.org/lager/v3"
)

type FakeContainerConfigurer struct {
	ApplyStub        func(lager.Logger, kawasaki.NetworkConfig, int) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 lager.Logger
		arg2 kawasaki.NetworkConfig
		arg3 int
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerConfigurer) Apply(arg1 lager.Logger, arg2 kawasaki.NetworkConfig, arg3 int) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 lager.Logger
		arg2 kawasaki.NetworkConfig
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContainerConfigurer) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeContainerConfigurer) ApplyCalls(stub func(lager.Logger, kawasaki.NetworkConfig, int) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeContainerConfigurer) ApplyArgsForCall(i int) (lager.Logger, kawasaki.NetworkConfig, int) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeContainerConfigurer) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerConfigurer) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerConfigurer) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.ContainerConfigurer = new(FakeContainerConfigurer)
