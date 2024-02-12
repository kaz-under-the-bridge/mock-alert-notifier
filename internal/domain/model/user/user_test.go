package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserToMap(t *testing.T) {

	user := User{
		ID:             1,
		FamilyName:     "鈴木",
		GivenName:      "テスト",
		Email:          "test@example.com",
		PhoneNumber:    "03-1234-5678",
		OrganizationID: 1,
	}

	want := map[string]string{
		"User_FamilyName":  "鈴木",
		"User_GivenName":   "テスト",
		"User_Email":       "test@example.com",
		"User_PhoneNumber": "03-1234-5678",
	}

	got := user.ToMap()
	fmt.Printf("want: %v\n", got)

	assert.Equal(t, want, got)

}
