package repository

import (
	"context"
	"user/domain"

	"github.com/google/uuid"
)

type UserCacheRepository struct {
	users map[uuid.UUID]domain.User
}

func (t UserCacheRepository) GetId(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return &t.users[id], nil
}

func (t UserCacheRepository) GetLogin(ctx context.Context, login domain.Login) (*domain.User, error) {
	for i, j := range t.users {
		if j.Login == login {
			return &t.users[i], nil
		}
	}
	return &domain.User{}, nil
}

func (t UserCacheRepository) GetEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	for i, j := range t.users {
		if j.Email == email {
			return &t.users[i], nil
		}
	}
	return &domain.User{}, nil
}

func (t *UserCacheRepository) Create(ctx context.Context, users domain.User) error {
	t.users[users.Id] = users
	return nil
}

func (t *UserCacheRepository) Update(ctx context.Context, users domain.User) error {
	t.users[users.Id] = users
	return nil
}

func (t *UserCacheRepository) Delete(ctx context.Context, users domain.User) error {
	delete(t.users, users.id)
	return nil
}
