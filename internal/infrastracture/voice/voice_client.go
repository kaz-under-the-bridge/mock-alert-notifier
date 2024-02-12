package voice

import model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"

type VoiceClient struct{}

type MockVoiceClient struct {
	*VoiceClient
}

func NewExportMockVoiceClient() VoiceClientInterface {
	return &MockVoiceClient{
		VoiceClient: &VoiceClient{},
	}
}

func (c *MockVoiceClient) Call(voice *model_voice.VoiceMessage) error {
	voice.UpdateSentAt()
	return nil
}
