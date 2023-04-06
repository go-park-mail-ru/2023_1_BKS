package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/repository"
	"github.com/google/uuid"
)

type PostUseCase struct {
	postRepo repository.PostConnection
}

func (p *PostUseCase) CreatePost(ctx context.Context, post *domain.Post) (string, error) {
	newPost, err := p.postRepo.SavePost(ctx, *post)
	if err != nil {
		return "", err
	}
	return newPost.ID.String(), nil
}

func (p *PostUseCase) GetPostByID(ctx context.Context, postID string) (*domain.Post, error) {
	id, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}
	post, err := p.postRepo.FindPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostUseCase) EditPost(ctx context.Context, post *domain.Post) (string, error) {
	_, err := p.postRepo.FindPostByID(ctx, post.ID)
	if err != nil {
		return "", err
	}
	editedPost, err := p.postRepo.UpdatePost(ctx, *post)
	if err != nil {
		return "", err
	}
	return editedPost.ID.String(), nil
}

func (p *PostUseCase) DeletePost(ctx context.Context, postID string) (string, error) {
	id, err := uuid.Parse(postID)
	if err != nil {
		return "", err
	}
	if err = p.postRepo.DeletePost(ctx, id); err != nil {
		return "", err
	}
	return id.String(), nil
}
