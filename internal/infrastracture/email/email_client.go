package email

import (
	"time"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
)

// verifyなどの共通の処理を書く
type EmailClient struct{}

func (c *EmailClient) updateSentAt(email *model.Email) error {
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

func (c *MockEmailClient) Send(email *model.Email) error {
	c.updateSentAt(email)
	return nil
}
