// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	models "todo-app/models"

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

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(input models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), input)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseUserIdFromToken mocks base method.
func (m *MockAuthorization) ParseUserIdFromToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUserIdFromToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseUserIdFromToken indicates an expected call of ParseUserIdFromToken.
func (mr *MockAuthorizationMockRecorder) ParseUserIdFromToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUserIdFromToken", reflect.TypeOf((*MockAuthorization)(nil).ParseUserIdFromToken), token)
}

// MockTodoList is a mock of TodoList interface.
type MockTodoList struct {
	ctrl     *gomock.Controller
	recorder *MockTodoListMockRecorder
}

// MockTodoListMockRecorder is the mock recorder for MockTodoList.
type MockTodoListMockRecorder struct {
	mock *MockTodoList
}

// NewMockTodoList creates a new mock instance.
func NewMockTodoList(ctrl *gomock.Controller) *MockTodoList {
	mock := &MockTodoList{ctrl: ctrl}
	mock.recorder = &MockTodoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoList) EXPECT() *MockTodoListMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoList) Create(userId int, title string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, title)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoListMockRecorder) Create(userId, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoList)(nil).Create), userId, title)
}

// Delete mocks base method.
func (m *MockTodoList) Delete(userId, listId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, listId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoListMockRecorder) Delete(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoList)(nil).Delete), userId, listId)
}

// GetAll mocks base method.
func (m *MockTodoList) GetAll(userId int) ([]models.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]models.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoListMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoList)(nil).GetAll), userId)
}

// GetById mocks base method.
func (m *MockTodoList) GetById(userId, listId int) (models.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, listId)
	ret0, _ := ret[0].(models.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTodoListMockRecorder) GetById(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTodoList)(nil).GetById), userId, listId)
}

// Update mocks base method.
func (m *MockTodoList) Update(userId int, input models.UpdateTodoListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoListMockRecorder) Update(userId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoList)(nil).Update), userId, input)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoItem) Create(userId int, input models.TodoItemInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoItemMockRecorder) Create(userId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoItem)(nil).Create), userId, input)
}

// Delete mocks base method.
func (m *MockTodoItem) Delete(userId, itemId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoItemMockRecorder) Delete(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoItem)(nil).Delete), userId, itemId)
}

// GetAllByListId mocks base method.
func (m *MockTodoItem) GetAllByListId(userId, listId int) ([]models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByListId", userId, listId)
	ret0, _ := ret[0].([]models.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByListId indicates an expected call of GetAllByListId.
func (mr *MockTodoItemMockRecorder) GetAllByListId(userId, listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByListId", reflect.TypeOf((*MockTodoItem)(nil).GetAllByListId), userId, listId)
}

// GetById mocks base method.
func (m *MockTodoItem) GetById(userId, itemId int) (models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, itemId)
	ret0, _ := ret[0].(models.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTodoItemMockRecorder) GetById(userId, itemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTodoItem)(nil).GetById), userId, itemId)
}

// Update mocks base method.
func (m *MockTodoItem) Update(userId, itemId int, input models.UpdateTodoItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, itemId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoItemMockRecorder) Update(userId, itemId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoItem)(nil).Update), userId, itemId, input)
}
