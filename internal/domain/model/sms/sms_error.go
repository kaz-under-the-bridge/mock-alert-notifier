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

func (e *InvalidSMSAttributeError) Error() string {
	return fmt.Sprintf("from: %s, to: %s のSMSで正しくないフィールドが存在します: %s", e.From, e.To, e.Cause)
}
