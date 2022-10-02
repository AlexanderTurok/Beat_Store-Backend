package model

type Product struct {
	Id        int64  `json:"id"         db:"id"`
	ArtistId  int64  `json:"artist_id"  db:"artist_id"`
	StripeId  string `json:"stripe_id"  db:"stripe_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type ProductBeat struct {
	Id        int64  `json:"id"         db:"id"`
	ArtistId  int64  `json:"artist_id"  db:"artist_id"`
	StripeId  string `json:"stripe_id"  db:"stripe_id"`
	Beat      Beat   `json:"beat"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
