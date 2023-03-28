package post

// Controller

type Storage interface {
	Create(post *Post) (*Post, error)
	Delete(post *Post) error
	GetAll() []*Post
	GetUsersPosts(id int) []*Post
}
