package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// FIXME: use configs
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
