package org

import model "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"

func GenerateDummyOrgs() *model.Organizations {
	return &model.Organizations{
		{
			ID:          1,
			Name:        "〇〇株式会社",
			Team:        "開発部",
			Email:       "dev@maru.example.com",
			PhoneNumber: "03-1234-5678",
		},
		{
			ID:          2,
			Name:        "△△合同会社",
			Team:        "開発部",
			Email:       "deveolop@sankaku.example.com",
			PhoneNumber: "090-1234-5678",
		},
		{
			ID:          3,
			Name:        "〇〇株式会社",
			Team:        "新規企画室",
			Email:       "plan@maru.example.com",
			PhoneNumber: "03-1234-5679",
		},
	}
}
