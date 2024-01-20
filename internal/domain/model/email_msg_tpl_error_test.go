package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var yamlEmailMsgTpl string = `
- name: welcome
  subject: welcom {{.Name}}-san!
  body: |
    {{.Name}}-san, welcome!
    This message is test!

- name: bye
  subject: bye {{.Name}}-san!
  body: |
    {{.Name}}-san
    This message is test!
    Bye-bye!
`

func TestNewEmailMsgTpl(t *testing.T) {
	tpls, err := NewEmailMessageTemplates([]byte(yamlEmailMsgTpl))
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := tpls.FindByName("welcome")
	assert.NoError(t, err)
	assert.Equal(t, "welcome", tpl.Name)
	assert.Equal(t, "welcom {{.Name}}-san!", tpl.Subject)
	assert.Equal(t, "{{.Name}}-san, welcome!\nThis message is test!\n", tpl.Body)

	tpl, err = tpls.FindByName("bye")
	assert.NoError(t, err)
	assert.Equal(t, "bye", tpl.Name)
	assert.Equal(t, "bye {{.Name}}-san!", tpl.Subject)
	assert.Equal(t, "{{.Name}}-san\nThis message is test!\nBye-bye!\n", tpl.Body)
}

var yamlEmailMsgTplExpectedError string = `
- subject: no name!
  body: |
    this message is no name sample!

- name: welcome
  #subjectがない
  body: |
    {{.Name}}-san, welcome!
    This message is test!

- name: bye
  subject: bye {{.Name}}-san!
  #bodyが空
  body: ""
`

func TestNewEmailMsgTplWithError(t *testing.T) {
	_, err := NewEmailMessageTemplates([]byte(yamlEmailMsgTplExpectedError))
	assert.NoError(t, err)

	errs := GetMessageTempalteErrors()
	want := `MessageTemplate(Email)のnameフィールドが空白/未指定のものが存在します
MessageTemplate(Email)のname(welcome)のsubjectが不正です: 空白です
MessageTemplate(Email)のname(bye)のbodyが不正です: 空白です
`

	assert.Equal(t, want, errs.Error())
}

func TestEmailMsgTplRenderer(t *testing.T) {
	tpls, err := NewEmailMessageTemplates([]byte(yamlEmailMsgTpl))
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := tpls.FindByName("welcome")
	assert.NoError(t, err)

	want := "welcom テスト太郎-san!"
	got, err := tpl.GetSubject(map[string]string{"Name": "テスト太郎"})
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	want = "テスト太郎-san, welcome!\nThis message is test!\n"
	got, err = tpl.GetBody(map[string]string{"Name": "テスト太郎"})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
