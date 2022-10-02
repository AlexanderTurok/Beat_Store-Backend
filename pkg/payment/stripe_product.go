package payment

import (
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/price"
	"github.com/stripe/stripe-go/v73/product"
)

type Product interface {
	CreateProduct() (string, error)
	CreatePrice(stripeId string) error
}

func (p *Payment) CreateProduct() (string, error) {
	params := &stripe.ProductParams{
		Name: stripe.String("Basic Dashboard"),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			UnitAmount: stripe.Int64(1000),
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval: stripe.String("month"),
			},
		},
	}
	params.AddExpand("default_price")
	result, err := product.New(params)

	return result.ID, err
}

func (p *Payment) CreatePrice(stripeId string) error {
	params := &stripe.PriceParams{
		Product:    stripe.String(stripeId),
		UnitAmount: stripe.Int64(1000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
	}
	_, err := price.New(params)

	return err
}
