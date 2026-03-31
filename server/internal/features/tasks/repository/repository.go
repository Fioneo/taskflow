package tasks_repository

import core_postgres_pool "github.com/Fioneo/taskflow/internal/core/repository/pool"

type TasksRepository struct {
	dbpool *core_postgres_pool.Pool
}

func NewTasksRepository(dbpool *core_postgres_pool.Pool) *TasksRepository {
	return &TasksRepository{
		dbpool: dbpool,
	}
}
