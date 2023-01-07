// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/Krabik6/meal-schedule/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// MockRecipes is a mock of Recipes interface.
type MockRecipes struct {
	ctrl     *gomock.Controller
	recorder *MockRecipesMockRecorder
}

// MockRecipesMockRecorder is the mock recorder for MockRecipes.
type MockRecipesMockRecorder struct {
	mock *MockRecipes
}

// NewMockRecipes creates a new mock instance.
func NewMockRecipes(ctrl *gomock.Controller) *MockRecipes {
	mock := &MockRecipes{ctrl: ctrl}
	mock.recorder = &MockRecipesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipes) EXPECT() *MockRecipesMockRecorder {
	return m.recorder
}

// CreateRecipe mocks base method.
func (m *MockRecipes) CreateRecipe(userId int, recipe models.Recipe) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecipe", userId, recipe)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecipe indicates an expected call of CreateRecipe.
func (mr *MockRecipesMockRecorder) CreateRecipe(userId, recipe interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecipe", reflect.TypeOf((*MockRecipes)(nil).CreateRecipe), userId, recipe)
}

// DeleteRecipe mocks base method.
func (m *MockRecipes) DeleteRecipe(userId, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecipe", userId, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecipe indicates an expected call of DeleteRecipe.
func (mr *MockRecipesMockRecorder) DeleteRecipe(userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecipe", reflect.TypeOf((*MockRecipes)(nil).DeleteRecipe), userId, id)
}

// GetAllRecipes mocks base method.
func (m *MockRecipes) GetAllRecipes(userId int) ([]models.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecipes", userId)
	ret0, _ := ret[0].([]models.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRecipes indicates an expected call of GetAllRecipes.
func (mr *MockRecipesMockRecorder) GetAllRecipes(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRecipes", reflect.TypeOf((*MockRecipes)(nil).GetAllRecipes), userId)
}

// GetRecipeById mocks base method.
func (m *MockRecipes) GetRecipeById(userId, id int) (models.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipeById", userId, id)
	ret0, _ := ret[0].(models.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipeById indicates an expected call of GetRecipeById.
func (mr *MockRecipesMockRecorder) GetRecipeById(userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipeById", reflect.TypeOf((*MockRecipes)(nil).GetRecipeById), userId, id)
}

// UpdateRecipe mocks base method.
func (m *MockRecipes) UpdateRecipe(userId, id int, input models.UpdateRecipeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecipe", userId, id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecipe indicates an expected call of UpdateRecipe.
func (mr *MockRecipesMockRecorder) UpdateRecipe(userId, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecipe", reflect.TypeOf((*MockRecipes)(nil).UpdateRecipe), userId, id, input)
}

// MockSchedule is a mock of Schedule interface.
type MockSchedule struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleMockRecorder
}

// MockScheduleMockRecorder is the mock recorder for MockSchedule.
type MockScheduleMockRecorder struct {
	mock *MockSchedule
}

// NewMockSchedule creates a new mock instance.
func NewMockSchedule(ctrl *gomock.Controller) *MockSchedule {
	mock := &MockSchedule{ctrl: ctrl}
	mock.recorder = &MockScheduleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedule) EXPECT() *MockScheduleMockRecorder {
	return m.recorder
}

// DeleteSchedule mocks base method.
func (m *MockSchedule) DeleteSchedule(userId int, date string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSchedule", userId, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSchedule indicates an expected call of DeleteSchedule.
func (mr *MockScheduleMockRecorder) DeleteSchedule(userId, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSchedule", reflect.TypeOf((*MockSchedule)(nil).DeleteSchedule), userId, date)
}

// FillSchedule mocks base method.
func (m *MockSchedule) FillSchedule(userId int, schedule models.Schedule) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillSchedule", userId, schedule)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FillSchedule indicates an expected call of FillSchedule.
func (mr *MockScheduleMockRecorder) FillSchedule(userId, schedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillSchedule", reflect.TypeOf((*MockSchedule)(nil).FillSchedule), userId, schedule)
}

// GetAllSchedule mocks base method.
func (m *MockSchedule) GetAllSchedule(userId int) ([]models.ScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSchedule", userId)
	ret0, _ := ret[0].([]models.ScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSchedule indicates an expected call of GetAllSchedule.
func (mr *MockScheduleMockRecorder) GetAllSchedule(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSchedule", reflect.TypeOf((*MockSchedule)(nil).GetAllSchedule), userId)
}

// GetScheduleByDate mocks base method.
func (m *MockSchedule) GetScheduleByDate(userId int, date string) (models.ScheduleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduleByDate", userId, date)
	ret0, _ := ret[0].(models.ScheduleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduleByDate indicates an expected call of GetScheduleByDate.
func (mr *MockScheduleMockRecorder) GetScheduleByDate(userId, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduleByDate", reflect.TypeOf((*MockSchedule)(nil).GetScheduleByDate), userId, date)
}

// UpdateSchedule mocks base method.
func (m *MockSchedule) UpdateSchedule(userId int, date string, input models.UpdateScheduleInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedule", userId, date, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchedule indicates an expected call of UpdateSchedule.
func (mr *MockScheduleMockRecorder) UpdateSchedule(userId, date, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedule", reflect.TypeOf((*MockSchedule)(nil).UpdateSchedule), userId, date, input)
}