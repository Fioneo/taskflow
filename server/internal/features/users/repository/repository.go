package users_repository

import core_postgres_pool "github.com/Fioneo/taskflow/internal/core/repository/pool"

type UsersRepository struct {
	dbpool *core_postgres_pool.Pool
}

func NewUsersRepository(dbpool *core_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		dbpool: dbpool,
	}
}
