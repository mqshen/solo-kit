// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/v1/resources/common/kubernetes/kube_namespace_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
)

// MockKubeNamespaceWatcher is a mock of KubeNamespaceWatcher interface
type MockKubeNamespaceWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockKubeNamespaceWatcherMockRecorder
}

// MockKubeNamespaceWatcherMockRecorder is the mock recorder for MockKubeNamespaceWatcher
type MockKubeNamespaceWatcherMockRecorder struct {
	mock *MockKubeNamespaceWatcher
}

// NewMockKubeNamespaceWatcher creates a new mock instance
func NewMockKubeNamespaceWatcher(ctrl *gomock.Controller) *MockKubeNamespaceWatcher {
	mock := &MockKubeNamespaceWatcher{ctrl: ctrl}
	mock.recorder = &MockKubeNamespaceWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeNamespaceWatcher) EXPECT() *MockKubeNamespaceWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockKubeNamespaceWatcher) Watch(opts clients.WatchOpts) (<-chan kubernetes.KubeNamespaceList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", opts)
	ret0, _ := ret[0].(<-chan kubernetes.KubeNamespaceList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockKubeNamespaceWatcherMockRecorder) Watch(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockKubeNamespaceWatcher)(nil).Watch), opts)
}

// MockKubeNamespaceClient is a mock of KubeNamespaceClient interface
type MockKubeNamespaceClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeNamespaceClientMockRecorder
}

// MockKubeNamespaceClientMockRecorder is the mock recorder for MockKubeNamespaceClient
type MockKubeNamespaceClientMockRecorder struct {
	mock *MockKubeNamespaceClient
}

// NewMockKubeNamespaceClient creates a new mock instance
func NewMockKubeNamespaceClient(ctrl *gomock.Controller) *MockKubeNamespaceClient {
	mock := &MockKubeNamespaceClient{ctrl: ctrl}
	mock.recorder = &MockKubeNamespaceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeNamespaceClient) EXPECT() *MockKubeNamespaceClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockKubeNamespaceClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockKubeNamespaceClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockKubeNamespaceClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockKubeNamespaceClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockKubeNamespaceClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockKubeNamespaceClient)(nil).Register))
}

// Read mocks base method
func (m *MockKubeNamespaceClient) Read(name string, opts clients.ReadOpts) (*kubernetes.KubeNamespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", name, opts)
	ret0, _ := ret[0].(*kubernetes.KubeNamespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockKubeNamespaceClientMockRecorder) Read(name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockKubeNamespaceClient)(nil).Read), name, opts)
}

// Write mocks base method
func (m *MockKubeNamespaceClient) Write(resource *kubernetes.KubeNamespace, opts clients.WriteOpts) (*kubernetes.KubeNamespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*kubernetes.KubeNamespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockKubeNamespaceClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockKubeNamespaceClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockKubeNamespaceClient) Delete(name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockKubeNamespaceClientMockRecorder) Delete(name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubeNamespaceClient)(nil).Delete), name, opts)
}

// List mocks base method
func (m *MockKubeNamespaceClient) List(opts clients.ListOpts) (kubernetes.KubeNamespaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", opts)
	ret0, _ := ret[0].(kubernetes.KubeNamespaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockKubeNamespaceClientMockRecorder) List(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubeNamespaceClient)(nil).List), opts)
}

// Watch mocks base method
func (m *MockKubeNamespaceClient) Watch(opts clients.WatchOpts) (<-chan kubernetes.KubeNamespaceList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", opts)
	ret0, _ := ret[0].(<-chan kubernetes.KubeNamespaceList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockKubeNamespaceClientMockRecorder) Watch(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockKubeNamespaceClient)(nil).Watch), opts)
}
