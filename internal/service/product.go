package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
)

type ProductService struct {
	repos repository.Product
}

func NewProductService(repos repository.Product) *ProductService {
	return &ProductService{
		repos: repos,
	}
}
