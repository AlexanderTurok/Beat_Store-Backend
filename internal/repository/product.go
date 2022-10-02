package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(artistId int, stripeId string) (int64, error) {
	var productId int64

	query := fmt.Sprintf(`INSERT INTO %s (artist_id, stripe_id, created_at) VALUES ($1, $2, $3) RETURNING id`, productTable)
	err := r.db.Get(&productId, query, artistId, stripeId, time.Now())

	return productId, err
}
