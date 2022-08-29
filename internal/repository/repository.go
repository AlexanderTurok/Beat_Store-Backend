package repository

import "database/sql"

type Authorization interface {
}

type Beat interface {
}

type Repository struct {
	Authorization
	Beat
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
