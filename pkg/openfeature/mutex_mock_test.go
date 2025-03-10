// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/openfeature/mutex.go

// Package openfeature is a generated GoMock package.
package openfeature

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockmutex is a mock of mutex interface.
type Mockmutex struct {
	ctrl     *gomock.Controller
	recorder *MockmutexMockRecorder
}

// MockmutexMockRecorder is the mock recorder for Mockmutex.
type MockmutexMockRecorder struct {
	mock *Mockmutex
}

// NewMockmutex creates a new mock instance.
func NewMockmutex(ctrl *gomock.Controller) *Mockmutex {
	mock := &Mockmutex{ctrl: ctrl}
	mock.recorder = &MockmutexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockmutex) EXPECT() *MockmutexMockRecorder {
	return m.recorder
}

// Lock mocks base method.
func (m *Mockmutex) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockmutexMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*Mockmutex)(nil).Lock))
}

// RLock mocks base method.
func (m *Mockmutex) RLock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RLock")
}

// RLock indicates an expected call of RLock.
func (mr *MockmutexMockRecorder) RLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RLock", reflect.TypeOf((*Mockmutex)(nil).RLock))
}

// RUnlock mocks base method.
func (m *Mockmutex) RUnlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RUnlock")
}

// RUnlock indicates an expected call of RUnlock.
func (mr *MockmutexMockRecorder) RUnlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RUnlock", reflect.TypeOf((*Mockmutex)(nil).RUnlock))
}

// Unlock mocks base method.
func (m *Mockmutex) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockmutexMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*Mockmutex)(nil).Unlock))
}
