package service

import "github.com/AlexanderTurok/beat-store-backend/internal/repository"

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
