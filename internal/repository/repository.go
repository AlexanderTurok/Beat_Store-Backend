package repository

import (
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateAccount(account beatstore.Account) (int, error)
	GetAccountId(email, password string) (int, error)
}

type Account interface {
	Get(accountId int) (beatstore.Account, error)
}

type Artist interface {
}

type Beat interface {
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthService(db),
		Beat:          NewBeatRepository(db),
		Account:       NewAccountRepository(db),
	}
}
