package beatstore

import "time"

type Artist struct {
	Id        int       `json:"-"          db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
