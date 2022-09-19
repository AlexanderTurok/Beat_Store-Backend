package repository

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Account interface {
	Create(account model.Account) (int, error)
	Confirm(username string) error
	GetId(email, password string) (int, error)
	Get(accountId int) (model.Account, error)
	Update(accountId int, input model.AccountUpdateInput) error
	GetPasswordHash(accountId int) (string, error)
	Delete(accountId int) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (model.Account, error)
	GetAll() ([]model.Account, error)
	GetPasswordHash(accountId int) (string, error)
	Delete(accountId int) error
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

type Repository struct {
	Account
	Artist
	Beat
	Playlist
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account:  NewAccountRepository(db),
		Artist:   NewArtistRepository(db),
		Beat:     NewBeatRepository(db),
		Playlist: NewPlaylistRepository(db),
		Payment:  NewPaymentRepository(db),
	}
}
