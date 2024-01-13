package model

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
