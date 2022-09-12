package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	accountTable = "account"
	artistTable  = "artist"

	accountBeatTable     = "account_beat"
	accountPlaylistTable = " account_playlist"

	playlistBeatTable = "playlist_beat"
	playlistTable     = "playlist"

	beatTable  = "beat"
	tagTable   = "tag"
	priceTable = "price"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "dbname=beat_store_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
