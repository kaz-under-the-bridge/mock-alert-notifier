package sms

import (
	"context"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/sms"
)

//var _ ServiceInterface = (*Service)(nil)

type ServiceInterface interface {
	Send(SMS model.SMSMessage) error
}

type Service struct {
	ctx  context.Context
	repo sms.RepositoryInterface
}

func NewSMSService(ctx context.Context, r sms.RepositoryInterface) *Service {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) Send(SMS *model.SMSMessage) error {
	return s.repo.Send(SMS)
}
