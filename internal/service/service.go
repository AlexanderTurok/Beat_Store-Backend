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
	Delete(accountId int, password string) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (beatstore.Account, error)
	Delete(accountId int, password string) error
}

type Beat interface {
	Create(artistId int, input beatstore.Beat) (int, error)
	// Get(beatId int) (beatstore.Beat, error)
	GetAll() ([]beatstore.Beat, error)
	GetAllArtistsBeats(artistId int) ([]beatstore.Beat, error)
	Update(beatId int, input beatstore.BeatUpdateInput) error
	Delete(beatId int) error
}

type Playlist interface {
	Create(accountId int, input beatstore.Playlist) (int, error)
	GetAllAccountsPlaylists(accountId int) ([]beatstore.Playlist, error)
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
		Playlist:      NewPlaylistService(repos.Playlist),
	}
}
