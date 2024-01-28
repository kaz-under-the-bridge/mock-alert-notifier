package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserVeificationWithError(t *testing.T) {
	_ = NewUser(1, "", "", "", "hoge", 1)
	assert.NotEqual(t, 0, ObjUserErrors.Len())
	want := `ID(1)のFamilyNameが不正です: 空白です
ID(1)のGivenName が不正です: 空白です
ID(1)のEmail()が不正です: 空白です
ID(1)のEmail()が不正です: フォーマットが不正です
ID(1)のPhoneNumber(hoge)が不正です: フォーマットが不正です
`
	assert.Equal(t, want, ObjUserErrors.Error())
}

func TestUserEmailVerificationWithOK(t *testing.T) {
	testCases := []struct {
		id             int
		familyName     string
		givenName      string
		email          string
		phoneNumber    string
		organizationID int
		errorLength    int // UserErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, familyName: "-", givenName: "-", email: "test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 0},
		{id: 2, familyName: "-", givenName: "-", email: "test123.test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 0},
		{id: 3, familyName: "-", givenName: "-", email: "test123..test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 0},
		{id: 4, familyName: "-", givenName: "-", email: "test123-test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 0},
		{id: 5, familyName: "-", givenName: "-", email: "test@test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 1},
		{id: 6, familyName: "-", givenName: "-", email: "test&test@example.com", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 1},
	}

	for _, tc := range testCases {
		ObjUserErrors = NewUserErrors() // ループの前にerrorをrefresh
		_ = NewUser(tc.id, tc.familyName, tc.givenName, tc.email, tc.phoneNumber, tc.organizationID)

		assert.Equal(t, tc.errorLength, ObjUserErrors.Len())
		if ObjUserErrors.Len() != 0 {
			assert.Equal(t, fmt.Sprintf("ID(%d)のEmail(%s)が不正です: フォーマットが不正です\n", tc.id, tc.email), ObjUserErrors.Error())
		}
	}
}

func TestUserPhoneNumberVerificationWithOK(t *testing.T) {
	testCases := []struct {
		id             int
		familyName     string
		givenName      string
		email          string
		phoneNumber    string
		organizationID int
		errorLength    int // UserErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, familyName: "-", givenName: "-", email: "t@e.c", phoneNumber: "080-1234-5678", organizationID: 1, errorLength: 0},
		{id: 2, familyName: "-", givenName: "-", email: "t@e.c", phoneNumber: "1234-5678-9012", organizationID: 1, errorLength: 0},
		{id: 3, familyName: "-", givenName: "-", email: "t@e.c", phoneNumber: "12345678910", organizationID: 1, errorLength: 1},
		{id: 4, familyName: "-", givenName: "-", email: "t@e.c", phoneNumber: "12345-678-9012", organizationID: 1, errorLength: 1},
	}

	for _, tc := range testCases {
		ObjUserErrors = NewUserErrors() // ループの前のerrorをrefresh
		_ = NewUser(tc.id, tc.familyName, tc.givenName, tc.email, tc.phoneNumber, tc.organizationID)

		assert.Equal(t, tc.errorLength, ObjUserErrors.Len())
		if ObjUserErrors.Len() != 0 {
			fmt.Printf("sample: %s", ObjUserErrors.Error())
			assert.Equal(t, fmt.Sprintf("ID(%d)のPhoneNumber(%s)が不正です: フォーマットが不正です\n", tc.id, tc.phoneNumber), ObjUserErrors.Error())
		}
	}
}
