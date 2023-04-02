package entity

import "image"

type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
	Image  []image.RGBA
}
