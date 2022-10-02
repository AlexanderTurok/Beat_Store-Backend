package service

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/payment"
)

type ProductService struct {
	repos     repository.Product
	paymenter payment.Product
}

func NewProductService(repos repository.Product, paymenter payment.Product) *ProductService {
	return &ProductService{
		repos: repos,
	}
}

func (s *ProductService) Create(input model.Beat) (string, error) {
	return "", nil
}

func (s *ProductService) Get(productId string) (model.Product, error) {
	return model.Product{}, nil
}

func (s *ProductService) GetAll(productId string) ([]model.Product, error) {
	return nil, nil
}

func (s *ProductService) Delete(productId string) error {
	return nil
}
