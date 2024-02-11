package voice

import "fmt"

type VoiceErrors []error

func NewVoiceErrors() *VoiceErrors {
	return &VoiceErrors{}
}

func GetVoiceErrors() *VoiceErrors {
	return ObjVoiceErrors
}

func (ves *VoiceErrors) Push(err error) {
	*ves = append(*ves, err)
}

func (ves VoiceErrors) Len() int {
	cnt := 0

	for range ves {
		cnt++
	}
	return cnt
}

func (ves VoiceErrors) Error() string {
	var errString string

	for _, err := range ves {
		errString += err.Error() + "\n"
	}
	return errString
}

type InvalidVoiceAttributeError struct {
	From  string
	To    string
	Cause string
}

func (e *InvalidVoiceAttributeError) Error() string {
	return fmt.Sprintf("from: %s, to: %s の電話発信定義で正しくないフィールドが存在します: %s", e.From, e.To, e.Cause)
}
