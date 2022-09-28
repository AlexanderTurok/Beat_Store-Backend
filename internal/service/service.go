package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
	"github.com/AlexanderTurok/beat-store-backend/pkg/payment"
)

type Auth interface {
	CreateAccount(account model.Account) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Account interface {
	Confirm(username string) error
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

type Product interface {
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

type Dependencies struct {
	Repositories *repository.Repositories
	Hasher       hash.SHA1Hasher
	Manager      auth.Manager
	Sender       email.Client
	Paymenter    payment.Payment
}

type Services struct {
	Auth     Auth
	Account  Account
	Artist   Artist
	Product  Product
	Beat     Beat
	Playlist Playlist
}

func NewServices(d Dependencies) *Services {
	authService := NewAuthService(d.Repositories.Auth, &d.Hasher, &d.Manager, NewEmailService(d.Sender))
	accountService := NewAccountService(d.Repositories.Account, &d.Hasher)
	artistService := NewArtistService(d.Repositories.Artist, &d.Hasher)
	productService := NewProductService(d.Repositories.Product)
	beatService := NewBeatService(d.Repositories.Beat)
	playlistService := NewPlaylistService(d.Repositories.Playlist)

	return &Services{
		Auth:     authService,
		Account:  accountService,
		Artist:   artistService,
		Product:  productService,
		Beat:     beatService,
		Playlist: playlistService,
	}
}
