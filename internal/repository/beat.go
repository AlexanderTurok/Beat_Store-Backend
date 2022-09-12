package repository

import (
	"fmt"
	"time"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jackskj/carta"
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

// func (r *BeatRepository) Get(beatId int) (beatstore.Beat, error) {
// 	beat := beatstore.Beat{}

// 	query := fmt.Sprintf(`SELECT beat.*, tag.id AS tag_id, tag.tag_name AS tag_name FROM %s LEFT OUTER JOIN %s ON beat.id = tag.beat_id WHERE beat.id=$1`, beatTable, tagTable)
// 	rows, err := r.db.Query(query, beatId)
// 	if err != nil {
// 		return beatstore.Beat{}, err
// 	}

// 	err = carta.Map(rows, &beat)

// 	return beat, err
// }

func (r *BeatRepository) GetAll() ([]beatstore.Beat, error) {
	beats := []beatstore.Beat{}

	query := fmt.Sprintf(`SELECT beat.*, tag.id AS tag_id, tag.tag_name AS tag_name FROM %s LEFT OUTER JOIN %s ON beat.id = tag.beat_id`, beatTable, tagTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}

func (r *BeatRepository) GetAllArtistsBeats(artistId int) ([]beatstore.Beat, error) {
	beats := []beatstore.Beat{}

	query := fmt.Sprintf(`SELECT beat.*, tag.id AS tag_id, tag.tag_name AS tag_name FROM %s LEFT OUTER JOIN %s ON beat.id = tag.beat_id WHERE beat.artist_id=$1`, beatTable, tagTable)
	rows, err := r.db.Query(query, artistId)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}
