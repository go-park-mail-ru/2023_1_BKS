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

func (t UserPostgressRepository) AuthGetId(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		email        string
		phoneNumber  string
		login        string
		password     string
		secondName   string
		firstName    string
		patronimic   string
		pathToAvatar string
	)

	row := t.users.QueryRow("SELECT email, phonenumber, login, password, secondname, firstname, patronimic, pathavatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &phoneNumber, &login, &password, &secondName, &firstName, &patronimic, &pathToAvatar)

	return domain.User{
		Email:       email,
		PhoneNumber: phoneNumber,

		Login:    login,
		Password: password,

		SecondName: secondName,
		FirstName:  firstName,
		Patronimic: patronimic,

		PathToAvatar: pathToAvatar,
	}, err
}

func (t UserPostgressRepository) GetId(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		email       string
		phoneNumber string

		secondName string
		firstName  string
		patronimic string

		pathToAvatar string
	)

	row := t.users.QueryRow("SELECT email, phonenumber, secondname, firstname, patronimic, pathavatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &phoneNumber, &secondName, &firstName, &patronimic, &pathToAvatar)

	return domain.User{
		Email:       email,
		PhoneNumber: phoneNumber,

		SecondName: secondName,
		FirstName:  firstName,
		Patronimic: patronimic,

		PathToAvatar: pathToAvatar,
	}, err
}

func (t *UserPostgressRepository) Create(ctx context.Context, user domain.User) error {
	_, err := t.users.Exec("insert into users (id, email,  phonenumber, login, password, firstname, secondname, patronimic, avatar) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		user.Id, user.Email, user.PhoneNumber, user.Login, user.Password, user.FirstName,
		user.SecondName, user.Patronimic, user.PathToAvatar)
	// Должен возвращаться или ссесия или jwt токен
	return err
}

func (t *UserPostgressRepository) Update(ctx context.Context, user domain.User) error {
	id, _ := uuid.Parse("978137d3-a263-4dc7-9308-43e35c0c83ff") // Тут должго быть получение значений из авторизированного пользователя
	_, err := t.users.Exec("update users set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9",
		user.Email, user.PhoneNumber, user.Login, user.Password, user.FirstName,
		user.SecondName, user.Patronimic, user.PathToAvatar, id)
	return err
}

func (t *UserPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.users.Exec("update users set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9", id)
	//logout
	return err
}
