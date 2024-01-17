package model

import "fmt"

type EmailErrors []error

func NewEmailErrors() *EmailErrors {
	return &EmailErrors{}
}

func GetEmailErrors() *EmailErrors {
	return ObjEmailErrors
}

func (ees *EmailErrors) Push(err error) {
	*ees = append(*ees, err)
}

func (ees EmailErrors) Len() int {
	cnt := 0

	for range ees {
		cnt++
	}
	return cnt
}

func (ees EmailErrors) Error() string {
	var errString string

	for _, err := range ees {
		errString += err.Error() + "\n"
	}
	return errString
}

// Name fieldが空の場合、NotFoundErrorが別のロジックで返すことになるので実装しない
//type InvalidEmailTemplateNameError struct {
//}

type InvalidEmailTemplateSubjectError struct {
	Name  string
	Cause string
}

type InvalidEmailTemplateBodyError struct {
	Name  string
	Cause string
}

func (e *InvalidEmailTemplateSubjectError) Error() string {
	return fmt.Sprintf("name(%s)のsubject が不正です: %s", e.Name, e.Cause)
}

func (e *InvalidEmailTemplateBodyError) Error() string {
	return fmt.Sprintf("name(%s)のbody が不正です: %s", e.Name, e.Cause)
}

type EmailTemplateNotFoundError struct {
	Name string
}

func (e *EmailTemplateNotFoundError) Error() string {
	return fmt.Sprintf("name(%s)のEmailTemplateが見つかりません", e.Name)
}
