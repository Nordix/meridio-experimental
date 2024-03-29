// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/nsp/v1/targetregistry.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/nordix/meridio/api/nsp/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockTargetRegistryClient is a mock of TargetRegistryClient interface.
type MockTargetRegistryClient struct {
	ctrl     *gomock.Controller
	recorder *MockTargetRegistryClientMockRecorder
}

// MockTargetRegistryClientMockRecorder is the mock recorder for MockTargetRegistryClient.
type MockTargetRegistryClientMockRecorder struct {
	mock *MockTargetRegistryClient
}

// NewMockTargetRegistryClient creates a new mock instance.
func NewMockTargetRegistryClient(ctrl *gomock.Controller) *MockTargetRegistryClient {
	mock := &MockTargetRegistryClient{ctrl: ctrl}
	mock.recorder = &MockTargetRegistryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetRegistryClient) EXPECT() *MockTargetRegistryClientMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockTargetRegistryClient) Register(ctx context.Context, in *v1.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockTargetRegistryClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockTargetRegistryClient)(nil).Register), varargs...)
}

// Unregister mocks base method.
func (m *MockTargetRegistryClient) Unregister(ctx context.Context, in *v1.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unregister", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister.
func (mr *MockTargetRegistryClientMockRecorder) Unregister(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockTargetRegistryClient)(nil).Unregister), varargs...)
}

// Update mocks base method.
func (m *MockTargetRegistryClient) Update(ctx context.Context, in *v1.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTargetRegistryClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTargetRegistryClient)(nil).Update), varargs...)
}

// Watch mocks base method.
func (m *MockTargetRegistryClient) Watch(ctx context.Context, in *v1.Target, opts ...grpc.CallOption) (v1.TargetRegistry_WatchClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(v1.TargetRegistry_WatchClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockTargetRegistryClientMockRecorder) Watch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockTargetRegistryClient)(nil).Watch), varargs...)
}

// MockTargetRegistry_WatchClient is a mock of TargetRegistry_WatchClient interface.
type MockTargetRegistry_WatchClient struct {
	ctrl     *gomock.Controller
	recorder *MockTargetRegistry_WatchClientMockRecorder
}

// MockTargetRegistry_WatchClientMockRecorder is the mock recorder for MockTargetRegistry_WatchClient.
type MockTargetRegistry_WatchClientMockRecorder struct {
	mock *MockTargetRegistry_WatchClient
}

// NewMockTargetRegistry_WatchClient creates a new mock instance.
func NewMockTargetRegistry_WatchClient(ctrl *gomock.Controller) *MockTargetRegistry_WatchClient {
	mock := &MockTargetRegistry_WatchClient{ctrl: ctrl}
	mock.recorder = &MockTargetRegistry_WatchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetRegistry_WatchClient) EXPECT() *MockTargetRegistry_WatchClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockTargetRegistry_WatchClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockTargetRegistry_WatchClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockTargetRegistry_WatchClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTargetRegistry_WatchClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).Context))
}

// Header mocks base method.
func (m *MockTargetRegistry_WatchClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockTargetRegistry_WatchClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockTargetRegistry_WatchClient) Recv() (*v1.TargetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.TargetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockTargetRegistry_WatchClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockTargetRegistry_WatchClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTargetRegistry_WatchClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockTargetRegistry_WatchClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTargetRegistry_WatchClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockTargetRegistry_WatchClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockTargetRegistry_WatchClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTargetRegistry_WatchClient)(nil).Trailer))
}

// MockTargetRegistryServer is a mock of TargetRegistryServer interface.
type MockTargetRegistryServer struct {
	ctrl     *gomock.Controller
	recorder *MockTargetRegistryServerMockRecorder
}

// MockTargetRegistryServerMockRecorder is the mock recorder for MockTargetRegistryServer.
type MockTargetRegistryServerMockRecorder struct {
	mock *MockTargetRegistryServer
}

// NewMockTargetRegistryServer creates a new mock instance.
func NewMockTargetRegistryServer(ctrl *gomock.Controller) *MockTargetRegistryServer {
	mock := &MockTargetRegistryServer{ctrl: ctrl}
	mock.recorder = &MockTargetRegistryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetRegistryServer) EXPECT() *MockTargetRegistryServerMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockTargetRegistryServer) Register(arg0 context.Context, arg1 *v1.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockTargetRegistryServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockTargetRegistryServer)(nil).Register), arg0, arg1)
}

// Unregister mocks base method.
func (m *MockTargetRegistryServer) Unregister(arg0 context.Context, arg1 *v1.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister.
func (mr *MockTargetRegistryServerMockRecorder) Unregister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockTargetRegistryServer)(nil).Unregister), arg0, arg1)
}

// Update mocks base method.
func (m *MockTargetRegistryServer) Update(arg0 context.Context, arg1 *v1.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTargetRegistryServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTargetRegistryServer)(nil).Update), arg0, arg1)
}

// Watch mocks base method.
func (m *MockTargetRegistryServer) Watch(arg0 *v1.Target, arg1 v1.TargetRegistry_WatchServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockTargetRegistryServerMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockTargetRegistryServer)(nil).Watch), arg0, arg1)
}

// MockTargetRegistry_WatchServer is a mock of TargetRegistry_WatchServer interface.
type MockTargetRegistry_WatchServer struct {
	ctrl     *gomock.Controller
	recorder *MockTargetRegistry_WatchServerMockRecorder
}

// MockTargetRegistry_WatchServerMockRecorder is the mock recorder for MockTargetRegistry_WatchServer.
type MockTargetRegistry_WatchServerMockRecorder struct {
	mock *MockTargetRegistry_WatchServer
}

// NewMockTargetRegistry_WatchServer creates a new mock instance.
func NewMockTargetRegistry_WatchServer(ctrl *gomock.Controller) *MockTargetRegistry_WatchServer {
	mock := &MockTargetRegistry_WatchServer{ctrl: ctrl}
	mock.recorder = &MockTargetRegistry_WatchServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTargetRegistry_WatchServer) EXPECT() *MockTargetRegistry_WatchServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockTargetRegistry_WatchServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTargetRegistry_WatchServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockTargetRegistry_WatchServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTargetRegistry_WatchServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockTargetRegistry_WatchServer) Send(arg0 *v1.TargetResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTargetRegistry_WatchServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockTargetRegistry_WatchServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockTargetRegistry_WatchServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockTargetRegistry_WatchServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTargetRegistry_WatchServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockTargetRegistry_WatchServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockTargetRegistry_WatchServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockTargetRegistry_WatchServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockTargetRegistry_WatchServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockTargetRegistry_WatchServer)(nil).SetTrailer), arg0)
}
