package user

import (
	"context"

	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/user"
)

var _ ServiceInterface = (*Service)(nil)

type ServiceInterface interface {
}

type Service struct {
	ctx  context.Context
	repo user.RepositoryInterface
}

func NewUserService(ctx context.Context, r user.RepositoryInterface) *Service {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) FindByEmail(email string) (*model.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	user, ok := users.FindByEmail(email)
	if !ok {
		return nil, &model.UserNotFoundError{Email: email}
	}
	return user, nil
}
