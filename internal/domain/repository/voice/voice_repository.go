package voice

import (
	"context"

	model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/voice"
)

type RepositoryInterface interface {
	Call(voice *model_voice.VoiceMessage) error
}

type Repository struct {
	ctx    context.Context
	client voice.VoiceClientInterface
}

func NewVoiceRepository(ctx context.Context, client voice.VoiceClientInterface) *Repository {
	return &Repository{
		ctx:    ctx,
		client: client,
	}
}

func (r *Repository) Call(voice *model_voice.VoiceMessage) error {
	if err := r.client.Call(voice); err != nil {
		return err
	}

	return nil
}
