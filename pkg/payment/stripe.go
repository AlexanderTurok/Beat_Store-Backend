package payment

import (
	"github.com/stripe/stripe-go/v72"
)

type Payment struct {
}

func NewPayment(apiKey string) *Payment {
	stripe.Key = apiKey
	return &Payment{}
}
