// This file was generated by counterfeiter
package user_fake

import (
	"os/user"
	"sync"

	"code.cloudfoundry.org/goshims/usershim"
)

type FakeUser struct {
	CurrentStub        func() (*user.User, error)
	currentMutex       sync.RWMutex
	currentArgsForCall []struct{}
	currentReturns     struct {
		result1 *user.User
		result2 error
	}
	LookupStub        func(username string) (*user.User, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		username string
	}
	lookupReturns struct {
		result1 *user.User
		result2 error
	}
	LookupIdStub        func(uid string) (*user.User, error)
	lookupIdMutex       sync.RWMutex
	lookupIdArgsForCall []struct {
		uid string
	}
	lookupIdReturns struct {
		result1 *user.User
		result2 error
	}
	LookupGroupStub        func(name string) (*user.Group, error)
	lookupGroupMutex       sync.RWMutex
	lookupGroupArgsForCall []struct {
		name string
	}
	lookupGroupReturns struct {
		result1 *user.Group
		result2 error
	}
	LookupGroupIdStub        func(gid string) (*user.Group, error)
	lookupGroupIdMutex       sync.RWMutex
	lookupGroupIdArgsForCall []struct {
		gid string
	}
	lookupGroupIdReturns struct {
		result1 *user.Group
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUser) Current() (*user.User, error) {
	fake.currentMutex.Lock()
	fake.currentArgsForCall = append(fake.currentArgsForCall, struct{}{})
	fake.recordInvocation("Current", []interface{}{})
	fake.currentMutex.Unlock()
	if fake.CurrentStub != nil {
		return fake.CurrentStub()
	} else {
		return fake.currentReturns.result1, fake.currentReturns.result2
	}
}

func (fake *FakeUser) CurrentCallCount() int {
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	return len(fake.currentArgsForCall)
}

func (fake *FakeUser) CurrentReturns(result1 *user.User, result2 error) {
	fake.CurrentStub = nil
	fake.currentReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Lookup(username string) (*user.User, error) {
	fake.lookupMutex.Lock()
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		username string
	}{username})
	fake.recordInvocation("Lookup", []interface{}{username})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(username)
	} else {
		return fake.lookupReturns.result1, fake.lookupReturns.result2
	}
}

func (fake *FakeUser) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUser) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].username
}

func (fake *FakeUser) LookupReturns(result1 *user.User, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupId(uid string) (*user.User, error) {
	fake.lookupIdMutex.Lock()
	fake.lookupIdArgsForCall = append(fake.lookupIdArgsForCall, struct {
		uid string
	}{uid})
	fake.recordInvocation("LookupId", []interface{}{uid})
	fake.lookupIdMutex.Unlock()
	if fake.LookupIdStub != nil {
		return fake.LookupIdStub(uid)
	} else {
		return fake.lookupIdReturns.result1, fake.lookupIdReturns.result2
	}
}

func (fake *FakeUser) LookupIdCallCount() int {
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	return len(fake.lookupIdArgsForCall)
}

func (fake *FakeUser) LookupIdArgsForCall(i int) string {
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	return fake.lookupIdArgsForCall[i].uid
}

func (fake *FakeUser) LookupIdReturns(result1 *user.User, result2 error) {
	fake.LookupIdStub = nil
	fake.lookupIdReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroup(name string) (*user.Group, error) {
	fake.lookupGroupMutex.Lock()
	fake.lookupGroupArgsForCall = append(fake.lookupGroupArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("LookupGroup", []interface{}{name})
	fake.lookupGroupMutex.Unlock()
	if fake.LookupGroupStub != nil {
		return fake.LookupGroupStub(name)
	} else {
		return fake.lookupGroupReturns.result1, fake.lookupGroupReturns.result2
	}
}

func (fake *FakeUser) LookupGroupCallCount() int {
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	return len(fake.lookupGroupArgsForCall)
}

func (fake *FakeUser) LookupGroupArgsForCall(i int) string {
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	return fake.lookupGroupArgsForCall[i].name
}

func (fake *FakeUser) LookupGroupReturns(result1 *user.Group, result2 error) {
	fake.LookupGroupStub = nil
	fake.lookupGroupReturns = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroupId(gid string) (*user.Group, error) {
	fake.lookupGroupIdMutex.Lock()
	fake.lookupGroupIdArgsForCall = append(fake.lookupGroupIdArgsForCall, struct {
		gid string
	}{gid})
	fake.recordInvocation("LookupGroupId", []interface{}{gid})
	fake.lookupGroupIdMutex.Unlock()
	if fake.LookupGroupIdStub != nil {
		return fake.LookupGroupIdStub(gid)
	} else {
		return fake.lookupGroupIdReturns.result1, fake.lookupGroupIdReturns.result2
	}
}

func (fake *FakeUser) LookupGroupIdCallCount() int {
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	return len(fake.lookupGroupIdArgsForCall)
}

func (fake *FakeUser) LookupGroupIdArgsForCall(i int) string {
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	return fake.lookupGroupIdArgsForCall[i].gid
}

func (fake *FakeUser) LookupGroupIdReturns(result1 *user.Group, result2 error) {
	fake.LookupGroupIdStub = nil
	fake.lookupGroupIdReturns = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUser) recordInvocation(key string, args []interface{}) {
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

var _ usershim.User = new(FakeUser)
