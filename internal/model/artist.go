package model

import "time"

type Artist struct {
	Id        int64     `json:"id"         db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
