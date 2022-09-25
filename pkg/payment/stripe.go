package payment

import (
	"github.com/stripe/stripe-go/v73"
)

type Payment struct {
}

func NewPayment(apiKey string) *Payment {
	stripe.Key = apiKey
	return &Payment{}
}
