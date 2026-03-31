package tasks_service

import (
	"context"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

func (s *TasksService) GetTasks(ctx context.Context, limit, offset, userID *int) ([]domain.Task, error) {
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
	if userID != nil {
		if *userID <= 0 {
			return nil, fmt.Errorf("userID must be bigger then 0: %w", core_errors.ErrInvalidArgument)
		}
	}

	domainTasks, err := s.tasksRepository.GetTasks(ctx, limit, offset, userID)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("failed to get tasks from database: %w", err)
	}

	return domainTasks, nil
}
