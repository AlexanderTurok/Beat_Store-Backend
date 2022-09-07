package service

import (
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
