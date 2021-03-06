// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeGetHealthCheckActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetApplicationProcessHealthChecksByNameAndSpaceStub        func(string, string) ([]v7action.ProcessHealthCheck, v7action.Warnings, error)
	getApplicationProcessHealthChecksByNameAndSpaceMutex       sync.RWMutex
	getApplicationProcessHealthChecksByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationProcessHealthChecksByNameAndSpaceReturns struct {
		result1 []v7action.ProcessHealthCheck
		result2 v7action.Warnings
		result3 error
	}
	getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall map[int]struct {
		result1 []v7action.ProcessHealthCheck
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetHealthCheckActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeGetHealthCheckActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeGetHealthCheckActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeGetHealthCheckActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeGetHealthCheckActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpace(arg1 string, arg2 string) ([]v7action.ProcessHealthCheck, v7action.Warnings, error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall[len(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall)]
	fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall = append(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationProcessHealthChecksByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationProcessHealthChecksByNameAndSpaceStub != nil {
		return fake.GetApplicationProcessHealthChecksByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationProcessHealthChecksByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceCallCount() int {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall)
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceCalls(stub func(string, string) ([]v7action.ProcessHealthCheck, v7action.Warnings, error)) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = stub
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationProcessHealthChecksByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceReturns(result1 []v7action.ProcessHealthCheck, result2 v7action.Warnings, result3 error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = nil
	fake.getApplicationProcessHealthChecksByNameAndSpaceReturns = struct {
		result1 []v7action.ProcessHealthCheck
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetHealthCheckActor) GetApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall(i int, result1 []v7action.ProcessHealthCheck, result2 v7action.Warnings, result3 error) {
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Lock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.Unlock()
	fake.GetApplicationProcessHealthChecksByNameAndSpaceStub = nil
	if fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 []v7action.ProcessHealthCheck
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationProcessHealthChecksByNameAndSpaceReturnsOnCall[i] = struct {
		result1 []v7action.ProcessHealthCheck
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeGetHealthCheckActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RLock()
	defer fake.getApplicationProcessHealthChecksByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGetHealthCheckActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.GetHealthCheckActor = new(FakeGetHealthCheckActor)
