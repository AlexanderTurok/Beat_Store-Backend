package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type Authorization interface {
	CreateAccount(account model.Account) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Account interface {
	Get(accountId int) (model.Account, error)
	Update(accountId int, input model.AccountUpdateInput) error
	Delete(accountId int, password string) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (model.Account, error)
	GetAll() ([]model.Account, error)
	Delete(accountId int, password string) error
}

type Beat interface {
	Create(artistId int, input model.Beat) (int, error)
	Get(beatId int) (model.Beat, error)
	GetAll() ([]model.Beat, error)
	GetAllArtistsBeats(artistId int) ([]model.Beat, error)
	Update(beatId int, input model.BeatUpdateInput) error
	Delete(beatId int) error
}

type Playlist interface {
	Create(accountId int, input model.Playlist) (int, error)
	GetAllAccountsPlaylists(accountId int) ([]model.Playlist, error)
	Update(playlistId int, input model.PlaylistUpdateInput) error
	Delete(playlistId int) error
	AddBeat(playlistId, beatId int) error
	GetAllBeats(playlistId int) ([]model.Beat, error)
	DeleteBeat(playlistId, beatId int) error
}

type Payment interface {
}

type Service struct {
	Authorization
	Account
	Artist
	Beat
	Playlist
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Account:       NewAccountService(repos.Account),
		Artist:        NewArtistService(repos.Artist),
		Beat:          NewBeatService(repos.Beat),
		Playlist:      NewPlaylistService(repos.Playlist),
		Payment:       NewPaymentService(repos.Payment),
	}
}
