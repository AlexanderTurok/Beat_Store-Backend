package beatstore

import "time"

type Artist struct {
	Id         int       `json:"-" db:"id"`
	ArtistName string    `json:"artist_name" db:"artist_name" binding:"required"`
	CreatedAt  time.Time `json:"craeted_at" db:"created_at" binding:"required"`
}
