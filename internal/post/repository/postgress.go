package repository

import (
	"context"
	"database/sql"
	"post/domain"
	"time"

	"github.com/google/uuid"
)

type PostPostgressRepository struct {
	posts *sql.DB
}

func (t PostPostgressRepository) GetId(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	var (
		userID      uuid.UUID
		title       string
		description string
		price       string
		close       bool
		time        time.Time
		tags        []string
		pathImages  []string
	)

	row := t.posts.QueryRow("SELECT userid, title, description, price, close, tags, images, time FROM posts WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&userID, &title, &description, &price, &close, &tags, &pathImages, &time)
	return domain.Post{
		Id:         id,
		UserID:     userID,
		Title:      title,
		Desciption: description,
		Price:      price,
		Tags:       tags,
		PathImages: pathImages,
		Time:       time,
	}, err
}

func (t PostPostgressRepository) GetSortNew(ctx context.Context, number uint) (domain.Post, error) {
	var (
		id          uuid.UUID
		userID      uuid.UUID
		title       string
		description string
		price       string
		close       bool
		time        time.Time
		tags        []string
		images      []string
	)

	row := t.posts.QueryRow("SELECT id, userid, title, description, price, close, tags, images, time FROM posts ORDERED BY time LIMIT $1, 1", number-1)
	err := row.Scan(&id, &userID, &title, &description, &price, &close, &tags, &images, &time)
	return domain.Post{
		Id:         id,
		UserID:     userID,
		Title:      title,
		Desciption: description,
		Price:      price,
		Tags:       tags,
		PathImages: images,
		Time:       time,
	}, err
}

func (t *PostPostgressRepository) Create(ctx context.Context, post domain.Post) error {
	_, err := t.posts.Exec("insert into posts (id, userid, title, description, price, close, tags, images, time) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		post.Id, post.UserID, post.Title, post.Desciption, post.Price, post.Close,
		post.Tags, post.PathImages, post.Time)
	return err
}

func (t *PostPostgressRepository) Update(ctx context.Context, post domain.Post) error {
	id, _ := uuid.Parse("978137d3-a263-4dc7-9308-43e35c0c83ff") // Тут должго быть получение значений из авторизированного пользователя
	_, err := t.posts.Exec("update posts set  userid = $1, title = $2, description = $3, price = $4, close = $5, tags = $6, images = $7, time = $8 where id = $9",
		post.UserID, post.Title, post.Desciption, post.Price, post.Close,
		post.Tags, post.PathImages, post.Time, id)
	return err
}

func (t *PostPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.posts.Exec("delete from posts where id = $1", id)
	return err
}

func (t *PostPostgressRepository) Close(ctx context.Context, id uuid.UUID) error {
	close := true
	_, err := t.posts.Exec("update posts set close = $1 where id = &2", close, id)
	return err
}
