package repository

import (
	"fmt"
	"time"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
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
	fmt.Println(playlistId)

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
