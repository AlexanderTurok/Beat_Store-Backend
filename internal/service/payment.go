package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/payment"
)

type PaymentService struct {
	repos   repository.Payment
	payment payment.Payment
}

func NewPaymentService(repos repository.Payment, payment payment.Payment) *PaymentService {
	return &PaymentService{
		repos:   repos,
		payment: payment,
	}
}

func (s *PaymentService) CreatePaymentIntent(input model.PaymentInfo) (payment.PaymentIntent, error) {
	paymentIntent, err := s.payment.CreatePaymentIntent(input.OrderAmount)
	if err != nil {
		return payment.PaymentIntent{}, err
	}

	return paymentIntent, nil
}
