package model

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var ObjEmailErrors *EmailErrors

func init() {
	ObjEmailErrors = NewEmailErrors()
}

type Emails []*Email

type Email struct {
	Subject     string
	Body        string
	FromAddress string
	ToAddresses []string
	SentAt      *time.Time
}

func NewEmail(body string, fromAddress string, toAddresses []string) (*Email, error) {
	email := Email{
		Body:        body,
		FromAddress: fromAddress,
		ToAddresses: toAddresses,
	}

	return &email, nil
}

// Emails has Push, Len, verify
func (es *Emails) Push(e *Email) {
	*es = append(*es, e)
}

func (es *Emails) Len() int {
	cnt := 0

	for range *es {
		cnt++
	}
	return cnt
}

type InvalidEmailAddressError struct {
	Email string
}

type InvalidEmailBodyError struct {
}

func (e *InvalidEmailAddressError) Error() string {
	return fmt.Sprintf("EmailAddress(%s)のフォーマットが不正です", e.Email)
}

func (e *InvalidEmailBodyError) Error() string {
	return "EmailBodyが指定されていません"
}

func (e Email) verify() {
	// veirfy body is not empty

	// verify fromAddress is not empty

	// verify toAddresses is not empty
}

type EmailTempaltes []*EmailTemplate

// New EmailTemplates from []byte of yaml
func NewEmailTemplates(yamlBytes []byte) (*EmailTempaltes, error) {
	var templates EmailTempaltes

	if err := yaml.Unmarshal(yamlBytes, &templates); err != nil {
		return nil, err
	}

	if templates.Len() == 0 {
		return nil, fmt.Errorf("EmailTemplateが空です")
	}

	if err := templates.verify(); err != nil {
		return nil, errors.Wrapf(err, "EmailTemplateの検証に失敗しました")
	}

	return &templates, nil
}

func (ets EmailTempaltes) FindByName(name string) (*EmailTemplate, bool) {
	for _, et := range ets {
		if et.Name == name {
			return et, true
		}
	}
	return nil, false
}

func (ets EmailTempaltes) Len() int {
	cnt := 0

	for range ets {
		cnt++
	}
	return cnt
}

func (ets EmailTempaltes) verify() error {
	for _, et := range ets {
		if err := et.verify(); err != nil {
			return err
		}
	}
	return nil
}

// subject,body might consist of template string
type EmailTemplate struct {
	Name    string `yaml:"name"`
	Subject string `yaml:"subject"`
	Body    string `yaml:"body"`
}

func (et *EmailTemplate) verify() error {
	if et.Subject == "" {
		return &InvalidEmailTemplateSubjectError{Name: et.Name, Cause: "空白です"}
	}

	if et.Body == "" {
		return &InvalidEmailTemplateBodyError{Name: et.Name, Cause: "空白です"}
	}

	return nil
}
