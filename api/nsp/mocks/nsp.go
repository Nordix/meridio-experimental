// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/nsp/nsp.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	nsp "github.com/nordix/meridio/api/nsp"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockNetworkServicePlateformServiceClient is a mock of NetworkServicePlateformServiceClient interface.
type MockNetworkServicePlateformServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServicePlateformServiceClientMockRecorder
}

// MockNetworkServicePlateformServiceClientMockRecorder is the mock recorder for MockNetworkServicePlateformServiceClient.
type MockNetworkServicePlateformServiceClientMockRecorder struct {
	mock *MockNetworkServicePlateformServiceClient
}

// NewMockNetworkServicePlateformServiceClient creates a new mock instance.
func NewMockNetworkServicePlateformServiceClient(ctrl *gomock.Controller) *MockNetworkServicePlateformServiceClient {
	mock := &MockNetworkServicePlateformServiceClient{ctrl: ctrl}
	mock.recorder = &MockNetworkServicePlateformServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkServicePlateformServiceClient) EXPECT() *MockNetworkServicePlateformServiceClientMockRecorder {
	return m.recorder
}

// GetTargets mocks base method.
func (m *MockNetworkServicePlateformServiceClient) GetTargets(ctx context.Context, in *nsp.TargetType, opts ...grpc.CallOption) (*nsp.GetTargetsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTargets", varargs...)
	ret0, _ := ret[0].(*nsp.GetTargetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargets indicates an expected call of GetTargets.
func (mr *MockNetworkServicePlateformServiceClientMockRecorder) GetTargets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargets", reflect.TypeOf((*MockNetworkServicePlateformServiceClient)(nil).GetTargets), varargs...)
}

// Monitor mocks base method.
func (m *MockNetworkServicePlateformServiceClient) Monitor(ctx context.Context, in *nsp.TargetType, opts ...grpc.CallOption) (nsp.NetworkServicePlateformService_MonitorClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Monitor", varargs...)
	ret0, _ := ret[0].(nsp.NetworkServicePlateformService_MonitorClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Monitor indicates an expected call of Monitor.
func (mr *MockNetworkServicePlateformServiceClientMockRecorder) Monitor(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Monitor", reflect.TypeOf((*MockNetworkServicePlateformServiceClient)(nil).Monitor), varargs...)
}

// Register mocks base method.
func (m *MockNetworkServicePlateformServiceClient) Register(ctx context.Context, in *nsp.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
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
func (mr *MockNetworkServicePlateformServiceClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockNetworkServicePlateformServiceClient)(nil).Register), varargs...)
}

// Unregister mocks base method.
func (m *MockNetworkServicePlateformServiceClient) Unregister(ctx context.Context, in *nsp.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
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
func (mr *MockNetworkServicePlateformServiceClientMockRecorder) Unregister(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockNetworkServicePlateformServiceClient)(nil).Unregister), varargs...)
}

// Update mocks base method.
func (m *MockNetworkServicePlateformServiceClient) Update(ctx context.Context, in *nsp.Target, opts ...grpc.CallOption) (*empty.Empty, error) {
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
func (mr *MockNetworkServicePlateformServiceClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetworkServicePlateformServiceClient)(nil).Update), varargs...)
}

// MockNetworkServicePlateformService_MonitorClient is a mock of NetworkServicePlateformService_MonitorClient interface.
type MockNetworkServicePlateformService_MonitorClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServicePlateformService_MonitorClientMockRecorder
}

// MockNetworkServicePlateformService_MonitorClientMockRecorder is the mock recorder for MockNetworkServicePlateformService_MonitorClient.
type MockNetworkServicePlateformService_MonitorClientMockRecorder struct {
	mock *MockNetworkServicePlateformService_MonitorClient
}

// NewMockNetworkServicePlateformService_MonitorClient creates a new mock instance.
func NewMockNetworkServicePlateformService_MonitorClient(ctrl *gomock.Controller) *MockNetworkServicePlateformService_MonitorClient {
	mock := &MockNetworkServicePlateformService_MonitorClient{ctrl: ctrl}
	mock.recorder = &MockNetworkServicePlateformService_MonitorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkServicePlateformService_MonitorClient) EXPECT() *MockNetworkServicePlateformService_MonitorClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockNetworkServicePlateformService_MonitorClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockNetworkServicePlateformService_MonitorClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).Context))
}

// Header mocks base method.
func (m *MockNetworkServicePlateformService_MonitorClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockNetworkServicePlateformService_MonitorClient) Recv() (*nsp.TargetEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*nsp.TargetEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockNetworkServicePlateformService_MonitorClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockNetworkServicePlateformService_MonitorClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockNetworkServicePlateformService_MonitorClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockNetworkServicePlateformService_MonitorClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorClient)(nil).Trailer))
}

// MockNetworkServicePlateformServiceServer is a mock of NetworkServicePlateformServiceServer interface.
type MockNetworkServicePlateformServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServicePlateformServiceServerMockRecorder
}

// MockNetworkServicePlateformServiceServerMockRecorder is the mock recorder for MockNetworkServicePlateformServiceServer.
type MockNetworkServicePlateformServiceServerMockRecorder struct {
	mock *MockNetworkServicePlateformServiceServer
}

// NewMockNetworkServicePlateformServiceServer creates a new mock instance.
func NewMockNetworkServicePlateformServiceServer(ctrl *gomock.Controller) *MockNetworkServicePlateformServiceServer {
	mock := &MockNetworkServicePlateformServiceServer{ctrl: ctrl}
	mock.recorder = &MockNetworkServicePlateformServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkServicePlateformServiceServer) EXPECT() *MockNetworkServicePlateformServiceServerMockRecorder {
	return m.recorder
}

// GetTargets mocks base method.
func (m *MockNetworkServicePlateformServiceServer) GetTargets(arg0 context.Context, arg1 *nsp.TargetType) (*nsp.GetTargetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargets", arg0, arg1)
	ret0, _ := ret[0].(*nsp.GetTargetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargets indicates an expected call of GetTargets.
func (mr *MockNetworkServicePlateformServiceServerMockRecorder) GetTargets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargets", reflect.TypeOf((*MockNetworkServicePlateformServiceServer)(nil).GetTargets), arg0, arg1)
}

// Monitor mocks base method.
func (m *MockNetworkServicePlateformServiceServer) Monitor(arg0 *nsp.TargetType, arg1 nsp.NetworkServicePlateformService_MonitorServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Monitor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Monitor indicates an expected call of Monitor.
func (mr *MockNetworkServicePlateformServiceServerMockRecorder) Monitor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Monitor", reflect.TypeOf((*MockNetworkServicePlateformServiceServer)(nil).Monitor), arg0, arg1)
}

// Register mocks base method.
func (m *MockNetworkServicePlateformServiceServer) Register(arg0 context.Context, arg1 *nsp.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockNetworkServicePlateformServiceServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockNetworkServicePlateformServiceServer)(nil).Register), arg0, arg1)
}

// Unregister mocks base method.
func (m *MockNetworkServicePlateformServiceServer) Unregister(arg0 context.Context, arg1 *nsp.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unregister indicates an expected call of Unregister.
func (mr *MockNetworkServicePlateformServiceServerMockRecorder) Unregister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockNetworkServicePlateformServiceServer)(nil).Unregister), arg0, arg1)
}

// Update mocks base method.
func (m *MockNetworkServicePlateformServiceServer) Update(arg0 context.Context, arg1 *nsp.Target) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockNetworkServicePlateformServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetworkServicePlateformServiceServer)(nil).Update), arg0, arg1)
}

// MockNetworkServicePlateformService_MonitorServer is a mock of NetworkServicePlateformService_MonitorServer interface.
type MockNetworkServicePlateformService_MonitorServer struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServicePlateformService_MonitorServerMockRecorder
}

// MockNetworkServicePlateformService_MonitorServerMockRecorder is the mock recorder for MockNetworkServicePlateformService_MonitorServer.
type MockNetworkServicePlateformService_MonitorServerMockRecorder struct {
	mock *MockNetworkServicePlateformService_MonitorServer
}

// NewMockNetworkServicePlateformService_MonitorServer creates a new mock instance.
func NewMockNetworkServicePlateformService_MonitorServer(ctrl *gomock.Controller) *MockNetworkServicePlateformService_MonitorServer {
	mock := &MockNetworkServicePlateformService_MonitorServer{ctrl: ctrl}
	mock.recorder = &MockNetworkServicePlateformService_MonitorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkServicePlateformService_MonitorServer) EXPECT() *MockNetworkServicePlateformService_MonitorServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNetworkServicePlateformService_MonitorServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockNetworkServicePlateformService_MonitorServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockNetworkServicePlateformService_MonitorServer) Send(arg0 *nsp.TargetEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNetworkServicePlateformService_MonitorServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockNetworkServicePlateformService_MonitorServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockNetworkServicePlateformService_MonitorServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNetworkServicePlateformService_MonitorServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNetworkServicePlateformService_MonitorServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNetworkServicePlateformService_MonitorServer)(nil).SetTrailer), arg0)
}
