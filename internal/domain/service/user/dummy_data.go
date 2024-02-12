package user

import model "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"

func GenerateDummyUsers() *model.Users {
	return &model.Users{
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
