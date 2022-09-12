package repository

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("beat_store_db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
