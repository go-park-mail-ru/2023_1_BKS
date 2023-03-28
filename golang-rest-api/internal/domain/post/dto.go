package post

import "image"

type CreatePostDTO struct {
	UserID int
	Title  string
	Body   string
	Image  []image.RGBA
}

type UpdatePostDTO struct {
	UserID int
	Title  string
	Body   string
	Image  []image.RGBA
}
