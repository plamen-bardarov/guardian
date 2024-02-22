// Code generated by counterfeiter. DO NOT EDIT.
package rundmcfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc"
)

type FakeCPUCgrouper struct {
	CreateBadCgroupStub        func(string) error
	createBadCgroupMutex       sync.RWMutex
	createBadCgroupArgsForCall []struct {
		arg1 string
	}
	createBadCgroupReturns struct {
		result1 error
	}
	createBadCgroupReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyBadCgroupStub        func(string) error
	destroyBadCgroupMutex       sync.RWMutex
	destroyBadCgroupArgsForCall []struct {
		arg1 string
	}
	destroyBadCgroupReturns struct {
		result1 error
	}
	destroyBadCgroupReturnsOnCall map[int]struct {
		result1 error
	}
	ReadBadCgroupUsageStub        func(string) (garden.ContainerCPUStat, error)
	readBadCgroupUsageMutex       sync.RWMutex
	readBadCgroupUsageArgsForCall []struct {
		arg1 string
	}
	readBadCgroupUsageReturns struct {
		result1 garden.ContainerCPUStat
		result2 error
	}
	readBadCgroupUsageReturnsOnCall map[int]struct {
		result1 garden.ContainerCPUStat
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCPUCgrouper) CreateBadCgroup(arg1 string) error {
	fake.createBadCgroupMutex.Lock()
	ret, specificReturn := fake.createBadCgroupReturnsOnCall[len(fake.createBadCgroupArgsForCall)]
	fake.createBadCgroupArgsForCall = append(fake.createBadCgroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateBadCgroupStub
	fakeReturns := fake.createBadCgroupReturns
	fake.recordInvocation("CreateBadCgroup", []interface{}{arg1})
	fake.createBadCgroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCPUCgrouper) CreateBadCgroupCallCount() int {
	fake.createBadCgroupMutex.RLock()
	defer fake.createBadCgroupMutex.RUnlock()
	return len(fake.createBadCgroupArgsForCall)
}

func (fake *FakeCPUCgrouper) CreateBadCgroupCalls(stub func(string) error) {
	fake.createBadCgroupMutex.Lock()
	defer fake.createBadCgroupMutex.Unlock()
	fake.CreateBadCgroupStub = stub
}

func (fake *FakeCPUCgrouper) CreateBadCgroupArgsForCall(i int) string {
	fake.createBadCgroupMutex.RLock()
	defer fake.createBadCgroupMutex.RUnlock()
	argsForCall := fake.createBadCgroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCPUCgrouper) CreateBadCgroupReturns(result1 error) {
	fake.createBadCgroupMutex.Lock()
	defer fake.createBadCgroupMutex.Unlock()
	fake.CreateBadCgroupStub = nil
	fake.createBadCgroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCPUCgrouper) CreateBadCgroupReturnsOnCall(i int, result1 error) {
	fake.createBadCgroupMutex.Lock()
	defer fake.createBadCgroupMutex.Unlock()
	fake.CreateBadCgroupStub = nil
	if fake.createBadCgroupReturnsOnCall == nil {
		fake.createBadCgroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createBadCgroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCPUCgrouper) DestroyBadCgroup(arg1 string) error {
	fake.destroyBadCgroupMutex.Lock()
	ret, specificReturn := fake.destroyBadCgroupReturnsOnCall[len(fake.destroyBadCgroupArgsForCall)]
	fake.destroyBadCgroupArgsForCall = append(fake.destroyBadCgroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DestroyBadCgroupStub
	fakeReturns := fake.destroyBadCgroupReturns
	fake.recordInvocation("DestroyBadCgroup", []interface{}{arg1})
	fake.destroyBadCgroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCPUCgrouper) DestroyBadCgroupCallCount() int {
	fake.destroyBadCgroupMutex.RLock()
	defer fake.destroyBadCgroupMutex.RUnlock()
	return len(fake.destroyBadCgroupArgsForCall)
}

func (fake *FakeCPUCgrouper) DestroyBadCgroupCalls(stub func(string) error) {
	fake.destroyBadCgroupMutex.Lock()
	defer fake.destroyBadCgroupMutex.Unlock()
	fake.DestroyBadCgroupStub = stub
}

func (fake *FakeCPUCgrouper) DestroyBadCgroupArgsForCall(i int) string {
	fake.destroyBadCgroupMutex.RLock()
	defer fake.destroyBadCgroupMutex.RUnlock()
	argsForCall := fake.destroyBadCgroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCPUCgrouper) DestroyBadCgroupReturns(result1 error) {
	fake.destroyBadCgroupMutex.Lock()
	defer fake.destroyBadCgroupMutex.Unlock()
	fake.DestroyBadCgroupStub = nil
	fake.destroyBadCgroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCPUCgrouper) DestroyBadCgroupReturnsOnCall(i int, result1 error) {
	fake.destroyBadCgroupMutex.Lock()
	defer fake.destroyBadCgroupMutex.Unlock()
	fake.DestroyBadCgroupStub = nil
	if fake.destroyBadCgroupReturnsOnCall == nil {
		fake.destroyBadCgroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyBadCgroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsage(arg1 string) (garden.ContainerCPUStat, error) {
	fake.readBadCgroupUsageMutex.Lock()
	ret, specificReturn := fake.readBadCgroupUsageReturnsOnCall[len(fake.readBadCgroupUsageArgsForCall)]
	fake.readBadCgroupUsageArgsForCall = append(fake.readBadCgroupUsageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReadBadCgroupUsageStub
	fakeReturns := fake.readBadCgroupUsageReturns
	fake.recordInvocation("ReadBadCgroupUsage", []interface{}{arg1})
	fake.readBadCgroupUsageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsageCallCount() int {
	fake.readBadCgroupUsageMutex.RLock()
	defer fake.readBadCgroupUsageMutex.RUnlock()
	return len(fake.readBadCgroupUsageArgsForCall)
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsageCalls(stub func(string) (garden.ContainerCPUStat, error)) {
	fake.readBadCgroupUsageMutex.Lock()
	defer fake.readBadCgroupUsageMutex.Unlock()
	fake.ReadBadCgroupUsageStub = stub
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsageArgsForCall(i int) string {
	fake.readBadCgroupUsageMutex.RLock()
	defer fake.readBadCgroupUsageMutex.RUnlock()
	argsForCall := fake.readBadCgroupUsageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsageReturns(result1 garden.ContainerCPUStat, result2 error) {
	fake.readBadCgroupUsageMutex.Lock()
	defer fake.readBadCgroupUsageMutex.Unlock()
	fake.ReadBadCgroupUsageStub = nil
	fake.readBadCgroupUsageReturns = struct {
		result1 garden.ContainerCPUStat
		result2 error
	}{result1, result2}
}

func (fake *FakeCPUCgrouper) ReadBadCgroupUsageReturnsOnCall(i int, result1 garden.ContainerCPUStat, result2 error) {
	fake.readBadCgroupUsageMutex.Lock()
	defer fake.readBadCgroupUsageMutex.Unlock()
	fake.ReadBadCgroupUsageStub = nil
	if fake.readBadCgroupUsageReturnsOnCall == nil {
		fake.readBadCgroupUsageReturnsOnCall = make(map[int]struct {
			result1 garden.ContainerCPUStat
			result2 error
		})
	}
	fake.readBadCgroupUsageReturnsOnCall[i] = struct {
		result1 garden.ContainerCPUStat
		result2 error
	}{result1, result2}
}

func (fake *FakeCPUCgrouper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBadCgroupMutex.RLock()
	defer fake.createBadCgroupMutex.RUnlock()
	fake.destroyBadCgroupMutex.RLock()
	defer fake.destroyBadCgroupMutex.RUnlock()
	fake.readBadCgroupUsageMutex.RLock()
	defer fake.readBadCgroupUsageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCPUCgrouper) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.CPUCgrouper = new(FakeCPUCgrouper)
