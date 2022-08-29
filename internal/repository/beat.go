package repository

import (
	"database/sql"
	"fmt"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type BeatRepository struct {
	db *sql.DB
}

func NewBeatRepository(db *sql.DB) *BeatRepository {
	return &BeatRepository{
		db: db,
	}
}

func (r *BeatRepository) GetById(id int) (beatstore.Beat, error) {
	var beat beatstore.Beat
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", beatTable)
	err := r.db.QueryRow(query, id).Scan(&beat)

	return beat, err
}
