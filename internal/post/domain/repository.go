package domain

import (
	"context"

	"github.com/google/uuid"
)

type CUDRepositoryPost interface {
	Create(ctx context.Context, post Post) error
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, id uuid.UUID) error
	Close(ctx context.Context, id uuid.UUID) error
}

type RRepositoryPost interface {
	GetId(ctx context.Context, id uuid.UUID) (Post, error)
	GetSortNew(ctx context.Context, number uint) (Post, error)
}

type CUDRepositoryCart interface {
	Add(ctx context.Context, cart Cart) error
	Remove(ctx context.Context, cart Cart) error
}

type RRepositoryCart interface {
	Get(ctx context.Context) ([]Cart, error)
}
