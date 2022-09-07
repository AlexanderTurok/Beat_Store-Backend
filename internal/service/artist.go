package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type ArtistService struct {
	repos repository.Artist
}

func NewArtistService(repos repository.Artist) *ArtistService {
	return &ArtistService{
		repos: repos,
	}
}

func (s *ArtistService) Create(accountId int) error {
	return s.repos.Create(accountId)
}

func (s *ArtistService) Get(accountId int) (beatstore.Account, error) {
	return s.repos.Get(accountId)
}
