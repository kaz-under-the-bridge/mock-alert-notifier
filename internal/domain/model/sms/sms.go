package sms

import (
	"fmt"
	"time"

	"github.com/nyaruka/phonenumbers"
)

// このmodule("github.com/nyaruka/phonenumbers")は勝手forkされたものではなく
// ↓のように本家が3rd party portsとして紹介している
// https://github.com/google/libphonenumber?tab=readme-ov-file#third-party-ports

var ObjSMSErrors *SMSErrors

func init() {
	ObjSMSErrors = NewSMSErrors()
}

type SMSMessages []*SMSMessage

type SMSMessage struct {
	// メッセージ本文
	body string
	// 03-1234-5678などの電話番号形式
	// ある程度フォーマットが揃っていなくてもphonenumbersでパースしてくれるので気にしなくていいかも
	strFrom string
	strTo   string
	// E.164形式の電話番号を使ったりするのでObject化された電話番号にする
	from *phonenumbers.PhoneNumber
	to   *phonenumbers.PhoneNumber
	// 送信した時間（ログなどに使う予定）
	sentAt *time.Time
}

func NewSMS(body string, from string, to string) *SMSMessage {

	sms := &SMSMessage{
		body:    body,
		strFrom: from,
		strTo:   to,
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
	if msg.strFrom == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.strFrom,
			To:    msg.strTo,
			Cause: "発信元電話番号が空白です",
		})
	} else {
		from, err := phonenumbers.Parse(msg.strFrom, "JP")
		if err != nil {
			ObjSMSErrors.Push(&InvalidSMSAttributeError{
				From:  msg.strFrom,
				To:    msg.strTo,
				Cause: fmt.Sprintf("fromアドレスが電話番号として正しくありません: %s", err.Error()),
			})
		}
		msg.from = from
	}

	if msg.strTo == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.strFrom,
			To:    msg.strTo,
			Cause: "発信先電話番号が空白です",
		})
	} else {
		to, err := phonenumbers.Parse(msg.strTo, "JP")
		if err != nil {
			ObjSMSErrors.Push(&InvalidSMSAttributeError{
				From:  msg.strFrom,
				To:    msg.strTo,
				Cause: fmt.Sprintf("toアドレスが電話番号として正しくありません: %s", err.Error()),
			})
		}
		msg.to = to
	}

	if msg.body == "" {
		ObjSMSErrors.Push(&InvalidSMSAttributeError{
			From:  msg.strFrom,
			To:    msg.strTo,
			Cause: "本文(body)が空白です",
		})
	}

	bodyLen := countBodyMessageLength(msg.body)
	if bodyLen > 1530 {
		ObjSMSErrors.Push(&InvalidSMSBodyLengthError{
			Body:   msg.body,
			Length: bodyLen,
		})
	}

}

func (msg SMSMessage) GetBody() string {
	return msg.body
}

func (msg SMSMessage) GetFromNumberE164() string {
	return phonenumbers.Format(msg.from, phonenumbers.E164)
}

func (msg SMSMessage) GetToNumberE164() string {
	return phonenumbers.Format(msg.to, phonenumbers.E164)
}

func (msg *SMSMessage) UpdateSentAt() {
	now := time.Now()
	msg.sentAt = &now
}

func (msg SMSMessage) GetSentAtJSTFormatRFC3339() string {
	timeAtJST := msg.sentAt.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	return timeAtJST.Format(time.RFC3339)
}

func countBodyMessageLength(body string) int {
	var cnt int
	var hankakuCnt int
	var zenkakuCnt int

	// 以下のコード区分では以下が半角としてカウントされる
	// 個別に細かく指定してもいいが、あまり多用される文字列ではなく誤差が小さいのでざっくり算出とする
	/*
			句読点
		。（句点）
		、（読点）
		．（中点）
		？
		！
		記号
		々（重ね点）
		〜（波ダッシュ）
		‖（縦線）
		¥（円記号）
		§（節記号）
		¶（段落記号）
		その他
		ー（長音符） *空格（半角スペース）
	*/
	for _, c := range body {
		if c >= '\u3040' && c <= '\u309F' ||
			c >= '\u30A0' && c <= '\u30FF' ||
			c >= '\u4E00' && c <= '\u9FFF' {
			zenkakuCnt++
		} else {
			//fmt.Printf("%c\n", c)
			hankakuCnt++
		}
	}

	cnt = zenkakuCnt*2 + hankakuCnt

	return cnt
}
