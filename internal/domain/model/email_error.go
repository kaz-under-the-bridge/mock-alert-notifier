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

// ToDo: Emailのerror実装はあとで
type InvalidEmailAddressError struct {
	Email string
}

type InvalidEmailBodyError struct {
}

func (e *InvalidEmailAddressError) Error() string {
	return fmt.Sprintf("EmailAddress(%s)のフォーマットが不正です", e.Email)
}

func (e *InvalidEmailBodyError) Error() string {
	return "EmailBodyが指定されていません"
}

//func (e Email) verify() {
//	// veirfy body is not empty
//
//	// verify fromAddress is not empty
//
//	// verify toAddresses is not empty
//}
