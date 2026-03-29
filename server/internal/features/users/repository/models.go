package users_repository

import "github.com/Fioneo/taskflow/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainFromModels(userModels []UserModel) []domain.User {
	userDomains := make([]domain.User, len(userModels))

	for i, model := range userModels {
		userDomains[i] = domain.NewUser(
			model.ID,
			model.Version,
			model.FullName,
			model.PhoneNumber,
		)
	}
	return userDomains
}
