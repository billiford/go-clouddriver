// Code generated by counterfeiter. DO NOT EDIT.
package kubernetesfakes

import (
	"sync"

	"github.com/billiford/go-clouddriver/pkg/http/core/kubernetes"
)

type FakeActionHandler struct {
	NewDeleteManifestActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newDeleteManifestActionMutex       sync.RWMutex
	newDeleteManifestActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newDeleteManifestActionReturns struct {
		result1 kubernetes.Action
	}
	newDeleteManifestActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewDeployManifestActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newDeployManifestActionMutex       sync.RWMutex
	newDeployManifestActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newDeployManifestActionReturns struct {
		result1 kubernetes.Action
	}
	newDeployManifestActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewPatchManifestActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newPatchManifestActionMutex       sync.RWMutex
	newPatchManifestActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newPatchManifestActionReturns struct {
		result1 kubernetes.Action
	}
	newPatchManifestActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewRollbackActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newRollbackActionMutex       sync.RWMutex
	newRollbackActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newRollbackActionReturns struct {
		result1 kubernetes.Action
	}
	newRollbackActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewRollingRestartActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newRollingRestartActionMutex       sync.RWMutex
	newRollingRestartActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newRollingRestartActionReturns struct {
		result1 kubernetes.Action
	}
	newRollingRestartActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewRunJobActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newRunJobActionMutex       sync.RWMutex
	newRunJobActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newRunJobActionReturns struct {
		result1 kubernetes.Action
	}
	newRunJobActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	NewScaleManifestActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newScaleManifestActionMutex       sync.RWMutex
	newScaleManifestActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newScaleManifestActionReturns struct {
		result1 kubernetes.Action
	}
	newScaleManifestActionReturnsOnCall map[int]struct {
		result1 kubernetes.Action
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActionHandler) NewDeleteManifestAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newDeleteManifestActionMutex.Lock()
	ret, specificReturn := fake.newDeleteManifestActionReturnsOnCall[len(fake.newDeleteManifestActionArgsForCall)]
	fake.newDeleteManifestActionArgsForCall = append(fake.newDeleteManifestActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewDeleteManifestAction", []interface{}{arg1})
	fake.newDeleteManifestActionMutex.Unlock()
	if fake.NewDeleteManifestActionStub != nil {
		return fake.NewDeleteManifestActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newDeleteManifestActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewDeleteManifestActionCallCount() int {
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	return len(fake.newDeleteManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewDeleteManifestActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newDeleteManifestActionMutex.Lock()
	defer fake.newDeleteManifestActionMutex.Unlock()
	fake.NewDeleteManifestActionStub = stub
}

func (fake *FakeActionHandler) NewDeleteManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	argsForCall := fake.newDeleteManifestActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewDeleteManifestActionReturns(result1 kubernetes.Action) {
	fake.newDeleteManifestActionMutex.Lock()
	defer fake.newDeleteManifestActionMutex.Unlock()
	fake.NewDeleteManifestActionStub = nil
	fake.newDeleteManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewDeleteManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newDeleteManifestActionMutex.Lock()
	defer fake.newDeleteManifestActionMutex.Unlock()
	fake.NewDeleteManifestActionStub = nil
	if fake.newDeleteManifestActionReturnsOnCall == nil {
		fake.newDeleteManifestActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newDeleteManifestActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewDeployManifestAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newDeployManifestActionMutex.Lock()
	ret, specificReturn := fake.newDeployManifestActionReturnsOnCall[len(fake.newDeployManifestActionArgsForCall)]
	fake.newDeployManifestActionArgsForCall = append(fake.newDeployManifestActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewDeployManifestAction", []interface{}{arg1})
	fake.newDeployManifestActionMutex.Unlock()
	if fake.NewDeployManifestActionStub != nil {
		return fake.NewDeployManifestActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newDeployManifestActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewDeployManifestActionCallCount() int {
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	return len(fake.newDeployManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewDeployManifestActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newDeployManifestActionMutex.Lock()
	defer fake.newDeployManifestActionMutex.Unlock()
	fake.NewDeployManifestActionStub = stub
}

func (fake *FakeActionHandler) NewDeployManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	argsForCall := fake.newDeployManifestActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewDeployManifestActionReturns(result1 kubernetes.Action) {
	fake.newDeployManifestActionMutex.Lock()
	defer fake.newDeployManifestActionMutex.Unlock()
	fake.NewDeployManifestActionStub = nil
	fake.newDeployManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewDeployManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newDeployManifestActionMutex.Lock()
	defer fake.newDeployManifestActionMutex.Unlock()
	fake.NewDeployManifestActionStub = nil
	if fake.newDeployManifestActionReturnsOnCall == nil {
		fake.newDeployManifestActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newDeployManifestActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewPatchManifestAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newPatchManifestActionMutex.Lock()
	ret, specificReturn := fake.newPatchManifestActionReturnsOnCall[len(fake.newPatchManifestActionArgsForCall)]
	fake.newPatchManifestActionArgsForCall = append(fake.newPatchManifestActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewPatchManifestAction", []interface{}{arg1})
	fake.newPatchManifestActionMutex.Unlock()
	if fake.NewPatchManifestActionStub != nil {
		return fake.NewPatchManifestActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newPatchManifestActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewPatchManifestActionCallCount() int {
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
	return len(fake.newPatchManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewPatchManifestActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newPatchManifestActionMutex.Lock()
	defer fake.newPatchManifestActionMutex.Unlock()
	fake.NewPatchManifestActionStub = stub
}

func (fake *FakeActionHandler) NewPatchManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
	argsForCall := fake.newPatchManifestActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewPatchManifestActionReturns(result1 kubernetes.Action) {
	fake.newPatchManifestActionMutex.Lock()
	defer fake.newPatchManifestActionMutex.Unlock()
	fake.NewPatchManifestActionStub = nil
	fake.newPatchManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewPatchManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newPatchManifestActionMutex.Lock()
	defer fake.newPatchManifestActionMutex.Unlock()
	fake.NewPatchManifestActionStub = nil
	if fake.newPatchManifestActionReturnsOnCall == nil {
		fake.newPatchManifestActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newPatchManifestActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollbackAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newRollbackActionMutex.Lock()
	ret, specificReturn := fake.newRollbackActionReturnsOnCall[len(fake.newRollbackActionArgsForCall)]
	fake.newRollbackActionArgsForCall = append(fake.newRollbackActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewRollbackAction", []interface{}{arg1})
	fake.newRollbackActionMutex.Unlock()
	if fake.NewRollbackActionStub != nil {
		return fake.NewRollbackActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newRollbackActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewRollbackActionCallCount() int {
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	return len(fake.newRollbackActionArgsForCall)
}

func (fake *FakeActionHandler) NewRollbackActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newRollbackActionMutex.Lock()
	defer fake.newRollbackActionMutex.Unlock()
	fake.NewRollbackActionStub = stub
}

func (fake *FakeActionHandler) NewRollbackActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	argsForCall := fake.newRollbackActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewRollbackActionReturns(result1 kubernetes.Action) {
	fake.newRollbackActionMutex.Lock()
	defer fake.newRollbackActionMutex.Unlock()
	fake.NewRollbackActionStub = nil
	fake.newRollbackActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollbackActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newRollbackActionMutex.Lock()
	defer fake.newRollbackActionMutex.Unlock()
	fake.NewRollbackActionStub = nil
	if fake.newRollbackActionReturnsOnCall == nil {
		fake.newRollbackActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newRollbackActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollingRestartAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newRollingRestartActionMutex.Lock()
	ret, specificReturn := fake.newRollingRestartActionReturnsOnCall[len(fake.newRollingRestartActionArgsForCall)]
	fake.newRollingRestartActionArgsForCall = append(fake.newRollingRestartActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewRollingRestartAction", []interface{}{arg1})
	fake.newRollingRestartActionMutex.Unlock()
	if fake.NewRollingRestartActionStub != nil {
		return fake.NewRollingRestartActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newRollingRestartActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewRollingRestartActionCallCount() int {
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	return len(fake.newRollingRestartActionArgsForCall)
}

func (fake *FakeActionHandler) NewRollingRestartActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newRollingRestartActionMutex.Lock()
	defer fake.newRollingRestartActionMutex.Unlock()
	fake.NewRollingRestartActionStub = stub
}

func (fake *FakeActionHandler) NewRollingRestartActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	argsForCall := fake.newRollingRestartActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewRollingRestartActionReturns(result1 kubernetes.Action) {
	fake.newRollingRestartActionMutex.Lock()
	defer fake.newRollingRestartActionMutex.Unlock()
	fake.NewRollingRestartActionStub = nil
	fake.newRollingRestartActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollingRestartActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newRollingRestartActionMutex.Lock()
	defer fake.newRollingRestartActionMutex.Unlock()
	fake.NewRollingRestartActionStub = nil
	if fake.newRollingRestartActionReturnsOnCall == nil {
		fake.newRollingRestartActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newRollingRestartActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRunJobAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newRunJobActionMutex.Lock()
	ret, specificReturn := fake.newRunJobActionReturnsOnCall[len(fake.newRunJobActionArgsForCall)]
	fake.newRunJobActionArgsForCall = append(fake.newRunJobActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewRunJobAction", []interface{}{arg1})
	fake.newRunJobActionMutex.Unlock()
	if fake.NewRunJobActionStub != nil {
		return fake.NewRunJobActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newRunJobActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewRunJobActionCallCount() int {
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	return len(fake.newRunJobActionArgsForCall)
}

func (fake *FakeActionHandler) NewRunJobActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newRunJobActionMutex.Lock()
	defer fake.newRunJobActionMutex.Unlock()
	fake.NewRunJobActionStub = stub
}

func (fake *FakeActionHandler) NewRunJobActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	argsForCall := fake.newRunJobActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewRunJobActionReturns(result1 kubernetes.Action) {
	fake.newRunJobActionMutex.Lock()
	defer fake.newRunJobActionMutex.Unlock()
	fake.NewRunJobActionStub = nil
	fake.newRunJobActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRunJobActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newRunJobActionMutex.Lock()
	defer fake.newRunJobActionMutex.Unlock()
	fake.NewRunJobActionStub = nil
	if fake.newRunJobActionReturnsOnCall == nil {
		fake.newRunJobActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newRunJobActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewScaleManifestAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newScaleManifestActionMutex.Lock()
	ret, specificReturn := fake.newScaleManifestActionReturnsOnCall[len(fake.newScaleManifestActionArgsForCall)]
	fake.newScaleManifestActionArgsForCall = append(fake.newScaleManifestActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewScaleManifestAction", []interface{}{arg1})
	fake.newScaleManifestActionMutex.Unlock()
	if fake.NewScaleManifestActionStub != nil {
		return fake.NewScaleManifestActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newScaleManifestActionReturns
	return fakeReturns.result1
}

func (fake *FakeActionHandler) NewScaleManifestActionCallCount() int {
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	return len(fake.newScaleManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewScaleManifestActionCalls(stub func(kubernetes.ActionConfig) kubernetes.Action) {
	fake.newScaleManifestActionMutex.Lock()
	defer fake.newScaleManifestActionMutex.Unlock()
	fake.NewScaleManifestActionStub = stub
}

func (fake *FakeActionHandler) NewScaleManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	argsForCall := fake.newScaleManifestActionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeActionHandler) NewScaleManifestActionReturns(result1 kubernetes.Action) {
	fake.newScaleManifestActionMutex.Lock()
	defer fake.newScaleManifestActionMutex.Unlock()
	fake.NewScaleManifestActionStub = nil
	fake.newScaleManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewScaleManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.newScaleManifestActionMutex.Lock()
	defer fake.newScaleManifestActionMutex.Unlock()
	fake.NewScaleManifestActionStub = nil
	if fake.newScaleManifestActionReturnsOnCall == nil {
		fake.newScaleManifestActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newScaleManifestActionReturnsOnCall[i] = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeActionHandler) recordInvocation(key string, args []interface{}) {
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

var _ kubernetes.ActionHandler = new(FakeActionHandler)
