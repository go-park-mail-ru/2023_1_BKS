package repository

import (
	"context"
	"database/sql"
	"user/domain"

	"github.com/google/uuid"
)

type UserPostgressRepository struct {
	users *sql.DB
}

func (t UserPostgressRepository) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		email       string
		login       string
		phoneNumber string
		secondName  string
		firstName   string
		patronimic  string
		avatar      string
	)

	row := t.users.QueryRow("SELECT email, login, phonenumber, secondname, firstname, patronimic, avatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &login, &phoneNumber, &secondName, &firstName, &patronimic, &avatar)
	return domain.User{
		Id:          id,
		Email:       domain.CreateEmail(email),
		Login:       domain.CreateLogin(login),
		PhoneNumber: domain.CreatePhoneNumber(phoneNumber),
		FullName:    domain.CreateFullName(secondName, firstName, patronimic),
		Avatar:      domain.CreateAvatar(avatar),
	}, err
}

func (t *UserPostgressRepository) Create(ctx context.Context, user domain.User) error {
	_, err := t.users.Exec("insert into users (id, email,  phonenumber, login, password, firstname, secondname, patronimic, avatar) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		user.Id, user.Email, user.PhoneNumber, user.Login, user.Password, user.FullName.FirstName(),
		user.FullName.SecondName(), user.FullName.Patronimic(), user.Avatar)
	// Должен возвращаться или ссесия или jwt токен
	return err
}

func (t *UserPostgressRepository) Update(ctx context.Context, user domain.User) error {
	id, _ := uuid.Parse("978137d3-a263-4dc7-9308-43e35c0c83ff") // Тут должго быть получение значений из авторизированного пользователя
	_, err := t.users.Exec("update users set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9",
		user.Email, user.PhoneNumber, user.Login, user.Password, user.FullName.FirstName(),
		user.FullName.SecondName(), user.FullName.Patronimic(), user.Avatar, id)
	return err
}

func (t *UserPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.users.Exec("update users set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9", id)
	//logout
	return err
}
