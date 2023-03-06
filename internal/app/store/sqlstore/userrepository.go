package sqlstore

import (
	"database/sql"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/store"
)

type UserRepository struct {
	store *Store
}

// Создаем пользователя
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return nil
	}

	if err := u.BeforeCreate(); err != nil {
		return nil
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (first_name, second_name, email, password, card_number, card_cvv, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		u.First_name, u.Second_name, u.Email, u.EncryptedPassword, u.Card_number, u.Card_cvv, nil,
	).Scan(&u.Id)
}

// Ищем пользователя
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, first_name, second_name, email, password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.First_name,
		&u.Second_name,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
