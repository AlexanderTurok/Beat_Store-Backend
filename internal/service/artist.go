package service

import (
	"errors"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
)

type ArtistService struct {
	repos  repository.Artist
	hasher hash.PasswordHasher
}

func NewArtistService(repos repository.Artist, hasher hash.PasswordHasher) *ArtistService {
	return &ArtistService{
		repos:  repos,
		hasher: hasher,
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

func (s *ArtistService) Delete(accountId int, inputPassword string) error {
	inputPassword = s.hasher.Hash(inputPassword)

	passwordHash, err := s.repos.GetPasswordHash(accountId)
	if err != nil {
		return err
	}

	if inputPassword != passwordHash {
		return errors.New("invalid password")
	}

	return s.repos.Delete(accountId)
}
