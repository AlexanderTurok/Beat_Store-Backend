package repository

import (
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user beatstore.User) (int, error)
	GetUser(email, password string) (int, error)
}

type Beat interface {
	Create(userId int, beat beatstore.Beat) (int, error)
	GetAll() ([]beatstore.Beat, error)
	GetUsersBeats(userId int) ([]beatstore.Beat, error)
	Update(userId, beatId int, input beatstore.BeatUpdateInput) error
	Delete(userId, beatId int) error
}

type User interface {
	Get(userId int) (beatstore.User, error)
	GetAll() ([]beatstore.User, error)
	Update(userId int, input beatstore.UserUpdateInput) error
	Delete(userId int) error
}

type Repository struct {
	Authorization
	Beat
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthService(db),
		Beat:          NewBeatRepository(db),
		User:          NewUserRepository(db),
	}
}
