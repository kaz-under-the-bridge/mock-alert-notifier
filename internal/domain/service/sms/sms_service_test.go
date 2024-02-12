package sms

import (
	"context"
	"testing"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	repo_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/sms"
	infra_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/sms"
	"github.com/stretchr/testify/assert"
)

func TestSMSSend(t *testing.T) {
	ctx := context.Background()
	client := infra_sms.NewExportMockSMSClient()
	repo := repo_sms.NewSMSRepository(ctx, client)

	sms := model_sms.SMSMessage{}

	err := repo.Send(&sms)

	assert.NoError(t, err)
	assert.NotEmpty(t, sms.GetSentAtJSTFormatRFC3339())
}
