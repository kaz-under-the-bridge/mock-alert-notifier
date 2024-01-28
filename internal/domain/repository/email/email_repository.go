package email

import (
	"context"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/email"
)

type RepositoryInterface interface {
	Send(email *model.Email) error
}

type Repository struct {
	ctx    context.Context
	client email.EmailClientInterface
}

func NewEmailRepository(ctx context.Context, client email.EmailClientInterface) *Repository {
	return &Repository{
		ctx:    ctx,
		client: client,
	}
}

func (r *Repository) Send(email *model.Email) error {
	err := r.client.Send(email)
	if err != nil {
		return err
	}

	return nil
}
