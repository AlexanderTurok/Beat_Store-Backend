package repository

import (
	"fmt"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type BeatRepository struct {
	db *sqlx.DB
}

func NewBeatRepository(db *sqlx.DB) *BeatRepository {
	return &BeatRepository{
		db: db,
	}
}

func (r *BeatRepository) Create(userId int, beat beatstore.Beat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var beatId int
	beatQuery := fmt.Sprintf("INSERT INTO %s (bpm, key, path, tag, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", beatTable)
	row := tx.QueryRow(beatQuery, beat.Bpm, beat.Key, beat.Path, beat.Tag, beat.Price)
	if err := row.Scan(&beatId); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersBeatQuery := fmt.Sprintf("INSERT INTO %s (user_id, beat_id) VALUES ($1, $2)", usersBeatTable)
	_, err = tx.Exec(usersBeatQuery, userId, beatId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return beatId, tx.Commit()
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
		if err := rows.Scan(&beat.Id, &beat.Bpm, &beat.Key, &beat.Path, &beat.Tag, &beat.Price); err != nil {
			return beats, err
		}
		beats = append(beats, beat)
	}

	return beats, rows.Err()
}
