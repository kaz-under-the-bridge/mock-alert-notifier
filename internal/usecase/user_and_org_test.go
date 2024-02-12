package usecase

import (
	"context"
	"testing"

	model_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"
	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	service_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/org"
	service_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetUsers() (*model_user.Users, error) {
	return service_user.GenerateDummyUsers(), nil
}

type MockOrganizationRepository struct{}

func (m *MockOrganizationRepository) GetOrganizations() (*model_org.Organizations, error) {
	return service_org.GenerateDummyOrgs(), nil
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
