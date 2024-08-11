package user

import (
	"context"
	"newproject/domain/entities"
)

// Repository of user
type Repository interface {
	Create(ctx context.Context, user entities.User) error
}
