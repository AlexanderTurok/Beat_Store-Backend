package repository

import (
	"fmt"
	"time"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	if err = row.Scan(&beatId); err != nil {
		tx.Rollback()
		return 0, err
	}

	insertPriceQuery := fmt.Sprintf(`
		INSERT INTO %s (beat_id, standart, premium, unlimited) 
		VALUES ($1, $2, $3, $4)`, priceTable)
	if _, err = tx.Exec(insertPriceQuery, beatId, input.Price.Standart, input.Price.Premium, input.Price.Unlimited); err != nil {
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
	var b beatstore.Beat
	query := fmt.Sprintf(`
	SELECT * FROM %s
  LEFT JOIN lateral (
    SELECT tag.id, tag.beat_id, tag.tag_name
    FROM %s
    WHERE tag.beat_id = beat.id
  ) c ON true
	WHERE beat.id = $1;
	`, beatTable, tagTable)
	row := r.db.QueryRow(query, beatId)
	err := row.Scan(
		&b.Id,
		&b.ArtistId,
		&b.Name,
		&b.Bpm,
		&b.Key,
		&b.PhotoPath,
		&b.MP3Path,
		&b.WavPath,
		&b.Likes,
		&b.Genre,
		&b.Mood,
		&b.CreatedAt,
		pq.Array(&b.Tags),
		pq.Array(&b.Tags),
		pq.Array(&b.Tags),
	)

	return b, err
}
