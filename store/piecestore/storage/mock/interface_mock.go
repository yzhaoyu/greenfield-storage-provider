// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	storage "github.com/bnb-chain/greenfield-storage-provider/store/piecestore/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockObjectStorage is a mock of ObjectStorage interface.
type MockObjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStorageMockRecorder
}

// MockObjectStorageMockRecorder is the mock recorder for MockObjectStorage.
type MockObjectStorageMockRecorder struct {
	mock *MockObjectStorage
}

// NewMockObjectStorage creates a new mock instance.
func NewMockObjectStorage(ctrl *gomock.Controller) *MockObjectStorage {
	mock := &MockObjectStorage{ctrl: ctrl}
	mock.recorder = &MockObjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStorage) EXPECT() *MockObjectStorageMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockObjectStorage) CreateBucket(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockObjectStorageMockRecorder) CreateBucket(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockObjectStorage)(nil).CreateBucket), ctx)
}

// DeleteObject mocks base method.
func (m *MockObjectStorage) DeleteObject(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStorageMockRecorder) DeleteObject(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStorage)(nil).DeleteObject), ctx, key)
}

// GetObject mocks base method.
func (m *MockObjectStorage) GetObject(ctx context.Context, key string, offset, limit int64) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, key, offset, limit)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockObjectStorageMockRecorder) GetObject(ctx, key, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockObjectStorage)(nil).GetObject), ctx, key, offset, limit)
}

// HeadBucket mocks base method.
func (m *MockObjectStorage) HeadBucket(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadBucket", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HeadBucket indicates an expected call of HeadBucket.
func (mr *MockObjectStorageMockRecorder) HeadBucket(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBucket", reflect.TypeOf((*MockObjectStorage)(nil).HeadBucket), ctx)
}

// HeadObject mocks base method.
func (m *MockObjectStorage) HeadObject(ctx context.Context, key string) (storage.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadObject", ctx, key)
	ret0, _ := ret[0].(storage.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject.
func (mr *MockObjectStorageMockRecorder) HeadObject(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockObjectStorage)(nil).HeadObject), ctx, key)
}

// ListAllObjects mocks base method.
func (m *MockObjectStorage) ListAllObjects(ctx context.Context, prefix, marker string) (<-chan storage.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllObjects", ctx, prefix, marker)
	ret0, _ := ret[0].(<-chan storage.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllObjects indicates an expected call of ListAllObjects.
func (mr *MockObjectStorageMockRecorder) ListAllObjects(ctx, prefix, marker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllObjects", reflect.TypeOf((*MockObjectStorage)(nil).ListAllObjects), ctx, prefix, marker)
}

// ListObjects mocks base method.
func (m *MockObjectStorage) ListObjects(ctx context.Context, prefix, marker, delimiter string, limit int64) ([]storage.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", ctx, prefix, marker, delimiter, limit)
	ret0, _ := ret[0].([]storage.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects.
func (mr *MockObjectStorageMockRecorder) ListObjects(ctx, prefix, marker, delimiter, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockObjectStorage)(nil).ListObjects), ctx, prefix, marker, delimiter, limit)
}

// PutObject mocks base method.
func (m *MockObjectStorage) PutObject(ctx context.Context, key string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", ctx, key, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject.
func (mr *MockObjectStorageMockRecorder) PutObject(ctx, key, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockObjectStorage)(nil).PutObject), ctx, key, reader)
}

// String mocks base method.
func (m *MockObjectStorage) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockObjectStorageMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockObjectStorage)(nil).String))
}

// MockObject is a mock of Object interface.
type MockObject struct {
	ctrl     *gomock.Controller
	recorder *MockObjectMockRecorder
}

// MockObjectMockRecorder is the mock recorder for MockObject.
type MockObjectMockRecorder struct {
	mock *MockObject
}

// NewMockObject creates a new mock instance.
func NewMockObject(ctrl *gomock.Controller) *MockObject {
	mock := &MockObject{ctrl: ctrl}
	mock.recorder = &MockObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObject) EXPECT() *MockObjectMockRecorder {
	return m.recorder
}

// IsSymlink mocks base method.
func (m *MockObject) IsSymlink() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSymlink")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSymlink indicates an expected call of IsSymlink.
func (mr *MockObjectMockRecorder) IsSymlink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSymlink", reflect.TypeOf((*MockObject)(nil).IsSymlink))
}

// Key mocks base method.
func (m *MockObject) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockObjectMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockObject)(nil).Key))
}

// ModTime mocks base method.
func (m *MockObject) ModTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// ModTime indicates an expected call of ModTime.
func (mr *MockObjectMockRecorder) ModTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModTime", reflect.TypeOf((*MockObject)(nil).ModTime))
}

// Size mocks base method.
func (m *MockObject) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockObjectMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockObject)(nil).Size))
}
