package repository

import (
	"fmt"
	"time"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
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

func (r *ArtistRepository) Get(accountId int) (model.Artist, error) {
	var artist model.Artist

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, artistTable)
	err := r.db.Get(&artist, query, accountId)

	return artist, err
}

func (r *ArtistRepository) GetAll() ([]model.Artist, error) {
	var artists []model.Artist

	query := fmt.Sprintf(`SELECT * FROM %s`, artistTable)
	err := r.db.Select(&artists, query)

	return artists, err
}

func (r *ArtistRepository) GetPasswordHash(accountId int) (string, error) {
	var passwordHash string
	query := fmt.Sprintf("SELECT password_hash FROM %s WHERE id=$1", accountTable)
	err := r.db.Get(&passwordHash, query, accountId)

	return passwordHash, err
}

func (r *ArtistRepository) Delete(accountId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", artistTable)
	_, err := r.db.Exec(query, accountId)

	return err
}
