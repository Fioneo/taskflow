package users_service

import (
	"context"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user: %w", err)
	}
	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}
