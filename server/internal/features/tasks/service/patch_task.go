package tasks_service

import (
	"context"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task from the database: %w", err)
	}
	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}
	patchedTack, err := s.tasksRepository.PatchTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to patch task: %w", err)
	}
	return patchedTack, nil
}
