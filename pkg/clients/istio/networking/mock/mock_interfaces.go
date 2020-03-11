// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_istio_networking is a generated GoMock package.
package mock_istio_networking

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockGatewayClient is a mock of GatewayClient interface
type MockGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClientMockRecorder
}

// MockGatewayClientMockRecorder is the mock recorder for MockGatewayClient
type MockGatewayClientMockRecorder struct {
	mock *MockGatewayClient
}

// NewMockGatewayClient creates a new mock instance
func NewMockGatewayClient(ctrl *gomock.Controller) *MockGatewayClient {
	mock := &MockGatewayClient{ctrl: ctrl}
	mock.recorder = &MockGatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayClient) EXPECT() *MockGatewayClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGatewayClient) Create(ctx context.Context, gateway *v1alpha3.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, gateway)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockGatewayClientMockRecorder) Create(ctx, gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGatewayClient)(nil).Create), ctx, gateway)
}

// Get mocks base method
func (m *MockGatewayClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha3.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha3.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGatewayClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGatewayClient)(nil).Get), ctx, objKey)
}

// Update mocks base method
func (m *MockGatewayClient) Update(ctx context.Context, gateway *v1alpha3.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, gateway)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGatewayClientMockRecorder) Update(ctx, gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGatewayClient)(nil).Update), ctx, gateway)
}

// MockEnvoyFilterClient is a mock of EnvoyFilterClient interface
type MockEnvoyFilterClient struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterClientMockRecorder
}

// MockEnvoyFilterClientMockRecorder is the mock recorder for MockEnvoyFilterClient
type MockEnvoyFilterClientMockRecorder struct {
	mock *MockEnvoyFilterClient
}

// NewMockEnvoyFilterClient creates a new mock instance
func NewMockEnvoyFilterClient(ctrl *gomock.Controller) *MockEnvoyFilterClient {
	mock := &MockEnvoyFilterClient{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvoyFilterClient) EXPECT() *MockEnvoyFilterClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockEnvoyFilterClient) Create(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, envoyFilter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockEnvoyFilterClientMockRecorder) Create(ctx, envoyFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEnvoyFilterClient)(nil).Create), ctx, envoyFilter)
}

// Get mocks base method
func (m *MockEnvoyFilterClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha3.EnvoyFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha3.EnvoyFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockEnvoyFilterClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEnvoyFilterClient)(nil).Get), ctx, objKey)
}

// MockServiceEntryClient is a mock of ServiceEntryClient interface
type MockServiceEntryClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryClientMockRecorder
}

// MockServiceEntryClientMockRecorder is the mock recorder for MockServiceEntryClient
type MockServiceEntryClientMockRecorder struct {
	mock *MockServiceEntryClient
}

// NewMockServiceEntryClient creates a new mock instance
func NewMockServiceEntryClient(ctrl *gomock.Controller) *MockServiceEntryClient {
	mock := &MockServiceEntryClient{ctrl: ctrl}
	mock.recorder = &MockServiceEntryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceEntryClient) EXPECT() *MockServiceEntryClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServiceEntryClient) Create(ctx context.Context, serviceEntry *v1alpha3.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, serviceEntry)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockServiceEntryClientMockRecorder) Create(ctx, serviceEntry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceEntryClient)(nil).Create), ctx, serviceEntry)
}

// Get mocks base method
func (m *MockServiceEntryClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha3.ServiceEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha3.ServiceEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceEntryClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceEntryClient)(nil).Get), ctx, objKey)
}

// MockVirtualServiceClient is a mock of VirtualServiceClient interface
type MockVirtualServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceClientMockRecorder
}

// MockVirtualServiceClientMockRecorder is the mock recorder for MockVirtualServiceClient
type MockVirtualServiceClientMockRecorder struct {
	mock *MockVirtualServiceClient
}

// NewMockVirtualServiceClient creates a new mock instance
func NewMockVirtualServiceClient(ctrl *gomock.Controller) *MockVirtualServiceClient {
	mock := &MockVirtualServiceClient{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualServiceClient) EXPECT() *MockVirtualServiceClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockVirtualServiceClient) Get(ctx context.Context, key client.ObjectKey) (*v1alpha3.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualServiceClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualServiceClient)(nil).Get), ctx, key)
}

// Create mocks base method
func (m *MockVirtualServiceClient) Create(ctx context.Context, virtualService *v1alpha3.VirtualService, options ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, virtualService}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockVirtualServiceClientMockRecorder) Create(ctx, virtualService interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, virtualService}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVirtualServiceClient)(nil).Create), varargs...)
}

// Update mocks base method
func (m *MockVirtualServiceClient) Update(ctx context.Context, virtualService *v1alpha3.VirtualService, options ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, virtualService}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockVirtualServiceClientMockRecorder) Update(ctx, virtualService interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, virtualService}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVirtualServiceClient)(nil).Update), varargs...)
}

// Upsert mocks base method
func (m *MockVirtualServiceClient) Upsert(ctx context.Context, virtualService *v1alpha3.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, virtualService)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockVirtualServiceClientMockRecorder) Upsert(ctx, virtualService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockVirtualServiceClient)(nil).Upsert), ctx, virtualService)
}

// MockDestinationRuleClient is a mock of DestinationRuleClient interface
type MockDestinationRuleClient struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleClientMockRecorder
}

// MockDestinationRuleClientMockRecorder is the mock recorder for MockDestinationRuleClient
type MockDestinationRuleClientMockRecorder struct {
	mock *MockDestinationRuleClient
}

// NewMockDestinationRuleClient creates a new mock instance
func NewMockDestinationRuleClient(ctrl *gomock.Controller) *MockDestinationRuleClient {
	mock := &MockDestinationRuleClient{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDestinationRuleClient) EXPECT() *MockDestinationRuleClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDestinationRuleClient) Get(ctx context.Context, key client.ObjectKey) (*v1alpha3.DestinationRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v1alpha3.DestinationRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDestinationRuleClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDestinationRuleClient)(nil).Get), ctx, key)
}

// Create mocks base method
func (m *MockDestinationRuleClient) Create(ctx context.Context, destinationRule *v1alpha3.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, destinationRule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockDestinationRuleClientMockRecorder) Create(ctx, destinationRule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDestinationRuleClient)(nil).Create), ctx, destinationRule)
}

// Update mocks base method
func (m *MockDestinationRuleClient) Update(ctx context.Context, destinationRule *v1alpha3.DestinationRule, options ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destinationRule}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDestinationRuleClientMockRecorder) Update(ctx, destinationRule interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destinationRule}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDestinationRuleClient)(nil).Update), varargs...)
}

// Upsert mocks base method
func (m *MockDestinationRuleClient) Upsert(ctx context.Context, destinationRule *v1alpha3.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, destinationRule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockDestinationRuleClientMockRecorder) Upsert(ctx, destinationRule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDestinationRuleClient)(nil).Upsert), ctx, destinationRule)
}