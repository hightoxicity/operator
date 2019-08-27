// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/operator/drivers/storage (interfaces: Driver)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/libopenstorage/operator/drivers/storage"
	v1alpha1 "github.com/libopenstorage/operator/pkg/apis/core/v1alpha1"
	v1 "k8s.io/api/core/v1"
	record "k8s.io/client-go/tools/record"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockDriver is a mock of Driver interface
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// DeleteStorage mocks base method
func (m *MockDriver) DeleteStorage(arg0 *v1alpha1.StorageCluster) (*v1alpha1.ClusterCondition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStorage", arg0)
	ret0, _ := ret[0].(*v1alpha1.ClusterCondition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStorage indicates an expected call of DeleteStorage
func (mr *MockDriverMockRecorder) DeleteStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStorage", reflect.TypeOf((*MockDriver)(nil).DeleteStorage), arg0)
}

// GetSelectorLabels mocks base method
func (m *MockDriver) GetSelectorLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectorLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetSelectorLabels indicates an expected call of GetSelectorLabels
func (mr *MockDriverMockRecorder) GetSelectorLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectorLabels", reflect.TypeOf((*MockDriver)(nil).GetSelectorLabels))
}

// GetStoragePodSpec mocks base method
func (m *MockDriver) GetStoragePodSpec(arg0 *v1alpha1.StorageCluster, arg1 string) (v1.PodSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoragePodSpec", arg0, arg1)
	ret0, _ := ret[0].(v1.PodSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoragePodSpec indicates an expected call of GetStoragePodSpec
func (mr *MockDriverMockRecorder) GetStoragePodSpec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePodSpec", reflect.TypeOf((*MockDriver)(nil).GetStoragePodSpec), arg0, arg1)
}

// GetStorkDriverName mocks base method
func (m *MockDriver) GetStorkDriverName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorkDriverName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorkDriverName indicates an expected call of GetStorkDriverName
func (mr *MockDriverMockRecorder) GetStorkDriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorkDriverName", reflect.TypeOf((*MockDriver)(nil).GetStorkDriverName))
}

// GetStorkEnvList mocks base method
func (m *MockDriver) GetStorkEnvList(arg0 *v1alpha1.StorageCluster) []v1.EnvVar {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorkEnvList", arg0)
	ret0, _ := ret[0].([]v1.EnvVar)
	return ret0
}

// GetStorkEnvList indicates an expected call of GetStorkEnvList
func (mr *MockDriverMockRecorder) GetStorkEnvList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorkEnvList", reflect.TypeOf((*MockDriver)(nil).GetStorkEnvList), arg0)
}

// Init mocks base method
func (m *MockDriver) Init(arg0 client.Client, arg1 record.EventRecorder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockDriverMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDriver)(nil).Init), arg0, arg1)
}

// PreInstall mocks base method
func (m *MockDriver) PreInstall(arg0 *v1alpha1.StorageCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreInstall", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreInstall indicates an expected call of PreInstall
func (mr *MockDriverMockRecorder) PreInstall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreInstall", reflect.TypeOf((*MockDriver)(nil).PreInstall), arg0)
}

// SetDefaultsOnStorageCluster mocks base method
func (m *MockDriver) SetDefaultsOnStorageCluster(arg0 *v1alpha1.StorageCluster) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDefaultsOnStorageCluster", arg0)
}

// SetDefaultsOnStorageCluster indicates an expected call of SetDefaultsOnStorageCluster
func (mr *MockDriverMockRecorder) SetDefaultsOnStorageCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultsOnStorageCluster", reflect.TypeOf((*MockDriver)(nil).SetDefaultsOnStorageCluster), arg0)
}

// String mocks base method
func (m *MockDriver) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockDriverMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockDriver)(nil).String))
}

// UpdateDriver mocks base method
func (m *MockDriver) UpdateDriver(arg0 *storage.UpdateDriverInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDriver", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDriver indicates an expected call of UpdateDriver
func (mr *MockDriverMockRecorder) UpdateDriver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDriver", reflect.TypeOf((*MockDriver)(nil).UpdateDriver), arg0)
}

// UpdateStorageClusterStatus mocks base method
func (m *MockDriver) UpdateStorageClusterStatus(arg0 *v1alpha1.StorageCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStorageClusterStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStorageClusterStatus indicates an expected call of UpdateStorageClusterStatus
func (mr *MockDriverMockRecorder) UpdateStorageClusterStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStorageClusterStatus", reflect.TypeOf((*MockDriver)(nil).UpdateStorageClusterStatus), arg0)
}
