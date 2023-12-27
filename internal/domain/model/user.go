package model

import "regexp"

type Users []*User

// declare regex matcher for email
// - @マークの前は英数字, ハイフン, ドット(.)を許可
// - ドメインは英小文字、数字、ハイフンを許可(トップレベルドメインは英小文字のみ)
var regexEmail = regexp.MustCompile(`^[\w+\-.]+@[a-z\d\-.]+\.[a-z]+$`)

// declare regex matcher for phone number
// - xxxx-xxxx-xxxx(4桁, ハイフンあり)
var regexPhoneNumber = regexp.MustCompile(`^\d{1,4}-\d{1,4}-\d{4}$`)

var UserErrors *UserErrrors

func init() {
	UserErrors = NewUserErrrors()
}

type User struct {
	ID           int    `json:"id"`
	FamilyName   string `json:"name"`
	GivenName    string `json:"given_name"`
	Organization string `json:"organization"`
	Email        string `json:"email"`        // regexEmailに正規表現を定義
	PhoneNumber  string `json:"phone_number"` // regexPhoneNumberに正規表現を定義
}

func NewUser(
	id int,
	familyName, givenName,
	organization,
	email,
	phoneNumber string,
) *User {
	u := &User{
		ID:           id,
		FamilyName:   familyName,
		GivenName:    givenName,
		Organization: organization,
		Email:        email,
		PhoneNumber:  phoneNumber,
	}
	u.verify()

	return u
}

func (us *Users) ToMap() map[int]*User {
	m := make(map[int]*User)
	for _, u := range *us {
		m[u.ID] = u
	}
	return m
}

func (us *Users) ToSlice() []*User {
	slice := make([]*User, 0, len(*us))

	for _, u := range *us {
		slice = append(slice, u)
	}
	return slice
}

func (us *Users) ToSliceWithID() []int {
	slice := make([]int, 0, len(*us))

	for _, u := range *us {
		slice = append(slice, u.ID)
	}
	return slice
}

func (us *Users) Push(u *User) {
	*us = append(*us, u)
}

func (us *Users) Len() int {
	cnt := 0

	for range *us {
		cnt++
	}
	return cnt
}

func (us *Users) FindByID(id int) *User {
	for _, u := range *us {
		if u.ID == id {
			return u
		}
	}
	return nil
}

func (us *Users) FindBySameOrganization(u *User) []*User {
	users := make([]*User, 0, len(*us))

	for _, user := range *us {
		if user.Organization == u.Organization {
			users = append(users, user)
		}
	}
	return users
}

func (u User) verify() {
	if u.FamilyName == "" {
		UserErrors.Push(&InvalidUserFamilyNameError{id: u.ID, cause: "空白です"})
	}

	if u.GivenName == "" {
		UserErrors.Push(&InvalidUserGivenNameError{id: u.ID, cause: "空白です"})
	}

	if u.Organization == "" {
		UserErrors.Push(&InvalidUserOrganizationError{id: u.ID, cause: "空白です"})
	}

	if u.Email == "" {
		UserErrors.Push(&InvalidUserEmailError{id: u.ID, cause: "空白です"})
	}

	// verify email format
	if !regexEmail.MatchString(u.Email) {
		UserErrors.Push(&InvalidUserEmailError{id: u.ID, original: u.Email, cause: "フォーマットが不正です"})
	}

	// verify phone number format
	if !regexPhoneNumber.MatchString(u.PhoneNumber) {
		UserErrors.Push(&InvalidUserPhoneNumberError{id: u.ID, original: u.Email, cause: "フォーマットが不正です"})
	}
}
