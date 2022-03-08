// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cassandra/cassandra.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
)

// MockSessionInterface is a mock of SessionInterface interface.
type MockSessionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSessionInterfaceMockRecorder
}

// MockSessionInterfaceMockRecorder is the mock recorder for MockSessionInterface.
type MockSessionInterfaceMockRecorder struct {
	mock *MockSessionInterface
}

// NewMockSessionInterface creates a new mock instance.
func NewMockSessionInterface(ctrl *gomock.Controller) *MockSessionInterface {
	mock := &MockSessionInterface{ctrl: ctrl}
	mock.recorder = &MockSessionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionInterface) EXPECT() *MockSessionInterfaceMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockSessionInterface) Query(arg0 string, arg1 ...interface{}) *gocql.Query {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*gocql.Query)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockSessionInterfaceMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSessionInterface)(nil).Query), varargs...)
}

// MockQueryInterface is a mock of QueryInterface interface.
type MockQueryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockQueryInterfaceMockRecorder
}

// MockQueryInterfaceMockRecorder is the mock recorder for MockQueryInterface.
type MockQueryInterfaceMockRecorder struct {
	mock *MockQueryInterface
}

// NewMockQueryInterface creates a new mock instance.
func NewMockQueryInterface(ctrl *gomock.Controller) *MockQueryInterface {
	mock := &MockQueryInterface{ctrl: ctrl}
	mock.recorder = &MockQueryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryInterface) EXPECT() *MockQueryInterfaceMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockQueryInterface) Exec() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec")
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockQueryInterfaceMockRecorder) Exec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockQueryInterface)(nil).Exec))
}
