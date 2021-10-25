// Code generated by counterfeiter. DO NOT EDIT.
package kubernetesfakes

import (
	"sync"

	"github.com/homedepot/go-clouddriver/internal/kubernetes"
	"k8s.io/client-go/rest"
)

type FakeController struct {
	NewClientStub        func(*rest.Config) (kubernetes.Client, error)
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		arg1 *rest.Config
	}
	newClientReturns struct {
		result1 kubernetes.Client
		result2 error
	}
	newClientReturnsOnCall map[int]struct {
		result1 kubernetes.Client
		result2 error
	}
	NewClientsetStub        func(*rest.Config) (kubernetes.Clientset, error)
	newClientsetMutex       sync.RWMutex
	newClientsetArgsForCall []struct {
		arg1 *rest.Config
	}
	newClientsetReturns struct {
		result1 kubernetes.Clientset
		result2 error
	}
	newClientsetReturnsOnCall map[int]struct {
		result1 kubernetes.Clientset
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeController) NewClient(arg1 *rest.Config) (kubernetes.Client, error) {
	fake.newClientMutex.Lock()
	ret, specificReturn := fake.newClientReturnsOnCall[len(fake.newClientArgsForCall)]
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewClientStub
	fakeReturns := fake.newClientReturns
	fake.recordInvocation("NewClient", []interface{}{arg1})
	fake.newClientMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeController) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

func (fake *FakeController) NewClientCalls(stub func(*rest.Config) (kubernetes.Client, error)) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = stub
}

func (fake *FakeController) NewClientArgsForCall(i int) *rest.Config {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	argsForCall := fake.newClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeController) NewClientReturns(result1 kubernetes.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 kubernetes.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeController) NewClientReturnsOnCall(i int, result1 kubernetes.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	if fake.newClientReturnsOnCall == nil {
		fake.newClientReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Client
			result2 error
		})
	}
	fake.newClientReturnsOnCall[i] = struct {
		result1 kubernetes.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeController) NewClientset(arg1 *rest.Config) (kubernetes.Clientset, error) {
	fake.newClientsetMutex.Lock()
	ret, specificReturn := fake.newClientsetReturnsOnCall[len(fake.newClientsetArgsForCall)]
	fake.newClientsetArgsForCall = append(fake.newClientsetArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewClientsetStub
	fakeReturns := fake.newClientsetReturns
	fake.recordInvocation("NewClientset", []interface{}{arg1})
	fake.newClientsetMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeController) NewClientsetCallCount() int {
	fake.newClientsetMutex.RLock()
	defer fake.newClientsetMutex.RUnlock()
	return len(fake.newClientsetArgsForCall)
}

func (fake *FakeController) NewClientsetCalls(stub func(*rest.Config) (kubernetes.Clientset, error)) {
	fake.newClientsetMutex.Lock()
	defer fake.newClientsetMutex.Unlock()
	fake.NewClientsetStub = stub
}

func (fake *FakeController) NewClientsetArgsForCall(i int) *rest.Config {
	fake.newClientsetMutex.RLock()
	defer fake.newClientsetMutex.RUnlock()
	argsForCall := fake.newClientsetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeController) NewClientsetReturns(result1 kubernetes.Clientset, result2 error) {
	fake.newClientsetMutex.Lock()
	defer fake.newClientsetMutex.Unlock()
	fake.NewClientsetStub = nil
	fake.newClientsetReturns = struct {
		result1 kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeController) NewClientsetReturnsOnCall(i int, result1 kubernetes.Clientset, result2 error) {
	fake.newClientsetMutex.Lock()
	defer fake.newClientsetMutex.Unlock()
	fake.NewClientsetStub = nil
	if fake.newClientsetReturnsOnCall == nil {
		fake.newClientsetReturnsOnCall = make(map[int]struct {
			result1 kubernetes.Clientset
			result2 error
		})
	}
	fake.newClientsetReturnsOnCall[i] = struct {
		result1 kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	fake.newClientsetMutex.RLock()
	defer fake.newClientsetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeController) recordInvocation(key string, args []interface{}) {
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

var _ kubernetes.Controller = new(FakeController)