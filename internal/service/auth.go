package service

import (
	"errors"
	"os"
	"time"

	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/hash"
	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repos  repository.Authorization
	hasher *hash.SHA1Hasher
}

func NewAuthService(repos repository.Authorization, hasher *hash.SHA1Hasher) *AuthService {
	return &AuthService{
		repos: repos,
	}
}

func (s *AuthService) CreateAccount(account model.Account) (int, error) {
	account.Password = s.hasher.Hash(account.Password)

	return s.repos.CreateAccount(account)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	userId, err := s.repos.GetAccountId(email, s.hasher.Hash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	signingKey := os.Getenv("SIGNING_KEY")

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	signingKey := os.Getenv("SIGNING_KEY")

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
