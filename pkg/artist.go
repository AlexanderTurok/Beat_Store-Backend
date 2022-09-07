package beatstore

import "time"

type Artist struct {
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
