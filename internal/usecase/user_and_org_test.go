package usecase

import (
	"context"
	"testing"

	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	service_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetUsers() (*model_user.Users, error) {
	users := &model_user.Users{}

	dummyUsers := service_user.GenerateDummyUsers()

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

func TestGetUsersAndOrgs(t *testing.T) {
	repo := &MockUserRepository{}
	users := service_user.NewUserService(context.Background(), repo)

	users.FindByEmail("test@example.com")
}
