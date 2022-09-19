package service

import (
	"fmt"
	"os"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
)

type EmailService struct {
	sender email.Client
}

func NewEmailService(sender email.Client) *EmailService {
	return &EmailService{
		sender: sender,
	}
}

func (s *EmailService) SendVerificationEmail(input model.Account) error {
	result, err := s.sender.AddEmailToList(email.AddEmailToList{
		ListId: os.Getenv("VERIFICATION_LIST_ID"),
		Emails: []email.EmailData{
			{
				Email: input.Email,
				Variables: map[string]string{
					"name": input.Name,
					"link": s.createVerificationLink(input.Id),
				},
			},
		},
	})
	if !result.Result {
		return fmt.Errorf("sendulse response: got: %t, expected: true", result.Result)
	}

	return err
}

func (s *EmailService) createVerificationLink(id int) string {
	domain := os.Getenv("DOMAIN")
	return fmt.Sprintf("%s/api/account/confirmation/%d", domain, id)
}
