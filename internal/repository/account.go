package repository

import (
	"fmt"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Get(accountId int) (beatstore.Account, error) {
	var account beatstore.Account

	query := fmt.Sprintf("SELECT name, username, email, photo_path, created_at FROM %s WHERE id=$1", accountTable)
	err := r.db.Get(&account, query, accountId)

	return account, err
}
