package model

import "fmt"

type UserErrrors []error

func NewUserErrrors() *UserErrrors {
	return &UserErrrors{}
}

func GetUserErrors() *UserErrrors {
	return UserErrors
}

func (ues *UserErrrors) Push(err error) {
	*ues = append(*ues, err)
}

func (ues UserErrrors) Len() int {
	cnt := 0

	for range ues {
		cnt++
	}
	return cnt
}

// print all UserErrrors as string
func (ues UserErrrors) Error() string {
	var errString string

	for _, err := range ues {
		errString += err.Error() + "\n"
	}
	return errString
}

// IDはRepositoryでチェック（空白）しているので、ここではチェックしない
//type InvalidUserIdError struct {
//	id    int
//	cause string
//}

type InvalidUserFamilyNameError struct {
	id    int
	cause string
}

type InvalidUserGivenNameError struct {
	id    int
	cause string
}

type InvalidUserOrganizationError struct {
	id    int
	cause string
}

type InvalidUserEmailError struct {
	id       int
	original string
	cause    string
}

type InvalidUserPhoneNumberError struct {
	id       int
	original string
	cause    string
}

func (e *InvalidUserFamilyNameError) Error() string {
	return fmt.Sprintf("ID(%d)のFamilyNameが不正です: %s", e.id, e.cause)
}

func (e *InvalidUserGivenNameError) Error() string {
	return fmt.Sprintf("ID(%d)のGivenName が不正です: %s", e.id, e.cause)
}

func (e *InvalidUserOrganizationError) Error() string {
	return fmt.Sprintf("ID(%d)のOrganizationが不正です: %s", e.id, e.cause)
}

func (e *InvalidUserEmailError) Error() string {
	return fmt.Sprintf("ID(%d)のEmail(%s)が不正です: %s", e.id, e.original, e.cause)
}

func (e *InvalidUserPhoneNumberError) Error() string {
	return fmt.Sprintf("ID(%d)のPhoneNumber(%s)が不正です: %s", e.id, e.original, e.cause)
}
