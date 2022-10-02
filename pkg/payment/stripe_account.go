package payment

import (
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/account"
	"github.com/stripe/stripe-go/v73/accountlink"
)

func (p *Payment) CreateRegistrationURL(stripeId, refreshURL, returnURL string) (string, error) {
	linkParams := &stripe.AccountLinkParams{
		Account:    stripe.String(stripeId),
		Type:       stripe.String("account_onboarding"),
		RefreshURL: stripe.String(refreshURL),
		ReturnURL:  stripe.String(returnURL),
	}
	link, err := accountlink.New(linkParams)

	return link.URL, err
}

func (p *Payment) CreateAccount() (string, error) {
	accountParams := &stripe.AccountParams{
		Type: stripe.String(string(stripe.AccountTypeExpress)),
	}
	account, err := account.New(accountParams)

	return account.ID, err
}
