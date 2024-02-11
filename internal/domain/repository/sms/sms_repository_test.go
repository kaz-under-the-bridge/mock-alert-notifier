package sms

import (
	"context"
	"testing"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	infra_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/sms"
	"github.com/stretchr/testify/assert"
)

func TestMockSMSSend(t *testing.T) {
	ctx := context.Background()

	mockclient := infra_sms.NewExportMockSMSClient()
	repo := NewSMSRepository(ctx, mockclient)

	sms := &model_sms.SMSMessage{}
	err := repo.Send(sms)

	assert.NoError(t, err)
	assert.NotEmpty(t, sms.GetSentAtJSTFormatRFC3339())
}
