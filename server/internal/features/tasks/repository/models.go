package tasks_repository

import (
	"time"

	"github.com/Fioneo/taskflow/internal/core/domain"
)

type TaskModel struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func taskDomainsFromModels(tasks []TaskModel) []domain.Task {
	tasksDomain := make([]domain.Task, len(tasks))

	for i, taskModel := range tasks {
		tasksDomain[i] = domain.NewTask(
			taskModel.ID,
			taskModel.Version,
			taskModel.Title,
			taskModel.Description,
			taskModel.Completed,
			taskModel.CreatedAt,
			taskModel.CompletedAt,
			taskModel.AuthorID,
		)
	}
	return tasksDomain
}
