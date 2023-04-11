package repository

import (
	"context"
	"database/sql"
	"post/domain"

	"github.com/google/uuid"
)

type UserPostgressRepository struct {
	posts *sql.DB
}

func (t UserPostgressRepository) Get(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	var (
		email       string
		login       string
		phoneNumber string
		secondName  string
		firstName   string
		patronimic  string
		avatar      string
	)

	row := t.posts.QueryRow("SELECT email, login, phonenumber, secondname, firstname, patronimic, avatar FROM posts WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &login, &phoneNumber, &secondName, &firstName, &patronimic, &avatar)
	return domain.Post{
		Id:          id,
		Email:       domain.CreateEmail(email),
		Login:       domain.CreateLogin(login),
		PhoneNumber: domain.CreatePhoneNumber(phoneNumber),
		FullName:    domain.CreateFullName(secondName, firstName, patronimic),
		Avatar:      domain.CreateAvatar(avatar),
	}, err
}

func (t *UserPostgressRepository) Create(ctx context.Context, post domain.User) error {
	_, err := t.posts.Exec("insert into posts (id, email,  phonenumber, login, password, firstname, secondname, patronimic, avatar) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		post.Id, post.Email, post.PhoneNumber, post.Login, post.Password, post.FullName.FirstName(),
		post.FullName.SecondName(), post.FullName.Patronimic(), post.Avatar)
	// Должен возвращаться или ссесия или jwt токен
	return err
}

func (t *UserPostgressRepository) Update(ctx context.Context, post domain.User) error {
	id, _ := uuid.Parse("978137d3-a263-4dc7-9308-43e35c0c83ff") // Тут должго быть получение значений из авторизированного пользователя
	_, err := t.posts.Exec("update posts set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9",
		post.Email, post.PhoneNumber, post.Login, post.Password, post.FullName.FirstName(),
		post.FullName.SecondName(), post.FullName.Patronimic(), post.Avatar, id)
	return err
}

func (t *UserPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.posts.Exec("update posts set email = $1,  phonenumber = $2, login = $3, password = $4, firstname = $5, secondname = $6, patronimic = $7, avatar = $8 where id = $9", id)
	//logout
	return err
}
