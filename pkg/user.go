package beatstore

import "errors"

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" db:"name"`
	Username  string `json:"username" db:"username" binding:"required"`
	PhotoPath string `json:"photo_path" db:"photo_path"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"password_hash"`
}

type UserUpdateInput struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Photo    *string `json:"photo" `
	Email    *string `json:"email"`
}

func (u *UserUpdateInput) Validate() error {
	if u.Name == nil && u.Username == nil &&
		u.Photo == nil && u.Email == nil {
		return errors.New("update strcture has no values")
	}

	return nil
}
