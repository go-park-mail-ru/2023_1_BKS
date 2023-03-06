package teststore

import (
	"errors"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return u, nil
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return nil
	}

	if err := u.BeforeCreate(); err != nil {
		return nil
	}

	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	// u, ok := r.users[email]
	// if !ok {
	// 	return nil, errors.New("not found")
	// }

	return nil, errors.New("Record not found")
}
