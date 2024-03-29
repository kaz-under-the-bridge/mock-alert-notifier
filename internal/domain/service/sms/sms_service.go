package sms

import (
	"context"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/sms"
)

var _ ServiceInterface = (*Service)(nil)

type ServiceInterface interface {
	Send(sms *model_sms.SMSMessage) error
}

type Service struct {
	ctx  context.Context
	repo sms.RepositoryInterface
}

func NewSMSService(ctx context.Context, r sms.RepositoryInterface) ServiceInterface {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) Send(sms *model_sms.SMSMessage) error {
	return s.repo.Send(sms)
}
