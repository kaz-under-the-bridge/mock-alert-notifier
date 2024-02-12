package voice

import (
	"context"

	model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/voice"
)

type ServiceInterface interface {
	Send(voice *model_voice.VoiceMessage) error
}

type Service struct {
	ctx  context.Context
	repo voice.RepositoryInterface
}

func NewVoiceService(ctx context.Context, r voice.RepositoryInterface) ServiceInterface {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) Send(voice *model_voice.VoiceMessage) error {
	return s.repo.Call(voice)
}
