package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type RRepository interface {
	GetId(ctx context.Context, id uuid.UUID) (User, error)
}
