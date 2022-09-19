package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
)

type Account interface {
	Create(account model.Account) (int, error)
	Confirm(accountId int) error
	GenerateToken(email, password string) (string, error)
	Get(accountId int) (model.Account, error)
	Update(accountId int, input model.AccountUpdateInput) error
	Delete(accountId int, inputPassword string) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (model.Account, error)
	GetAll() ([]model.Account, error)
	Delete(accountId int, inputPassword string) error
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
	Account
	Artist
	Beat
	Playlist
	Payment
}

func NewService(repos *repository.Repository, hasher hash.SHA1Hasher, manager auth.Manager, sender email.Client) *Service {
	return &Service{
		Account:  NewAccountService(repos.Account, hasher, manager, NewEmailService(sender)),
		Artist:   NewArtistService(repos.Artist, hasher),
		Beat:     NewBeatService(repos.Beat),
		Playlist: NewPlaylistService(repos.Playlist),
		Payment:  NewPaymentService(repos.Payment),
	}
}
