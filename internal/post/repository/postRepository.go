package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
	"github.com/google/uuid"
)

type PostPostgreSQL struct { // Связь с базой данных
	connection *sql.DB
}

type PostRepository interface {
	FindPostByID(ctx context.Context, ID uuid.UUID) (*domain.Post, error)
	FindPostsByTitle(ctx context.Context, title string) (*domain.Post, error)
	SavePost(ctx context.Context, post domain.Post) (*domain.Post, error)
	DeletePost(ctx context.Context, post domain.Post) (*domain.Post, error)
	UpdatePost(ctx context.Context, ID uuid.UUID) (*domain.Post, error)
}

func (p *PostPostgreSQL) ToEntity() (domain.Post, error) {
	return domain.Post{}, nil
}
