package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type BeatService struct {
	repos repository.Beat
}

func NewBeatService(repos repository.Beat) *BeatService {
	return &BeatService{
		repos: repos,
	}
}

func (s *BeatService) Create(artistId int, input beatstore.Beat) (int, error) {
	return s.repos.Create(artistId, input)
}

func (s *BeatService) Get(beatId int) (beatstore.Beat, error) {
	return s.repos.Get(beatId)
}
