// Code generated by counterfeiter. DO NOT EDIT.
package kawasakifakes

import (
	"net"
	"sync"

	"code.cloudfoundry.org/guardian/kawasaki"
)

type FakeResolvCompiler struct {
	DetermineStub        func(string, net.IP, []net.IP, []net.IP, []net.IP, []string) []string
	determineMutex       sync.RWMutex
	determineArgsForCall []struct {
		arg1 string
		arg2 net.IP
		arg3 []net.IP
		arg4 []net.IP
		arg5 []net.IP
		arg6 []string
	}
	determineReturns struct {
		result1 []string
	}
	determineReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResolvCompiler) Determine(arg1 string, arg2 net.IP, arg3 []net.IP, arg4 []net.IP, arg5 []net.IP, arg6 []string) []string {
	var arg3Copy []net.IP
	if arg3 != nil {
		arg3Copy = make([]net.IP, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []net.IP
	if arg4 != nil {
		arg4Copy = make([]net.IP, len(arg4))
		copy(arg4Copy, arg4)
	}
	var arg5Copy []net.IP
	if arg5 != nil {
		arg5Copy = make([]net.IP, len(arg5))
		copy(arg5Copy, arg5)
	}
	var arg6Copy []string
	if arg6 != nil {
		arg6Copy = make([]string, len(arg6))
		copy(arg6Copy, arg6)
	}
	fake.determineMutex.Lock()
	ret, specificReturn := fake.determineReturnsOnCall[len(fake.determineArgsForCall)]
	fake.determineArgsForCall = append(fake.determineArgsForCall, struct {
		arg1 string
		arg2 net.IP
		arg3 []net.IP
		arg4 []net.IP
		arg5 []net.IP
		arg6 []string
	}{arg1, arg2, arg3Copy, arg4Copy, arg5Copy, arg6Copy})
	stub := fake.DetermineStub
	fakeReturns := fake.determineReturns
	fake.recordInvocation("Determine", []interface{}{arg1, arg2, arg3Copy, arg4Copy, arg5Copy, arg6Copy})
	fake.determineMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResolvCompiler) DetermineCallCount() int {
	fake.determineMutex.RLock()
	defer fake.determineMutex.RUnlock()
	return len(fake.determineArgsForCall)
}

func (fake *FakeResolvCompiler) DetermineCalls(stub func(string, net.IP, []net.IP, []net.IP, []net.IP, []string) []string) {
	fake.determineMutex.Lock()
	defer fake.determineMutex.Unlock()
	fake.DetermineStub = stub
}

func (fake *FakeResolvCompiler) DetermineArgsForCall(i int) (string, net.IP, []net.IP, []net.IP, []net.IP, []string) {
	fake.determineMutex.RLock()
	defer fake.determineMutex.RUnlock()
	argsForCall := fake.determineArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeResolvCompiler) DetermineReturns(result1 []string) {
	fake.determineMutex.Lock()
	defer fake.determineMutex.Unlock()
	fake.DetermineStub = nil
	fake.determineReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeResolvCompiler) DetermineReturnsOnCall(i int, result1 []string) {
	fake.determineMutex.Lock()
	defer fake.determineMutex.Unlock()
	fake.DetermineStub = nil
	if fake.determineReturnsOnCall == nil {
		fake.determineReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.determineReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeResolvCompiler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.determineMutex.RLock()
	defer fake.determineMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResolvCompiler) recordInvocation(key string, args []interface{}) {
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

var _ kawasaki.ResolvCompiler = new(FakeResolvCompiler)
