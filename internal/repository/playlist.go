package repository

import (
	"fmt"
	"time"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
)

type PlaylistRepository struct {
	db *sqlx.DB
}

func NewPlaylistRepository(db *sqlx.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (r *PlaylistRepository) Create(accountId int, input model.Playlist) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var playlistId int
	playlistInsertQuery := fmt.Sprintf("INSERT INTO %s (name, created_at) VALUES ($1, $2) RETURNING id", playlistTable)
	if err := tx.QueryRow(playlistInsertQuery, input.Name, time.Now()).Scan(&playlistId); err != nil {
		tx.Rollback()
		return 0, err
	}

	accountPlaylistInsertQuery := fmt.Sprintf(`INSERT INTO %s (account_id, playlist_id)
		VALUES ($1, $2)`, accountPlaylistTable)
	if _, err := tx.Exec(accountPlaylistInsertQuery, accountId, playlistId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return playlistId, tx.Commit()
}

func (r *PlaylistRepository) Get(playlistId int) (model.Playlist, error) {
	var playlist model.Playlist

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", playlistTable)
	err := r.db.Get(&playlist, query, playlistId)

	return playlist, err
}

func (r *PlaylistRepository) GetAll() ([]model.Playlist, error) {
	var playlists []model.Playlist

	query := fmt.Sprintf("SELECT * FROM %s", playlistTable)
	err := r.db.Select(&playlists, query)

	return playlists, err
}

func (r *PlaylistRepository) GetAccountsPlaylist(accountId, playlistId int) (model.Playlist, error) {
	var playlists model.Playlist

	query := fmt.Sprintf(` 
		SELECT playlist.* FROM %s 
			LEFT JOIN playlist ON playlist.id = account_playlist.playlist_id
		WHERE account_playlist.account_id=$1 AND playlist.id=$2`, accountPlaylistTable)
	err := r.db.Get(&playlists, query, accountId, playlistId)

	return playlists, err
}

func (r *PlaylistRepository) GetAllAccountsPlaylists(accountId int) ([]model.Playlist, error) {
	var playlists []model.Playlist

	query := fmt.Sprintf(` 
		SELECT playlist.* FROM %s 
			LEFT JOIN playlist ON playlist.id = account_playlist.playlist_id
		WHERE account_playlist.account_id=$1`, accountPlaylistTable)
	err := r.db.Select(&playlists, query, accountId)

	return playlists, err
}

func (r *PlaylistRepository) Update(playlistId int, input model.PlaylistUpdateInput) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", playlistTable)
	_, err := r.db.Exec(query, input.Name, playlistId)

	return err
}

func (r *PlaylistRepository) Delete(playlistId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", playlistTable)
	_, err := r.db.Exec(query, playlistId)

	return err
}

func (r *PlaylistRepository) AddBeat(playlistId, beatId int) error {
	query := fmt.Sprintf("INSERT INTO %s (playlist_id, beat_id) VALUES ($1, $2)", playlistBeatTable)
	_, err := r.db.Exec(query, playlistId, beatId)

	return err
}

func (r *PlaylistRepository) GetBeat(playlistId, beatId int) (model.Beat, error) {
	var beat model.Beat

	query := fmt.Sprintf(`
	SELECT 
		product.created_at,
		beat.*, 
		tag.id AS tag_id, 
		tag.name AS tag_name,
		price.* 
	FROM %s 
		LEFT OUTER JOIN %s ON playlist_beat.beat_id = beat.id
		LEFT OUTER JOIN %s ON beat.product_id = product.id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id
		LEFT OUTER JOIN %s ON beat.id = price.id
	WHERE playlist_beat.playlist_id=$1 AND beat.id=$2`,
		playlistBeatTable, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, playlistId, beatId)
	if err != nil {
		return model.Beat{}, err
	}

	err = carta.Map(rows, &beat)

	return beat, err
}

func (r *PlaylistRepository) GetAllBeats(playlistId int) ([]model.Beat, error) {
	var beats []model.Beat

	query := fmt.Sprintf(`
	SELECT 
		product.created_at,
		beat.*, 
		tag.id AS tag_id, 
		tag.name AS tag_name,
		price.* 
	FROM %s 
		LEFT OUTER JOIN %s ON playlist_beat.beat_id = beat.id
		LEFT OUTER JOIN %s ON beat.product_id = product.id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id
		LEFT OUTER JOIN %s ON beat.id = price.id
	WHERE playlist_beat.playlist_id=$1`,
		playlistBeatTable, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, playlistId)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}

func (r *PlaylistRepository) GetBeatFromAccountsPlaylists(accountId, playlistId, beatId int) (model.Beat, error) {
	var beat model.Beat

	query := fmt.Sprintf(`
	SELECT 
		product.created_at,
		beat.*, 
		tag.id AS tag_id, 
		tag.name AS tag_name,
		price.* 
	FROM %s 
		LEFT OUTER JOIN account_playlist ON account_playlist.account_id = $2
		LEFT OUTER JOIN %s ON playlist_beat.playlist_id = $1
		LEFT OUTER JOIN %s ON beat.product_id = product.id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id
		LEFT OUTER JOIN %s ON beat.id = price.id
	WHERE beat.id = $3`,
		playlistBeatTable, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, playlistId, accountId, beatId)
	if err != nil {
		return model.Beat{}, err
	}

	err = carta.Map(rows, &beat)

	return beat, err
}

func (r *PlaylistRepository) GetAllBeatsFromAccountsPlaylists(accountId, playlistId int) ([]model.Beat, error) {
	var beats []model.Beat

	query := fmt.Sprintf(`
	SELECT 
		product.created_at,
		beat.*, 
		tag.id AS tag_id, 
		tag.name AS tag_name,
		price.* 
	FROM %s 
		LEFT OUTER JOIN account_playlist ON account_playlist.account_id = $2
		LEFT OUTER JOIN %s ON playlist_beat.playlist_id = $1
		LEFT OUTER JOIN %s ON beat.product_id = product.id
		LEFT OUTER JOIN %s ON beat.id = tag.beat_id
		LEFT OUTER JOIN %s ON beat.id = price.id`,
		playlistBeatTable, beatTable, productTable, tagTable, priceTable)
	rows, err := r.db.Query(query, playlistId, accountId)
	if err != nil {
		return nil, err
	}

	err = carta.Map(rows, &beats)

	return beats, err
}

func (r *PlaylistRepository) DeleteBeat(playlistId, beatId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE playlist_id = $1 AND beat_id = $2", playlistBeatTable)
	_, err := r.db.Exec(query, playlistId, beatId)

	return err
}
