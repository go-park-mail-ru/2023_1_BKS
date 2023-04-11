package repository

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
)

type PostCacheRepository struct {
	posts map[uuid.UUID]domain.Post
}

func (t *PostCacheRepository) New() {
	t.posts = make(map[uuid.UUID]domain.Post)
}

func (t PostCacheRepository) GetId(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	return t.posts[id], nil
}

func (t PostCacheRepository) GetTitle(ctx context.Context, title domain.Title) (domain.Post, error) {
	for i, j := range t.posts {
		if j.Title == title {
			return t.posts[i], nil
		}
	}
	return domain.Post{}, nil
}

func (t *PostCacheRepository) Create(ctx context.Context, posts domain.Post) error {
	t.posts[posts.Id] = posts
	return nil
}

func (t *PostCacheRepository) Update(ctx context.Context, posts domain.Post) error {
	t.posts[posts.Id] = posts
	return nil
}

func (t *PostCacheRepository) Delete(ctx context.Context, id uuid.UUID) error {
	delete(t.posts, id)
	return nil
}
