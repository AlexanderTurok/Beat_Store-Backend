package service

import (
	"errors"

	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type AccountService struct {
	repos repository.Account
}

func NewAccountService(repos repository.Account) *AccountService {
	return &AccountService{
		repos: repos,
	}
}

func (s *AccountService) Get(accountId int) (beatstore.Account, error) {
	return s.repos.Get(accountId)
}

func (s *AccountService) Update(accountId int, input beatstore.AccountUpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repos.Update(accountId, input)
}

func (s *AccountService) Delete(accountId int, password beatstore.Password) error {
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
