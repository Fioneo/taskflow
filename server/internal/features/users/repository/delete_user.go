package users_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.dbpool.OpTimeOut())
	defer cancel()

	query := `DELETE FROM taskflow.users WHERE id=$1;`

	cmdTag, err := r.dbpool.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d' not found: %w", id, core_errors.ErrNotFound)
	}
	return nil
}
