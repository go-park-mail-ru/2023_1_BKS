package store

import "github.com/go-park-mail-ru/2023_1_BKS/internal/app/model"

type UserRepository struct {
	store *Store
}

// Создаем пользователя
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (first_name, second_name, email, password, card_number, card_cvv, image) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		u.First_name, u.Second_name, u.Email, u.EncryptedPassword, u.Card_number, u.Card_cvv, nil,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return u, nil
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
		return nil, err
	}

	return u, nil
}
