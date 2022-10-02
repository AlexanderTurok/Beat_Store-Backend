package repository

import (
	"fmt"
	"strings"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
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

func (r *BeatRepository) Create(productId int64, input model.Beat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var beatId int
	insertBeatQuery := fmt.Sprintf(`INSERT INTO %s (
			product_id, name, bpm, key, photo_path, mp3_path, wav_path, genre, mood)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		beatTable,
	)

	if err := tx.QueryRow(insertBeatQuery,
		productId,
		input.Name,
		input.Bpm,
		input.Key,
		input.PhotoPath,
		input.MP3Path,
		input.WavPath,
		input.Genre,
		input.Mood,
	).Scan(&beatId); err != nil {
		tx.Rollback()
		return 0, err
	}

	insertPriceQuery := fmt.Sprintf(`
		INSERT INTO %s (id, standart, premium, ultimate)
		VALUES ($1, $2, $3, $4)`, priceTable)
	if _, err := tx.Exec(insertPriceQuery, beatId, input.Price.Standart,
		input.Price.Premium, input.Price.Ultimate); err != nil {
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

func (r *BeatRepository) Get(beatId int) (model.Beat, error) {
	beat := model.Beat{}

	query := fmt.Sprintf(`SELECT product.created_at, beat.*, price.*, tag.id AS tag_id, tag.name AS tag_name FROM %s 
		LEFT OUTER JOIN %s ON product.id = beat.product_id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id 
		LEFT OUTER JOIN %s ON beat.id = price.id
		WHERE beat.id=$1`, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, beatId)
	if err != nil {
		return model.Beat{}, err
	}

	err = carta.Map(rows, &beat)

	return beat, err
}

func (r *BeatRepository) GetAll() ([]model.Beat, error) {
	var beats []model.Beat

	query := fmt.Sprintf(`SELECT product.created_at, beat.*, price.*, tag.id AS tag_id, tag.name AS tag_name FROM %s 
		LEFT OUTER JOIN %s ON product.id = beat.product_id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id 
		LEFT OUTER JOIN %s ON beat.id = price.id`, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}

func (r *BeatRepository) GetArtistsBeat(artistId, beatId int) (model.Beat, error) {
	var beat model.Beat

	query := fmt.Sprintf(`SELECT product.created_at, beat.*, price.*, tag.id AS tag_id, tag.name AS tag_name FROM %s 
		LEFT OUTER JOIN %s ON product.id = beat.product_id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id 
		LEFT OUTER JOIN %s ON beat.id = price.id
		WHERE beat.id=$1 AND product.artist_id=$2`,
		beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, beatId, artistId)
	if err != nil {
		return beat, err
	}

	err = carta.Map(rows, &beat)

	return beat, err
}

func (r *BeatRepository) GetAllArtistsBeats(artistId int) ([]model.Beat, error) {
	var beats []model.Beat

	query := fmt.Sprintf(`SELECT product.created_at, beat.*, price.*, tag.id AS tag_id, tag.name AS tag_name FROM %s 
		LEFT OUTER JOIN %s ON product.id = beat.product_id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id 
		LEFT OUTER JOIN %s ON beat.id = price.id
		WHERE product.artist_id=$1`,
		beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, artistId)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}

func (r *BeatRepository) Update(beatId int, input model.BeatUpdateInput) error {
	query, args := createBeatUpdateQuery(beatId, input)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *BeatRepository) Delete(beatId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", beatTable)
	_, err := r.db.Exec(query, beatId)

	return err
}

func createBeatUpdateQuery(beatId int, input model.BeatUpdateInput) (string, []interface{}) {
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

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", beatTable, setQuery, argId)
	args = append(args, beatId)

	return query, args
}
