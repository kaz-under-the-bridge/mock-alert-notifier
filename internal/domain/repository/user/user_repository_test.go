package user

import (
	"context"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/spreadsheet"
	"github.com/stretchr/testify/assert"
)

type MockSpreadsheetDatastore struct{}

func (m *MockSpreadsheetDatastore) Values(ctx context.Context, spreadsheetId, readRange string) (*spreadsheet.ValueResponse, error) {
	return &spreadsheet.ValueResponse{
		Range:          spreadsheetId,
		MajorDimension: "test_major_dimension",
		Values: [][]interface{}{
			{1, "test_family_name", "test_given_name", "test_email", "test_phone_number", 1},
		},
	}, nil
}

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

func TestUserRepository_GetUsers(t *testing.T) {
	testCases := []struct {
		id             int
		familyName     string
		givenName      string
		email          string
		phoneNumber    string
		organizationID int
	}{
		{
			id:             1,
			familyName:     "test_family_name",
			givenName:      "test_given_name",
			email:          "test_email",
			phoneNumber:    "test_phone_number",
			organizationID: 1,
		},
	}

	ds := &MockSpreadsheetDatastore{}
	repo := NewUserRepository(context.Background(), ds)
	users, err := repo.GetUsers()

	assert.NoError(t, err)

	user, ok := users.FindByID(1)
	assert.Equal(t, true, ok)

	for _, tc := range testCases {
		assert.Equal(t, tc.id, user.ID)
		assert.Equal(t, tc.familyName, user.FamilyName)
		assert.Equal(t, tc.givenName, user.GivenName)
		assert.Equal(t, tc.email, user.Email)
		assert.Equal(t, tc.phoneNumber, user.PhoneNumber)
		assert.Equal(t, tc.organizationID, user.OrganizationID)
	}

}

/*
以下のような表形式のデータのvalidation
1 A B    C      D     E             F
2 1 test taro   a@b.c 080-1234-5678 1 <- OK
3 2 test hanako d@e.f 090-1234-5678 2 <- OK
4 a test hanako d@e.f 090-1234-5678 2 <- NG
5 4 1    hanako d@e.f 090-1234-5678 2 <- NG
6 5 test taro   1     080-1234-5678 1 <- NG
7 6 test taro   a@b.c a             1 <- NG
8 7 test hanako d@e.f 090-1234-5678 a <- NG
*/
func TestValidateRowDataType(t *testing.T) {
	testCases := []struct {
		data        []interface{}
		isError     bool
		msgExpected string
	}{
		{data: []interface{}{1, "test", "taro", "a@b.c", "080-1234-5678", 1}, isError: false, msgExpected: ""},
		{data: []interface{}{2, "test", "hanako", "d@e.f", "090-1234-5678", 2}, isError: false, msgExpected: ""},
		{data: []interface{}{"a", "test", "hanako", "d@e.f", "090-1234-5678", 2}, isError: true, msgExpected: "Cell: A5: Invalid Data"},
		{data: []interface{}{4, 1, "hanako", "d@e.f", "090-1234-5678", 2}, isError: true, msgExpected: "Cell: B6: Invalid Data"},
		{data: []interface{}{5, "test", 1, "d@e.f", "090-1234-5678", 2}, isError: true, msgExpected: "Cell: C7: Invalid Data"},
		{data: []interface{}{6, "test", "taro", 1, "080-1234-5678", 1}, isError: true, msgExpected: "Cell: D8: Invalid Data"},
		{data: []interface{}{7, "test", "taro", "a@b.c", 1, 1}, isError: true, msgExpected: "Cell: E9: Invalid Data"},
		{data: []interface{}{8, "test", "hanako", "d@e.f", "090-1234-5678", "a"}, isError: true, msgExpected: "Cell: F10: Invalid Data"},
	}

	for r, tc := range testCases {
		err := validateRowDataType(r, tc.data)
		if tc.isError {
			assert.Error(t, err)
			assert.Equal(t, tc.msgExpected, err.Error())
		} else {
			assert.NoError(t, err)
		}
	}
}
