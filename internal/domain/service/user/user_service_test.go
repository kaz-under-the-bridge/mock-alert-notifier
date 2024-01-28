package user

import (
	"context"
	"testing"

	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/stretchr/testify/assert"
)

func GenerateDummyUsers() []*model_user.User {
	return []*model_user.User{
		{
			ID:             1,
			FamilyName:     "テスト",
			GivenName:      "太郎",
			Email:          "taro@example.com",
			PhoneNumber:    "03-1234-5678",
			OrganizationID: 1,
		},
		{
			ID:             2,
			FamilyName:     "テスト",
			GivenName:      "花子",
			Email:          "hanako@example.com",
			PhoneNumber:    "06-8765-4321",
			OrganizationID: 2,
		},
		{
			ID:             3,
			FamilyName:     "Jhon",
			GivenName:      "Doe",
			Email:          "jd@example.com",
			PhoneNumber:    "310-555-3067",
			OrganizationID: 1,
		},
	}
}

type MockUserRepository struct{}

func (m *MockUserRepository) GetUsers() (*model_user.Users, error) {
	users := &model_user.Users{}

	dummyUsers := GenerateDummyUsers()

	for _, elm := range dummyUsers {
		users.Push(elm)
	}

	return users, nil
}

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

func TestUserService_FindByEmail(t *testing.T) {
	testCase := []struct {
		email           string
		wantID          int
		wantFamilyName  string
		wantGivenName   string
		wantPhoneNumber string
		isErr           bool
	}{
		{
			email:           "taro@example.com",
			wantID:          1,
			wantFamilyName:  "テスト",
			wantGivenName:   "太郎",
			wantPhoneNumber: "03-1234-5678",
			isErr:           false,
		},
		{
			email:           "hanako@example.com",
			wantID:          2,
			wantFamilyName:  "テスト",
			wantGivenName:   "花子",
			wantPhoneNumber: "06-8765-4321",
			isErr:           false,
		},
		{
			email:           "unknown@example.com",
			wantID:          0,
			wantFamilyName:  "",
			wantGivenName:   "",
			wantPhoneNumber: "",
			isErr:           true,
		},
	}

	repo := &MockUserRepository{}
	users := NewUserService(context.Background(), repo)

	for _, tc := range testCase {
		user, err := users.FindByEmail(tc.email)

		if tc.isErr {
			assert.Error(t, err)
			assert.Equal(t, &model_user.UserNotFoundError{Email: tc.email}, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.wantID, user.ID)
			assert.Equal(t, tc.wantFamilyName, user.FamilyName)
			assert.Equal(t, tc.wantGivenName, user.GivenName)
			assert.Equal(t, tc.wantPhoneNumber, user.PhoneNumber)
		}
	}
}
