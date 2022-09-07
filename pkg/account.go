package beatstore

import (
	"errors"
	"time"
)

type Account struct {
	Id        int       `json:"-"`
	Name      string    `json:"name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	PhotoPath string    `json:"photo_path"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountUpdateInput struct {
	Name      *string `json:"name"`
	Username  *string `json:"username"`
	PhotoPath *string `json:"photo" `
	Email     *string `json:"email"`
}

func (a *AccountUpdateInput) Validate() error {
	if a.Name == nil && a.Username == nil &&
		a.PhotoPath == nil && a.Email == nil {
		return errors.New("update strcture has no values")
	}

	return nil
}
