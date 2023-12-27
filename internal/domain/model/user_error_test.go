package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserVeificationWithError(t *testing.T) {
	_ = NewUser(1, "", "", "", "hoge", "")
	assert.NotEqual(t, 0, UserErrors.Len())
	want := `ID(1)のFamilyNameが不正です: 空白です
ID(1)のGivenName が不正です: 空白です
ID(1)のOrganizationが不正です: 空白です
ID(1)のEmailが不正です: フォーマットが不正です
ID(1)のPhoneNumberが不正です: フォーマットが不正です
`
	assert.Equal(t, want, UserErrors.Error())
}

func TestUserEmailVerificationWithOK(t *testing.T) {
	testCases := []struct {
		id           int
		familyName   string
		givenName    string
		organization string
		email        string
		phoneNumber  string
		errorLength  int // UserErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, familyName: "-", givenName: "-", organization: "-", email: "test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 2, familyName: "-", givenName: "-", organization: "-", email: "test123.test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 3, familyName: "-", givenName: "-", organization: "-", email: "test123..test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 4, familyName: "-", givenName: "-", organization: "-", email: "test123-test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 5, familyName: "-", givenName: "-", organization: "-", email: "test@test@example.com", phoneNumber: "080-1234-5678", errorLength: 1},
		{id: 6, familyName: "-", givenName: "-", organization: "-", email: "test&test@example.com", phoneNumber: "080-1234-5678", errorLength: 1},
	}

	for _, tc := range testCases {
		UserErrors = NewUserErrrors()
		_ = NewUser(tc.id, tc.familyName, tc.givenName, tc.organization, tc.email, tc.phoneNumber)

		assert.Equal(t, tc.errorLength, UserErrors.Len())
		if UserErrors.Len() != 0 {
			assert.Equal(t, fmt.Sprintf("ID(%d)のEmailが不正です: フォーマットが不正です\n", tc.id), UserErrors.Error())
		}
	}
}

func TestUserPhoneNumberVerificationWithOK(t *testing.T) {
	testCases := []struct {
		id           int
		familyName   string
		givenName    string
		organization string
		email        string
		phoneNumber  string
		errorLength  int // UserErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, familyName: "-", givenName: "-", organization: "-", email: "t@e.c", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 2, familyName: "-", givenName: "-", organization: "-", email: "t@e.c", phoneNumber: "1234-5678-9012", errorLength: 0},
		{id: 3, familyName: "-", givenName: "-", organization: "-", email: "t@e.c", phoneNumber: "12345678910", errorLength: 1},
		{id: 3, familyName: "-", givenName: "-", organization: "-", email: "t@e.c", phoneNumber: "12345-678-9012", errorLength: 1},
	}

	for _, tc := range testCases {
		UserErrors = NewUserErrrors()
		_ = NewUser(tc.id, tc.familyName, tc.givenName, tc.organization, tc.email, tc.phoneNumber)

		assert.Equal(t, tc.errorLength, UserErrors.Len())
		if UserErrors.Len() != 0 {
			assert.Equal(t, fmt.Sprintf("ID(%d)のEmailが不正です: フォーマットが不正です\n", tc.id), UserErrors.Error())
		}
	}
}
