// Code generated by counterfeiter. DO NOT EDIT.
package artifactfakes

import (
	"net/http"
	"sync"

	"github.com/billiford/go-clouddriver/pkg/artifact"
	"github.com/billiford/go-clouddriver/pkg/helm"
	"github.com/google/go-github/v32/github"
)

type FakeCredentialsController struct {
	GitClientForAccountNameStub        func(string) (*github.Client, error)
	gitClientForAccountNameMutex       sync.RWMutex
	gitClientForAccountNameArgsForCall []struct {
		arg1 string
	}
	gitClientForAccountNameReturns struct {
		result1 *github.Client
		result2 error
	}
	gitClientForAccountNameReturnsOnCall map[int]struct {
		result1 *github.Client
		result2 error
	}
	GitRepoClientForAccountNameStub        func(string) (*http.Client, error)
	gitRepoClientForAccountNameMutex       sync.RWMutex
	gitRepoClientForAccountNameArgsForCall []struct {
		arg1 string
	}
	gitRepoClientForAccountNameReturns struct {
		result1 *http.Client
		result2 error
	}
	gitRepoClientForAccountNameReturnsOnCall map[int]struct {
		result1 *http.Client
		result2 error
	}
	HTTPClientForAccountNameStub        func(string) (*http.Client, error)
	hTTPClientForAccountNameMutex       sync.RWMutex
	hTTPClientForAccountNameArgsForCall []struct {
		arg1 string
	}
	hTTPClientForAccountNameReturns struct {
		result1 *http.Client
		result2 error
	}
	hTTPClientForAccountNameReturnsOnCall map[int]struct {
		result1 *http.Client
		result2 error
	}
	HelmClientForAccountNameStub        func(string) (helm.Client, error)
	helmClientForAccountNameMutex       sync.RWMutex
	helmClientForAccountNameArgsForCall []struct {
		arg1 string
	}
	helmClientForAccountNameReturns struct {
		result1 helm.Client
		result2 error
	}
	helmClientForAccountNameReturnsOnCall map[int]struct {
		result1 helm.Client
		result2 error
	}
	ListArtifactCredentialsNamesAndTypesStub        func() []artifact.Credentials
	listArtifactCredentialsNamesAndTypesMutex       sync.RWMutex
	listArtifactCredentialsNamesAndTypesArgsForCall []struct {
	}
	listArtifactCredentialsNamesAndTypesReturns struct {
		result1 []artifact.Credentials
	}
	listArtifactCredentialsNamesAndTypesReturnsOnCall map[int]struct {
		result1 []artifact.Credentials
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredentialsController) GitClientForAccountName(arg1 string) (*github.Client, error) {
	fake.gitClientForAccountNameMutex.Lock()
	ret, specificReturn := fake.gitClientForAccountNameReturnsOnCall[len(fake.gitClientForAccountNameArgsForCall)]
	fake.gitClientForAccountNameArgsForCall = append(fake.gitClientForAccountNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GitClientForAccountName", []interface{}{arg1})
	fake.gitClientForAccountNameMutex.Unlock()
	if fake.GitClientForAccountNameStub != nil {
		return fake.GitClientForAccountNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.gitClientForAccountNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredentialsController) GitClientForAccountNameCallCount() int {
	fake.gitClientForAccountNameMutex.RLock()
	defer fake.gitClientForAccountNameMutex.RUnlock()
	return len(fake.gitClientForAccountNameArgsForCall)
}

func (fake *FakeCredentialsController) GitClientForAccountNameCalls(stub func(string) (*github.Client, error)) {
	fake.gitClientForAccountNameMutex.Lock()
	defer fake.gitClientForAccountNameMutex.Unlock()
	fake.GitClientForAccountNameStub = stub
}

func (fake *FakeCredentialsController) GitClientForAccountNameArgsForCall(i int) string {
	fake.gitClientForAccountNameMutex.RLock()
	defer fake.gitClientForAccountNameMutex.RUnlock()
	argsForCall := fake.gitClientForAccountNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredentialsController) GitClientForAccountNameReturns(result1 *github.Client, result2 error) {
	fake.gitClientForAccountNameMutex.Lock()
	defer fake.gitClientForAccountNameMutex.Unlock()
	fake.GitClientForAccountNameStub = nil
	fake.gitClientForAccountNameReturns = struct {
		result1 *github.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) GitClientForAccountNameReturnsOnCall(i int, result1 *github.Client, result2 error) {
	fake.gitClientForAccountNameMutex.Lock()
	defer fake.gitClientForAccountNameMutex.Unlock()
	fake.GitClientForAccountNameStub = nil
	if fake.gitClientForAccountNameReturnsOnCall == nil {
		fake.gitClientForAccountNameReturnsOnCall = make(map[int]struct {
			result1 *github.Client
			result2 error
		})
	}
	fake.gitClientForAccountNameReturnsOnCall[i] = struct {
		result1 *github.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) GitRepoClientForAccountName(arg1 string) (*http.Client, error) {
	fake.gitRepoClientForAccountNameMutex.Lock()
	ret, specificReturn := fake.gitRepoClientForAccountNameReturnsOnCall[len(fake.gitRepoClientForAccountNameArgsForCall)]
	fake.gitRepoClientForAccountNameArgsForCall = append(fake.gitRepoClientForAccountNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GitRepoClientForAccountName", []interface{}{arg1})
	fake.gitRepoClientForAccountNameMutex.Unlock()
	if fake.GitRepoClientForAccountNameStub != nil {
		return fake.GitRepoClientForAccountNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.gitRepoClientForAccountNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredentialsController) GitRepoClientForAccountNameCallCount() int {
	fake.gitRepoClientForAccountNameMutex.RLock()
	defer fake.gitRepoClientForAccountNameMutex.RUnlock()
	return len(fake.gitRepoClientForAccountNameArgsForCall)
}

func (fake *FakeCredentialsController) GitRepoClientForAccountNameCalls(stub func(string) (*http.Client, error)) {
	fake.gitRepoClientForAccountNameMutex.Lock()
	defer fake.gitRepoClientForAccountNameMutex.Unlock()
	fake.GitRepoClientForAccountNameStub = stub
}

func (fake *FakeCredentialsController) GitRepoClientForAccountNameArgsForCall(i int) string {
	fake.gitRepoClientForAccountNameMutex.RLock()
	defer fake.gitRepoClientForAccountNameMutex.RUnlock()
	argsForCall := fake.gitRepoClientForAccountNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredentialsController) GitRepoClientForAccountNameReturns(result1 *http.Client, result2 error) {
	fake.gitRepoClientForAccountNameMutex.Lock()
	defer fake.gitRepoClientForAccountNameMutex.Unlock()
	fake.GitRepoClientForAccountNameStub = nil
	fake.gitRepoClientForAccountNameReturns = struct {
		result1 *http.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) GitRepoClientForAccountNameReturnsOnCall(i int, result1 *http.Client, result2 error) {
	fake.gitRepoClientForAccountNameMutex.Lock()
	defer fake.gitRepoClientForAccountNameMutex.Unlock()
	fake.GitRepoClientForAccountNameStub = nil
	if fake.gitRepoClientForAccountNameReturnsOnCall == nil {
		fake.gitRepoClientForAccountNameReturnsOnCall = make(map[int]struct {
			result1 *http.Client
			result2 error
		})
	}
	fake.gitRepoClientForAccountNameReturnsOnCall[i] = struct {
		result1 *http.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) HTTPClientForAccountName(arg1 string) (*http.Client, error) {
	fake.hTTPClientForAccountNameMutex.Lock()
	ret, specificReturn := fake.hTTPClientForAccountNameReturnsOnCall[len(fake.hTTPClientForAccountNameArgsForCall)]
	fake.hTTPClientForAccountNameArgsForCall = append(fake.hTTPClientForAccountNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HTTPClientForAccountName", []interface{}{arg1})
	fake.hTTPClientForAccountNameMutex.Unlock()
	if fake.HTTPClientForAccountNameStub != nil {
		return fake.HTTPClientForAccountNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.hTTPClientForAccountNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredentialsController) HTTPClientForAccountNameCallCount() int {
	fake.hTTPClientForAccountNameMutex.RLock()
	defer fake.hTTPClientForAccountNameMutex.RUnlock()
	return len(fake.hTTPClientForAccountNameArgsForCall)
}

func (fake *FakeCredentialsController) HTTPClientForAccountNameCalls(stub func(string) (*http.Client, error)) {
	fake.hTTPClientForAccountNameMutex.Lock()
	defer fake.hTTPClientForAccountNameMutex.Unlock()
	fake.HTTPClientForAccountNameStub = stub
}

func (fake *FakeCredentialsController) HTTPClientForAccountNameArgsForCall(i int) string {
	fake.hTTPClientForAccountNameMutex.RLock()
	defer fake.hTTPClientForAccountNameMutex.RUnlock()
	argsForCall := fake.hTTPClientForAccountNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredentialsController) HTTPClientForAccountNameReturns(result1 *http.Client, result2 error) {
	fake.hTTPClientForAccountNameMutex.Lock()
	defer fake.hTTPClientForAccountNameMutex.Unlock()
	fake.HTTPClientForAccountNameStub = nil
	fake.hTTPClientForAccountNameReturns = struct {
		result1 *http.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) HTTPClientForAccountNameReturnsOnCall(i int, result1 *http.Client, result2 error) {
	fake.hTTPClientForAccountNameMutex.Lock()
	defer fake.hTTPClientForAccountNameMutex.Unlock()
	fake.HTTPClientForAccountNameStub = nil
	if fake.hTTPClientForAccountNameReturnsOnCall == nil {
		fake.hTTPClientForAccountNameReturnsOnCall = make(map[int]struct {
			result1 *http.Client
			result2 error
		})
	}
	fake.hTTPClientForAccountNameReturnsOnCall[i] = struct {
		result1 *http.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) HelmClientForAccountName(arg1 string) (helm.Client, error) {
	fake.helmClientForAccountNameMutex.Lock()
	ret, specificReturn := fake.helmClientForAccountNameReturnsOnCall[len(fake.helmClientForAccountNameArgsForCall)]
	fake.helmClientForAccountNameArgsForCall = append(fake.helmClientForAccountNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HelmClientForAccountName", []interface{}{arg1})
	fake.helmClientForAccountNameMutex.Unlock()
	if fake.HelmClientForAccountNameStub != nil {
		return fake.HelmClientForAccountNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.helmClientForAccountNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredentialsController) HelmClientForAccountNameCallCount() int {
	fake.helmClientForAccountNameMutex.RLock()
	defer fake.helmClientForAccountNameMutex.RUnlock()
	return len(fake.helmClientForAccountNameArgsForCall)
}

func (fake *FakeCredentialsController) HelmClientForAccountNameCalls(stub func(string) (helm.Client, error)) {
	fake.helmClientForAccountNameMutex.Lock()
	defer fake.helmClientForAccountNameMutex.Unlock()
	fake.HelmClientForAccountNameStub = stub
}

func (fake *FakeCredentialsController) HelmClientForAccountNameArgsForCall(i int) string {
	fake.helmClientForAccountNameMutex.RLock()
	defer fake.helmClientForAccountNameMutex.RUnlock()
	argsForCall := fake.helmClientForAccountNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredentialsController) HelmClientForAccountNameReturns(result1 helm.Client, result2 error) {
	fake.helmClientForAccountNameMutex.Lock()
	defer fake.helmClientForAccountNameMutex.Unlock()
	fake.HelmClientForAccountNameStub = nil
	fake.helmClientForAccountNameReturns = struct {
		result1 helm.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) HelmClientForAccountNameReturnsOnCall(i int, result1 helm.Client, result2 error) {
	fake.helmClientForAccountNameMutex.Lock()
	defer fake.helmClientForAccountNameMutex.Unlock()
	fake.HelmClientForAccountNameStub = nil
	if fake.helmClientForAccountNameReturnsOnCall == nil {
		fake.helmClientForAccountNameReturnsOnCall = make(map[int]struct {
			result1 helm.Client
			result2 error
		})
	}
	fake.helmClientForAccountNameReturnsOnCall[i] = struct {
		result1 helm.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeCredentialsController) ListArtifactCredentialsNamesAndTypes() []artifact.Credentials {
	fake.listArtifactCredentialsNamesAndTypesMutex.Lock()
	ret, specificReturn := fake.listArtifactCredentialsNamesAndTypesReturnsOnCall[len(fake.listArtifactCredentialsNamesAndTypesArgsForCall)]
	fake.listArtifactCredentialsNamesAndTypesArgsForCall = append(fake.listArtifactCredentialsNamesAndTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("ListArtifactCredentialsNamesAndTypes", []interface{}{})
	fake.listArtifactCredentialsNamesAndTypesMutex.Unlock()
	if fake.ListArtifactCredentialsNamesAndTypesStub != nil {
		return fake.ListArtifactCredentialsNamesAndTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listArtifactCredentialsNamesAndTypesReturns
	return fakeReturns.result1
}

func (fake *FakeCredentialsController) ListArtifactCredentialsNamesAndTypesCallCount() int {
	fake.listArtifactCredentialsNamesAndTypesMutex.RLock()
	defer fake.listArtifactCredentialsNamesAndTypesMutex.RUnlock()
	return len(fake.listArtifactCredentialsNamesAndTypesArgsForCall)
}

func (fake *FakeCredentialsController) ListArtifactCredentialsNamesAndTypesCalls(stub func() []artifact.Credentials) {
	fake.listArtifactCredentialsNamesAndTypesMutex.Lock()
	defer fake.listArtifactCredentialsNamesAndTypesMutex.Unlock()
	fake.ListArtifactCredentialsNamesAndTypesStub = stub
}

func (fake *FakeCredentialsController) ListArtifactCredentialsNamesAndTypesReturns(result1 []artifact.Credentials) {
	fake.listArtifactCredentialsNamesAndTypesMutex.Lock()
	defer fake.listArtifactCredentialsNamesAndTypesMutex.Unlock()
	fake.ListArtifactCredentialsNamesAndTypesStub = nil
	fake.listArtifactCredentialsNamesAndTypesReturns = struct {
		result1 []artifact.Credentials
	}{result1}
}

func (fake *FakeCredentialsController) ListArtifactCredentialsNamesAndTypesReturnsOnCall(i int, result1 []artifact.Credentials) {
	fake.listArtifactCredentialsNamesAndTypesMutex.Lock()
	defer fake.listArtifactCredentialsNamesAndTypesMutex.Unlock()
	fake.ListArtifactCredentialsNamesAndTypesStub = nil
	if fake.listArtifactCredentialsNamesAndTypesReturnsOnCall == nil {
		fake.listArtifactCredentialsNamesAndTypesReturnsOnCall = make(map[int]struct {
			result1 []artifact.Credentials
		})
	}
	fake.listArtifactCredentialsNamesAndTypesReturnsOnCall[i] = struct {
		result1 []artifact.Credentials
	}{result1}
}

func (fake *FakeCredentialsController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gitClientForAccountNameMutex.RLock()
	defer fake.gitClientForAccountNameMutex.RUnlock()
	fake.gitRepoClientForAccountNameMutex.RLock()
	defer fake.gitRepoClientForAccountNameMutex.RUnlock()
	fake.hTTPClientForAccountNameMutex.RLock()
	defer fake.hTTPClientForAccountNameMutex.RUnlock()
	fake.helmClientForAccountNameMutex.RLock()
	defer fake.helmClientForAccountNameMutex.RUnlock()
	fake.listArtifactCredentialsNamesAndTypesMutex.RLock()
	defer fake.listArtifactCredentialsNamesAndTypesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredentialsController) recordInvocation(key string, args []interface{}) {
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

var _ artifact.CredentialsController = new(FakeCredentialsController)
