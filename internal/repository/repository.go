package repository

import (
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user beatstore.User) (int, error)
	GetUser(email, password string) (beatstore.User, error)
}

type Beat interface {
	Create(userId int, beat beatstore.Beat) (int, error)
	GetById(id int) (beatstore.Beat, error)
	GetAll() ([]beatstore.Beat, error)
}

type Repository struct {
	Authorization
	Beat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthService(db),
		Beat:          NewBeatRepository(db),
	}
}
