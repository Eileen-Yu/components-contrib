// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb-client-go/api (interfaces: WriteAPIBlocking,QueryAPI)

// Package mock_api is a generated GoMock package.
package influx

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influxdb-client-go/api"
	write "github.com/influxdata/influxdb-client-go/api/write"
	domain "github.com/influxdata/influxdb-client-go/domain"
)

// MockWriteAPIBlocking is a mock of WriteAPIBlocking interface.
type MockWriteAPIBlocking struct {
	ctrl     *gomock.Controller
	recorder *MockWriteAPIBlockingMockRecorder
}

// MockWriteAPIBlockingMockRecorder is the mock recorder for MockWriteAPIBlocking.
type MockWriteAPIBlockingMockRecorder struct {
	mock *MockWriteAPIBlocking
}

// NewMockWriteAPIBlocking creates a new mock instance.
func NewMockWriteAPIBlocking(ctrl *gomock.Controller) *MockWriteAPIBlocking {
	mock := &MockWriteAPIBlocking{ctrl: ctrl}
	mock.recorder = &MockWriteAPIBlockingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteAPIBlocking) EXPECT() *MockWriteAPIBlockingMockRecorder {
	return m.recorder
}

// WritePoint mocks base method.
func (m *MockWriteAPIBlocking) WritePoint(arg0 context.Context, arg1 ...*write.Point) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WritePoint", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePoint indicates an expected call of WritePoint.
func (mr *MockWriteAPIBlockingMockRecorder) WritePoint(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePoint", reflect.TypeOf((*MockWriteAPIBlocking)(nil).WritePoint), varargs...)
}

// WriteRecord mocks base method.
func (m *MockWriteAPIBlocking) WriteRecord(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteRecord", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRecord indicates an expected call of WriteRecord.
func (mr *MockWriteAPIBlockingMockRecorder) WriteRecord(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRecord", reflect.TypeOf((*MockWriteAPIBlocking)(nil).WriteRecord), varargs...)
}

// MockQueryAPI is a mock of QueryAPI interface.
type MockQueryAPI struct {
	ctrl     *gomock.Controller
	recorder *MockQueryAPIMockRecorder
}

// MockQueryAPIMockRecorder is the mock recorder for MockQueryAPI.
type MockQueryAPIMockRecorder struct {
	mock *MockQueryAPI
}

// NewMockQueryAPI creates a new mock instance.
func NewMockQueryAPI(ctrl *gomock.Controller) *MockQueryAPI {
	mock := &MockQueryAPI{ctrl: ctrl}
	mock.recorder = &MockQueryAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryAPI) EXPECT() *MockQueryAPIMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockQueryAPI) Query(arg0 context.Context, arg1 string) (*api.QueryTableResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(*api.QueryTableResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockQueryAPIMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockQueryAPI)(nil).Query), arg0, arg1)
}

// QueryRaw mocks base method.
func (m *MockQueryAPI) QueryRaw(arg0 context.Context, arg1 string, arg2 *domain.Dialect) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRaw", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRaw indicates an expected call of QueryRaw.
func (mr *MockQueryAPIMockRecorder) QueryRaw(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRaw", reflect.TypeOf((*MockQueryAPI)(nil).QueryRaw), arg0, arg1, arg2)
}