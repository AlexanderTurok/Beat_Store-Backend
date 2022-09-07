package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type Authorization interface {
	CreateAccount(account beatstore.Account) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Account interface {
}

type Artist interface {
}

type Beat interface {
}

type Playlist interface {
}

type Service struct {
	Authorization
	Account
	Artist
	Beat
	Playlist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Beat:          NewBeatService(repos.Beat),
		Account:       NewAccountService(repos.Account),
	}
}
