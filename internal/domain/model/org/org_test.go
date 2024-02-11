package org

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrgToMap(t *testing.T) {
	org := Organization{
		ID:          1,
		Name:        "test org name",
		Team:        "test team name",
		Email:       "test@example.com",
		PhoneNumber: "03-1234-5678",
	}

	want := map[string]string{
		"Name":        "test org name",
		"Team":        "test team name",
		"Email":       "test@example.com",
		"PhoneNumber": "03-1234-5678",
	}

	got := org.ToMap()

	assert.Equal(t, want, got)
}
