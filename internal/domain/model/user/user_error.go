package user

import "fmt"

type UserErrors []error

func NewUserErrors() *UserErrors {
	return &UserErrors{}
}

func GetUserErrors() *UserErrors {
	return ObjUserErrors
}

func (ues *UserErrors) Push(err error) {
	*ues = append(*ues, err)
}

func (ues UserErrors) Len() int {
	cnt := 0

	for range ues {
		cnt++
	}
	return cnt
}

// print all UserErrrors as string
func (ues UserErrors) Error() string {
	var errString string

	for _, err := range ues {
		errString += err.Error() + "\n"
	}
	return errString
}

// IDはRepositoryでチェック（空白）しているので、ここではチェックしない
//type InvalidUserIdError struct {
//	ID   int
//	Cause string
//}

type InvalidUserFamilyNameError struct {
	ID    int
	Cause string
}

type InvalidUserGivenNameError struct {
	ID    int
	Cause string
}

type InvalidUserOrganizationError struct {
	ID    int
	Cause string
}

type InvalidUserEmailError struct {
	ID       int
	Original string
	Cause    string
}

type InvalidUserPhoneNumberError struct {
	ID       int
	original string
	Cause    string
}

func (e *InvalidUserFamilyNameError) Error() string {
	return fmt.Sprintf("ID(%d)のFamilyNameが不正です: %s", e.ID, e.Cause)
}

func (e *InvalidUserGivenNameError) Error() string {
	return fmt.Sprintf("ID(%d)のGivenName が不正です: %s", e.ID, e.Cause)
}

func (e *InvalidUserOrganizationError) Error() string {
	return fmt.Sprintf("ID(%d)のOrganizationが不正です: %s", e.ID, e.Cause)
}

func (e *InvalidUserEmailError) Error() string {
	return fmt.Sprintf("ID(%d)のEmail(%s)が不正です: %s", e.ID, e.Original, e.Cause)
}

func (e *InvalidUserPhoneNumberError) Error() string {
	return fmt.Sprintf("ID(%d)のPhoneNumber(%s)が不正です: %s", e.ID, e.original, e.Cause)
}

type UserNotFoundError struct {
	ID    int
	Email string
}

func (e *UserNotFoundError) Error() string {
	switch {
	case e.ID != 0 && e.Email != "":
		return fmt.Sprintf("ID(%d) or Email(%s) is not found", e.ID, e.Email)
	case e.ID != 0 && e.Email == "":
		return fmt.Sprintf("ID(%d) is not found", e.ID)
	case e.ID == 0 && e.Email != "":
		return fmt.Sprintf("Email(%s) is not found", e.Email)
	default:
		return "User is not found, any identifier is not specified"
	}
}
