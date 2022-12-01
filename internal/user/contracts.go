package user

import (
	"context"
	"github.com/farzanehshahi/user-kit/internal/entity"
)

type Service interface {
	Create(ctx context.Context, reqUser *entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	Update(ctx context.Context, id string, username, password string) error
	Delete(ctx context.Context, id string) error
}

// repository contracts
type Repository interface {
	Create(ctx context.Context, req *entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	Update(ctx context.Context, id string, updatedUser *entity.User) error
	Delete(ctx context.Context, id string) error
}
