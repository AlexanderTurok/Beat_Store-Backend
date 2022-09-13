package repository

import (
	"fmt"
	"time"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
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

func (r *PlaylistRepository) Create(accountId int, input beatstore.Playlist) (int, error) {
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

func (r *PlaylistRepository) GetAllAccountsPlaylists(accountId int) ([]beatstore.Playlist, error) {
	var playlists []beatstore.Playlist

	selectPlaylistsId := fmt.Sprintf(` 
		SELECT playlist.* FROM %s 
		LEFT JOIN playlist ON playlist.id = account_playlist.playlist_id
		WHERE account_playlist.account_id = $1`, accountPlaylistTable)
	err := r.db.Select(&playlists, selectPlaylistsId, accountId)

	return playlists, err
}

func (r *PlaylistRepository) Update(playlistId int, input beatstore.PlaylistUpdateInput) error {
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

func (r *PlaylistRepository) GetAllBeats(playlistId int) ([]beatstore.Beat, error) {
	beats := []beatstore.Beat{}

	query := fmt.Sprintf(`
	SELECT 
		beat.*, 
		tag.id AS tag_id, 
		tag.tag_name AS tag_name 
	FROM %s 
	LEFT OUTER JOIN %s ON playlist_beat.beat_id = beat.id
	LEFT OUTER JOIN %s ON beat.id = tag.beat_id
	WHERE playlist_beat.playlist_id = $1`,
		playlistBeatTable, beatTable, tagTable)
	rows, err := r.db.Query(query, playlistId)
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
