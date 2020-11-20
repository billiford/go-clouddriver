// Code generated by counterfeiter. DO NOT EDIT.
package kubernetesfakes

import (
	"sync"

	"github.com/homedepot/go-clouddriver/pkg/http/core/kubernetes"
)

type FakeActionHandler struct {
	NewCleanupArtifactsActionStub        func(kubernetes.ActionConfig) kubernetes.Action
	newCleanupArtifactsActionMutex       sync.RWMutex
	newCleanupArtifactsActionArgsForCall []struct {
		arg1 kubernetes.ActionConfig
	}
	newCleanupArtifactsActionReturns struct {
		result1 kubernetes.Action
	}
	newCleanupArtifactsActionReturnsOnCall map[int]struct {
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeActionHandler) NewCleanupArtifactsAction(arg1 kubernetes.ActionConfig) kubernetes.Action {
	fake.newCleanupArtifactsActionMutex.Lock()
	ret, specificReturn := fake.newCleanupArtifactsActionReturnsOnCall[len(fake.newCleanupArtifactsActionArgsForCall)]
	fake.newCleanupArtifactsActionArgsForCall = append(fake.newCleanupArtifactsActionArgsForCall, struct {
		arg1 kubernetes.ActionConfig
	}{arg1})
	fake.recordInvocation("NewCleanupArtifactsAction", []interface{}{arg1})
	fake.newCleanupArtifactsActionMutex.Unlock()
	if fake.NewCleanupArtifactsActionStub != nil {
		return fake.NewCleanupArtifactsActionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newCleanupArtifactsActionReturns.result1
}

func (fake *FakeActionHandler) NewCleanupArtifactsActionCallCount() int {
	fake.newCleanupArtifactsActionMutex.RLock()
	defer fake.newCleanupArtifactsActionMutex.RUnlock()
	return len(fake.newCleanupArtifactsActionArgsForCall)
}

func (fake *FakeActionHandler) NewCleanupArtifactsActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newCleanupArtifactsActionMutex.RLock()
	defer fake.newCleanupArtifactsActionMutex.RUnlock()
	return fake.newCleanupArtifactsActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewCleanupArtifactsActionReturns(result1 kubernetes.Action) {
	fake.NewCleanupArtifactsActionStub = nil
	fake.newCleanupArtifactsActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewCleanupArtifactsActionReturnsOnCall(i int, result1 kubernetes.Action) {
	fake.NewCleanupArtifactsActionStub = nil
	if fake.newCleanupArtifactsActionReturnsOnCall == nil {
		fake.newCleanupArtifactsActionReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Action
		})
	}
	fake.newCleanupArtifactsActionReturnsOnCall[i] = struct {
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
	return fake.newDeployManifestActionReturns.result1
}

func (fake *FakeActionHandler) NewDeployManifestActionCallCount() int {
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	return len(fake.newDeployManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewDeployManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	return fake.newDeployManifestActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewDeployManifestActionReturns(result1 kubernetes.Action) {
	fake.NewDeployManifestActionStub = nil
	fake.newDeployManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewDeployManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newDeleteManifestActionReturns.result1
}

func (fake *FakeActionHandler) NewDeleteManifestActionCallCount() int {
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	return len(fake.newDeleteManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewDeleteManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	return fake.newDeleteManifestActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewDeleteManifestActionReturns(result1 kubernetes.Action) {
	fake.NewDeleteManifestActionStub = nil
	fake.newDeleteManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewDeleteManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newRollingRestartActionReturns.result1
}

func (fake *FakeActionHandler) NewRollingRestartActionCallCount() int {
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	return len(fake.newRollingRestartActionArgsForCall)
}

func (fake *FakeActionHandler) NewRollingRestartActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	return fake.newRollingRestartActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewRollingRestartActionReturns(result1 kubernetes.Action) {
	fake.NewRollingRestartActionStub = nil
	fake.newRollingRestartActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollingRestartActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newRollbackActionReturns.result1
}

func (fake *FakeActionHandler) NewRollbackActionCallCount() int {
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	return len(fake.newRollbackActionArgsForCall)
}

func (fake *FakeActionHandler) NewRollbackActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	return fake.newRollbackActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewRollbackActionReturns(result1 kubernetes.Action) {
	fake.NewRollbackActionStub = nil
	fake.newRollbackActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRollbackActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newRunJobActionReturns.result1
}

func (fake *FakeActionHandler) NewRunJobActionCallCount() int {
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	return len(fake.newRunJobActionArgsForCall)
}

func (fake *FakeActionHandler) NewRunJobActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	return fake.newRunJobActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewRunJobActionReturns(result1 kubernetes.Action) {
	fake.NewRunJobActionStub = nil
	fake.newRunJobActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewRunJobActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newScaleManifestActionReturns.result1
}

func (fake *FakeActionHandler) NewScaleManifestActionCallCount() int {
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	return len(fake.newScaleManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewScaleManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	return fake.newScaleManifestActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewScaleManifestActionReturns(result1 kubernetes.Action) {
	fake.NewScaleManifestActionStub = nil
	fake.newScaleManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewScaleManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
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
	return fake.newPatchManifestActionReturns.result1
}

func (fake *FakeActionHandler) NewPatchManifestActionCallCount() int {
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
	return len(fake.newPatchManifestActionArgsForCall)
}

func (fake *FakeActionHandler) NewPatchManifestActionArgsForCall(i int) kubernetes.ActionConfig {
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
	return fake.newPatchManifestActionArgsForCall[i].arg1
}

func (fake *FakeActionHandler) NewPatchManifestActionReturns(result1 kubernetes.Action) {
	fake.NewPatchManifestActionStub = nil
	fake.newPatchManifestActionReturns = struct {
		result1 kubernetes.Action
	}{result1}
}

func (fake *FakeActionHandler) NewPatchManifestActionReturnsOnCall(i int, result1 kubernetes.Action) {
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

func (fake *FakeActionHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newCleanupArtifactsActionMutex.RLock()
	defer fake.newCleanupArtifactsActionMutex.RUnlock()
	fake.newDeployManifestActionMutex.RLock()
	defer fake.newDeployManifestActionMutex.RUnlock()
	fake.newDeleteManifestActionMutex.RLock()
	defer fake.newDeleteManifestActionMutex.RUnlock()
	fake.newRollingRestartActionMutex.RLock()
	defer fake.newRollingRestartActionMutex.RUnlock()
	fake.newRollbackActionMutex.RLock()
	defer fake.newRollbackActionMutex.RUnlock()
	fake.newRunJobActionMutex.RLock()
	defer fake.newRunJobActionMutex.RUnlock()
	fake.newScaleManifestActionMutex.RLock()
	defer fake.newScaleManifestActionMutex.RUnlock()
	fake.newPatchManifestActionMutex.RLock()
	defer fake.newPatchManifestActionMutex.RUnlock()
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
