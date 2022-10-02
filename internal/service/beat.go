package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
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

func (s *BeatService) Create(artistId int64, input model.Beat) (int, error) {
	return s.repos.Create(artistId, input)
}

func (s *BeatService) Get(beatId int) (model.Beat, error) {
	return s.repos.Get(beatId)
}

func (s *BeatService) GetAll() ([]model.Beat, error) {
	return s.repos.GetAll()
}

func (s *BeatService) GetArtistsBeat(artistId, beatId int) (model.Beat, error) {
	return s.repos.GetArtistsBeat(artistId, beatId)
}

func (s *BeatService) GetAllArtistsBeats(artistId int) ([]model.Beat, error) {
	return s.repos.GetAllArtistsBeats(artistId)
}

func (s *BeatService) Update(beatId int, input model.BeatUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(beatId, input)
}

func (s *BeatService) Delete(beatId int) error {
	return s.repos.Delete(beatId)
}
