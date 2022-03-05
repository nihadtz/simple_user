package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/nihadtz/simple_user/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r UserRepositoryMock) FindUserByID(id int64) (*models.User, error) {
	user := models.User{
		ID:          1,
		Name:        "Nihad",
		YearOfBirth: 1984,
		Updated:     1646472110,
	}
	return &user, nil
}

func (r UserRepositoryMock) CreateUser(user models.User) (*int64, error) {
	user = models.User{
		ID:          1,
		Name:        "Nihad",
		YearOfBirth: 1984,
		Updated:     1646472110,
	}
	return &user.ID, nil
}

func (r UserRepositoryMock) UpdateUser(user models.User) (*models.User, error) {
	updatedUser := models.User{
		ID:          user.ID,
		Name:        user.Name,
		YearOfBirth: user.YearOfBirth,
		Updated:     time.Now().Unix(),
	}
	return &updatedUser, nil
}

type TestCases struct {
	User              models.User
	ErrExpectedCreate bool
	ErrExpectedUpdate bool
}

var testUsers = []TestCases{
	{
		User: models.User{
			ID:          1,
			Name:        "Nihad",
			YearOfBirth: 1984,
			Updated:     1646472110,
		},
		ErrExpectedCreate: false,
		ErrExpectedUpdate: false,
	},
	{
		User: models.User{
			ID:          2,
			Name:        "",
			YearOfBirth: 1984,
			Updated:     1646472110,
		},
		ErrExpectedCreate: true,
		ErrExpectedUpdate: true,
	},
	{
		User: models.User{
			ID:          3,
			Name:        "Nihad",
			YearOfBirth: 0,
			Updated:     1646472110,
		},
		ErrExpectedCreate: true,
		ErrExpectedUpdate: true,
	},
	{
		User: models.User{
			ID:          0,
			Name:        "Nihad",
			YearOfBirth: 1984,
			Updated:     1646472110,
		},
		ErrExpectedCreate: false,
		ErrExpectedUpdate: true,
	},
}

func TestService_RegisterUser(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("CreateUser").Return(1, nil)

	userService := UserService{repository}

	for _, test := range testUsers {
		_, err := userService.RegisterUser(test.User)

		if test.ErrExpectedCreate {
			assert.Error(t, err, fmt.Sprintf("Expected error %v", test.ErrExpectedCreate))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Expected error %v", test.ErrExpectedCreate))
		}
	}
}

func TestService_FindUserByID(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("FindUserByID").Return(models.User{}, nil)

	userService := UserService{repository}
	user, _ := userService.RetrieveUserByID(1)

	assert.Equal(t, user.ID, int64(1), "User id matches")
}

func TestService_UpdateUser(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("CreateUser").Return(models.User{}, nil)

	userService := UserService{repository}

	for _, test := range testUsers {
		user, err := userService.UpdateUser(test.User)

		if test.ErrExpectedUpdate {
			assert.Error(t, err, fmt.Sprintf("Expected error %v", test.ErrExpectedUpdate))
		} else {
			assert.NoError(t, err, fmt.Sprintf("Expected error %v", test.ErrExpectedUpdate))

			assert.Greater(t, user.Updated, test.User.Updated, "Update Date must be newer")
		}
	}
}
