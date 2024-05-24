package userservice

import "github.com/kardecDev/api_go_back/internal/repository/userrepository"

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser() error
}
