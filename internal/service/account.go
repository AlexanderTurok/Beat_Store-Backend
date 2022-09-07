package service

import (
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
