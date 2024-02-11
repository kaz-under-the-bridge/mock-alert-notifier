package user

import (
	"reflect"
	"regexp"
)

// declare regex matcher for email
// - @マークの前は英数字, ハイフン, ドット(.)を許可
// - ドメインは英小文字、数字、ハイフンを許可(トップレベルドメインは英小文字のみ)
var regexEmail = regexp.MustCompile(`^[\w+\-.]+@[a-z\d\-.]+\.[a-z]+$`)

// declare regex matcher for phone number
// - xxxx-xxxx-xxxx(4桁, ハイフンあり)
var regexPhoneNumber = regexp.MustCompile(`^\d{1,4}-\d{1,4}-\d{4}$`)

var ObjUserErrors *UserErrors

func init() {
	ObjUserErrors = NewUserErrors()
}

type Users []*User

type User struct {
	ID             int    `json:"id"`
	FamilyName     string `json:"name"`
	GivenName      string `json:"given_name"`
	Email          string `json:"email"`        // regexEmailに正規表現を定義
	PhoneNumber    string `json:"phone_number"` // regexPhoneNumberに正規表現を定義
	OrganizationID int    `json:"organization_id"`
}

func NewUser(
	id int,
	familyName, givenName string,
	email, phoneNumber string,
	organizationId int,
) *User {
	u := &User{
		ID:             id,
		FamilyName:     familyName,
		GivenName:      givenName,
		Email:          email,
		PhoneNumber:    phoneNumber,
		OrganizationID: organizationId,
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

func (us Users) FindByID(id int) (*User, bool) {
	for _, u := range us {
		if u.ID == id {
			return u, true
		}
	}
	return nil, false
}

func (us Users) FindByEmail(email string) (*User, bool) {
	for _, u := range us {
		if u.Email == email {
			return u, true
		}
	}
	return nil, false
}

func (us *Users) FindBySameOrganization(u *User) (*Users, bool) {
	users := Users{}
	flag := false

	for _, user := range *us {
		if user.OrganizationID == u.OrganizationID {
			users.Push(user)
			flag = true
		}
	}
	return &users, flag
}

func (u User) verify() {
	// ID列（A列・D列
	if u.FamilyName == "" {
		ObjUserErrors.Push(&InvalidUserFamilyNameError{ID: u.ID, Cause: "空白です"})
	}

	if u.GivenName == "" {
		ObjUserErrors.Push(&InvalidUserGivenNameError{ID: u.ID, Cause: "空白です"})
	}

	if u.Email == "" {
		ObjUserErrors.Push(&InvalidUserEmailError{ID: u.ID, Cause: "空白です"})
	}

	// verify email format
	if !regexEmail.MatchString(u.Email) {
		ObjUserErrors.Push(&InvalidUserEmailError{ID: u.ID, Original: u.Email, Cause: "フォーマットが不正です"})
	}

	// verify phone number format
	if !regexPhoneNumber.MatchString(u.PhoneNumber) {
		ObjUserErrors.Push(&InvalidUserPhoneNumberError{ID: u.ID, original: u.PhoneNumber, Cause: "フォーマットが不正です"})
	}
}

// string fieldのみをmapに変換する
func (u User) ToMap() map[string]string {
	data := map[string]string{}

	v := reflect.ValueOf(u)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.String {
			data[t.Field(i).Name] = v.Field(i).String()
		}
	}

	return data
}
