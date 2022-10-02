package service

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/AlexanderTurok/beat-store-backend/pkg/email"
)

type EmailService struct {
	sender *email.Client
}

func NewEmailService(sender *email.Client) *EmailService {
	return &EmailService{
		sender: sender,
	}
}

func (s *EmailService) SendVerificationEmail(input model.Account) error {
	listId, _ := strconv.Atoi(os.Getenv("VERIFICATION_LIST_ID"))
	result, err := s.sender.AddEmailToList(email.AddEmailToList{
		ListId: listId,
		Emails: []email.EmailData{
			{
				Email: input.Email,
				Variables: map[string]string{
					"Name": input.Name,
					"Link": s.createVerificationLink(input.Username),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	if !result.Result {
		return fmt.Errorf("sendulse response: got: %t, expected: true", result.Result)
	}

	return nil
}

func (s *EmailService) createVerificationLink(value string) string {
	domain := os.Getenv("DOMAIN")
	return fmt.Sprintf("%s/api/accounts/%s", domain, value)
}
