package post

import "context"

// UseCase

type Service interface {
	GetPostsByUserID(ctx context.Context, id int) []*Post
	GetAllPosts(ctx context.Context) []*Post
	CreatePost(ctx context.Context, dto *CreatePostDTO) *Post
}

type service struct {
	storage Storage
}

func NewService(storage Storage) service {
	return service{storage: storage}
}

func (s *service) GetPostByUserID(ctx context.Context, id int) []*Post {
	return s.storage.GetUsersPosts(id)
}

func (s *service) GetAllPosts(ctx context.Context) []*Post {
	return s.storage.GetAll()
}

func (s *service) CreatePost(ctx context.Context, dto *CreatePostDTO) *Post {
	return nil
}
