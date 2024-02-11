package usecase

import (
	"fmt"

	model_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"
	model_tpl "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/template"
	service_email "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/email"
)

// ToDo: あとで実装
// 対象の組織にメールを送信する
func SendEmailToOrgs(
	orgs model_org.Organizations,
	tpl *model_tpl.EmailMessageTemplate,
	service service_email.ServiceInterface,
) {
	for _, org := range orgs {
		// tplにorg.ToMapを渡してtemplateのredneringを行う
		fmt.Println(org.ToMap())

		// tplとorgからemailを作成する

		// emailを送信する
	}
}
