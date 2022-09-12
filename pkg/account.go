package beatstore

import (
	"errors"
	"time"
)

type Account struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"       binding:"required"`
	Username  string    `json:"username"   binding:"required"`
	Email     string    `json:"email"      binding:"required"`
	PhotoPath string    `json:"photo_path"`
	Confirmed bool      `json:"confirmed"`
	Password  string    `json:"password"   binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountUpdateInput struct {
	Name      *string `json:"name"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	Confirmed *string `json:"confirmed"`
	PhotoPath *string `json:"photo_path"`
}

type AccountPassword struct {
	Password string `json:"password"`
}

func (a *AccountUpdateInput) Validate() error {
	if a.Name == nil && a.Username == nil &&
		a.PhotoPath == nil && a.Email == nil && a.Confirmed == nil {
		return errors.New("update strcture has no values")
	}

	return nil
}
