package sms

import (
	"time"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
)

type SMSClient struct{}

func (c *SMSClient) updateSentAt(SMS *model.SMSMessage) error {
	now := time.Now()
	SMS.SentAt = &now

	return nil
}

type MockSMSClient struct {
	*SMSClient
}

func NewExportMockSMSClient() SMSClientInterface {
	return &MockSMSClient{
		SMSClient: &SMSClient{},
	}
}

func (c *MockSMSClient) Send(SMS *model.SMSMessage) error {
	c.updateSentAt(SMS)
	return nil
}
