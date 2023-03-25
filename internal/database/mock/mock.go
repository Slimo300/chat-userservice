// Code generated by mockery v2.14.1. DO NOT EDIT.

package mock

import (
	models "github.com/Slimo300/chat-userservice/internal/models"
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockUsersDB is an autogenerated mock type for the DBLayer type
type MockUsersDB struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: userID, oldPassword, newPassword
func (_m *MockUsersDB) ChangePassword(userID uuid.UUID, oldPassword string, newPassword string) error {
	ret := _m.Called(userID, oldPassword, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string, string) error); ok {
		r0 = rf(userID, oldPassword, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProfilePicture provides a mock function with given fields: userID
func (_m *MockUsersDB) DeleteProfilePicture(userID uuid.UUID) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(uuid.UUID) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfilePictureURL provides a mock function with given fields: userID
func (_m *MockUsersDB) GetProfilePictureURL(userID uuid.UUID) (string, bool, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(uuid.UUID) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uuid.UUID) bool); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uuid.UUID) error); ok {
		r2 = rf(userID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserById provides a mock function with given fields: uid
func (_m *MockUsersDB) GetUserById(uid uuid.UUID) (models.User, error) {
	ret := _m.Called(uid)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) models.User); ok {
		r0 = rf(uid)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewResetPasswordCode provides a mock function with given fields: email
func (_m *MockUsersDB) NewResetPasswordCode(email string) (*models.User, *models.ResetCode, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 *models.ResetCode
	if rf, ok := ret.Get(1).(func(string) *models.ResetCode); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.ResetCode)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterUser provides a mock function with given fields: _a0
func (_m *MockUsersDB) RegisterUser(_a0 models.User) (*models.User, *models.VerificationCode, error) {
	ret := _m.Called(_a0)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(models.User) *models.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 *models.VerificationCode
	if rf, ok := ret.Get(1).(func(models.User) *models.VerificationCode); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.VerificationCode)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(models.User) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResetPassword provides a mock function with given fields: code, newPassword
func (_m *MockUsersDB) ResetPassword(code string, newPassword string) error {
	ret := _m.Called(code, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(code, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignIn provides a mock function with given fields: email, password
func (_m *MockUsersDB) SignIn(email string, password string) (models.User, error) {
	ret := _m.Called(email, password)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string, string) models.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyCode provides a mock function with given fields: code
func (_m *MockUsersDB) VerifyCode(code string) (*models.User, error) {
	ret := _m.Called(code)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockUsersDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUsersDB creates a new instance of MockUsersDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUsersDB(t mockConstructorTestingTNewMockUsersDB) *MockUsersDB {
	mock := &MockUsersDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}