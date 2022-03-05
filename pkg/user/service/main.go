package service

import (
	"errors"

	"github.com/nihadtz/simple_user/models"
	usr "github.com/nihadtz/simple_user/pkg/user/repository"
)

type UserService struct {
	repo usr.IUserRepository
}

func NewUserService(repo *usr.UserRepository) *UserService {
	service := &UserService{repo}
	return service
}

// IUserService interacts with users
type IUserService interface {
	RetrieveUserByID(id int64) (*models.User, error)
	RegisterUser(user models.User) (*int64, error)
	UpdateUser(user models.User) (*models.User, error)
}

func (svc UserService) RetrieveUserByID(id int64) (*models.User, error) {
	return svc.repo.FindUserByID(id)
}

func (svc UserService) RegisterUser(user models.User) (*int64, error) {
	if len(user.Name) == 0 {
		return nil, errors.New("name is not set")
	}

	if user.YearOfBirth == 0 {
		return nil, errors.New("year of birth is not set")
	}

	return svc.repo.CreateUser(user)
}

func (svc UserService) UpdateUser(user models.User) (*models.User, error) {
	if len(user.Name) == 0 {
		return nil, errors.New("name is not set")
	}

	if user.YearOfBirth == 0 {
		return nil, errors.New("year of birth is not set")
	}

	if user.ID == 0 {
		return nil, errors.New("user ID is not provided")
	}

	return svc.repo.UpdateUser(user)
}
