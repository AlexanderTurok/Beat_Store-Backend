package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ArtistRepository struct {
	db *sqlx.DB
}

func NewArtistRepository(db *sqlx.DB) *ArtistRepository {
	return &ArtistRepository{
		db: db,
	}
}

func (r *ArtistRepository) Create(accountId int) error {
	query := fmt.Sprintf("INSERT INTO %s (id, created_at) VALUES ($1, $2)", artistTable)
	_, err := r.db.Exec(query, accountId, time.Now())

	return err
}
