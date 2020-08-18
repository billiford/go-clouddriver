// Code generated by counterfeiter. DO NOT EDIT.
package sqlfakes

import (
	"sync"

	"github.com/billiford/go-clouddriver/pkg/kubernetes"
	"github.com/billiford/go-clouddriver/pkg/sql"
	_ "github.com/mattn/go-sqlite3"
)

type FakeClient struct {
	CreateKubernetesProviderStub        func(kubernetes.Provider) error
	createKubernetesProviderMutex       sync.RWMutex
	createKubernetesProviderArgsForCall []struct {
		arg1 kubernetes.Provider
	}
	createKubernetesProviderReturns struct {
		result1 error
	}
	createKubernetesProviderReturnsOnCall map[int]struct {
		result1 error
	}
	GetKubernetesProviderStub        func(string) (kubernetes.Provider, error)
	getKubernetesProviderMutex       sync.RWMutex
	getKubernetesProviderArgsForCall []struct {
		arg1 string
	}
	getKubernetesProviderReturns struct {
		result1 kubernetes.Provider
		result2 error
	}
	getKubernetesProviderReturnsOnCall map[int]struct {
		result1 kubernetes.Provider
		result2 error
	}
	CreateKubernetesResourceStub        func(kubernetes.Resource) error
	createKubernetesResourceMutex       sync.RWMutex
	createKubernetesResourceArgsForCall []struct {
		arg1 kubernetes.Resource
	}
	createKubernetesResourceReturns struct {
		result1 error
	}
	createKubernetesResourceReturnsOnCall map[int]struct {
		result1 error
	}
	ListKubernetesResourcesStub        func(string) ([]kubernetes.Resource, error)
	listKubernetesResourcesMutex       sync.RWMutex
	listKubernetesResourcesArgsForCall []struct {
		arg1 string
	}
	listKubernetesResourcesReturns struct {
		result1 []kubernetes.Resource
		result2 error
	}
	listKubernetesResourcesReturnsOnCall map[int]struct {
		result1 []kubernetes.Resource
		result2 error
	}
	ListKubernetesAccountsBySpinnakerAppStub        func(string) ([]string, error)
	listKubernetesAccountsBySpinnakerAppMutex       sync.RWMutex
	listKubernetesAccountsBySpinnakerAppArgsForCall []struct {
		arg1 string
	}
	listKubernetesAccountsBySpinnakerAppReturns struct {
		result1 []string
		result2 error
	}
	listKubernetesAccountsBySpinnakerAppReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateKubernetesProvider(arg1 kubernetes.Provider) error {
	fake.createKubernetesProviderMutex.Lock()
	ret, specificReturn := fake.createKubernetesProviderReturnsOnCall[len(fake.createKubernetesProviderArgsForCall)]
	fake.createKubernetesProviderArgsForCall = append(fake.createKubernetesProviderArgsForCall, struct {
		arg1 kubernetes.Provider
	}{arg1})
	fake.recordInvocation("CreateKubernetesProvider", []interface{}{arg1})
	fake.createKubernetesProviderMutex.Unlock()
	if fake.CreateKubernetesProviderStub != nil {
		return fake.CreateKubernetesProviderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createKubernetesProviderReturns.result1
}

func (fake *FakeClient) CreateKubernetesProviderCallCount() int {
	fake.createKubernetesProviderMutex.RLock()
	defer fake.createKubernetesProviderMutex.RUnlock()
	return len(fake.createKubernetesProviderArgsForCall)
}

func (fake *FakeClient) CreateKubernetesProviderArgsForCall(i int) kubernetes.Provider {
	fake.createKubernetesProviderMutex.RLock()
	defer fake.createKubernetesProviderMutex.RUnlock()
	return fake.createKubernetesProviderArgsForCall[i].arg1
}

func (fake *FakeClient) CreateKubernetesProviderReturns(result1 error) {
	fake.CreateKubernetesProviderStub = nil
	fake.createKubernetesProviderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateKubernetesProviderReturnsOnCall(i int, result1 error) {
	fake.CreateKubernetesProviderStub = nil
	if fake.createKubernetesProviderReturnsOnCall == nil {
		fake.createKubernetesProviderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createKubernetesProviderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetKubernetesProvider(arg1 string) (kubernetes.Provider, error) {
	fake.getKubernetesProviderMutex.Lock()
	ret, specificReturn := fake.getKubernetesProviderReturnsOnCall[len(fake.getKubernetesProviderArgsForCall)]
	fake.getKubernetesProviderArgsForCall = append(fake.getKubernetesProviderArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetKubernetesProvider", []interface{}{arg1})
	fake.getKubernetesProviderMutex.Unlock()
	if fake.GetKubernetesProviderStub != nil {
		return fake.GetKubernetesProviderStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getKubernetesProviderReturns.result1, fake.getKubernetesProviderReturns.result2
}

func (fake *FakeClient) GetKubernetesProviderCallCount() int {
	fake.getKubernetesProviderMutex.RLock()
	defer fake.getKubernetesProviderMutex.RUnlock()
	return len(fake.getKubernetesProviderArgsForCall)
}

func (fake *FakeClient) GetKubernetesProviderArgsForCall(i int) string {
	fake.getKubernetesProviderMutex.RLock()
	defer fake.getKubernetesProviderMutex.RUnlock()
	return fake.getKubernetesProviderArgsForCall[i].arg1
}

func (fake *FakeClient) GetKubernetesProviderReturns(result1 kubernetes.Provider, result2 error) {
	fake.GetKubernetesProviderStub = nil
	fake.getKubernetesProviderReturns = struct {
		result1 kubernetes.Provider
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetKubernetesProviderReturnsOnCall(i int, result1 kubernetes.Provider, result2 error) {
	fake.GetKubernetesProviderStub = nil
	if fake.getKubernetesProviderReturnsOnCall == nil {
		fake.getKubernetesProviderReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Provider
			result2 error
		})
	}
	fake.getKubernetesProviderReturnsOnCall[i] = struct {
		result1 kubernetes.Provider
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateKubernetesResource(arg1 kubernetes.Resource) error {
	fake.createKubernetesResourceMutex.Lock()
	ret, specificReturn := fake.createKubernetesResourceReturnsOnCall[len(fake.createKubernetesResourceArgsForCall)]
	fake.createKubernetesResourceArgsForCall = append(fake.createKubernetesResourceArgsForCall, struct {
		arg1 kubernetes.Resource
	}{arg1})
	fake.recordInvocation("CreateKubernetesResource", []interface{}{arg1})
	fake.createKubernetesResourceMutex.Unlock()
	if fake.CreateKubernetesResourceStub != nil {
		return fake.CreateKubernetesResourceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createKubernetesResourceReturns.result1
}

func (fake *FakeClient) CreateKubernetesResourceCallCount() int {
	fake.createKubernetesResourceMutex.RLock()
	defer fake.createKubernetesResourceMutex.RUnlock()
	return len(fake.createKubernetesResourceArgsForCall)
}

func (fake *FakeClient) CreateKubernetesResourceArgsForCall(i int) kubernetes.Resource {
	fake.createKubernetesResourceMutex.RLock()
	defer fake.createKubernetesResourceMutex.RUnlock()
	return fake.createKubernetesResourceArgsForCall[i].arg1
}

func (fake *FakeClient) CreateKubernetesResourceReturns(result1 error) {
	fake.CreateKubernetesResourceStub = nil
	fake.createKubernetesResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CreateKubernetesResourceReturnsOnCall(i int, result1 error) {
	fake.CreateKubernetesResourceStub = nil
	if fake.createKubernetesResourceReturnsOnCall == nil {
		fake.createKubernetesResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createKubernetesResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ListKubernetesResources(arg1 string) ([]kubernetes.Resource, error) {
	fake.listKubernetesResourcesMutex.Lock()
	ret, specificReturn := fake.listKubernetesResourcesReturnsOnCall[len(fake.listKubernetesResourcesArgsForCall)]
	fake.listKubernetesResourcesArgsForCall = append(fake.listKubernetesResourcesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListKubernetesResources", []interface{}{arg1})
	fake.listKubernetesResourcesMutex.Unlock()
	if fake.ListKubernetesResourcesStub != nil {
		return fake.ListKubernetesResourcesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listKubernetesResourcesReturns.result1, fake.listKubernetesResourcesReturns.result2
}

func (fake *FakeClient) ListKubernetesResourcesCallCount() int {
	fake.listKubernetesResourcesMutex.RLock()
	defer fake.listKubernetesResourcesMutex.RUnlock()
	return len(fake.listKubernetesResourcesArgsForCall)
}

func (fake *FakeClient) ListKubernetesResourcesArgsForCall(i int) string {
	fake.listKubernetesResourcesMutex.RLock()
	defer fake.listKubernetesResourcesMutex.RUnlock()
	return fake.listKubernetesResourcesArgsForCall[i].arg1
}

func (fake *FakeClient) ListKubernetesResourcesReturns(result1 []kubernetes.Resource, result2 error) {
	fake.ListKubernetesResourcesStub = nil
	fake.listKubernetesResourcesReturns = struct {
		result1 []kubernetes.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListKubernetesResourcesReturnsOnCall(i int, result1 []kubernetes.Resource, result2 error) {
	fake.ListKubernetesResourcesStub = nil
	if fake.listKubernetesResourcesReturnsOnCall == nil {
		fake.listKubernetesResourcesReturnsOnCall = make(map[int]struct {
			result1 []kubernetes.Resource
			result2 error
		})
	}
	fake.listKubernetesResourcesReturnsOnCall[i] = struct {
		result1 []kubernetes.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListKubernetesAccountsBySpinnakerApp(arg1 string) ([]string, error) {
	fake.listKubernetesAccountsBySpinnakerAppMutex.Lock()
	ret, specificReturn := fake.listKubernetesAccountsBySpinnakerAppReturnsOnCall[len(fake.listKubernetesAccountsBySpinnakerAppArgsForCall)]
	fake.listKubernetesAccountsBySpinnakerAppArgsForCall = append(fake.listKubernetesAccountsBySpinnakerAppArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListKubernetesAccountsBySpinnakerApp", []interface{}{arg1})
	fake.listKubernetesAccountsBySpinnakerAppMutex.Unlock()
	if fake.ListKubernetesAccountsBySpinnakerAppStub != nil {
		return fake.ListKubernetesAccountsBySpinnakerAppStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listKubernetesAccountsBySpinnakerAppReturns.result1, fake.listKubernetesAccountsBySpinnakerAppReturns.result2
}

func (fake *FakeClient) ListKubernetesAccountsBySpinnakerAppCallCount() int {
	fake.listKubernetesAccountsBySpinnakerAppMutex.RLock()
	defer fake.listKubernetesAccountsBySpinnakerAppMutex.RUnlock()
	return len(fake.listKubernetesAccountsBySpinnakerAppArgsForCall)
}

func (fake *FakeClient) ListKubernetesAccountsBySpinnakerAppArgsForCall(i int) string {
	fake.listKubernetesAccountsBySpinnakerAppMutex.RLock()
	defer fake.listKubernetesAccountsBySpinnakerAppMutex.RUnlock()
	return fake.listKubernetesAccountsBySpinnakerAppArgsForCall[i].arg1
}

func (fake *FakeClient) ListKubernetesAccountsBySpinnakerAppReturns(result1 []string, result2 error) {
	fake.ListKubernetesAccountsBySpinnakerAppStub = nil
	fake.listKubernetesAccountsBySpinnakerAppReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListKubernetesAccountsBySpinnakerAppReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ListKubernetesAccountsBySpinnakerAppStub = nil
	if fake.listKubernetesAccountsBySpinnakerAppReturnsOnCall == nil {
		fake.listKubernetesAccountsBySpinnakerAppReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listKubernetesAccountsBySpinnakerAppReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createKubernetesProviderMutex.RLock()
	defer fake.createKubernetesProviderMutex.RUnlock()
	fake.getKubernetesProviderMutex.RLock()
	defer fake.getKubernetesProviderMutex.RUnlock()
	fake.createKubernetesResourceMutex.RLock()
	defer fake.createKubernetesResourceMutex.RUnlock()
	fake.listKubernetesResourcesMutex.RLock()
	defer fake.listKubernetesResourcesMutex.RUnlock()
	fake.listKubernetesAccountsBySpinnakerAppMutex.RLock()
	defer fake.listKubernetesAccountsBySpinnakerAppMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ sql.Client = new(FakeClient)