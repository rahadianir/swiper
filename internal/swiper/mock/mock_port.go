// Code generated by MockGen. DO NOT EDIT.
// Source: port.go
//
// Generated by this command:
//
//	mockgen -source port.go -destination mock/mock_port.go
//

// Package mock_swiper is a generated GoMock package.
package mock_swiper

import (
	context "context"
	reflect "reflect"
	time "time"

	models "github.com/rahadianir/swiper/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockSwiperLogicInterface is a mock of SwiperLogicInterface interface.
type MockSwiperLogicInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSwiperLogicInterfaceMockRecorder
	isgomock struct{}
}

// MockSwiperLogicInterfaceMockRecorder is the mock recorder for MockSwiperLogicInterface.
type MockSwiperLogicInterfaceMockRecorder struct {
	mock *MockSwiperLogicInterface
}

// NewMockSwiperLogicInterface creates a new mock instance.
func NewMockSwiperLogicInterface(ctrl *gomock.Controller) *MockSwiperLogicInterface {
	mock := &MockSwiperLogicInterface{ctrl: ctrl}
	mock.recorder = &MockSwiperLogicInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwiperLogicInterface) EXPECT() *MockSwiperLogicInterfaceMockRecorder {
	return m.recorder
}

// GetTargetProfile mocks base method.
func (m *MockSwiperLogicInterface) GetTargetProfile(ctx context.Context, userID int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetProfile", ctx, userID)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetProfile indicates an expected call of GetTargetProfile.
func (mr *MockSwiperLogicInterfaceMockRecorder) GetTargetProfile(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetProfile", reflect.TypeOf((*MockSwiperLogicInterface)(nil).GetTargetProfile), ctx, userID)
}

// SwipeLeft mocks base method.
func (m *MockSwiperLogicInterface) SwipeLeft(ctx context.Context, userID, targetId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwipeLeft", ctx, userID, targetId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwipeLeft indicates an expected call of SwipeLeft.
func (mr *MockSwiperLogicInterfaceMockRecorder) SwipeLeft(ctx, userID, targetId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwipeLeft", reflect.TypeOf((*MockSwiperLogicInterface)(nil).SwipeLeft), ctx, userID, targetId)
}

// SwipeRight mocks base method.
func (m *MockSwiperLogicInterface) SwipeRight(ctx context.Context, userID, targetId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwipeRight", ctx, userID, targetId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwipeRight indicates an expected call of SwipeRight.
func (mr *MockSwiperLogicInterfaceMockRecorder) SwipeRight(ctx, userID, targetId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwipeRight", reflect.TypeOf((*MockSwiperLogicInterface)(nil).SwipeRight), ctx, userID, targetId)
}

// MockSwiperRepositoryInterface is a mock of SwiperRepositoryInterface interface.
type MockSwiperRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSwiperRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockSwiperRepositoryInterfaceMockRecorder is the mock recorder for MockSwiperRepositoryInterface.
type MockSwiperRepositoryInterfaceMockRecorder struct {
	mock *MockSwiperRepositoryInterface
}

// NewMockSwiperRepositoryInterface creates a new mock instance.
func NewMockSwiperRepositoryInterface(ctrl *gomock.Controller) *MockSwiperRepositoryInterface {
	mock := &MockSwiperRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockSwiperRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwiperRepositoryInterface) EXPECT() *MockSwiperRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetUserLikedUserIDs mocks base method.
func (m *MockSwiperRepositoryInterface) GetUserLikedUserIDs(ctx context.Context, userID int, params models.LikedUserParams) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLikedUserIDs", ctx, userID, params)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLikedUserIDs indicates an expected call of GetUserLikedUserIDs.
func (mr *MockSwiperRepositoryInterfaceMockRecorder) GetUserLikedUserIDs(ctx, userID, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLikedUserIDs", reflect.TypeOf((*MockSwiperRepositoryInterface)(nil).GetUserLikedUserIDs), ctx, userID, params)
}

// StoreUserLike mocks base method.
func (m *MockSwiperRepositoryInterface) StoreUserLike(ctx context.Context, userID, targetID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUserLike", ctx, userID, targetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUserLike indicates an expected call of StoreUserLike.
func (mr *MockSwiperRepositoryInterfaceMockRecorder) StoreUserLike(ctx, userID, targetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUserLike", reflect.TypeOf((*MockSwiperRepositoryInterface)(nil).StoreUserLike), ctx, userID, targetID)
}

// UpdateMatchStatus mocks base method.
func (m *MockSwiperRepositoryInterface) UpdateMatchStatus(ctx context.Context, userID, targetID int, status bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMatchStatus", ctx, userID, targetID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMatchStatus indicates an expected call of UpdateMatchStatus.
func (mr *MockSwiperRepositoryInterfaceMockRecorder) UpdateMatchStatus(ctx, userID, targetID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMatchStatus", reflect.TypeOf((*MockSwiperRepositoryInterface)(nil).UpdateMatchStatus), ctx, userID, targetID, status)
}

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetRandomUser mocks base method.
func (m *MockUserRepositoryInterface) GetRandomUser(ctx context.Context, excludeList []int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomUser", ctx, excludeList)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomUser indicates an expected call of GetRandomUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetRandomUser(ctx, excludeList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetRandomUser), ctx, excludeList)
}

// GetUserByUserID mocks base method.
func (m *MockUserRepositoryInterface) GetUserByUserID(ctx context.Context, id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUserID", ctx, id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUserID indicates an expected call of GetUserByUserID.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUserByUserID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUserID", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUserByUserID), ctx, id)
}

// MockCacheInterface is a mock of CacheInterface interface.
type MockCacheInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCacheInterfaceMockRecorder
	isgomock struct{}
}

// MockCacheInterfaceMockRecorder is the mock recorder for MockCacheInterface.
type MockCacheInterfaceMockRecorder struct {
	mock *MockCacheInterface
}

// NewMockCacheInterface creates a new mock instance.
func NewMockCacheInterface(ctrl *gomock.Controller) *MockCacheInterface {
	mock := &MockCacheInterface{ctrl: ctrl}
	mock.recorder = &MockCacheInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheInterface) EXPECT() *MockCacheInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCacheInterface) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheInterfaceMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheInterface)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockCacheInterface) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheInterfaceMockRecorder) Set(ctx, key, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheInterface)(nil).Set), ctx, key, value, ttl)
}

// Update mocks base method.
func (m *MockCacheInterface) Update(ctx context.Context, key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCacheInterfaceMockRecorder) Update(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCacheInterface)(nil).Update), ctx, key, value)
}