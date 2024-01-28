package sms

import (
	"context"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	repo_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/sms"
	infra_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/sms"
	"github.com/stretchr/testify/assert"
)

func TestSMSSend(t *testing.T) {
	ctx := context.Background()
	client := infra_sms.NewExportMockSMSClient()
	repo := repo_sms.NewSMSRepository(ctx, client)

	SMS := model.SMSMessage{}

	err := repo.Send(&SMS)

	assert.NoError(t, err)
	assert.NotNil(t, SMS.SentAt)
}
