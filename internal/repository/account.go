package repository

import (
	"fmt"
	"strings"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
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

func (r *AccountRepository) Confirm(username string) error {
	query := fmt.Sprintf("UPDATE %s SET confirmed=true WHERE username=$1", accountTable)
	_, err := r.db.Exec(query, username)

	return err
}

func (r *AccountRepository) Get(accountId int) (model.Account, error) {
	var account model.Account

	query := fmt.Sprintf("SELECT name, username, email, photo_path, confirmed, created_at FROM %s WHERE id=$1", accountTable)
	err := r.db.Get(&account, query, accountId)

	return account, err
}

func (r *AccountRepository) Update(accountId int, input model.AccountUpdateInput) error {
	query, args := createAccountUpdateQuery(accountId, input)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *AccountRepository) GetPasswordHash(accountId int) (string, error) {
	var passwordHash string

	query := fmt.Sprintf("SELECT password_hash FROM %s WHERE id=$1", accountTable)
	err := r.db.Get(&passwordHash, query, accountId)

	return passwordHash, err
}

func (r *AccountRepository) Delete(accountId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", accountTable)
	_, err := r.db.Exec(query, accountId)

	return err
}

func createAccountUpdateQuery(accountId int, input model.AccountUpdateInput) (string, []interface{}) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.PhotoPath != nil {
		setValues = append(setValues, fmt.Sprintf("photo_path=$%d", argId))
		args = append(args, *input.PhotoPath)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", accountTable, setQuery, argId)
	args = append(args, accountId)

	return query, args
}
