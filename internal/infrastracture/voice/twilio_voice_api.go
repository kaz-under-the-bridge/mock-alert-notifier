package voice

import (
	"os"

	"github.com/twilio/twilio-go"
)

func NewTwilioVoiceClient(accountSid, authToken string) *twilio.RestClient {
	os.Setenv("TWILIO_ACCOUNT_SID", accountSid)
	os.Setenv("TWILIO_AUTH_TOKEN", authToken)
	return twilio.NewRestClient()
}

/*
	params := &twilio_api.CreateCallParams{}
	params.SetMethod("GET")
	params.SetSendDigits("1234#")
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")
	params.SetTo("+14155551212")
	params.SetFrom("+18668675310")

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
*/

/*
package main

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

ステータスコールバックを用いたコードサンプル
func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	toPhoneNumber := "送信先の電話番号"
	statusCallbackUrl := "ステータスコールバックURL"

	client := twilio.NewRestClient(accountSid, authToken)

	params := &openapi.CreateCallParams{}
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")
	params.SetStatusCallback(statusCallbackUrl)
	params.SetStatusCallbackEvent([]string{"initiated", "ringing", "answered", "completed"})

	resp, err := client.ApiV2010.CreateCall(fromPhoneNumber, toPhoneNumber, params)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp.Sid)
}

*/
