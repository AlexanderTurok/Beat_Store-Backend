package service

import (
	"github.com/AlexanderTurok/beat-store-backend/internal/repository"
	"github.com/AlexanderTurok/beat-store-backend/pkg/payment"
)

const (
	refreshURL = "http://localhost:8000/api/accounts/artists"
	returnURL  = "http://localhost:8000/api/accounts/artists"
)

type PaymentService struct {
	repos     repository.Payment
	paymenter payment.Payment
}

func NewPaymentService(repos repository.Payment, paymenter payment.Payment) *PaymentService {
	return &PaymentService{
		repos:     repos,
		paymenter: paymenter,
	}
}

func (s *PaymentService) CreatePaymentAccount(accountId int) (string, error) {
	stripeId, err := s.paymenter.CreateAccount()
	if err != nil {
		return "", err
	}

	if err := s.repos.CreatePaymentAccount(accountId, stripeId); err != nil {
		return "", err
	}

	return s.paymenter.CreateRegistrationURL(stripeId, refreshURL, returnURL)
}

func (s *PaymentService) DeletePaymentAccount(accountId int) (string, error) {
	return "", nil
}
