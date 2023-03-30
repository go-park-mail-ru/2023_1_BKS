package model

import "image"

type User struct {
	ID       int
	Email    string
	Password string
	Fullname [3]string
	Login    string
	Wallet   string
	Image    image.RGBA
}
