package message_tpl

import "fmt"

var ObjMessageTemplateErrors *MessageTemplateErrors

// error
type MessageTemplateErrors []error

func NewMessageTemplateErrors() {
	ObjMessageTemplateErrors = &MessageTemplateErrors{}
}

func GetMessageTempalteErrors() *MessageTemplateErrors {
	return ObjMessageTemplateErrors
}

// MessageTemplateErrors has functions Push, Len, Error
func (mtes *MessageTemplateErrors) Push(err error) {
	*mtes = append(*mtes, err)
}

func (mtes MessageTemplateErrors) Len() int {
	cnt := 0

	for range mtes {
		cnt++
	}
	return cnt
}

func (mtes MessageTemplateErrors) Error() string {
	var errString string

	for _, err := range mtes {
		errString += err.Error() + "\n"
	}
	return errString
}

type InvalidMessageTemplateNameError struct {
	Type string
}

type InvalidMessageTemplateSubjectError struct {
	Type  string
	Name  string
	Cause string
}

type InvalidMessageTemplateBodyError struct {
	Type  string
	Name  string
	Cause string
}

// Error functio for above error
func (e *InvalidMessageTemplateNameError) Error() string {
	return fmt.Sprintf("MessageTemplate(%s)のnameフィールドが空白/未指定のものが存在します", e.Type)
}

func (e *InvalidMessageTemplateSubjectError) Error() string {
	return fmt.Sprintf("MessageTemplate(%s)のname(%s)のsubjectが不正です: %s", e.Type, e.Name, e.Cause)
}

func (e *InvalidMessageTemplateBodyError) Error() string {
	return fmt.Sprintf("MessageTemplate(%s)のname(%s)のbodyが不正です: %s", e.Type, e.Name, e.Cause)
}

// not found error
type MessageTemplateNotFoundError struct {
	Type string
	Name string
}

func (e *MessageTemplateNotFoundError) Error() string {
	return fmt.Sprintf("MessageTemplate(%s)のname(%s)が見つかりません", e.Type, e.Name)
}
