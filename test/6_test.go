package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type UserDetails struct {
	ID   int
	Name string
}

type UserAPI interface {
	GetUserDetails(userId int) (UserDetails, error)
}

type MockUserAPI struct {
	mock.Mock
}

func (m *MockUserAPI) GetUserDetails(userId int) (UserDetails, error) {
	args := m.Called(userId)
	if res := args.Get(0); res != nil {
		return res.(UserDetails), nil
	}

	return UserDetails{}, fmt.Errorf("not implemented Id: %d", userId)
}

func GetUserDetails(userId int, api UserAPI) (UserDetails, error) {
	return api.GetUserDetails(userId)
}

func TestGetUserDetails(t *testing.T) {
	mockAPI := &MockUserAPI{}

	mockAPI.On("GetUserDetails", 1).Return(UserDetails{
		ID:   1,
		Name: "John Doe",
	}, nil)

	userDetails, err := GetUserDetails(1, mockAPI)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if userDetails.ID != 1 || userDetails.Name != "John Doe" {
		t.Errorf("Incorrect user details returned")
	}
}
