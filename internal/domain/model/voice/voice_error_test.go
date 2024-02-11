package voice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoiceVerificationWithError(t *testing.T) {
	cases := []struct {
		voiceURL    string
		from        string
		to          string
		errorString string
	}{
		{
			voiceURL: "example.com",
			from:     "あ",
			to:       "0901234567811111111", //15桁以上
			errorString: `from: あ, to: 0901234567811111111 の電話発信定義で正しくないフィールドが存在します: voiceURLがURLの形式ではありません
from: あ, to: 0901234567811111111 の電話発信定義で正しくないフィールドが存在します: fromが電話番号として正しくありません
from: あ, to: 0901234567811111111 の電話発信定義で正しくないフィールドが存在します: toが電話番号として正しくありません
`,
		},
		{
			voiceURL: "",
			from:     "",
			to:       "",
			errorString: `from: , to:  の電話発信定義で正しくないフィールドが存在します: voiceURLが空です
from: , to:  の電話発信定義で正しくないフィールドが存在します: voiceURLがURLの形式ではありません
from: , to:  の電話発信定義で正しくないフィールドが存在します: fromが空です
from: , to:  の電話発信定義で正しくないフィールドが存在します: toが空です
`,
		},
	}

	for _, tc := range cases {
		ObjVoiceErrors = NewVoiceErrors()
		_ = NewVoice(tc.voiceURL, tc.from, tc.to)

		//fmt.Println(ObjVoiceErrors.Error())

		assert.NotEqual(t, 0, ObjVoiceErrors.Len())
		assert.Equal(t, tc.errorString, ObjVoiceErrors.Error())
	}
}

func TestVoiceVerificationWithOK(t *testing.T) {
	cases := []struct {
		voiceURL string
		from     string
		to       string
	}{
		{
			voiceURL: "https://example.com/params?param1=1&param2=2",
			from:     "090-1234-5678",
			to:       "090-1234-5678",
		},
		{
			voiceURL: "https://example.com",
			from:     "09012345678",
			to:       "09012345678",
		},
		{
			voiceURL: "http://example.com",
			from:     "+819012345678",
			to:       "+819012345678",
		},
	}

	for _, tc := range cases {
		ObjVoiceErrors = NewVoiceErrors()
		_ = NewVoice(tc.voiceURL, tc.from, tc.to)

		assert.Equal(t, 0, ObjVoiceErrors.Len())
		assert.Equal(t, "", ObjVoiceErrors.Error())
	}
}
