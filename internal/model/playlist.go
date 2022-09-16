package model

import (
	"errors"
	"time"
)

type Playlist struct {
	Id        int       `json:"id"         db:"id"`
	Name      string    `json:"name"       db:"name"       binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type PlaylistUpdateInput struct {
	Name *string `json:"name" db:"name"`
}

func (i *PlaylistUpdateInput) Validate() error {
	if i.Name == nil {
		return errors.New("update input is empty")
	}

	return nil
}
