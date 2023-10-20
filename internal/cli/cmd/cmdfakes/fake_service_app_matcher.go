// Copyright © 2021 - 2023 SUSE LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/epinio/epinio/internal/cli/cmd"
	"github.com/epinio/epinio/internal/cli/usercmd"
)

type FakeServiceAppMatcher struct {
	AppsMatchingStub        func(string) []string
	appsMatchingMutex       sync.RWMutex
	appsMatchingArgsForCall []struct {
		arg1 string
	}
	appsMatchingReturns struct {
		result1 []string
	}
	appsMatchingReturnsOnCall map[int]struct {
		result1 []string
	}
	GetAPIStub        func() usercmd.APIClient
	getAPIMutex       sync.RWMutex
	getAPIArgsForCall []struct {
	}
	getAPIReturns struct {
		result1 usercmd.APIClient
	}
	getAPIReturnsOnCall map[int]struct {
		result1 usercmd.APIClient
	}
	ServiceMatchingStub        func(string) []string
	serviceMatchingMutex       sync.RWMutex
	serviceMatchingArgsForCall []struct {
		arg1 string
	}
	serviceMatchingReturns struct {
		result1 []string
	}
	serviceMatchingReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAppMatcher) AppsMatching(arg1 string) []string {
	fake.appsMatchingMutex.Lock()
	ret, specificReturn := fake.appsMatchingReturnsOnCall[len(fake.appsMatchingArgsForCall)]
	fake.appsMatchingArgsForCall = append(fake.appsMatchingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppsMatchingStub
	fakeReturns := fake.appsMatchingReturns
	fake.recordInvocation("AppsMatching", []interface{}{arg1})
	fake.appsMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAppMatcher) AppsMatchingCallCount() int {
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	return len(fake.appsMatchingArgsForCall)
}

func (fake *FakeServiceAppMatcher) AppsMatchingCalls(stub func(string) []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = stub
}

func (fake *FakeServiceAppMatcher) AppsMatchingArgsForCall(i int) string {
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	argsForCall := fake.appsMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceAppMatcher) AppsMatchingReturns(result1 []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = nil
	fake.appsMatchingReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeServiceAppMatcher) AppsMatchingReturnsOnCall(i int, result1 []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = nil
	if fake.appsMatchingReturnsOnCall == nil {
		fake.appsMatchingReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.appsMatchingReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeServiceAppMatcher) GetAPI() usercmd.APIClient {
	fake.getAPIMutex.Lock()
	ret, specificReturn := fake.getAPIReturnsOnCall[len(fake.getAPIArgsForCall)]
	fake.getAPIArgsForCall = append(fake.getAPIArgsForCall, struct {
	}{})
	stub := fake.GetAPIStub
	fakeReturns := fake.getAPIReturns
	fake.recordInvocation("GetAPI", []interface{}{})
	fake.getAPIMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAppMatcher) GetAPICallCount() int {
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	return len(fake.getAPIArgsForCall)
}

func (fake *FakeServiceAppMatcher) GetAPICalls(stub func() usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = stub
}

func (fake *FakeServiceAppMatcher) GetAPIReturns(result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	fake.getAPIReturns = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeServiceAppMatcher) GetAPIReturnsOnCall(i int, result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	if fake.getAPIReturnsOnCall == nil {
		fake.getAPIReturnsOnCall = make(map[int]struct {
			result1 usercmd.APIClient
		})
	}
	fake.getAPIReturnsOnCall[i] = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeServiceAppMatcher) ServiceMatching(arg1 string) []string {
	fake.serviceMatchingMutex.Lock()
	ret, specificReturn := fake.serviceMatchingReturnsOnCall[len(fake.serviceMatchingArgsForCall)]
	fake.serviceMatchingArgsForCall = append(fake.serviceMatchingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ServiceMatchingStub
	fakeReturns := fake.serviceMatchingReturns
	fake.recordInvocation("ServiceMatching", []interface{}{arg1})
	fake.serviceMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeServiceAppMatcher) ServiceMatchingCallCount() int {
	fake.serviceMatchingMutex.RLock()
	defer fake.serviceMatchingMutex.RUnlock()
	return len(fake.serviceMatchingArgsForCall)
}

func (fake *FakeServiceAppMatcher) ServiceMatchingCalls(stub func(string) []string) {
	fake.serviceMatchingMutex.Lock()
	defer fake.serviceMatchingMutex.Unlock()
	fake.ServiceMatchingStub = stub
}

func (fake *FakeServiceAppMatcher) ServiceMatchingArgsForCall(i int) string {
	fake.serviceMatchingMutex.RLock()
	defer fake.serviceMatchingMutex.RUnlock()
	argsForCall := fake.serviceMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceAppMatcher) ServiceMatchingReturns(result1 []string) {
	fake.serviceMatchingMutex.Lock()
	defer fake.serviceMatchingMutex.Unlock()
	fake.ServiceMatchingStub = nil
	fake.serviceMatchingReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeServiceAppMatcher) ServiceMatchingReturnsOnCall(i int, result1 []string) {
	fake.serviceMatchingMutex.Lock()
	defer fake.serviceMatchingMutex.Unlock()
	fake.ServiceMatchingStub = nil
	if fake.serviceMatchingReturnsOnCall == nil {
		fake.serviceMatchingReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.serviceMatchingReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeServiceAppMatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	fake.serviceMatchingMutex.RLock()
	defer fake.serviceMatchingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceAppMatcher) recordInvocation(key string, args []interface{}) {
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

var _ cmd.ServiceAppMatcher = new(FakeServiceAppMatcher)