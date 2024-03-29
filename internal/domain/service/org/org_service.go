package org

import (
	"context"

	model_org "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/org"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/org"
)

type ServiceInterface interface {
	FindByID(id int) (*model_org.Organization, error)
}

type Service struct {
	ctx  context.Context
	repo org.RepositoryInterface
}

func NewOrganizationService(ctx context.Context, r org.RepositoryInterface) ServiceInterface {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) FindByID(id int) (*model_org.Organization, error) {
	organizations, err := s.repo.GetOrganizations()
	if err != nil {
		return nil, err
	}

	organization, ok := organizations.FindByID(id)
	if !ok {
		return nil, &model_org.OrganizationNotFoundError{ID: id}
	}
	return organization, nil
}
