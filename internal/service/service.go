package service

import "github.com/AlexanderTurok/beat-store-backend/internal/repository"

type Authorization interface {
}

type Beat interface {
}

type Service struct {
	Authorization
	Beat
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
