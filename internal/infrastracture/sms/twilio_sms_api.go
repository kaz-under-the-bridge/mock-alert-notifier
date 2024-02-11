package sms

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/pkg/errors"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSClientInterface interface {
	Send(SMS *model_sms.SMSMessage) error
}

type TwilioSMSClient struct {
	ctx context.Context
	*SMSClient
}

func NewTwilioSMSClient(ctx context.Context) SMSClientInterface {
	return &TwilioSMSClient{
		ctx:       ctx,
		SMSClient: &SMSClient{},
	}
}

func (e *TwilioSMSClient) Send(sms *model_sms.SMSMessage) error {
	if err := sendSMSByTwilioAPI(sms, e.ctx); err != nil {
		return err
	}
	e.updateSentAt(sms)
	return nil
}

func sendSMSByTwilioAPI(sms *model_sms.SMSMessage, ctx context.Context) error {
	sid := helper.GetTwilioSid(ctx)
	token := helper.GetTwilioAuthToken(ctx)

	if sid == "" || token == "" {
		return errors.New("twilio sid or token is not set")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: sid,
		Password: token,
	})

	params := &api.CreateMessageParams{}
	params.SetTo(sms.GetToNumberE164())
	params.SetFrom(sms.GetFromNumberE164())
	params.SetBody(sms.GetBody())

	resp, err := client.Api.CreateMessage(params)

	response, _ := json.Marshal(*resp)
	slogfrom := slog.String("from", sms.GetFromNumberE164())
	slogto := slog.String("to", sms.GetToNumberE164())

	if err != nil {
		helper.Logger.Error("SMS sent failed", slogfrom, slogto, slog.String("api_response", string(response)))
		return errors.Wrap(err, "twilio sms message creation api error")
	}

	helper.Logger.Info("SMS sent successfully", slogfrom, slogto)

	return nil
}
