package sms

import "time"

var ObjSMSErrors *SMSErrors

func init() {
	ObjSMSErrors = NewSMSErrors()
}

type SMSMessages []*SMSMessage

type SMSMessage struct {
	Body   string
	From   string
	To     string
	SentAt *time.Time
}

func NewSMS(body string, from string, to string) *SMSMessage {
	sms := &SMSMessage{
		Body: body,
		From: from,
		To:   to,
	}

	sms.verify()

	return sms
}

// Emails has Push, Len, verify
func (msgs *SMSMessages) Push(msg *SMSMessage) {
	*msgs = append(*msgs, msg)
}

func (msgs *SMSMessages) Len() int {
	cnt := 0

	for range *msgs {
		cnt++
	}
	return cnt
}

func (msg *SMSMessage) verify() {
	if msg.From == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.From,
			To:    msg.To,
			Cause: "fromアドレスが空白です",
		})
	}

	if msg.To == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.From,
			To:    msg.To,
			Cause: "toアドレスが空白です",
		})
	}

	if msg.Body == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.From,
			To:    msg.To,
			Cause: "本文(body)が空白です",
		})
	}

	// ToDo: 国コード(+81など)を含む電話番号かどうかのチェックを入れる
	// 国コードは一旦は日本だけでいいので、+81と桁数チェックにする
}
