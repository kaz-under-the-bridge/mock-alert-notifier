package org

import (
	"context"
	"testing"

	model_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/stretchr/testify/assert"
)

type MockOrganizationRepository struct{}

func (r *MockOrganizationRepository) GetOrganizations() (*model_org.Organizations, error) {
	return GenerateDummyOrgs(), nil
}

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

func TestUserService_FindByID(t *testing.T) {
	testCases := []struct {
		ID          int
		Name        string
		Team        string
		Email       string
		PhoneNumber string
		isErr       bool
	}{
		{
			ID:          1,
			Name:        "〇〇株式会社",
			Team:        "開発部",
			Email:       "dev@maru.example.com",
			PhoneNumber: "03-1234-5678",
			isErr:       false,
		},
		{
			ID:          2,
			Name:        "△△合同会社",
			Team:        "開発部",
			Email:       "deveolop@sankaku.example.com",
			PhoneNumber: "090-1234-5678",
			isErr:       false,
		},
		{
			ID:          0,
			Name:        "",
			Team:        "",
			Email:       "unknown@example.com",
			PhoneNumber: "",
			isErr:       true,
		},
	}

	repo := &MockOrganizationRepository{}
	orgs := NewOrganizationService(context.Background(), repo)

	for _, tc := range testCases {
		org, err := orgs.FindByID(tc.ID)

		if tc.isErr {
			assert.Error(t, err)
			assert.Equal(t, &model_org.OrganizationNotFoundError{ID: tc.ID}, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.ID, org.ID)
			assert.Equal(t, tc.Name, org.Name)
			assert.Equal(t, tc.Team, org.Team)
			assert.Equal(t, tc.Email, org.Email)
			assert.Equal(t, tc.PhoneNumber, org.PhoneNumber)
		}
	}
}
