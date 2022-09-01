package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type Authorization interface {
	CreateUser(user beatstore.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Beat interface {
	Create(userId int, beat beatstore.Beat) (int, error)
	GetAll() ([]beatstore.Beat, error)
	GetUsersBeats(userId int) ([]beatstore.Beat, error)
	Update(userId, beatId int, input beatstore.BeatUpdateInput) error
	Delete(userId, beatId int) error
}

type User interface {
	Get(userId int) (beatstore.User, error)
	GetAll() ([]beatstore.User, error)
	Update(userId int, input beatstore.UserUpdateInput) error
	Delete(userId int) error
}

type Service struct {
	Authorization
	Beat
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Beat:          NewBeatService(repos.Beat),
		User:          NewUserService(repos.User),
	}
}
