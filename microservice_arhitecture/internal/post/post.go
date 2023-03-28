package post

import "image"

type Post struct {
	UserID int
	Title  string
	Body   string
	Image  []image.RGBA
}
