package usecase

import (
	"context"
	"testing"

	model_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"
	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	service_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/org"
	service_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/service/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/helper"
	"github.com/stretchr/testify/assert"
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
	userRepo := &MockUserRepository{}
	users := service_user.NewUserService(context.Background(), userRepo)

	orgRepo := &MockOrganizationRepository{}
	orgs := service_org.NewOrganizationService(context.Background(), orgRepo)

	_, err := users.FindByEmail("test@example.com")
	assert.Error(t, err)

	user, _ := users.FindByEmail("taro@example.com")
	assert.Equal(t, "テスト 太郎", user.FullName())

	_, err = orgs.FindByID(0)
	assert.Error(t, err)

	org, _ := orgs.FindByID(1)
	assert.Equal(t, "〇〇株式会社", org.Name)
}
