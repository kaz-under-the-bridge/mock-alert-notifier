package org

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrgToMap(t *testing.T) {
	org := Organization{
		ID:          1,
		Name:        "◯◯株式会社",
		Team:        "開発部",
		Email:       "dev@example.com",
		PhoneNumber: "03-1234-5678",
	}

	want := map[string]string{
		"Organization_Name":        "◯◯株式会社",
		"Organization_Team":        "開発部",
		"Organization_Email":       "dev@example.com",
		"Organization_PhoneNumber": "03-1234-5678",
	}

	got := org.ToMap()

	assert.Equal(t, want, got)
}
