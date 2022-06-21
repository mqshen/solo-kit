// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/multicluster/v1/kubeconfigs_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1 "github.com/solo-io/solo-kit/pkg/multicluster/v1"
)

// MockKubeconfigsSnapshotEmitter is a mock of KubeconfigsSnapshotEmitter interface
type MockKubeconfigsSnapshotEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockKubeconfigsSnapshotEmitterMockRecorder
}

// MockKubeconfigsSnapshotEmitterMockRecorder is the mock recorder for MockKubeconfigsSnapshotEmitter
type MockKubeconfigsSnapshotEmitterMockRecorder struct {
	mock *MockKubeconfigsSnapshotEmitter
}

// NewMockKubeconfigsSnapshotEmitter creates a new mock instance
func NewMockKubeconfigsSnapshotEmitter(ctrl *gomock.Controller) *MockKubeconfigsSnapshotEmitter {
	mock := &MockKubeconfigsSnapshotEmitter{ctrl: ctrl}
	mock.recorder = &MockKubeconfigsSnapshotEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeconfigsSnapshotEmitter) EXPECT() *MockKubeconfigsSnapshotEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockKubeconfigsSnapshotEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.KubeconfigsSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.KubeconfigsSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockKubeconfigsSnapshotEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockKubeconfigsSnapshotEmitter)(nil).Snapshots), watchNamespaces, opts)
}

// MockKubeconfigsEmitter is a mock of KubeconfigsEmitter interface
type MockKubeconfigsEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockKubeconfigsEmitterMockRecorder
}

// MockKubeconfigsEmitterMockRecorder is the mock recorder for MockKubeconfigsEmitter
type MockKubeconfigsEmitterMockRecorder struct {
	mock *MockKubeconfigsEmitter
}

// NewMockKubeconfigsEmitter creates a new mock instance
func NewMockKubeconfigsEmitter(ctrl *gomock.Controller) *MockKubeconfigsEmitter {
	mock := &MockKubeconfigsEmitter{ctrl: ctrl}
	mock.recorder = &MockKubeconfigsEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeconfigsEmitter) EXPECT() *MockKubeconfigsEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockKubeconfigsEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v1.KubeconfigsSnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v1.KubeconfigsSnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockKubeconfigsEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockKubeconfigsEmitter)(nil).Snapshots), watchNamespaces, opts)
}

// Register mocks base method
func (m *MockKubeconfigsEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockKubeconfigsEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockKubeconfigsEmitter)(nil).Register))
}

// KubeConfig mocks base method
func (m *MockKubeconfigsEmitter) KubeConfig() v1.KubeConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeConfig")
	ret0, _ := ret[0].(v1.KubeConfigClient)
	return ret0
}

// KubeConfig indicates an expected call of KubeConfig
func (mr *MockKubeconfigsEmitterMockRecorder) KubeConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeConfig", reflect.TypeOf((*MockKubeconfigsEmitter)(nil).KubeConfig))
}
