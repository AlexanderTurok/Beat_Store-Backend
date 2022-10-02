package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
)

type AuthService struct {
	repos   repository.Auth
	hasher  hash.PasswordHasher
	manager auth.TokenManager
}

func NewAuthService(repos repository.Auth, hasher hash.PasswordHasher, manager auth.TokenManager) *AuthService {
	return &AuthService{
		repos:   repos,
		hasher:  hasher,
		manager: manager,
	}
}

func (s *AuthService) CreateAccount(account model.Account) (int, error) {
	account.Password = s.hasher.Hash(account.Password)
	return s.repos.Create(account)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	userId, err := s.repos.GetId(email, s.hasher.Hash(password))
	if err != nil {
		return "", err
	}

	token, err := s.manager.NewJWT(userId, tokenTTL)

	return token, err
}
