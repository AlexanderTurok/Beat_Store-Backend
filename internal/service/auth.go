package service

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{
		repos: repos,
	}
}

func (s *AuthService) CreateUser(user beatstore.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func generatePasswordHash(password string) string {
	salt := os.Getenv("SALT")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
