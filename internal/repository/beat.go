package repository

import (
	"fmt"
	"time"

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

func (r *BeatRepository) Create(artistId int, input beatstore.Beat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var beatId int
	insertBeatQuery := fmt.Sprintf(`INSERT INTO %s (
			artist_id,
			name,
			bpm,
			key,
			photo_path,
			mp3_path,
			wav_path,
			genre,
			mood,
			created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		beatTable,
	)

	row := tx.QueryRow(insertBeatQuery,
		artistId,
		input.Name,
		input.Bpm,
		input.Key,
		input.PhotoPath,
		input.MP3Path,
		input.WavPath,
		input.Genre,
		input.Mood,
		time.Now(),
	)
	err = row.Scan(&beatId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	insertPriceQuery := fmt.Sprintf(`INSERT INTO %s (
			beat_id, 
			standart_price, 
			premium_price, 
			unlimited_price) 
		VALUES ($1, $2, $3, $4)`,
		priceTable,
	)
	_, err = tx.Exec(insertPriceQuery,
		beatId,
		input.Price.StandartPrice,
		input.Price.PremiumPrice,
		input.Price.UnlimitedPrice,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	insertTagQuery := fmt.Sprintf(`INSERT INTO %s (beat_id, name) VALUES ($1, $2)`, tagTable)
	for _, tag := range input.Tags {
		if _, err := tx.Exec(insertTagQuery, beatId, tag.Name); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return beatId, tx.Commit()
}
