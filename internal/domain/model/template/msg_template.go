package message_tpl

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type TemplateType int

const (
	TemplateTypeEmail TemplateType = iota
	TemplateTypeSMS
)

func (tt TemplateType) String() string {
	switch tt {
	case TemplateTypeEmail:
		return "Email"
	case TemplateTypeSMS:
		return "SMS"
	default:
		panic("invalid TemplateType")
	}
}

type MessageTemplates []*MessageTemplate

// unmarshal yaml from data(bytes)
func NewMessageTemplates(data []byte) (*MessageTemplates, error) {
	yamlTpl := MessageTemplates{}

	if err := yaml.Unmarshal(data, &yamlTpl); err != nil {
		return nil, err
	}

	return &yamlTpl, nil
}

type MessageTemplate struct {
	Name    string `yaml:"name"`
	Subject string `yaml:"subject"`
	Body    string `yaml:"body"`
}

func (mt MessageTemplate) verify(check TemplateType) {
	switch check {

	case TemplateTypeEmail:
		if mt.Name == "" {
			ObjMessageTemplateErrors.Push(&InvalidMessageTemplateNameError{Type: check.String()})
		}
		if mt.Subject == "" {
			ObjMessageTemplateErrors.Push(&InvalidMessageTemplateSubjectError{Type: check.String(), Name: mt.Name, Cause: "空白です"})
		}
		if mt.Body == "" {
			ObjMessageTemplateErrors.Push(&InvalidMessageTemplateBodyError{Type: check.String(), Name: mt.Name, Cause: "空白です"})
		}

	case TemplateTypeSMS:
		if mt.Name == "" {
			ObjMessageTemplateErrors.Push(&InvalidMessageTemplateNameError{Type: check.String()})
		}
		// SMS TemplateはSubjectがない
		if mt.Body == "" {
			ObjMessageTemplateErrors.Push(&InvalidMessageTemplateBodyError{Type: check.String(), Name: mt.Name, Cause: "空白です"})
		}

	default:
		panic("invalid TemplateType")
	}
}

func (mt *MessageTemplate) GetSubject(data map[string]string) (string, error) {
	t := template.Must(template.New("subject").Parse(mt.Subject))

	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, data); err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}

func (mt *MessageTemplate) GetBody(data map[string]string) (string, error) {
	t := template.Must(template.New("body").Parse(mt.Body))

	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, data); err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}
