package org

import "regexp"

// declare regex matcher for email
// - @マークの前は英数字, ハイフン, ドット(.)を許可
// - ドメインは英小文字、数字、ハイフンを許可(トップレベルドメインは英小文字のみ)
var regexEmail = regexp.MustCompile(`^[\w+\-.]+@[a-z\d\-.]+\.[a-z]+$`)

// declare regex matcher for phone number
// - xxxx-xxxx-xxxx(4桁, ハイフンあり)
var regexPhoneNumber = regexp.MustCompile(`^\d{1,4}-\d{1,4}-\d{4}$`)

var ObjOrganizationErrors *OrganizationErrors

func init() {
	ObjOrganizationErrors = NewOrganizationErrors()
}

type Organizations []*Organization

type Organization struct {
	ID          int
	Name        string
	Team        string
	Email       string
	PhoneNumber string
}

func NewOrganization(
	id int,
	name, team, email, phoneNumber string,
) *Organization {
	o := &Organization{
		ID:          id,
		Name:        name,
		Team:        team,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
	o.verify()

	return o
}

func (os *Organizations) ToMap() map[int]*Organization {
	m := make(map[int]*Organization)
	for _, o := range *os {
		m[o.ID] = o
	}
	return m
}

func (os *Organizations) ToSlice() []*Organization {
	slice := make([]*Organization, 0, len(*os))

	for _, o := range *os {
		slice = append(slice, o)
	}
	return slice
}

func (os *Organizations) ToSliceWithID() []int {
	slice := make([]int, 0, len(*os))

	for _, o := range *os {
		slice = append(slice, o.ID)
	}
	return slice
}

func (os *Organizations) Push(o *Organization) {
	*os = append(*os, o)
}

func (os *Organizations) Len() int {
	cnt := 0

	for range *os {
		cnt++
	}
	return cnt
}

func (os *Organizations) FindByID(id int) (*Organization, bool) {
	for _, o := range *os {
		if o.ID == id {
			return o, true
		}
	}
	return nil, false
}

func (os *Organizations) FindByName(name string) (*Organization, bool) {
	for _, o := range *os {
		if o.Name == name {
			return o, true
		}
	}
	return nil, false
}

func (os *Organizations) FindByTeam(team string) (*Organization, bool) {
	for _, o := range *os {
		if o.Team == team {
			return o, true
		}
	}
	return nil, false
}

func (os *Organizations) FindByEmail(email string) (*Organization, bool) {
	for _, o := range *os {
		if o.Email == email {
			return o, true
		}
	}
	return nil, false
}

func (os *Organizations) FindByPhoneNumber(phoneNumber string) (*Organization, bool) {
	for _, o := range *os {
		if o.PhoneNumber == phoneNumber {
			return o, true
		}
	}
	return nil, false
}

func (o *Organization) verify() {
	if o.Name == "" {
		ObjOrganizationErrors.Push(&InvalidOrganizationNameError{ID: o.ID, Cause: "空白です"})
	}

	if o.Team == "" {
		ObjOrganizationErrors.Push(&InvalidOrganizationTeamError{ID: o.ID, Cause: "空白です"})
	}

	if !regexEmail.MatchString(o.Email) {
		ObjOrganizationErrors.Push(&InvalidOrganizationEmailError{ID: o.ID, Original: o.Email, Cause: "フォーマットが不正です"})
	}

	if !regexPhoneNumber.MatchString(o.PhoneNumber) {
		ObjOrganizationErrors.Push(&InvalidOrganizationPhoneNumberError{ID: o.ID, Original: o.PhoneNumber, Cause: "フォーマットが不正です"})
	}
}
