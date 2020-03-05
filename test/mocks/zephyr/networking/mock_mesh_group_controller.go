// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/controller (interfaces: MeshGroupController,TrafficPolicyController)

// Package mock_zephyr_networking is a generated GoMock package.
package mock_zephyr_networking

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMeshGroupController is a mock of MeshGroupController interface
type MockMeshGroupController struct {
	ctrl     *gomock.Controller
	recorder *MockMeshGroupControllerMockRecorder
}

// MockMeshGroupControllerMockRecorder is the mock recorder for MockMeshGroupController
type MockMeshGroupControllerMockRecorder struct {
	mock *MockMeshGroupController
}

// NewMockMeshGroupController creates a new mock instance
func NewMockMeshGroupController(ctrl *gomock.Controller) *MockMeshGroupController {
	mock := &MockMeshGroupController{ctrl: ctrl}
	mock.recorder = &MockMeshGroupControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshGroupController) EXPECT() *MockMeshGroupControllerMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockMeshGroupController) AddEventHandler(arg0 context.Context, arg1 controller.MeshGroupEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockMeshGroupControllerMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshGroupController)(nil).AddEventHandler), varargs...)
}

// MockTrafficPolicyController is a mock of TrafficPolicyController interface
type MockTrafficPolicyController struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyControllerMockRecorder
}

// MockTrafficPolicyControllerMockRecorder is the mock recorder for MockTrafficPolicyController
type MockTrafficPolicyControllerMockRecorder struct {
	mock *MockTrafficPolicyController
}

// NewMockTrafficPolicyController creates a new mock instance
func NewMockTrafficPolicyController(ctrl *gomock.Controller) *MockTrafficPolicyController {
	mock := &MockTrafficPolicyController{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrafficPolicyController) EXPECT() *MockTrafficPolicyControllerMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockTrafficPolicyController) AddEventHandler(arg0 context.Context, arg1 controller.TrafficPolicyEventHandler, arg2 ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockTrafficPolicyControllerMockRecorder) AddEventHandler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockTrafficPolicyController)(nil).AddEventHandler), varargs...)
}
