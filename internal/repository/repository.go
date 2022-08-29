package repository

type Authorization interface {
}

type Beat interface {
}

type Repository struct {
	Authorization
	Beat
}

func NewRepository() *Repository {
	return &Repository{}
}
