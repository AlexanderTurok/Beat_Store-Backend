package repository

import (
	"fmt"
	"strings"
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

	if err := tx.QueryRow(insertBeatQuery,
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
	).Scan(&beatId); err != nil {
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
	beat := beatstore.Beat{}

	query := fmt.Sprintf(`SELECT beat.*, tag.id AS tag_id, tag.tag_name AS tag_name FROM %s LEFT OUTER JOIN %s ON beat.id = tag.beat_id WHERE beat.id=$1`, beatTable, tagTable)
	rows, err := r.db.Query(query, beatId)
	if err != nil {
		return beatstore.Beat{}, err
	}

	err = carta.Map(rows, &beat)

	return beat, err
}

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

func (r *BeatRepository) Update(beatId int, input beatstore.BeatUpdateInput) error {
	query, args := createBeatUpdateQuery(beatId, input)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *BeatRepository) Delete(beatId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", beatTable)
	_, err := r.db.Exec(query, beatId)

	return err
}

func createBeatUpdateQuery(beatId int, input beatstore.BeatUpdateInput) (string, []interface{}) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Bpm != nil {
		setValues = append(setValues, fmt.Sprintf("bpm=$%d", argId))
		args = append(args, *input.Bpm)
		argId++
	}

	if input.Key != nil {
		setValues = append(setValues, fmt.Sprintf("key=$%d", argId))
		args = append(args, *input.Key)
		argId++
	}

	if input.PhotoPath != nil {
		setValues = append(setValues, fmt.Sprintf("photo_path=$%d", argId))
		args = append(args, *input.PhotoPath)
		argId++
	}

	if input.MP3Path != nil {
		setValues = append(setValues, fmt.Sprintf("mp3_path=$%d", argId))
		args = append(args, *input.MP3Path)
		argId++
	}

	if input.WavPath != nil {
		setValues = append(setValues, fmt.Sprintf("wav_path=$%d", argId))
		args = append(args, *input.WavPath)
		argId++
	}

	if input.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *input.Genre)
		argId++
	}

	if input.Mood != nil {
		setValues = append(setValues, fmt.Sprintf("mood=$%d", argId))
		args = append(args, *input.Mood)
		argId++
	}

	if input.StandartPrice != nil {
		setValues = append(setValues, fmt.Sprintf("standart_price=$%d", argId))
		args = append(args, *input.StandartPrice)
		argId++
	}

	if input.PremiumPrice != nil {
		setValues = append(setValues, fmt.Sprintf("premium_price=$%d", argId))
		args = append(args, *input.PremiumPrice)
		argId++
	}

	if input.UnlimitedPrice != nil {
		setValues = append(setValues, fmt.Sprintf("unlimited_price=$%d", argId))
		args = append(args, *input.UnlimitedPrice)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", beatTable, setQuery, argId)
	args = append(args, beatId)

	return query, args
}
