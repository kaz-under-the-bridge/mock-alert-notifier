package user

import (
	"context"
	"testing"

	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetUsers() (*model_user.Users, error) {
	return GenerateDummyUsers(), nil
}

func TestMain(m *testing.M) {
	if err := helper.GetNewLogger(helper.SetLogType(context.Background(), "test")); err != nil {
		panic(err)
	}

	m.Run()
}

func TestUserService_FindByEmail(t *testing.T) {
	testCases := []struct {
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

	for _, tc := range testCases {
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
