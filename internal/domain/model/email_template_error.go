package model

import "fmt"

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