// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/types.go

// Package storagemock is a generated GoMock package.
package storagemock

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/rakoo/raft/internal/storage"
	raftpb "go.etcd.io/etcd/raft/v3/raftpb"
)

// MockSnapshotter is a mock of Snapshotter interface.
type MockSnapshotter struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotterMockRecorder
}

// MockSnapshotterMockRecorder is the mock recorder for MockSnapshotter.
type MockSnapshotterMockRecorder struct {
	mock *MockSnapshotter
}

// NewMockSnapshotter creates a new mock instance.
func NewMockSnapshotter(ctrl *gomock.Controller) *MockSnapshotter {
	mock := &MockSnapshotter{ctrl: ctrl}
	mock.recorder = &MockSnapshotterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotter) EXPECT() *MockSnapshotterMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSnapshotter) Read(arg0, arg1 uint64) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSnapshotterMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSnapshotter)(nil).Read), arg0, arg1)
}

// ReadFrom mocks base method.
func (m *MockSnapshotter) ReadFrom(arg0 string) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", arg0)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFrom indicates an expected call of ReadFrom.
func (mr *MockSnapshotterMockRecorder) ReadFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockSnapshotter)(nil).ReadFrom), arg0)
}

// Reader mocks base method.
func (m *MockSnapshotter) Reader(arg0, arg1 uint64) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader.
func (mr *MockSnapshotterMockRecorder) Reader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockSnapshotter)(nil).Reader), arg0, arg1)
}

// Write mocks base method.
func (m *MockSnapshotter) Write(arg0 *storage.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockSnapshotterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSnapshotter)(nil).Write), arg0)
}

// Writer mocks base method.
func (m *MockSnapshotter) Writer(arg0, arg1 uint64) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer", arg0, arg1)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Writer indicates an expected call of Writer.
func (mr *MockSnapshotterMockRecorder) Writer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockSnapshotter)(nil).Writer), arg0, arg1)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Boot mocks base method.
func (m *MockStorage) Boot(arg0 []byte) ([]byte, raftpb.HardState, []raftpb.Entry, *storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Boot", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(raftpb.HardState)
	ret2, _ := ret[2].([]raftpb.Entry)
	ret3, _ := ret[3].(*storage.Snapshot)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// Boot indicates an expected call of Boot.
func (mr *MockStorageMockRecorder) Boot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockStorage)(nil).Boot), arg0)
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// Exist mocks base method.
func (m *MockStorage) Exist() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist.
func (mr *MockStorageMockRecorder) Exist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockStorage)(nil).Exist))
}

// SaveEntries mocks base method.
func (m *MockStorage) SaveEntries(arg0 raftpb.HardState, arg1 []raftpb.Entry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEntries", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveEntries indicates an expected call of SaveEntries.
func (mr *MockStorageMockRecorder) SaveEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEntries", reflect.TypeOf((*MockStorage)(nil).SaveEntries), arg0, arg1)
}

// SaveSnapshot mocks base method.
func (m *MockStorage) SaveSnapshot(arg0 raftpb.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSnapshot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSnapshot indicates an expected call of SaveSnapshot.
func (mr *MockStorageMockRecorder) SaveSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSnapshot", reflect.TypeOf((*MockStorage)(nil).SaveSnapshot), arg0)
}

// Snapshotter mocks base method.
func (m *MockStorage) Snapshotter() storage.Snapshotter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshotter")
	ret0, _ := ret[0].(storage.Snapshotter)
	return ret0
}

// Snapshotter indicates an expected call of Snapshotter.
func (mr *MockStorageMockRecorder) Snapshotter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshotter", reflect.TypeOf((*MockStorage)(nil).Snapshotter))
}
