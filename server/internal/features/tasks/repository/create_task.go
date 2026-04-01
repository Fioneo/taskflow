package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbpool.OpTimeOut())
	defer cancel()

	query := `
	  INSERT INTO
	  taskflow.tasks 
	  (title,description,completed,completed_at,author_id)
	  VALUES ($1,$2,$3,$4,$5) 
	  RETURNING id,
	  version,
	  title,
	  description,
	  completed,
	  created_at,
	  completed_at,
	  author_id;
	  `

	var taskModel TaskModel

	row := r.dbpool.QueryRow(ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		task.AuthorID,
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
		var pgErr *pgconn.PgError
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("not found: %w", core_errors.ErrNotFound)
		}
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return domain.Task{}, fmt.Errorf("foreign key violation:user with id='%d' not found : %w", task.AuthorID, err)
			}
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
