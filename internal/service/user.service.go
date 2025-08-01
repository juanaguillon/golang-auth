package service

import (
	"golang-auth/internal/models"
	"golang-auth/internal/repository"
)

type UserServiceIface interface {
	CreateUser(models.UserCreateRequest) *models.User
}

type UserService struct {
	userRepo repository.UserRepoIface
}

func NewUserService(userRepo repository.UserRepoIface) UserServiceIface {
	return &UserService{
		userRepo: userRepo,
	}
}

func (userService *UserService) CreateUser(userRequest models.UserCreateRequest) *models.User {
	newUser := &models.User{
		Name:     userRequest.Name,
		Lastname: userRequest.Lastname,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Age:      *userRequest.Age,
	}

	return userService.userRepo.CreateUser(newUser)

}
