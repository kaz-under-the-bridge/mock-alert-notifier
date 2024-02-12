package org

import (
	"context"
	"testing"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/infrastracture/spreadsheet"
	"github.com/stretchr/testify/assert"
)

type MockOrganizationRepository struct{}

func (m *MockOrganizationRepository) Values(ctx context.Context, spreasheetId, readRange string) (*spreadsheet.ValueResponse, error) {
	return &spreadsheet.ValueResponse{
		Range:          spreadsheetId,
		MajorDimension: "test_major_dimension",
		Values: [][]interface{}{
			{1, "test_organization_name", "test_team_name", "test@example.com", "03-1234-5678"},
		},
	}, nil
}

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

func TestOrganizationRepository_GetOrganizations(t *testing.T) {
	testCases := []struct {
		id          int
		name        string
		team        string
		email       string
		phoneNumber string
	}{
		{
			id:          1,
			name:        "test_organization_name",
			team:        "test_team_name",
			email:       "test@example.com",
			phoneNumber: "03-1234-5678",
		},
	}

	ds := &MockOrganizationRepository{}
	repo := NewOrganizationRepository(context.Background(), ds)
	organizations, err := repo.GetOrganizations()

	assert.NoError(t, err)

	organization, ok := organizations.FindByID(1)
	assert.Equal(t, true, ok)

	for _, tc := range testCases {
		assert.Equal(t, tc.id, organization.ID)
		assert.Equal(t, tc.name, organization.Name)
		assert.Equal(t, tc.team, organization.Team)
		assert.Equal(t, tc.email, organization.Email)
		assert.Equal(t, tc.phoneNumber, organization.PhoneNumber)
	}
}

/*
以下のような表形式のデータのvalidation
1 A B        C       D                       E
2 1 ◯株式会社 チーム太郎 team-taro@example.com   090-1234-5678 <- OK
3 2 株式会社△ チーム花子 team-hanako@example.com 080-9876-5432 <- OK
4 a test     test    a@b.c                   123-4567-8901 <- NG
5 4 1        test    a@b.c                   123-4567-8901 <- NG
6 5 test     2       a@b.c                   123-4567-8901 <- NG
7 6 test     test    3                       123-4567-8901 <- NG
8 7 test     test    a@b.c                   a             <- NG
*/
func TestValidateRowDataType(t *testing.T) {
	testCases := []struct {
		data        []interface{}
		isError     bool
		msgExpected string
	}{
		{data: []interface{}{1, "◯株式会社", "チーム太郎", "team-taro@example.com", "090-1234-5678"}, isError: false, msgExpected: ""},
		{data: []interface{}{2, "株式会社△", "チーム花子", "team-hanako@example.com", "080-9876-5432"}, isError: false, msgExpected: ""},
		{data: []interface{}{"a", "test", "test", "a@b.c", "123-4567-8901"}, isError: true, msgExpected: "Cell: A5: Invalid Data"},
		{data: []interface{}{4, 1, "test", "a@b.c", "123-4567-8901"}, isError: true, msgExpected: "Cell: B6: Invalid Data"},
		{data: []interface{}{5, "test", 2, "a@b.c", "123-4567-8901"}, isError: true, msgExpected: "Cell: C7: Invalid Data"},
		{data: []interface{}{6, "test", "test", 3, "123-4567-8901"}, isError: true, msgExpected: "Cell: D8: Invalid Data"},
		{data: []interface{}{7, "test", "test", "a@b.c", 1}, isError: true, msgExpected: "Cell: E9: Invalid Data"},
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
