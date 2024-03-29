package sms

import (
	"context"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/sms"
)

type RepositoryInterface interface {
	Send(email *model_sms.SMSMessage) error
}

type Repository struct {
	ctx    context.Context
	client sms.SMSClientInterface
}

func NewSMSRepository(ctx context.Context, client sms.SMSClientInterface) *Repository {
	return &Repository{
		ctx:    ctx,
		client: client,
	}
}

func (r *Repository) Send(sms *model_sms.SMSMessage) error {
	err := r.client.Send(sms)
	if err != nil {
		return err
	}

	return nil
}
