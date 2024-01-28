package model

import "github.com/pkg/errors"

// MessageTemplate is embedded in SMSMessageTemplate

type SMSMessageTemplates []*SMSMessageTemplate

// Push, Len, Verify, FindByName
func (emts *SMSMessageTemplates) Push(emt *SMSMessageTemplate) {
	*emts = append(*emts, emt)
}

func (emts *SMSMessageTemplates) Len() int {
	cnt := 0

	for range *emts {
		cnt++
	}
	return cnt
}

func (emts *SMSMessageTemplates) verify() error {
	NewMessageTemplateErrors()

	for _, emt := range *emts {
		emt.verify(TemplateTypeSMS)
	}

	return nil
}

func (emts *SMSMessageTemplates) FindByName(name string) (*SMSMessageTemplate, error) {
	for _, emt := range *emts {
		if emt.Name == name {
			return emt, nil
		}
	}

	return nil, &MessageTemplateNotFoundError{Type: TemplateTypeSMS.String(), Name: name}
}

type SMSMessageTemplate struct {
	MessageTemplate
}

func NewSMSMessageTemplates(data []byte) (*SMSMessageTemplates, error) {
	SMSTpls := SMSMessageTemplates{}

	tpls, err := NewMessageTemplates(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewMessageTemplates at NewSMSMessageTemplates")
	}

	for _, tpl := range *tpls {
		SMSTpls = append(SMSTpls, &SMSMessageTemplate{*tpl})
	}

	SMSTpls.verify()

	return &SMSTpls, nil
}
