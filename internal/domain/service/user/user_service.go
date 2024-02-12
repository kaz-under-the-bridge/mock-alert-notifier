package user

import (
	"context"

	model_user "github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/model/user"
	"github.com/kaz-under-the-bridge/mock-alert-notifier/internal/domain/repository/user"
)

var _ ServiceInterface = (*Service)(nil)

type ServiceInterface interface {
	FindByEmail(email string) (*model_user.User, error)
}

type Service struct {
	ctx  context.Context
	repo user.RepositoryInterface
}

func NewUserService(ctx context.Context, r user.RepositoryInterface) ServiceInterface {
	return &Service{ctx: ctx, repo: r}
}

func (s *Service) FindByEmail(email string) (*model_user.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	user, ok := users.FindByEmail(email)
	if !ok {
		return nil, &model_user.UserNotFoundError{Email: email}
	}
	return user, nil
}
