package payment

import (
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/paymentintent"
)

type PaymentIntent struct {
	ClientSecret string `json:"clientSecret"`
}

func (p *Payment) CreatePaymentIntent(orderAmount int64) (PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(orderAmount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)

	paymentIntent := PaymentIntent{
		ClientSecret: pi.ClientSecret,
	}

	return paymentIntent, err
}
