package model

import "github.com/pkg/errors"

// MessageTemplate is embedded in EmailMessageTemplate

type EmailMessageTemplates []*EmailMessageTemplate

// Push, Len, Verify, FindByName
func (emts *EmailMessageTemplates) Push(emt *EmailMessageTemplate) {
	*emts = append(*emts, emt)
}

func (emts *EmailMessageTemplates) Len() int {
	cnt := 0

	for range *emts {
		cnt++
	}
	return cnt
}

func (emts *EmailMessageTemplates) verify() error {
	NewMessageTemplateErrors()

	for _, emt := range *emts {
		emt.verify(TemplateTypeEmail)
	}

	return nil
}

func (emts *EmailMessageTemplates) FindByName(name string) (*EmailMessageTemplate, error) {
	for _, emt := range *emts {
		if emt.Name == name {
			return emt, nil
		}
	}

	return nil, &MessageTemplateNotFoundError{Type: TemplateTypeEmail.String(), Name: name}
}

type EmailMessageTemplate struct {
	MessageTemplate
}

func NewEmailMessageTemplates(data []byte) (*EmailMessageTemplates, error) {
	emailTpls := EmailMessageTemplates{}

	tpls, err := NewMessageTemplates(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewMessageTemplates at NewEmailMessageTemplates")
	}

	for _, tpl := range *tpls {
		emailTpls = append(emailTpls, &EmailMessageTemplate{*tpl})
	}

	emailTpls.verify()

	return &emailTpls, nil
}
