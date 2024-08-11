package user

import (
	"context"
	"newproject/domain/entities"
)

type UseCases interface {
	// Create new User.
	Create(ctx context.Context, user entities.User) error
}
