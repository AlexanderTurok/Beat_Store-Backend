package service

import "github.com/AlexanderTurok/beat-store-backend/internal/repository"

type PaymentService struct {
	repos repository.Payment
}

func NewPaymentService(repos repository.Payment) *PaymentService {
	return &PaymentService{
		repos: repos,
	}
}
