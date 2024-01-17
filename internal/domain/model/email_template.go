package model

import (
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type EmailTemplates []*EmailTemplate

func NewEmailTemplates(yamlBytes []byte) (*EmailTemplates, error) {
	var templates EmailTemplates

	if err := yaml.Unmarshal(yamlBytes, &templates); err != nil {
		return nil, errors.Wrapf(err, "EmailTemplateのUnmarshalに失敗しました")
	}

	if templates.Len() == 0 {
		return nil, fmt.Errorf("EmailTemplateが空です")
	}

	if err := templates.verify(); err != nil {
		return nil, errors.Wrapf(err, "EmailTemplateの検証に失敗しました")
	}

	return &templates, nil
}

func (ets EmailTemplates) FindByName(name string) (*EmailTemplate, bool) {
	for _, et := range ets {
		if et.Name == name {
			return et, true
		}
	}
	return nil, false
}

func (ets EmailTemplates) Len() int {
	cnt := 0

	for range ets {
		cnt++
	}
	return cnt
}

func (ets EmailTemplates) verify() error {
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
