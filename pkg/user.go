package beatstore

import (
	"errors"
	"time"
)

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" db:"name"`
	Username  string `json:"username" db:"username" binding:"required"`
	PhotoPath string `json:"photo_path" db:"photo_path"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"password_hash"`
}

type UsersPlaylist struct {
	Id        int       `json:"-" db:"id"`
	UserId    int       `json:"user_id" db:"user_id" binding:"required"`
	BeatId    int       `json:"beat_id" db:"beat_id" binding:"required"`
	Name      string    `json:"name" db:"name" binding:"required"`
	CreatedAt time.Time `json:"craeted_at" db:"created_at" binding:"required"`
}

type UsersBoughtBeat struct {
	Id        int       `json:"-" db:"id"`
	UserId    int       `json:"user_id" db:"user_id" binding:"required"`
	BeatId    int       `json:"beat_id" db:"beat_id" binding:"required"`
	CreatedAt time.Time `json:"craeted_at" db:"created_at" binding:"required"`
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
