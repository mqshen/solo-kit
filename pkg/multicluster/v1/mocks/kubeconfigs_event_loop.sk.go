// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/multicluster/v1/kubeconfigs_event_loop.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/solo-kit/pkg/multicluster/v1"
)

// MockKubeconfigsSyncer is a mock of KubeconfigsSyncer interface
type MockKubeconfigsSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockKubeconfigsSyncerMockRecorder
}

// MockKubeconfigsSyncerMockRecorder is the mock recorder for MockKubeconfigsSyncer
type MockKubeconfigsSyncerMockRecorder struct {
	mock *MockKubeconfigsSyncer
}

// NewMockKubeconfigsSyncer creates a new mock instance
func NewMockKubeconfigsSyncer(ctrl *gomock.Controller) *MockKubeconfigsSyncer {
	mock := &MockKubeconfigsSyncer{ctrl: ctrl}
	mock.recorder = &MockKubeconfigsSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeconfigsSyncer) EXPECT() *MockKubeconfigsSyncerMockRecorder {
	return m.recorder
}

// Sync mocks base method
func (m *MockKubeconfigsSyncer) Sync(arg0 context.Context, arg1 *v1.KubeconfigsSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockKubeconfigsSyncerMockRecorder) Sync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockKubeconfigsSyncer)(nil).Sync), arg0, arg1)
}
