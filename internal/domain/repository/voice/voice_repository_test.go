package voice

import (
	"context"
	"testing"

	model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"
	infra_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/voice"
	"github.com/stretchr/testify/assert"
)

func TestMockVoiceSend(t *testing.T) {
	ctx := context.Background()

	mockClient := infra_voice.NewExportMockVoiceClient()
	repo := NewVoiceRepository(ctx, mockClient)

	voice := &model_voice.VoiceMessage{}
	err := repo.Call(voice)

	assert.NoError(t, err)
	assert.NotEmpty(t, voice.GetSentAtJSTFormatRFC3339())
}
