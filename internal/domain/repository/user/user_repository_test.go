package user

import (
	"context"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/spreadsheet"
	"github.com/stretchr/testify/assert"
)

type MockSpreadsheetDatastore struct{}

func (m *MockSpreadsheetDatastore) Values(ctx context.Context, spreadsheetId, readRange string) (*spreadsheet.ValueResponse, error) {
	return &spreadsheet.ValueResponse{
		Range:          spreadsheetId,
		MajorDimension: "test_major_dimension",
		Values: [][]interface{}{
			{1, "test_family_name", "test_given_name", "test_organization", "test_email", "test_phone_number"},
		},
	}, nil
}

func TestUserRepository_GetUsers(t *testing.T) {
	testCases := []struct {
		id           int
		familyName   string
		givenName    string
		organization string
		email        string
		phoneNumber  string
	}{
		{
			id:           1,
			familyName:   "test_family_name",
			givenName:    "test_given_name",
			organization: "test_organization",
			email:        "test_email",
			phoneNumber:  "test_phone_number",
		},
	}

	ds := &MockSpreadsheetDatastore{}
	repo := NewUserRepository(ds)
	users, err := repo.GetUsers(context.Background(), "test_spreadsheet_id", "test_read_range")

	assert.NoError(t, err)

	for i, tc := range testCases {
		assert.Equal(t, tc.id, users[i].ID)
		assert.Equal(t, tc.familyName, users[i].FamilyName)
		assert.Equal(t, tc.givenName, users[i].GivenName)
		assert.Equal(t, tc.organization, users[i].Organization)
		assert.Equal(t, tc.email, users[i].Email)
		assert.Equal(t, tc.phoneNumber, users[i].PhoneNumber)
	}

}
