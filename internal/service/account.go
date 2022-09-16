package service

import (
	"errors"

	model "github.com/AlexanderTurok/beat-store-backend/internal"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type AccountService struct {
	repos repository.Account
}

func NewAccountService(repos repository.Account) *AccountService {
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

func (s *AccountService) Delete(accountId int, password string) error {
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
