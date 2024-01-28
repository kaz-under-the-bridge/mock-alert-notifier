package email

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// テストしたいときは以下のようなenv.goを同じファイルに配置(そんなに動かさないので雑なファイル配置パターンで)
/*
package email

import "os"

func setenv() {
	os.Setenv("TEST_SENDGRID_API_KEY", "SG.xxx")
}
*/

func TestSendgridEmail(t *testing.T) {
	key := os.Getenv("TEST_SENDGRID_API_KEY")
	if key == "" {
		t.Skip("TEST_SENDGRID_API_KEY is not set")
	}
	from := mail.NewEmail("Kazuhiro Hashimoto", "kaz@under-the-bridge.work")
	to := mail.NewEmail("Kazuhiro Hashimoto", "kaz@under-the-bridge.work")

	dt := time.Now().Format("2006-01-02 15:04:05")
	subject := fmt.Sprintf("Test Email %s", dt)

	plainTextContent := fmt.Sprintf("this mail is test mail at %s", dt)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
