package sms

import (
	"context"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
)

type SMSClientInterface interface {
	Send(SMS *model_sms.SMSMessage) error
}

type TwilioSMSClient struct {
	ctx context.Context
	*SMSClient
}

//type OtherServiceSMSClient struct {
//}
//func (e *OtherServiceSMSClient) Send(SMS model.SMS) error {
//	return nil
//}
//func NewOtherServiceSMSClient(ctx context.Context) SMSClientInterface {
//	return &OtherServiceSMSClient{}
//}

func NewTwilioSMSClient(ctx context.Context) SMSClientInterface {
	return &TwilioSMSClient{
		ctx:       ctx,
		SMSClient: &SMSClient{},
	}
}

func (e *TwilioSMSClient) Send(SMS *model_sms.SMSMessage) error {
	if err := sendSMSByTwilioAPI(SMS, e.ctx); err != nil {
		return err
	}
	e.updateSentAt(SMS)
	return nil
}

func sendSMSByTwilioAPI(SMS *model_sms.SMSMessage, ctx context.Context) error {
	// ToDo: あとで実装

	return nil
}
