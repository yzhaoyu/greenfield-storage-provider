// Code generated by MockGen. DO NOT EDIT.
// Source: ./syncer_client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v1 "github.com/bnb-chain/greenfield-storage-provider/service/types/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSyncerAPI is a mock of SyncerAPI interface.
type MockSyncerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSyncerAPIMockRecorder
}

// MockSyncerAPIMockRecorder is the mock recorder for MockSyncerAPI.
type MockSyncerAPIMockRecorder struct {
	mock *MockSyncerAPI
}

// NewMockSyncerAPI creates a new mock instance.
func NewMockSyncerAPI(ctrl *gomock.Controller) *MockSyncerAPI {
	mock := &MockSyncerAPI{ctrl: ctrl}
	mock.recorder = &MockSyncerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncerAPI) EXPECT() *MockSyncerAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSyncerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncerAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncerAPI)(nil).Close))
}

// SyncPiece mocks base method.
func (m *MockSyncerAPI) SyncPiece(ctx context.Context, opts ...grpc.CallOption) (v1.SyncerService_SyncPieceClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPiece", varargs...)
	ret0, _ := ret[0].(v1.SyncerService_SyncPieceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPiece indicates an expected call of SyncPiece.
func (mr *MockSyncerAPIMockRecorder) SyncPiece(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPiece", reflect.TypeOf((*MockSyncerAPI)(nil).SyncPiece), varargs...)
}
