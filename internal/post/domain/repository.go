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
	AddFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error
	RemoveFavorite(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error
}

type RRepository interface {
	GetId(ctx context.Context, id uuid.UUID) (Post, error)
	GetByUserIdOpen(ctx context.Context, idUser uuid.UUID, number int) ([]Post, error)
	GetByUserIdClose(ctx context.Context, idUser uuid.UUID, number int) ([]Post, error)
	GetByTag(ctx context.Context, tag string, number int) ([]Post, error)
	GetSortNew(ctx context.Context, number int) ([]Post, error)
	GetFavorite(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error)
	GetByArray(ctx context.Context, postId []uuid.UUID) ([]Post, error)
	Search(ctx context.Context, search string) ([]uuid.UUID, error)
}

type CUDCartRepository interface {
	Add(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error
	Remove(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error
}

type RCartRepository interface {
	Get(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error)
}
