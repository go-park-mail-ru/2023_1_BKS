package repository

import (
	"context"
	"database/sql"
	"post/domain"

	"github.com/google/uuid"
)

type PostPostgressRepository struct {
	posts *sql.DB
}

func (t PostPostgressRepository) Get(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	var (
		title       string
		description string
		images      []string
		tags        []string
	)

	row := t.posts.QueryRow("SELECT title, description, images, tags FROM posts WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&title, &description, &images, &tags)
	return domain.Post{
		Id:         id,
		Title:      domain.CreateTitle(title),
		Desciption: domain.CreateDescription(description),
		Images:     domain.CreateImages(images),
		Tags:       domain.CreateTags(tags),
	}, err
}

func (t *PostPostgressRepository) Create(ctx context.Context, post domain.Post) error {
	_, err := t.posts.Exec("insert into posts (id, title,  description, images, tags) values ($1, $2, $3, $4, $5)",
		post.Id, post.Title, post.Desciption, post.Images, post.Tags)
	// Должен возвращаться или ссесия или jwt токен
	return err
}

func (t *PostPostgressRepository) Update(ctx context.Context, post domain.Post) error {
	id, _ := uuid.Parse("978137d3-a263-4dc7-9308-43e35c0c83ff") // Тут должго быть получение значений из авторизированного пользователя
	_, err := t.posts.Exec("update posts set title = $1,  description = $2, images = $3, tags = $4 where id = $5",
		post.Title, post.Desciption, post.Images, post.Tags, id)
	return err
}

func (t *PostPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.posts.Exec("delete from posts where id=$1", id)
	//logout
	return err
}
