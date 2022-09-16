package repository

import (
	"fmt"
	"strings"
	"time"

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

func (r *AccountRepository) CreateAccount(account model.Account) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, photo_path, password_hash, created_at) values ($1, $2, $3, $4, $5, $6) RETURNING id", accountTable)
	row := r.db.QueryRow(query, account.Name, account.Username, account.Email, account.PhotoPath, account.Password, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AccountRepository) GetAccountId(email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", accountTable)
	err := r.db.Get(&id, query, email, password)

	return id, err
}

func (r *AccountRepository) Get(accountId int) (model.Account, error) {
	var account model.Account

	query := fmt.Sprintf("SELECT name, username, email, photo_path, created_at FROM %s WHERE id=$1", accountTable)
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
