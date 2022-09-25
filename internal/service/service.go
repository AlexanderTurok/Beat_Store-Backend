package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
	"github.com/AlexanderTurok/beat-store-backend/pkg/payment"
)

type Account interface {
	Create(account model.Account) (int, error)
	Confirm(username string) error
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
	CreatePaymentIntent(input model.PaymentInfo) (payment.PaymentIntent, error)
}

type Dependencies struct {
	Repositories *repository.Repositories
	Hasher       hash.SHA1Hasher
	Manager      auth.Manager
	Sender       email.Client
	Paymenter    payment.Payment
}

type Services struct {
	Account  Account
	Artist   Artist
	Beat     Beat
	Playlist Playlist
	Payment  Payment
}

func NewServices(d Dependencies) *Services {
	accountService := NewAccountService(d.Repositories.Account, d.Hasher, d.Manager, NewEmailService(d.Sender))
	artistService := NewArtistService(d.Repositories.Artist, d.Hasher)
	beatService := NewBeatService(d.Repositories.Beat)
	playlistService := NewPlaylistService(d.Repositories.Playlist)
	paymentService := NewPaymentService(d.Repositories.Payment, d.Paymenter)

	return &Services{
		Account:  accountService,
		Artist:   artistService,
		Beat:     beatService,
		Playlist: playlistService,
		Payment:  paymentService,
	}
}
