// Code generated by MockGen. DO NOT EDIT.
// Source: ./helper.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	resource "github.com/openshift/hive/pkg/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockHelper is a mock of Helper interface
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockHelper) Apply(obj []byte) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", obj)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply
func (mr *MockHelperMockRecorder) Apply(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockHelper)(nil).Apply), obj)
}

// ApplyRuntimeObject mocks base method
func (m *MockHelper) ApplyRuntimeObject(obj runtime.Object, scheme *runtime.Scheme) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyRuntimeObject", obj, scheme)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyRuntimeObject indicates an expected call of ApplyRuntimeObject
func (mr *MockHelperMockRecorder) ApplyRuntimeObject(obj, scheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRuntimeObject", reflect.TypeOf((*MockHelper)(nil).ApplyRuntimeObject), obj, scheme)
}

// CreateOrUpdate mocks base method
func (m *MockHelper) CreateOrUpdate(obj []byte) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", obj)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockHelperMockRecorder) CreateOrUpdate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockHelper)(nil).CreateOrUpdate), obj)
}

// CreateOrUpdateRuntimeObject mocks base method
func (m *MockHelper) CreateOrUpdateRuntimeObject(obj runtime.Object, scheme *runtime.Scheme) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRuntimeObject", obj, scheme)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateRuntimeObject indicates an expected call of CreateOrUpdateRuntimeObject
func (mr *MockHelperMockRecorder) CreateOrUpdateRuntimeObject(obj, scheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRuntimeObject", reflect.TypeOf((*MockHelper)(nil).CreateOrUpdateRuntimeObject), obj, scheme)
}

// Create mocks base method
func (m *MockHelper) Create(obj []byte) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", obj)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockHelperMockRecorder) Create(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHelper)(nil).Create), obj)
}

// CreateRuntimeObject mocks base method
func (m *MockHelper) CreateRuntimeObject(obj runtime.Object, scheme *runtime.Scheme) (resource.ApplyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRuntimeObject", obj, scheme)
	ret0, _ := ret[0].(resource.ApplyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRuntimeObject indicates an expected call of CreateRuntimeObject
func (mr *MockHelperMockRecorder) CreateRuntimeObject(obj, scheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRuntimeObject", reflect.TypeOf((*MockHelper)(nil).CreateRuntimeObject), obj, scheme)
}

// Info mocks base method
func (m *MockHelper) Info(obj []byte) (*resource.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", obj)
	ret0, _ := ret[0].(*resource.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockHelperMockRecorder) Info(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockHelper)(nil).Info), obj)
}

// Patch mocks base method
func (m *MockHelper) Patch(name types.NamespacedName, kind, apiVersion string, patch []byte, patchType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", name, kind, apiVersion, patch, patchType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch
func (mr *MockHelperMockRecorder) Patch(name, kind, apiVersion, patch, patchType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockHelper)(nil).Patch), name, kind, apiVersion, patch, patchType)
}
