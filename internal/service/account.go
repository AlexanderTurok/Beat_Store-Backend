package service

import (
	"errors"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
)

type AccountService struct {
	repos  repository.Account
	hasher *hash.SHA1Hasher
}

func NewAccountService(repos repository.Account, hasher *hash.SHA1Hasher) *AccountService {
	return &AccountService{
		repos: repos,
	}
}

func (s *AccountService) Get(accountId int) (model.Account, error) {
	return s.repos.Get(accountId)
}

func (s *AccountService) Update(accountId int, input model.AccountUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(accountId, input)
}

func (s *AccountService) Delete(accountId int, inputPassword string) error {
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
