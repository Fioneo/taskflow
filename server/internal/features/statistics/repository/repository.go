package statistics_repository

import core_postgres_pool "github.com/Fioneo/taskflow/internal/core/repository/pool"

type StatisticsRepository struct {
	dbpool *core_postgres_pool.Pool
}

func NewStatisticsRepository(dbpool *core_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{
		dbpool: dbpool,
	}
}
