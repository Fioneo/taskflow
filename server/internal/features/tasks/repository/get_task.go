package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *TasksRepository) GetTask(ctx context.Context, id int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbpool.OpTimeOut())
	defer cancel()

	query := `
	SELECT 
	id,
	version,
	title,
	description,
	completed,
	created_at,
	completed_at,
	author_id 
	FROM taskflow.tasks 
	WHERE id=$1;
	`
	var taskModel TaskModel

	row := r.dbpool.QueryRow(ctx, query, id)

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("failed to find task with id='%d': %w", id, core_errors.ErrNotFound)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	domainTask := domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorID,
	)

	return domainTask, nil
}
