package repository

import (
	"fmt"
	"golang-auth/internal/models"
)

type UserRepoIface interface {
	CreateUser(userModel *models.User) *models.User
}

type UserRepo struct {
	db string
}

func NewUserRepo(db string) UserRepoIface {
	return &UserRepo{db: db}
}

func (repo *UserRepo) CreateUser(userModel *models.User) *models.User {
	fmt.Println("Creating a new user from repo ->", repo.db)
	return userModel
}
