package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserNameByID(id int64) (string, error) {
	ret := m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func TestGetUserName(t *testing.T) {
	repo := new(MockRepository)
	userService := NewUserService(repo)

	repo.On("GetUserNameByID", int64(1)).Return("Kemal", nil)
	name, err := userService.GetUsername(1)

	assert.Nil(t, err)
	assert.Equal(t, "Kemal", name)

	repo.On("GetUserNameByID", int64(2)).Return("", errors.New("fail"))
	name, err = userService.GetUsername(2)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Empty(t, name)

	repo.AssertExpectations(t)
}
