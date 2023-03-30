package repository

import (
	"image"

	"github.com/2023_1_BKS/internal/app/model"
	"github.com/2023_1_BKS/internal/app/store"
)

type PostRepository struct {
	store *store.Store
}

func (r *PostRepository) CreatePost(p *model.Post) (*model.Post, error) {
	panic("unimplemented!")
}

func (r *PostRepository) DeletePost(id int) error {
	panic("unimplemented!")
}

func (r *PostRepository) FindByID(id int) (*model.Post, error) {
	panic("unimplemented!")
}

func (r *PostRepository) FindByHeader(header string) (*model.Post, error) {
	panic("unimplemented!")
}

func (r *PostRepository) AddImage(image image.RGBA) error {
	panic("unimplemented!")
}
