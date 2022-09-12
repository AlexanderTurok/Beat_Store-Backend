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
			standart_price,
			premium_price,
			unlimited_price,
			created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
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
		input.StandartPrice,
		input.PremiumPrice,
		input.UnlimitedPrice,
		time.Now(),
	)
	if err = row.Scan(&beatId); err != nil {
		tx.Rollback()
		return 0, err
	}

	insertTagQuery := fmt.Sprintf(`INSERT INTO %s (beat_id, tag_name) VALUES ($1, $2)`, tagTable)
	for _, tag := range input.Tags {
		if _, err := tx.Exec(insertTagQuery, beatId, tag.TagName); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return beatId, tx.Commit()
}

func (r *BeatRepository) Get(beatId int) (beatstore.Beat, error) {
	var beat beatstore.Beat

	beatQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", beatTable)
	if err := r.db.Get(&beat, beatQuery, beatId); err != nil {
		return beatstore.Beat{}, err
	}

	tagsQuery := fmt.Sprintf("SELECT * FROM %s WHERE beat_id=$1", tagTable)
	err := r.db.Select(&beat.Tags, tagsQuery, beatId)

	return beat, err
}
