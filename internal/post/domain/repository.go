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
	GetByUserId(ctx context.Context, idUser uuid.UUID, number int) (Post, error)
	GetSortNew(ctx context.Context, number int) (Post, error)
}
