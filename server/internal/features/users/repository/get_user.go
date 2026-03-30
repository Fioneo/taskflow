package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fioneo/taskflow/internal/core/domain"
	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.dbpool.OpTimeOut())
	defer cancel()
	var userModel UserModel
	query := `SELECT id, version, full_name, phone_number FROM taskflow.users WHERE id=$1;`

	row := r.dbpool.QueryRow(ctx, query, id)

	err := row.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%d': %w", id, core_errors.ErrNotFould)
		}
		return domain.User{}, fmt.Errorf("get user by id: %w", err)
	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
