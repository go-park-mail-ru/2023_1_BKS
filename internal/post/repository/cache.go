package repository

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
)

type UserCacheRepository struct {
	posts map[uuid.UUID]domain.User
}

func (t *UserCacheRepository) New() {
	t.posts = make(map[uuid.UUID]domain.User)
}

func (t UserCacheRepository) GetId(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return t.posts[id], nil
}

func (t UserCacheRepository) GetLogin(ctx context.Context, login domain.Login) (domain.User, error) {
	for i, j := range t.posts {
		if j.Login == login {
			return t.posts[i], nil
		}
	}
	return domain.User{}, nil
}

func (t UserCacheRepository) GetEmail(ctx context.Context, email domain.Email) (domain.User, error) {
	for i, j := range t.posts {
		if j.Email == email {
			return t.posts[i], nil
		}
	}
	return domain.User{}, nil
}

func (t *UserCacheRepository) Create(ctx context.Context, posts domain.User) error {
	t.posts[posts.Id] = posts
	return nil
}

func (t *UserCacheRepository) Update(ctx context.Context, posts domain.User) error {
	t.posts[posts.Id] = posts
	return nil
}

func (t *UserCacheRepository) Delete(ctx context.Context, id uuid.UUID) error {
	delete(t.posts, id)
	return nil
}
