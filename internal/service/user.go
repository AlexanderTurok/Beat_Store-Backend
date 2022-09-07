package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type UserService struct {
	repos repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{
		repos: repos,
	}
}
