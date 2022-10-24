package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// FIXME: use configs
	db, err := sqlx.Open("postgres", "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
