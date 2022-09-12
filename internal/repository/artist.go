package repository

import (
	"fmt"
	"time"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
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

func (r *ArtistRepository) Get(accountId int) (beatstore.Account, error) {
	var artist beatstore.Account

	query := fmt.Sprintf(`
		SELECT 
			artist.created_at,
			account.name,
			account.username, 
			account.email, 
			account.photo_path
		FROM %s 
		JOIN %s ON account.id = artist.id
		WHERE account.id=$1`, artistTable, accountTable,
	)
	err := r.db.Get(&artist, query, accountId)

	return artist, err
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
