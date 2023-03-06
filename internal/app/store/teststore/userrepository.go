package teststore

import (
	"errors"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return nil
	}

	if err := u.BeforeCreate(); err != nil {
		return nil
	}

	r.users[u.Email] = u
	u.Id = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, errors.New("not found")
	}

	return u, nil
}
