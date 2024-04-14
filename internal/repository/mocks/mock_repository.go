// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/repository.go -destination=/home/evgeny/avito/internal/repository/mocks/mock_repository.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/loveavoider/avito-banners/internal/model"
	merror "github.com/loveavoider/avito-banners/merror"
	gomock "go.uber.org/mock/gomock"
)

// MockBannerRepository is a mock of BannerRepository interface.
type MockBannerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBannerRepositoryMockRecorder
}

// MockBannerRepositoryMockRecorder is the mock recorder for MockBannerRepository.
type MockBannerRepositoryMockRecorder struct {
	mock *MockBannerRepository
}

// NewMockBannerRepository creates a new mock instance.
func NewMockBannerRepository(ctrl *gomock.Controller) *MockBannerRepository {
	mock := &MockBannerRepository{ctrl: ctrl}
	mock.recorder = &MockBannerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBannerRepository) EXPECT() *MockBannerRepositoryMockRecorder {
	return m.recorder
}

// CheckUnique mocks base method.
func (m *MockBannerRepository) CheckUnique(arg0 int) ([]uint, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUnique", arg0)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// CheckUnique indicates an expected call of CheckUnique.
func (mr *MockBannerRepositoryMockRecorder) CheckUnique(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUnique", reflect.TypeOf((*MockBannerRepository)(nil).CheckUnique), arg0)
}

// CreateBanner mocks base method.
func (m *MockBannerRepository) CreateBanner(arg0 model.Banner) (uint, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBanner", arg0)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// CreateBanner indicates an expected call of CreateBanner.
func (mr *MockBannerRepositoryMockRecorder) CreateBanner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBanner", reflect.TypeOf((*MockBannerRepository)(nil).CreateBanner), arg0)
}

// DeleteBanner mocks base method.
func (m *MockBannerRepository) DeleteBanner(arg0 model.Banner) *merror.MError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", arg0)
	ret0, _ := ret[0].(*merror.MError)
	return ret0
}

// DeleteBanner indicates an expected call of DeleteBanner.
func (mr *MockBannerRepositoryMockRecorder) DeleteBanner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockBannerRepository)(nil).DeleteBanner), arg0)
}

// GetBanners mocks base method.
func (m *MockBannerRepository) GetBanners(arg0 model.GetBanners) ([]model.BannerResponse, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanners", arg0)
	ret0, _ := ret[0].([]model.BannerResponse)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// GetBanners indicates an expected call of GetBanners.
func (mr *MockBannerRepositoryMockRecorder) GetBanners(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanners", reflect.TypeOf((*MockBannerRepository)(nil).GetBanners), arg0)
}

// GetBannersByFeature mocks base method.
func (m *MockBannerRepository) GetBannersByFeature(arg0 model.GetBanners) ([]model.BannerResponse, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBannersByFeature", arg0)
	ret0, _ := ret[0].([]model.BannerResponse)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// GetBannersByFeature indicates an expected call of GetBannersByFeature.
func (mr *MockBannerRepositoryMockRecorder) GetBannersByFeature(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBannersByFeature", reflect.TypeOf((*MockBannerRepository)(nil).GetBannersByFeature), arg0)
}

// GetBannersByTag mocks base method.
func (m *MockBannerRepository) GetBannersByTag(arg0 model.GetBanners) ([]model.BannerResponse, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBannersByTag", arg0)
	ret0, _ := ret[0].([]model.BannerResponse)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// GetBannersByTag indicates an expected call of GetBannersByTag.
func (mr *MockBannerRepositoryMockRecorder) GetBannersByTag(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBannersByTag", reflect.TypeOf((*MockBannerRepository)(nil).GetBannersByTag), arg0)
}

// GetUserBanner mocks base method.
func (m *MockBannerRepository) GetUserBanner(arg0 model.GetUserBanner, arg1 bool) (model.BannerContent, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBanner", arg0, arg1)
	ret0, _ := ret[0].(model.BannerContent)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// GetUserBanner indicates an expected call of GetUserBanner.
func (mr *MockBannerRepositoryMockRecorder) GetUserBanner(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBanner", reflect.TypeOf((*MockBannerRepository)(nil).GetUserBanner), arg0, arg1)
}

// GetUserBannerWithTags mocks base method.
func (m *MockBannerRepository) GetUserBannerWithTags(arg0 model.GetBanners, arg1 bool) (model.BannerResponse, *merror.MError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBannerWithTags", arg0, arg1)
	ret0, _ := ret[0].(model.BannerResponse)
	ret1, _ := ret[1].(*merror.MError)
	return ret0, ret1
}

// GetUserBannerWithTags indicates an expected call of GetUserBannerWithTags.
func (mr *MockBannerRepositoryMockRecorder) GetUserBannerWithTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBannerWithTags", reflect.TypeOf((*MockBannerRepository)(nil).GetUserBannerWithTags), arg0, arg1)
}

// UpdateBanner mocks base method.
func (m *MockBannerRepository) UpdateBanner(arg0 model.UpdateBanner) *merror.MError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", arg0)
	ret0, _ := ret[0].(*merror.MError)
	return ret0
}

// UpdateBanner indicates an expected call of UpdateBanner.
func (mr *MockBannerRepositoryMockRecorder) UpdateBanner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBanner", reflect.TypeOf((*MockBannerRepository)(nil).UpdateBanner), arg0)
}
