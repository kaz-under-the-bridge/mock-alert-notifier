package sms

import (
	"fmt"
	"os"
	"testing"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func TestSendSMS(t *testing.T) {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	client := twilio.NewRestClient()
	to := os.Getenv("PHONE_NUMBER")

	params := &api.CreateMessageParams{}
	params.SetBody("テストメッセージです by Twilio API")
	if to == "" {
		t.Log("環境変数PHONE_NUMBERが設定されていません. skipします.")
		t.Skip()
	}
	params.SetFrom("+12019490946")
	params.SetTo(to)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}
