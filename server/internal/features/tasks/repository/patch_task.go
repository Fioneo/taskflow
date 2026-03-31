package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r TasksRepository) PatchTask(ctx context.Context, id int, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbpool.OpTimeOut())
	defer cancel()

	query := `UPDATE taskflow.tasks
	SET 
	title=$1, 
	description=$2, 
	completed=$3, 
	completed_at=$4, 
	version=version+1 

	WHERE id=$5 AND version=$6

	RETURNING 
	id,
	version,
	title,
	description,
	completed,
	created_at,
	completed_at,
	author_id;`

	var taskModel TaskModel

	row := r.dbpool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.ID,
		task.Version,
	)

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
			return domain.Task{}, fmt.Errorf("task with id='%d' concurrently accessed: %w", id, core_errors.ErrConflict)
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
	return domainTask, err
}
