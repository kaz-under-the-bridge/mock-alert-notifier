package email

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	model_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/email"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClientInterface interface {
	Send(email *model_email.Email) error
}

type TwilioEmailClient struct {
	ctx context.Context
	*EmailClient
}

//type OtherServiceEmailClient struct {
//}
//func (e *OtherServiceEmailClient) Send(email model_email.Email) error {
//	return nil
//}
//func NewOtherServiceEmailClient(ctx context.Context) EmailClientInterface {
//	return &OtherServiceEmailClient{}
//}

func NewTwilioEmailClient(ctx context.Context) EmailClientInterface {
	return &TwilioEmailClient{
		ctx:         ctx,
		EmailClient: &EmailClient{},
	}
}

func (e *TwilioEmailClient) Send(email *model_email.Email) error {
	if err := sendEmailBySendgrid(email, e.ctx); err != nil {
		return err
	}
	e.updateSentAt(email)
	return nil
}

func sendEmailBySendgrid(email *model_email.Email, ctx context.Context) error {
	_ = fmt.Sprintf("%s", email.CcAddresses)
	_ = fmt.Sprintf("%s", email.BccAddresses)

	from := mail.NewEmail("", email.FromAddress)
	var tos []*mail.Email

	for _, addr := range email.ToAddresses {
		tos = append(tos, mail.NewEmail("", addr))
	}

	content := mail.NewContent("text/plain", email.Body)

	message := mail.NewV3MailInit(from, email.Subject, tos[0], content)

	if len(tos) > 1 {
		personalization := mail.NewPersonalization()
		for _, to := range tos[1:] {
			personalization.AddTos(to)
		}
		message.AddPersonalizations(personalization)
	}

	key := helper.GetSendgridKey(ctx)
	if key == "" {
		return errors.New("unexpected error, sendgrid key is not set")
	}

	request := sendgrid.GetRequest(key, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(message)
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		helper.Logger.Warn("sendgrid response status code is not 2xx", slog.Int("StatusCode", response.StatusCode), slog.String("Body", response.Body), slog.String("Headers", fmt.Sprintf("%s", response.Headers)))
	} else {
		helper.Logger.Info("sendgrid response status code is 2xx", slog.Int("StatusCode", response.StatusCode))
	}

	return nil
}
