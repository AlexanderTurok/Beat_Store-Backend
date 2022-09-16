package service

import (
	"errors"
	"time"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/auth"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
)

const (
	tokenTTL = time.Hour * 24
)

type AccountService struct {
	repos   repository.Account
	hasher  hash.SHA1Hasher
	manager auth.Manager
}

func NewAccountService(repos repository.Account, hasher hash.SHA1Hasher, manager auth.Manager) *AccountService {
	return &AccountService{
		repos:   repos,
		hasher:  hasher,
		manager: manager,
	}
}

func (s *AccountService) CreateAccount(account model.Account) (int, error) {
	account.Password = s.hasher.Hash(account.Password)

	return s.repos.CreateAccount(account)
}

func (s *AccountService) GenerateToken(email, password string) (string, error) {
	userId, err := s.repos.GetAccountId(email, s.hasher.Hash(password))
	if err != nil {
		return "", err
	}

	token, err := s.manager.NewJWT(userId, tokenTTL)

	return token, err
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
