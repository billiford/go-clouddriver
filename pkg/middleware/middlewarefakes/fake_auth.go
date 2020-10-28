// Code generated by counterfeiter. DO NOT EDIT.
package middlewarefakes

import (
	"sync"

	"github.com/billiford/go-clouddriver/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type FakeAuth struct {
	AuthApplicationStub        func() gin.HandlerFunc
	authApplicationMutex       sync.RWMutex
	authApplicationArgsForCall []struct {
	}
	authApplicationReturns struct {
		result1 gin.HandlerFunc
	}
	authApplicationReturnsOnCall map[int]struct {
		result1 gin.HandlerFunc
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuth) AuthApplication() gin.HandlerFunc {
	fake.authApplicationMutex.Lock()
	ret, specificReturn := fake.authApplicationReturnsOnCall[len(fake.authApplicationArgsForCall)]
	fake.authApplicationArgsForCall = append(fake.authApplicationArgsForCall, struct {
	}{})
	fake.recordInvocation("AuthApplication", []interface{}{})
	fake.authApplicationMutex.Unlock()
	if fake.AuthApplicationStub != nil {
		return fake.AuthApplicationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authApplicationReturns
	return fakeReturns.result1
}

func (fake *FakeAuth) AuthApplicationCallCount() int {
	fake.authApplicationMutex.RLock()
	defer fake.authApplicationMutex.RUnlock()
	return len(fake.authApplicationArgsForCall)
}

func (fake *FakeAuth) AuthApplicationCalls(stub func() gin.HandlerFunc) {
	fake.authApplicationMutex.Lock()
	defer fake.authApplicationMutex.Unlock()
	fake.AuthApplicationStub = stub
}

func (fake *FakeAuth) AuthApplicationReturns(result1 gin.HandlerFunc) {
	fake.authApplicationMutex.Lock()
	defer fake.authApplicationMutex.Unlock()
	fake.AuthApplicationStub = nil
	fake.authApplicationReturns = struct {
		result1 gin.HandlerFunc
	}{result1}
}

func (fake *FakeAuth) AuthApplicationReturnsOnCall(i int, result1 gin.HandlerFunc) {
	fake.authApplicationMutex.Lock()
	defer fake.authApplicationMutex.Unlock()
	fake.AuthApplicationStub = nil
	if fake.authApplicationReturnsOnCall == nil {
		fake.authApplicationReturnsOnCall = make(map[int]struct {
			result1 gin.HandlerFunc
		})
	}
	fake.authApplicationReturnsOnCall[i] = struct {
		result1 gin.HandlerFunc
	}{result1}
}

func (fake *FakeAuth) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authApplicationMutex.RLock()
	defer fake.authApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuth) recordInvocation(key string, args []interface{}) {
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

var _ middleware.Auth = new(FakeAuth)
