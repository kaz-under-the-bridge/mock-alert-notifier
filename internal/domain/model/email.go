package model

import (
	"fmt"
	"time"
)

var ObjEmailErrors *EmailErrors

func init() {
	ObjEmailErrors = NewEmailErrors()
}

type Emails []*Email

type Email struct {
	Subject      string
	Body         string
	FromAddress  string
	ToAddresses  []string
	CcAddresses  []string
	BccAddresses []string
	SentAt       *time.Time
}

func NewEmail(body string, fromAddress string, toAddresses []string) (*Email, error) {
	email := Email{
		Body:        body,
		FromAddress: fromAddress,
		ToAddresses: toAddresses,
	}

	return &email, nil
}

// Emails has Push, Len, verify
func (es *Emails) Push(e *Email) {
	*es = append(*es, e)
}

func (es *Emails) Len() int {
	cnt := 0

	for range *es {
		cnt++
	}
	return cnt
}

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
