// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/admin.enterprise.mesh.gloo.solo.io/v1alpha1"
	controller "github.com/solo-io/gloo-mesh/pkg/api/admin.enterprise.mesh.gloo.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockIstioInstallationEventHandler is a mock of IstioInstallationEventHandler interface.
type MockIstioInstallationEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationEventHandlerMockRecorder
}

// MockIstioInstallationEventHandlerMockRecorder is the mock recorder for MockIstioInstallationEventHandler.
type MockIstioInstallationEventHandlerMockRecorder struct {
	mock *MockIstioInstallationEventHandler
}

// NewMockIstioInstallationEventHandler creates a new mock instance.
func NewMockIstioInstallationEventHandler(ctrl *gomock.Controller) *MockIstioInstallationEventHandler {
	mock := &MockIstioInstallationEventHandler{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioInstallationEventHandler) EXPECT() *MockIstioInstallationEventHandlerMockRecorder {
	return m.recorder
}

// CreateIstioInstallation mocks base method.
func (m *MockIstioInstallationEventHandler) CreateIstioInstallation(obj *v1alpha1.IstioInstallation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIstioInstallation", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIstioInstallation indicates an expected call of CreateIstioInstallation.
func (mr *MockIstioInstallationEventHandlerMockRecorder) CreateIstioInstallation(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIstioInstallation", reflect.TypeOf((*MockIstioInstallationEventHandler)(nil).CreateIstioInstallation), obj)
}

// DeleteIstioInstallation mocks base method.
func (m *MockIstioInstallationEventHandler) DeleteIstioInstallation(obj *v1alpha1.IstioInstallation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIstioInstallation", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIstioInstallation indicates an expected call of DeleteIstioInstallation.
func (mr *MockIstioInstallationEventHandlerMockRecorder) DeleteIstioInstallation(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIstioInstallation", reflect.TypeOf((*MockIstioInstallationEventHandler)(nil).DeleteIstioInstallation), obj)
}

// GenericIstioInstallation mocks base method.
func (m *MockIstioInstallationEventHandler) GenericIstioInstallation(obj *v1alpha1.IstioInstallation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericIstioInstallation", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericIstioInstallation indicates an expected call of GenericIstioInstallation.
func (mr *MockIstioInstallationEventHandlerMockRecorder) GenericIstioInstallation(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericIstioInstallation", reflect.TypeOf((*MockIstioInstallationEventHandler)(nil).GenericIstioInstallation), obj)
}

// UpdateIstioInstallation mocks base method.
func (m *MockIstioInstallationEventHandler) UpdateIstioInstallation(old, new *v1alpha1.IstioInstallation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIstioInstallation", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIstioInstallation indicates an expected call of UpdateIstioInstallation.
func (mr *MockIstioInstallationEventHandlerMockRecorder) UpdateIstioInstallation(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIstioInstallation", reflect.TypeOf((*MockIstioInstallationEventHandler)(nil).UpdateIstioInstallation), old, new)
}

// MockIstioInstallationEventWatcher is a mock of IstioInstallationEventWatcher interface.
type MockIstioInstallationEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockIstioInstallationEventWatcherMockRecorder
}

// MockIstioInstallationEventWatcherMockRecorder is the mock recorder for MockIstioInstallationEventWatcher.
type MockIstioInstallationEventWatcherMockRecorder struct {
	mock *MockIstioInstallationEventWatcher
}

// NewMockIstioInstallationEventWatcher creates a new mock instance.
func NewMockIstioInstallationEventWatcher(ctrl *gomock.Controller) *MockIstioInstallationEventWatcher {
	mock := &MockIstioInstallationEventWatcher{ctrl: ctrl}
	mock.recorder = &MockIstioInstallationEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIstioInstallationEventWatcher) EXPECT() *MockIstioInstallationEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockIstioInstallationEventWatcher) AddEventHandler(ctx context.Context, h controller.IstioInstallationEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockIstioInstallationEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockIstioInstallationEventWatcher)(nil).AddEventHandler), varargs...)
}