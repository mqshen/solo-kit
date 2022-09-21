// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/v1/resources/common/kubernetes/custom_resource_definition_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
)

// MockCustomResourceDefinitionReconciler is a mock of CustomResourceDefinitionReconciler interface
type MockCustomResourceDefinitionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionReconcilerMockRecorder
}

// MockCustomResourceDefinitionReconcilerMockRecorder is the mock recorder for MockCustomResourceDefinitionReconciler
type MockCustomResourceDefinitionReconcilerMockRecorder struct {
	mock *MockCustomResourceDefinitionReconciler
}

// NewMockCustomResourceDefinitionReconciler creates a new mock instance
func NewMockCustomResourceDefinitionReconciler(ctrl *gomock.Controller) *MockCustomResourceDefinitionReconciler {
	mock := &MockCustomResourceDefinitionReconciler{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionReconciler) EXPECT() *MockCustomResourceDefinitionReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockCustomResourceDefinitionReconciler) Reconcile(namespace string, desiredResources kubernetes.CustomResourceDefinitionList, transition kubernetes.TransitionCustomResourceDefinitionFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockCustomResourceDefinitionReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockCustomResourceDefinitionReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}