package beatstore

import (
	"errors"
	"time"
)

type Account struct {
	Id        int       `json:"-" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	PhotoPath string    `json:"photo_path" db:"photo_path"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type AccountUpdateInput struct {
	Name      *string `json:"name"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	PhotoPath *string `json:"photo_path"`
}

type AccountPassword struct {
	Password string `json:"password" db:"password_hash"`
}

func (a *AccountUpdateInput) Validate() error {
	if a.Name == nil && a.Username == nil &&
		a.PhotoPath == nil && a.Email == nil {
		return errors.New("update strcture has no values")
	}

	return nil
}
