package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type BeatService struct {
	repos repository.Beat
}

func NewBeatService(repos repository.Beat) *BeatService {
	return &BeatService{
		repos: repos,
	}
}
