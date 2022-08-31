package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type UserService struct {
	repos repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{
		repos: repos,
	}
}

func (s *UserService) Get(userId int) (beatstore.User, error) {
	return s.repos.Get(userId)
}

func (s *UserService) GetAll() ([]beatstore.User, error) {
	return s.repos.GetAll()
}

func (s *UserService) Update(userId int, input beatstore.UserUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(userId, input)
}

func (s *UserService) Delete(userId int) error {
	return s.repos.Delete(userId)
}
