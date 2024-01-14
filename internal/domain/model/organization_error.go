package model

import "fmt"

type OrganizationErrors []error

func NewOrganizationErrors() *OrganizationErrors {
	return &OrganizationErrors{}
}

func GetOrganizationErrors() *OrganizationErrors {
	return ObjOrganizationErrors
}

func (oes *OrganizationErrors) Push(err error) {
	*oes = append(*oes, err)
}

func (oes OrganizationErrors) Len() int {
	cnt := 0

	for range oes {
		cnt++
	}
	return cnt
}

func (oes OrganizationErrors) Error() string {
	var errString string

	for _, err := range oes {
		errString += err.Error() + "\n"
	}
	return errString
}

// IDはRepositoryでチェック（空白）しているので、ここではチェックしない
//type InvalidOrganizationIdError struct {
//	ID   int
//	Cause string
//}

type InvalidOrganizationNameError struct {
	ID    int
	Cause string
}

type InvalidOrganizationTeamError struct {
	ID    int
	Cause string
}

type InvalidOrganizationEmailError struct {
	ID       int
	Original string
	Cause    string
}

type InvalidOrganizationPhoneNumberError struct {
	ID       int
	Original string
	Cause    string
}

func (e *InvalidOrganizationNameError) Error() string {
	return fmt.Sprintf("ID(%d)のOrg名が不正です: %s", e.ID, e.Cause)
}

func (e *InvalidOrganizationTeamError) Error() string {
	return fmt.Sprintf("ID(%d)のTeam名が不正です: %s", e.ID, e.Cause)
}

func (e *InvalidOrganizationEmailError) Error() string {
	return fmt.Sprintf("ID(%d)のEmail(%s)が不正です: %s", e.ID, e.Original, e.Cause)
}

func (e *InvalidOrganizationPhoneNumberError) Error() string {
	return fmt.Sprintf("ID(%d)のPhoneNumber(%s)が不正です: %s", e.ID, e.Original, e.Cause)
}

type OrganizationNotFoundError struct {
	ID int
}

func (e *OrganizationNotFoundError) Error() string {
	return fmt.Sprintf("ID(%d)のOrganizationが見つかりません", e.ID)
}
