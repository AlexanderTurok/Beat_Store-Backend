package service

import (
	"errors"

	model "github.com/AlexanderTurok/beat-store-backend/internal"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
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

func (s *ArtistService) Get(accountId int) (model.Account, error) {
	return s.repos.Get(accountId)
}

func (s *ArtistService) GetAll() ([]model.Account, error) {
	return s.repos.GetAll()
}

func (s *ArtistService) Delete(accountId int, password string) error {
	password = generatePasswordHash(password)
	passwordHash, err := s.repos.GetPasswordHash(accountId)
	if err != nil {
		return err
	}

	if password != passwordHash {
		return errors.New("invalid password")
	}

	return s.repos.Delete(accountId)
}
