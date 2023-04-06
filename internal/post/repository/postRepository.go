package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
	"github.com/google/uuid"
)

type PostConnection struct { // Связь с базой данных
	connection *sql.DB
}

type IPostRepository interface {
	FindPostByID(ctx context.Context, ID uuid.UUID) (*domain.Post, error)
	FindPostsByTitle(ctx context.Context, title string) (*domain.Post, error)
	SavePost(ctx context.Context, post domain.Post) (*domain.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) error
	UpdatePost(ctx context.Context, ID uuid.UUID) (*domain.Post, error)
}

func NewRow(post domain.Post) PostPostgreSQL {
	return PostPostgreSQL{
		ID:     uuid.NewString(),
		UserID: post.UserID,
		Title:  string(post.Title),
		Body:   string(post.Body),
		Image:  post.Image,
	}
}

func (c *PostConnection) FindPostByID(ctx context.Context, ID uuid.UUID) (*domain.Post, error) {
	row := c.connection.QueryRowContext(ctx, "SELECT * FROM posts WHERE id=$1", ID.String())

	post := &PostPostgreSQL{}
	if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.Image); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	entity, err := post.ToEntity()
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (c *PostConnection) FindPostsByTitle(ctx context.Context, title string) ([]domain.Post, error) {
	rows, err := c.connection.QueryContext(ctx, "SELECT * FROM posts WHERE title=$1", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		post := &PostPostgreSQL{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.Image); err != nil {
			return nil, err
		}

		entity, err := post.ToEntity()
		if err != nil {
			return nil, err
		}

		posts = append(posts, entity)
	}

	return posts, nil
}

func (c *PostConnection) SavePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	row := NewRow(post)

	_, err := c.connection.ExecContext(ctx, "INSERT INTO posts (id, user_id, title, body, image) VALUES ($1, $2, $3, $4, $5)",
		row.ID, row.UserID, row.Title, row.Body, row.Image)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (c *PostConnection) DeletePost(ctx context.Context, id uuid.UUID) error {
	_, err := c.connection.ExecContext(ctx, "DELETE FROM posts WHERE id=$1", id.String())
	return err
}

func (c *PostConnection) UpdatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	row := NewRow(post)

	_, err := c.connection.ExecContext(ctx, "UPDATE posts SET user_id=$1, title=$2, body=$3, image=$4 WHERE id=$5",
		row.UserID, row.Title, row.Body, row.Image, row.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
