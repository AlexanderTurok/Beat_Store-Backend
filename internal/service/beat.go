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

func (s *BeatService) Create(userId int, beat beatstore.Beat) (int, error) {
	return s.repos.Create(userId, beat)
}

func (s *BeatService) GetAll() ([]beatstore.Beat, error) {
	return s.repos.GetAll()
}

func (s *BeatService) GetUsersBeats(userId int) ([]beatstore.Beat, error) {
	return s.repos.GetUsersBeats(userId)
}

func (s *BeatService) Update(userId, beatId int, input beatstore.BeatUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(userId, beatId, input)
}

func (s *BeatService) Delete(userId, beatId int) error {
	return s.repos.Delete(userId, beatId)
}
