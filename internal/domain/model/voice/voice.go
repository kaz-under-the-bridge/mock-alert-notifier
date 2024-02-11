package voice

import (
	"regexp"
	"time"

	"github.com/nyaruka/phonenumbers"
)

// regexpでURLの形式をチェックする
var regexVoiceURL = regexp.MustCompile(`^https?://.*$`)

var ObjVoiceErrors *VoiceErrors

func init() {
	ObjVoiceErrors = NewVoiceErrors()
}

type VoiceMessages []*VoiceMessage

type VoiceMessage struct {
	// 音声ファイルのURL
	voiceURL string
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

func NewVoice(voiceURL string, from string, to string) *VoiceMessage {

	voice := &VoiceMessage{
		voiceURL: voiceURL,
		strFrom:  from,
		strTo:    to,
	}

	voice.verify()

	return voice
}

func (msgs *VoiceMessages) Push(msg *VoiceMessage) {
	*msgs = append(*msgs, msg)
}

func (msgs *VoiceMessages) Len() int {
	cnt := 0

	for range *msgs {
		cnt++
	}
	return cnt
}

func (voice *VoiceMessage) verify() {
	if voice.voiceURL == "" {
		ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
			From:  voice.strFrom,
			To:    voice.strTo,
			Cause: "voiceURLが空です",
		})
	}

	// voiceURLがhttps://xxxxという形式であるかをチェック
	if !regexVoiceURL.MatchString(voice.voiceURL) {
		ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
			From:  voice.strFrom,
			To:    voice.strTo,
			Cause: "voiceURLがURLの形式ではありません",
		})
	}

	if voice.strFrom == "" {
		ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
			From:  voice.strFrom,
			To:    voice.strTo,
			Cause: "fromが空です",
		})
	} else {
		from, err := phonenumbers.Parse(voice.strFrom, "JP")
		if err != nil {
			ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
				From:  voice.strFrom,
				To:    voice.strTo,
				Cause: "fromが電話番号として正しくありません",
			})
		}
		voice.from = from
	}

	if voice.strTo == "" {
		ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
			From:  voice.strFrom,
			To:    voice.strTo,
			Cause: "toが空です",
		})
	} else {
		to, err := phonenumbers.Parse(voice.strTo, "JP")
		if err != nil {
			ObjVoiceErrors.Push(&InvalidVoiceAttributeError{
				From:  voice.strFrom,
				To:    voice.strTo,
				Cause: "toが電話番号として正しくありません",
			})
		}
		voice.to = to
	}
}

func (voice *VoiceMessage) GetVoiceURL() string {
	return voice.voiceURL
}

func (voice *VoiceMessage) GetFromNumberE164() string {
	return voice.strFrom
}

func (voice *VoiceMessage) GetToNumberE164() string {
	return voice.strTo
}

func (voice *VoiceMessage) UpdateSentAt() {
	now := time.Now()
	voice.sentAt = &now
}

func (voice *VoiceMessage) GetSentAtJSTFormatRFC3339() string {
	timeAtJST := voice.sentAt.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	return timeAtJST.Format(time.RFC3339)
}
