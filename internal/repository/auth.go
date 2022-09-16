package repository

import (
	"fmt"
	"time"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
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

func (r *AuthRepository) CreateAccount(account model.Account) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, photo_path, password_hash, created_at) values ($1, $2, $3, $4, $5, $6) RETURNING id", accountTable)
	row := r.db.QueryRow(query, account.Name, account.Username, account.Email, account.PhotoPath, account.Password, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetAccountId(email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", accountTable)
	err := r.db.Get(&id, query, email, password)

	return id, err
}
