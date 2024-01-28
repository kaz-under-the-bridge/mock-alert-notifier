package email

import (
	"context"

	model_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/email"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/email"
)

//var _ ServiceInterface = (*Service)(nil)

type ServiceInterface interface {
	Send(email model_email.Email) error
}

type Service struct {
	ctx  context.Context
	repo email.RepositoryInterface
}

func NewEmailService(ctx context.Context, r email.RepositoryInterface) *Service {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) Send(email *model_email.Email) error {
	return s.repo.Send(email)
}
