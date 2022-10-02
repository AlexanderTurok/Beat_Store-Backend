package repository

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Create(account model.Account) (int, error)
	GetId(email, password string) (int, error)
}

type Account interface {
	Confirm(username string) error
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

type Payment interface {
	CreatePaymentAccount(accountId int, stripeId string) error
}

type Product interface {
}

type Beat interface {
	Create(artistId int, input model.Beat) (int, error)
	Get(beatId int) (model.Beat, error)
	GetAll() ([]model.Beat, error)
	GetArtistsBeat(beatId, artistId int) (model.Beat, error)
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

type Repositories struct {
	Auth     Auth
	Account  Account
	Artist   Artist
	Payment  Payment
	Product  Product
	Beat     Beat
	Playlist Playlist
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Auth:     NewAuthRepository(db),
		Account:  NewAccountRepository(db),
		Artist:   NewArtistRepository(db),
		Payment:  NewPaymentRepository(db),
		Product:  NewProductRepository(db),
		Beat:     NewBeatRepository(db),
		Playlist: NewPlaylistRepository(db),
	}
}
