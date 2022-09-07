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
	Get(accountId int) (beatstore.Account, error)
	Update(accountId int, input beatstore.AccountUpdateInput) error
	Delete(accountId int, password beatstore.AccountPassword) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (beatstore.Account, error)
	Delete(accountId int, password beatstore.AccountPassword) error
}

type Beat interface {
	Create(artistId int, input beatstore.Beat) (int, error)
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
		Account:       NewAccountService(repos.Account),
		Artist:        NewArtistService(repos.Artist),
		Beat:          NewBeatService(repos.Beat),
	}
}
