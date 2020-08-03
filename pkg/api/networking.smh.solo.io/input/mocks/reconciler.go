// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconciler.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	v1alpha20 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
)

// MockmultiClusterReconciler is a mock of multiClusterReconciler interface.
type MockmultiClusterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockmultiClusterReconcilerMockRecorder
}

// MockmultiClusterReconcilerMockRecorder is the mock recorder for MockmultiClusterReconciler.
type MockmultiClusterReconcilerMockRecorder struct {
	mock *MockmultiClusterReconciler
}

// NewMockmultiClusterReconciler creates a new mock instance.
func NewMockmultiClusterReconciler(ctrl *gomock.Controller) *MockmultiClusterReconciler {
	mock := &MockmultiClusterReconciler{ctrl: ctrl}
	mock.recorder = &MockmultiClusterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmultiClusterReconciler) EXPECT() *MockmultiClusterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshService mocks base method.
func (m *MockmultiClusterReconciler) ReconcileMeshService(clusterName string, obj *v1alpha2.MeshService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshService indicates an expected call of ReconcileMeshService.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileMeshService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshService", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileMeshService), clusterName, obj)
}

// ReconcileMeshWorkload mocks base method.
func (m *MockmultiClusterReconciler) ReconcileMeshWorkload(clusterName string, obj *v1alpha2.MeshWorkload) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshWorkload", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshWorkload indicates an expected call of ReconcileMeshWorkload.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileMeshWorkload(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshWorkload", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileMeshWorkload), clusterName, obj)
}

// ReconcileMesh mocks base method.
func (m *MockmultiClusterReconciler) ReconcileMesh(clusterName string, obj *v1alpha2.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileMesh), clusterName, obj)
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockmultiClusterReconciler) ReconcileTrafficPolicy(clusterName string, obj *v1alpha20.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileTrafficPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileTrafficPolicy), clusterName, obj)
}

// ReconcileAccessPolicy mocks base method.
func (m *MockmultiClusterReconciler) ReconcileAccessPolicy(clusterName string, obj *v1alpha20.AccessPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicy", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessPolicy indicates an expected call of ReconcileAccessPolicy.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileAccessPolicy(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicy", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileAccessPolicy), clusterName, obj)
}

// ReconcileVirtualMesh mocks base method.
func (m *MockmultiClusterReconciler) ReconcileVirtualMesh(clusterName string, obj *v1alpha20.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileVirtualMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileVirtualMesh), clusterName, obj)
}

// ReconcileFailoverService mocks base method.
func (m *MockmultiClusterReconciler) ReconcileFailoverService(clusterName string, obj *v1alpha20.FailoverService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFailoverService indicates an expected call of ReconcileFailoverService.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileFailoverService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverService", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileFailoverService), clusterName, obj)
}

// ReconcileKubernetesCluster mocks base method.
func (m *MockmultiClusterReconciler) ReconcileKubernetesCluster(clusterName string, obj *v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileKubernetesCluster", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileKubernetesCluster indicates an expected call of ReconcileKubernetesCluster.
func (mr *MockmultiClusterReconcilerMockRecorder) ReconcileKubernetesCluster(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileKubernetesCluster", reflect.TypeOf((*MockmultiClusterReconciler)(nil).ReconcileKubernetesCluster), clusterName, obj)
}

// MocksingleClusterReconciler is a mock of singleClusterReconciler interface.
type MocksingleClusterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MocksingleClusterReconcilerMockRecorder
}

// MocksingleClusterReconcilerMockRecorder is the mock recorder for MocksingleClusterReconciler.
type MocksingleClusterReconcilerMockRecorder struct {
	mock *MocksingleClusterReconciler
}

// NewMocksingleClusterReconciler creates a new mock instance.
func NewMocksingleClusterReconciler(ctrl *gomock.Controller) *MocksingleClusterReconciler {
	mock := &MocksingleClusterReconciler{ctrl: ctrl}
	mock.recorder = &MocksingleClusterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksingleClusterReconciler) EXPECT() *MocksingleClusterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshService mocks base method.
func (m *MocksingleClusterReconciler) ReconcileMeshService(obj *v1alpha2.MeshService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshService indicates an expected call of ReconcileMeshService.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileMeshService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshService", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileMeshService), obj)
}

// ReconcileMeshWorkload mocks base method.
func (m *MocksingleClusterReconciler) ReconcileMeshWorkload(obj *v1alpha2.MeshWorkload) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshWorkload", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshWorkload indicates an expected call of ReconcileMeshWorkload.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshWorkload", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileMeshWorkload), obj)
}

// ReconcileMesh mocks base method.
func (m *MocksingleClusterReconciler) ReconcileMesh(obj *v1alpha2.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileMesh), obj)
}

// ReconcileTrafficPolicy mocks base method.
func (m *MocksingleClusterReconciler) ReconcileTrafficPolicy(obj *v1alpha20.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileTrafficPolicy), obj)
}

// ReconcileAccessPolicy mocks base method.
func (m *MocksingleClusterReconciler) ReconcileAccessPolicy(obj *v1alpha20.AccessPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessPolicy indicates an expected call of ReconcileAccessPolicy.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileAccessPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessPolicy", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileAccessPolicy), obj)
}

// ReconcileVirtualMesh mocks base method.
func (m *MocksingleClusterReconciler) ReconcileVirtualMesh(obj *v1alpha20.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileVirtualMesh), obj)
}

// ReconcileFailoverService mocks base method.
func (m *MocksingleClusterReconciler) ReconcileFailoverService(obj *v1alpha20.FailoverService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFailoverService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFailoverService indicates an expected call of ReconcileFailoverService.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileFailoverService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFailoverService", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileFailoverService), obj)
}

// ReconcileKubernetesCluster mocks base method.
func (m *MocksingleClusterReconciler) ReconcileKubernetesCluster(obj *v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileKubernetesCluster", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileKubernetesCluster indicates an expected call of ReconcileKubernetesCluster.
func (mr *MocksingleClusterReconcilerMockRecorder) ReconcileKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileKubernetesCluster", reflect.TypeOf((*MocksingleClusterReconciler)(nil).ReconcileKubernetesCluster), obj)
}