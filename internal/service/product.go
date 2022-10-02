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
		repos:     repos,
		paymenter: paymenter,
	}
}

func (s *ProductService) Create(artistId int) (int64, error) {
	stripeId, err := s.paymenter.CreateProduct()
	if err != nil {
		return 0, err
	}

	if err := s.paymenter.CreatePrice(); err != nil {
		return 0, err
	}

	return s.repos.Create(artistId, stripeId)
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
