package voice

import (
	"context"
	"encoding/json"
	"log/slog"

	model_voice "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/voice"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/pkg/errors"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type VoiceClientInterface interface {
	Call(voice *model_voice.VoiceMessage) error
}

type TwilioVoiceClient struct {
	ctx context.Context
	*VoiceClient
}

func NewTwilioVoiceClient(ctx context.Context) VoiceClientInterface {
	return &TwilioVoiceClient{
		ctx:         ctx,
		VoiceClient: &VoiceClient{},
	}
}

func (c *TwilioVoiceClient) Call(voice *model_voice.VoiceMessage) error {
	if err := callByTwilioAPI(voice, c.ctx); err != nil {
		return err
	}
	voice.UpdateSentAt()
	return nil
}

func callByTwilioAPI(voice *model_voice.VoiceMessage, ctx context.Context) error {
	sid := helper.GetTwilioSid(ctx)
	token := helper.GetTwilioAuthToken(ctx)

	if sid == "" || token == "" {
		return errors.New("twilio sid or token is not set")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: sid,
		Password: token,
	})

	params := &api.CreateCallParams{}
	params.SetTo(voice.GetToNumberE164())
	params.SetFrom(voice.GetFromNumberE164())
	params.SetUrl(voice.GetVoiceURL())

	resp, err := client.Api.CreateCall(params)

	response, _ := json.Marshal(*resp)
	slogfrom := slog.String("from", voice.GetFromNumberE164())
	slogto := slog.String("to", voice.GetToNumberE164())

	if err != nil {
		helper.Logger.Error("Voice call failed", slogfrom, slogto, slog.String("api_response", string(response)))
		return errors.Wrap(err, "twilio voice message creation api error")
	}

	helper.Logger.Info("Voice call successfully", slogfrom, slogto)

	return nil
}
