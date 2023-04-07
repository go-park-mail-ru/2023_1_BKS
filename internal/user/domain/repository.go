package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}

type RRepository interface {
	GetId(ctx context.Context, id uuid.UUID) (*User, error)
	GetEmail(ctx context.Context, email Email) (*User, error)
	GetLogin(ctx context.Context, login Login) (*User, error)
}
