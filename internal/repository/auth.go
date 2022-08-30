package repository

import (
	"fmt"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user beatstore.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	err := r.db.Get(&id, query, user.Name, user.Username, user.Password)

	return id, err
}

func (r *AuthRepository) GetUser(email, password string) (beatstore.User, error) {
	var user beatstore.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", userTable)
	err := r.db.QueryRow(query, email, password).Scan(&user)

	return user, err
}
