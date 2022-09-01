package repository

import (
	"fmt"
	"strings"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Get(userId int) (beatstore.User, error) {
	var user beatstore.User
	query := fmt.Sprintf("SELECT name, photo, username, email FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserRepository) GetAll() ([]beatstore.User, error) {
	var users []beatstore.User

	query := fmt.Sprintf("SELECT name, username, photo, email FROM %s", userTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user beatstore.User
		if err := rows.Scan(&user.Name, &user.Username, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *UserRepository) Update(userId int, input beatstore.UserUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Photo != nil {
		setValues = append(setValues, fmt.Sprintf("photo=$%d", argId))
		args = append(args, *input.Photo)
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

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		userTable, setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UserRepository) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}
