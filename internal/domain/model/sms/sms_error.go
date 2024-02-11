package sms

import "fmt"

type SMSErrors []error

func NewSMSErrors() *SMSErrors {
	return &SMSErrors{}
}

func GetSMSErrors() *SMSErrors {
	return ObjSMSErrors
}

func (ees *SMSErrors) Push(err error) {
	*ees = append(*ees, err)
}

func (ees SMSErrors) Len() int {
	cnt := 0

	for range ees {
		cnt++
	}
	return cnt
}

func (ees SMSErrors) Error() string {
	var errString string

	for _, err := range ees {
		errString += err.Error() + "\n"
	}
	return errString
}

type InvalidSMSAttributeError struct {
	From  string
	To    string
	Cause string
}

type InvalidSMSBodyLengthError struct {
	Body   string
	Length int
}

func (e *InvalidSMSAttributeError) Error() string {
	return fmt.Sprintf("from: %s, to: %s のSMSで正しくないフィールドが存在します: %s", e.From, e.To, e.Cause)
}

// bodyの長さが半角換算で1530文字以上の場合エラーにする
// Bodyの頭30文字を表示する
func (e *InvalidSMSBodyLengthError) Error() string {
	return fmt.Sprintf("SMSの本文が長すぎます(全角2文字, 半角1文字換算で1530文字以内必須): %s...(%d文字)", e.Body[:30], e.Length)
}
