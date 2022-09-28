package payment

import (
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/account"
	"github.com/stripe/stripe-go/v73/accountlink"
)

func (p *Payment) CreateSellerRegistrationURL(refreshURL, returnURL string) (*stripe.AccountLink, error) {
	account, err := p.createAccount()
	if err == nil {
		linkParams := &stripe.AccountLinkParams{
			Account:    stripe.String(account.ID),
			RefreshURL: stripe.String("https://example.com/reauth"),
			ReturnURL:  stripe.String("https://example.com/return"),
			Type:       stripe.String("account_onboarding"),
		}

		return accountlink.New(linkParams)
	}

	return nil, err
}

func (p *Payment) createAccount() (*stripe.Account, error) {
	accountParams := &stripe.AccountParams{
		Type: stripe.String(string(stripe.AccountTypeExpress)),
	}

	return account.New(accountParams)
}
