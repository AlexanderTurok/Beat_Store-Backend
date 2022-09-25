package model

import (
	"errors"
	"time"
)

type Account struct {
	Id        int64     `json:"id"          db:"id"`
	Name      string    `json:"name"       db:"name"          binding:"required"`
	Username  string    `json:"username"   db:"username"      binding:"required"`
	Email     string    `json:"email"      db:"email"         binding:"required"`
	PhotoPath string    `json:"photo_path" db:"photo_path"`
	Confirmed bool      `json:"confirmed"  db:"confirmed"`
	Password  string    `json:"password"   db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type AccountUpdateInput struct {
	Name      *string `json:"name"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	Confirmed *string `json:"confirmed"`
	PhotoPath *string `json:"photo_path"`
}

func (a *AccountUpdateInput) Validate() error {
	if a.Name == nil && a.Username == nil &&
		a.PhotoPath == nil && a.Email == nil && a.Confirmed == nil {
		return errors.New("update strcture has no values")
	}

	return nil
}
