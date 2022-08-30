package repository

import (
	"database/sql"
	"fmt"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(user beatstore.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)

	return id, err
}
