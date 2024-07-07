// Code generated by MockGen. DO NOT EDIT.
// Source: users_grpc.pb.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserServiceClient) GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserById", varargs...)
	ret0, _ := ret[0].(*UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceClientMockRecorder) GetUserById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserById), varargs...)
}

// GetUsersByIds mocks base method.
func (m *MockUserServiceClient) GetUsersByIds(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_GetUsersByIdsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersByIds", varargs...)
	ret0, _ := ret[0].(UserService_GetUsersByIdsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByIds indicates an expected call of GetUsersByIds.
func (mr *MockUserServiceClientMockRecorder) GetUsersByIds(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIds", reflect.TypeOf((*MockUserServiceClient)(nil).GetUsersByIds), varargs...)
}

// SearchUsers mocks base method.
func (m *MockUserServiceClient) SearchUsers(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (UserService_SearchUsersClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchUsers", varargs...)
	ret0, _ := ret[0].(UserService_SearchUsersClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockUserServiceClientMockRecorder) SearchUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockUserServiceClient)(nil).SearchUsers), varargs...)
}

// MockUserService_GetUsersByIdsClient is a mock of UserService_GetUsersByIdsClient interface.
type MockUserService_GetUsersByIdsClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersByIdsClientMockRecorder
}

// MockUserService_GetUsersByIdsClientMockRecorder is the mock recorder for MockUserService_GetUsersByIdsClient.
type MockUserService_GetUsersByIdsClientMockRecorder struct {
	mock *MockUserService_GetUsersByIdsClient
}

// NewMockUserService_GetUsersByIdsClient creates a new mock instance.
func NewMockUserService_GetUsersByIdsClient(ctrl *gomock.Controller) *MockUserService_GetUsersByIdsClient {
	mock := &MockUserService_GetUsersByIdsClient{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersByIdsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_GetUsersByIdsClient) EXPECT() *MockUserService_GetUsersByIdsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockUserService_GetUsersByIdsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockUserService_GetUsersByIdsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockUserService_GetUsersByIdsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockUserService_GetUsersByIdsClient) Recv() (*UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_GetUsersByIdsClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_GetUsersByIdsClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockUserService_GetUsersByIdsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockUserService_GetUsersByIdsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUserService_GetUsersByIdsClient)(nil).Trailer))
}

// MockUserService_SearchUsersClient is a mock of UserService_SearchUsersClient interface.
type MockUserService_SearchUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_SearchUsersClientMockRecorder
}

// MockUserService_SearchUsersClientMockRecorder is the mock recorder for MockUserService_SearchUsersClient.
type MockUserService_SearchUsersClientMockRecorder struct {
	mock *MockUserService_SearchUsersClient
}

// NewMockUserService_SearchUsersClient creates a new mock instance.
func NewMockUserService_SearchUsersClient(ctrl *gomock.Controller) *MockUserService_SearchUsersClient {
	mock := &MockUserService_SearchUsersClient{ctrl: ctrl}
	mock.recorder = &MockUserService_SearchUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_SearchUsersClient) EXPECT() *MockUserService_SearchUsersClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockUserService_SearchUsersClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockUserService_SearchUsersClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockUserService_SearchUsersClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_SearchUsersClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).Context))
}

// Header mocks base method.
func (m *MockUserService_SearchUsersClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockUserService_SearchUsersClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockUserService_SearchUsersClient) Recv() (*UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockUserService_SearchUsersClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_SearchUsersClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_SearchUsersClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_SearchUsersClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_SearchUsersClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockUserService_SearchUsersClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockUserService_SearchUsersClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUserService_SearchUsersClient)(nil).Trailer))
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserServiceServer) GetUserById(arg0 context.Context, arg1 *UserRequest) (*UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(*UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceServerMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserServiceServer)(nil).GetUserById), arg0, arg1)
}

// GetUsersByIds mocks base method.
func (m *MockUserServiceServer) GetUsersByIds(arg0 *UsersRequest, arg1 UserService_GetUsersByIdsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByIds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUsersByIds indicates an expected call of GetUsersByIds.
func (mr *MockUserServiceServerMockRecorder) GetUsersByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIds", reflect.TypeOf((*MockUserServiceServer)(nil).GetUsersByIds), arg0, arg1)
}

// SearchUsers mocks base method.
func (m *MockUserServiceServer) SearchUsers(arg0 *SearchRequest, arg1 UserService_SearchUsersServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockUserServiceServerMockRecorder) SearchUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockUserServiceServer)(nil).SearchUsers), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUserService_GetUsersByIdsServer is a mock of UserService_GetUsersByIdsServer interface.
type MockUserService_GetUsersByIdsServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_GetUsersByIdsServerMockRecorder
}

// MockUserService_GetUsersByIdsServerMockRecorder is the mock recorder for MockUserService_GetUsersByIdsServer.
type MockUserService_GetUsersByIdsServerMockRecorder struct {
	mock *MockUserService_GetUsersByIdsServer
}

// NewMockUserService_GetUsersByIdsServer creates a new mock instance.
func NewMockUserService_GetUsersByIdsServer(ctrl *gomock.Controller) *MockUserService_GetUsersByIdsServer {
	mock := &MockUserService_GetUsersByIdsServer{ctrl: ctrl}
	mock.recorder = &MockUserService_GetUsersByIdsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_GetUsersByIdsServer) EXPECT() *MockUserService_GetUsersByIdsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockUserService_GetUsersByIdsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_GetUsersByIdsServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockUserService_GetUsersByIdsServer) Send(arg0 *UserResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockUserService_GetUsersByIdsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_GetUsersByIdsServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockUserService_GetUsersByIdsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockUserService_GetUsersByIdsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockUserService_GetUsersByIdsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockUserService_GetUsersByIdsServer)(nil).SetTrailer), arg0)
}

// MockUserService_SearchUsersServer is a mock of UserService_SearchUsersServer interface.
type MockUserService_SearchUsersServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserService_SearchUsersServerMockRecorder
}

// MockUserService_SearchUsersServerMockRecorder is the mock recorder for MockUserService_SearchUsersServer.
type MockUserService_SearchUsersServerMockRecorder struct {
	mock *MockUserService_SearchUsersServer
}

// NewMockUserService_SearchUsersServer creates a new mock instance.
func NewMockUserService_SearchUsersServer(ctrl *gomock.Controller) *MockUserService_SearchUsersServer {
	mock := &MockUserService_SearchUsersServer{ctrl: ctrl}
	mock.recorder = &MockUserService_SearchUsersServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService_SearchUsersServer) EXPECT() *MockUserService_SearchUsersServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockUserService_SearchUsersServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserService_SearchUsersServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockUserService_SearchUsersServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserService_SearchUsersServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockUserService_SearchUsersServer) Send(arg0 *UserResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockUserService_SearchUsersServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockUserService_SearchUsersServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockUserService_SearchUsersServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockUserService_SearchUsersServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserService_SearchUsersServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockUserService_SearchUsersServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockUserService_SearchUsersServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockUserService_SearchUsersServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockUserService_SearchUsersServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockUserService_SearchUsersServer)(nil).SetTrailer), arg0)
}
