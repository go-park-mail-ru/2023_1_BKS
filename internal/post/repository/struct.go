package repository

import (
	"image"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
	"github.com/google/uuid"
)

type PostPostgreSQL struct {
	ID     string
	UserID int
	Title  string
	Body   string
	Image  []image.RGBA
}

func (p *PostPostgreSQL) ToEntity() (domain.Post, error) {
	if p == nil {
		return domain.Post{}, nil
	}

	parsed, err := uuid.Parse(p.ID)
	if err != nil {
		return domain.Post{}, err
	}

	return domain.Post{
		ID:     parsed,
		UserID: p.UserID,
		Title:  domain.Title(p.Title),
		Body:   domain.Body(p.Body),
		Image:  domain.Image(p.Image),
	}, nil
}
