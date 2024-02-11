package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMSVerificationWithError(t *testing.T) {
	cases := []struct {
		body        string
		from        string
		to          string
		errorString string
	}{
		{
			body: "",
			from: "あ",
			to:   "0901234567811111111", //15桁以上
			errorString: `from: あ, to: 0901234567811111111 のSMSで正しくないフィールドが存在します: fromアドレスが電話番号として正しくありません: the phone number supplied is not a number
from: あ, to: 0901234567811111111 のSMSで正しくないフィールドが存在します: toアドレスが電話番号として正しくありません: the string supplied is too long to be a phone number
from: あ, to: 0901234567811111111 のSMSで正しくないフィールドが存在します: 本文(body)が空白です
`,
		},
		{
			body:        long800wordNihongo,
			from:        "09012345678",
			to:          "09012345678",
			errorString: "SMSの本文が長すぎます(全角2文字, 半角1文字換算で1530文字以内必須): 恥の多い生涯を送って...(1546文字)\n",
		},
	}

	for _, tc := range cases {
		ObjSMSErrors = NewSMSErrors()
		_ = NewSMS(tc.body, tc.from, tc.to)

		//fmt.Println(ObjSMSErrors.Error())

		assert.NotEqual(t, 0, ObjSMSErrors.Len())
		assert.Equal(t, tc.errorString, ObjSMSErrors.Error())
	}
}

func TestVoiceVerificationWithOK(t *testing.T) {
	cases := []struct {
		body string
		from string
		to   string
	}{
		{
			body: "hello",
			from: "090-1234-5678",
			to:   "090-1234-5678",
		},
	}

	for _, tc := range cases {
		ObjSMSErrors = NewSMSErrors()
		_ = NewSMS(tc.body, tc.from, tc.to)

		assert.Equal(t, 0, ObjSMSErrors.Len())
		assert.Equal(t, "", ObjSMSErrors.Error())
	}
}

func TestCountBodyLen(t *testing.T) {
	// この場合、"。"と"々"が半角としてカウントされて誤差が出る(400文字の日本語=800の期待値に対して、783になる)
	msg := `吾輩は猫である。名前はまだ無い。どこで生れたかとんと見当がつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪な種族であったそうだ。この書生というのは時々我々を捕えて煮て食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌に載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始であろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶だ。その後猫にもだいぶ逢ったがこんな片輪には一度も出会わした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙を`

	cnt := countBodyMessageLength(msg)
	assert.Equal(t, 783, cnt)
}

var long800wordNihongo = `恥の多い生涯を送って来ました。自分には、人間の生活というものが、見当つかないのです。自分は東北の田舎に生れましたので、汽車をはじめて見たのは、よほど大きくなってからでした。自分は停車場のブリッジを、上って、降りて、そうしてそれが線路をまたぎ越えるために造られたものだという事には全然気づかず、ただそれは停車場の構内を外国の遊戯場みたいに、複雑に楽しく、ハイカラにするためにのみ、設備せられてあるものだとばかり思っていました。しかも、かなり永い間そう思っていたのです。ブリッジの上ったり降りたりは、自分にはむしろ、ずいぶん垢抜けのした遊戯で、それは鉄道のサーヴィスの中でも、最も気のきいたサーヴィスの一つだと思っていたのですが、のちにそれはただ旅客が線路をまたぎ越えるための頗る実利的な階段に過ぎないのを発見して、にわかに興が覚めました。また、自分は子供の頃、絵本で地下鉄道というものを見て、これもやはり、実利的な必要から案出せられたものではなく、地上の車に乗るよりは、地下の車に乗ったほうが風がわりで面白い遊びだから、とばかり思っていました。恥の多い生涯を送って来ました。自分には、人間の生活というものが、見当つかないのです。自分は東北の田舎に生れましたので、汽車をはじめて見たのは、よほど大きくなってからでした。自分は停車場のブリッジを、上って、降りて、そうしてそれが線路をまたぎ越えるために造られたものだという事には全然気づかず、ただそれは停車場の構内を外国の遊戯場みたいに、複雑に楽しく、ハイカラにするためにのみ、設備せられてあるものだとばかり思っていました。しかも、かなり永い間そう思っていたのです。ブリッジの上ったり降りたりは、自分にはむしろ、ずいぶん垢抜けのした遊戯で、それは鉄道のサーヴィスの中でも、最も気のきいたサーヴィスの一つだと思っていたのですが、のちにそれはただ旅客が線路をまた`
