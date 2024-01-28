package email

import (
	"context"

	model_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/email"
	infra_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/email"
)

type RepositoryInterface interface {
	Send(email *model_email.Email) error
}

type Repository struct {
	ctx    context.Context
	client infra_email.EmailClientInterface
}

func NewEmailRepository(ctx context.Context, client infra_email.EmailClientInterface) *Repository {
	return &Repository{
		ctx:    ctx,
		client: client,
	}
}

func (r *Repository) Send(email *model_email.Email) error {
	err := r.client.Send(email)
	if err != nil {
		return err
	}

	return nil
}
