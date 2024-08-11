package user

import (
	"context"
	"newproject/domain/entities"
	"newproject/infrastructure/repositories/user"
	"newproject/view/http_error"
	"strings"
)

type useCases struct {
	userRepo user.Repository
}

func NewUseCases(userRepo user.Repository) UseCases {
	return &useCases{
		userRepo: userRepo,
	}
}
func (u useCases) Create(ctx context.Context, user entities.User) error {

	user.Name = strings.TrimSpace(user.Name)

	if user.Name == "" {
		return http_error.NewBadRequestError("Nome não definido.")
	}
	if len(user.Name) > 100 {
		return http_error.NewBadRequestError("Nome não pode conter mais de 100 caracteres.")
	}

	return u.userRepo.Create(ctx, user)
}
