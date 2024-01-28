package email

import (
	"time"

	model_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/email"
)

// verifyなどの共通の処理を書く
type EmailClient struct{}

func (c *EmailClient) updateSentAt(email *model_email.Email) error {
	now := time.Now()
	email.SentAt = &now

	return nil
}

type MockEmailClient struct {
	*EmailClient
}

func NewExportMockEmailClient() EmailClientInterface {
	return &MockEmailClient{
		EmailClient: &EmailClient{},
	}
}

func (c *MockEmailClient) Send(email *model_email.Email) error {
	c.updateSentAt(email)
	return nil
}
