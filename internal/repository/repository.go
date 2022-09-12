package repository

import (
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateAccount(account beatstore.Account) (int, error)
	GetAccountId(email, password string) (int, error)
}

type Account interface {
	Get(accountId int) (beatstore.Account, error)
	Update(accountId int, input beatstore.AccountUpdateInput) error
	GetPasswordHash(accountId int) (beatstore.AccountPassword, error)
	Delete(accountId int) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (beatstore.Account, error)
	GetPasswordHash(accountId int) (beatstore.AccountPassword, error)
	Delete(accountId int) error
}

type Beat interface {
	Create(artistId int, input beatstore.Beat) (int, error)
	Get(beatId int) (beatstore.Beat, error)
}

type Playlist interface {
}

type Repository struct {
	Authorization
	Account
	Artist
	Beat
	Playlist
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthService(db),
		Account:       NewAccountRepository(db),
		Artist:        NewArtistRepository(db),
		Beat:          NewBeatRepository(db),
	}
}
