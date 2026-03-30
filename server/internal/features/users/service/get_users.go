package users_service

import (
	"context"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil {
		if *limit <= 0 || *limit > 100 {
			return nil, fmt.Errorf("limit must be between 1 and 100: %w", core_errors.ErrInvalidArgument)
		}
	}

	if offset != nil {
		if *offset < 0 {
			return nil, fmt.Errorf("offset must be non-negative: %w", core_errors.ErrInvalidArgument)
		}
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get users from database: %w", err)
	}
	return users, nil

}
