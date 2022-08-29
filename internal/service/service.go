package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type Authorization interface {
}

type Beat interface {
	GetById(id int) (beatstore.Beat, error)
	GetAll() ([]beatstore.Beat, error)
}

type Service struct {
	Authorization
	Beat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Beat: NewBeatService(repos.Beat),
	}
}
