package model

import "time"

type Artist struct {
	Id        int64     `json:"id"         db:"id"`
	StripeId  string    `json:"stripe_id"  db:"stripe_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
