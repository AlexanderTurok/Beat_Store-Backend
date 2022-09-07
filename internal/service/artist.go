package service

import (
	"errors"

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

func (s *ArtistService) Delete(accountId int, password beatstore.AccountPassword) error {
	password.Password = generatePasswordHash(password.Password)
	passwordHash, err := s.repos.GetPasswordHash(accountId)
	if err != nil {
		return err
	}

	if password.Password != passwordHash.Password {
		return errors.New("invalid password")
	}

	return s.repos.Delete(accountId)
}
