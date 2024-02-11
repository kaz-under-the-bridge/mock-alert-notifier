package sms

import (
	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
)

type SMSClient struct{}

func (c *SMSClient) updateSentAt(SMS *model_sms.SMSMessage) error {
	SMS.UpdateSentAt()

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

func (c *MockSMSClient) Send(SMS *model_sms.SMSMessage) error {
	c.updateSentAt(SMS)
	return nil
}
