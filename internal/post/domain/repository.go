package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepository interface {
	Create(ctx context.Context, post Post) error
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, id uuid.UUID) error
	Close(ctx context.Context, id uuid.UUID) error
}

type RRepository interface {
	GetId(ctx context.Context, id uuid.UUID) (Post, error)
	GetSortNew(ctx context.Context, number uint) (Post, error)
}
