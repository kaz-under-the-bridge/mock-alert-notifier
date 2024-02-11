package voice

import (
	"fmt"
	"os"
	"testing"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func TestTwilioCall(t *testing.T) {
	to := os.Getenv("PHONE_NUMBER")

	if to == "" {
		t.Log("PHONE_NUMBER is not set")
		t.Skip()
	}

	client := twilio.NewRestClient()

	params := &api.CreateCallParams{}
	//params.SetUrl("https://twilio-voice-data.s3.ap-northeast-1.amazonaws.com/speech_20240124061814724.mp3")
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")
	params.SetTo(to)
	params.SetFrom("+12019490946")

	resp, err := client.Api.CreateCall(params)
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
