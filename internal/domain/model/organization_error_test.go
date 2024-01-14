package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationVerificationWithError(t *testing.T) {
	_ = NewOrganization(1, "", "", "hoge", "")
	assert.NotEqual(t, 0, ObjOrganizationErrors.Len())
	want := `ID(1)のOrg名が不正です: 空白です
ID(1)のTeam名が不正です: 空白です
ID(1)のEmail(hoge)が不正です: フォーマットが不正です
ID(1)のPhoneNumber()が不正です: フォーマットが不正です
`
	assert.Equal(t, want, ObjOrganizationErrors.Error())
}

func TestOrganizationEmailVericationWithOK(t *testing.T) {
	// generate test data most likely TestUserEmailVerificationWithOK
	testCases := []struct {
		id          int
		name        string
		team        string
		email       string
		phoneNumber string
		errorLength int // OrganizationErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, name: "-", team: "-", email: "test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 2, name: "-", team: "-", email: "test123.test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 3, name: "-", team: "-", email: "test123..test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 4, name: "-", team: "-", email: "test123-test@example.com", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 5, name: "-", team: "-", email: "test@test@example.com", phoneNumber: "080-1234-5678", errorLength: 1},
		{id: 6, name: "-", team: "-", email: "test&test@example.com", phoneNumber: "080-1234-5678", errorLength: 1},
	}

	for _, tc := range testCases {
		ObjOrganizationErrors = NewOrganizationErrors() // ループの前にerrorをrefresh
		_ = NewOrganization(tc.id, tc.name, tc.team, tc.email, tc.phoneNumber)

		assert.Equal(t, tc.errorLength, ObjOrganizationErrors.Len())
		if ObjOrganizationErrors.Len() != 0 {
			assert.Equal(t, fmt.Sprintf("ID(%d)のEmail(%s)が不正です: フォーマットが不正です\n", tc.id, tc.email), ObjOrganizationErrors.Error())
		}
	}
}

func TestOrganizationPhoneNumberVerificationWithOK(t *testing.T) {
	// generate test data most likely TestUserPhoneNumberVerificationWithOK
	testCases := []struct {
		id          int
		name        string
		team        string
		email       string
		phoneNumber string
		errorLength int // OrganizationErrorsの長さ（0ならのエラーなし）
	}{
		{id: 1, name: "-", team: "-", email: "t@e.c", phoneNumber: "080-1234-5678", errorLength: 0},
		{id: 2, name: "-", team: "-", email: "t@e.c", phoneNumber: "1234-5678-9012", errorLength: 0},
		{id: 3, name: "-", team: "-", email: "t@e.c", phoneNumber: "12345678910", errorLength: 1},
		{id: 4, name: "-", team: "-", email: "t@e.c", phoneNumber: "12345-678-9012", errorLength: 1},
	}

	for _, tc := range testCases {
		ObjOrganizationErrors = NewOrganizationErrors() // ループの前にerrorをrefresh
		_ = NewOrganization(tc.id, tc.name, tc.team, tc.email, tc.phoneNumber)

		assert.Equal(t, tc.errorLength, ObjOrganizationErrors.Len())
		if ObjOrganizationErrors.Len() != 0 {
			assert.Equal(t, fmt.Sprintf("ID(%d)のPhoneNumber(%s)が不正です: フォーマットが不正です\n", tc.id, tc.phoneNumber), ObjOrganizationErrors.Error())
		}
	}
}
