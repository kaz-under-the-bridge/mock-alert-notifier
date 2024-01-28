package email

import (
	"context"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	repo_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/email"
	infra_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/email"
	"github.com/stretchr/testify/assert"
)

func TestEmailSend(t *testing.T) {
	ctx := context.Background()
	client := infra_email.NewExportMockEmailClient()
	repo := repo_email.NewEmailRepository(ctx, client)

	email := model.Email{}

	err := repo.Send(&email)

	assert.NoError(t, err)
	assert.NotNil(t, email.SentAt)
}
