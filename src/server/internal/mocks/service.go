// Code generated by MockGen. DO NOT EDIT.
// Source: file_transfer/server/internal/service (interfaces: Service)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	repository "file_transfer/server/internal/repository"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetAllFilesNames mocks base method.
func (m *MockService) GetAllFilesNames(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilesNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilesNames indicates an expected call of GetAllFilesNames.
func (mr *MockServiceMockRecorder) GetAllFilesNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilesNames", reflect.TypeOf((*MockService)(nil).GetAllFilesNames), arg0)
}

// GetFile mocks base method.
func (m *MockService) GetFile(arg0 context.Context, arg1 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", arg0, arg1)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockServiceMockRecorder) GetFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockService)(nil).GetFile), arg0, arg1)
}

// GetFileInfo mocks base method.
func (m *MockService) GetFileInfo(arg0 context.Context, arg1 string) (*repository.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileInfo", arg0, arg1)
	ret0, _ := ret[0].(*repository.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileInfo indicates an expected call of GetFileInfo.
func (mr *MockServiceMockRecorder) GetFileInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileInfo", reflect.TypeOf((*MockService)(nil).GetFileInfo), arg0, arg1)
}
