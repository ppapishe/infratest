// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/k8s/util.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
)

// Mockkubernetes is a mock of kubernetes interface.
type Mockkubernetes struct {
	ctrl     *gomock.Controller
	recorder *MockkubernetesMockRecorder
}

// MockkubernetesMockRecorder is the mock recorder for Mockkubernetes.
type MockkubernetesMockRecorder struct {
	mock *Mockkubernetes
}

// NewMockkubernetes creates a new mock instance.
func NewMockkubernetes(ctrl *gomock.Controller) *Mockkubernetes {
	mock := &Mockkubernetes{ctrl: ctrl}
	mock.recorder = &MockkubernetesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockkubernetes) EXPECT() *MockkubernetesMockRecorder {
	return m.recorder
}

// NewForConfig mocks base method.
func (m *Mockkubernetes) NewForConfig(arg0 *rest.Config) (*kubernetes.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewForConfig", arg0)
	ret0, _ := ret[0].(*kubernetes.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewForConfig indicates an expected call of NewForConfig.
func (mr *MockkubernetesMockRecorder) NewForConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewForConfig", reflect.TypeOf((*Mockkubernetes)(nil).NewForConfig), arg0)
}
