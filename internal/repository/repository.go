package repository

import (
	"database/sql"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type Authorization interface {
	CreateUser(user beatstore.User) (int, error)
}

type Beat interface {
	GetById(id int) (beatstore.Beat, error)
	GetAll() ([]beatstore.Beat, error)
}

type Repository struct {
	Authorization
	Beat
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Beat: NewBeatRepository(db),
	}
}
