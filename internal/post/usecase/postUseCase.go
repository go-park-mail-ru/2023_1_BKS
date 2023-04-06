package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
)

type IPostUseCase interface {
	CreatePost(context.Context, *domain.Post) (string, error)
	GetPostByID(context.Context, string) (*domain.Post, error)
	GetPostsByUserID(context.Context, int) ([]*domain.Post, error)
	EditPost(context.Context, *domain.Post) (string, error)
	DeletePost(context.Context, string) (string, error)
}
