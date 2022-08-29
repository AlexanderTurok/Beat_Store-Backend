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

func (r *BeatRepository) GetAll() ([]beatstore.Beat, error) {
	var beats []beatstore.Beat

	query := fmt.Sprintf("SELECT * FROM %s", beatTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var beat beatstore.Beat
		if err := rows.Scan(&beat.Id, &beat.Bpm, &beat.Key, &beat.Path, &beat.Tags, &beat.Price); err != nil {
			return beats, err
		}
		beats = append(beats, beat)
	}

	return beats, rows.Err()
}
