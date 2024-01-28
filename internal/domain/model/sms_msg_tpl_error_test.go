package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var yamlSMSMsgTpl string = `
- name: welcome
  body: |
    {{.Name}}-san, welcome!
    This message is test!

- name: bye
  body: |
    {{.Name}}-san
    This message is test!
    Bye-bye!
`

func TestNewSMSMsgTpl(t *testing.T) {
	tpls, err := NewSMSMessageTemplates([]byte(yamlSMSMsgTpl))
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := tpls.FindByName("welcome")
	assert.NoError(t, err)
	assert.Equal(t, "welcome", tpl.Name)
	assert.Equal(t, "{{.Name}}-san, welcome!\nThis message is test!\n", tpl.Body)

	tpl, err = tpls.FindByName("bye")
	assert.NoError(t, err)
	assert.Equal(t, "bye", tpl.Name)
	assert.Equal(t, "{{.Name}}-san\nThis message is test!\nBye-bye!\n", tpl.Body)
}

var yamlSMSMsgTplExpectedError string = `
- subject: no name!
  body: |
    this message is no name sample!

- name: bye
  subject: bye {{.Name}}-san!
  #bodyが空
  body: ""
`

func TestNewSMSMsgTplWithError(t *testing.T) {
	_, err := NewSMSMessageTemplates([]byte(yamlSMSMsgTplExpectedError))
	assert.NoError(t, err)

	errs := GetMessageTempalteErrors()
	want := `MessageTemplate(SMS)のnameフィールドが空白/未指定のものが存在します
MessageTemplate(SMS)のname(bye)のbodyが不正です: 空白です
`

	assert.Equal(t, want, errs.Error())
}

func TestSMSMsgTplRenderer(t *testing.T) {
	tpls, err := NewSMSMessageTemplates([]byte(yamlSMSMsgTpl))
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := tpls.FindByName("welcome")
	assert.NoError(t, err)

	want := "テスト太郎-san, welcome!\nThis message is test!\n"
	got, err := tpl.GetBody(map[string]string{"Name": "テスト太郎"})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
