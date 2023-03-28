package post

import (
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/internal/domain/post"
)

type postStorage struct {
}

func (ps *postStorage) Create(post *post.Post) (*post.Post, error) {
	return nil, nil
}
func (ps *postStorage) Delete(post *post.Post) error {
	return nil
}
func (ps *postStorage) GetAll() []*post.Post {
	return nil
}
func (ps *postStorage) GetUsersPosts(id int) []*post.Post {
	return nil
}
