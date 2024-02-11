package sms

import (
	"context"
	"os"
	"testing"

	model_sms "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/sms"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

// MockではなくてAPIじかたたきのテスト（実際にSMSが送信される）
// このテストは環境変数TEST_TWILIO_TO_PHONE_NUMBERが設定されている場合のみ実行
func TestSendSMSSendFunctionActually(t *testing.T) {
	ctx := context.Background()
	to := os.Getenv("TEST_TWILIO_TO_PHONENUMBER")

	if to == "" {
		t.Log("環境変数TEST_TWILIO_TO_PHONE_NUMBERが設定されていません. skipします.")
		t.Skip()
	}

	sid := os.Getenv("TEST_TWILIO_SID")
	token := os.Getenv("TEST_TWILIO_AUTH_TOKEN")

	ctx = helper.SetTwilioSid(ctx, sid)
	ctx = helper.SetTwilioAuthToken(ctx, token)

	sms := model_sms.NewSMS(
		"テストメッセージ本文",
		"+12019490946",
		to,
	)

	client := NewTwilioSMSClient(ctx)
	err := client.Send(sms)

	assert.NoError(t, err)
}
