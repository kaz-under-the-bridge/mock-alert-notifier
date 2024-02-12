package voice

import (
	"context"
	"testing"

	model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"
	repo_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/voice"
	infra_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/voice"
	"github.com/stretchr/testify/assert"
)

func TestVoiceCall(t *testing.T) {
	ctx := context.Background()

	mockClient := infra_voice.NewExportMockVoiceClient()
	repo := repo_voice.NewVoiceRepository(ctx, mockClient)
	service := NewVoiceService(ctx, repo)

	voice := &model_voice.VoiceMessage{}
	err := service.Send(voice)

	assert.NoError(t, err)
	assert.NotEmpty(t, voice.GetSentAtJSTFormatRFC3339())
}
