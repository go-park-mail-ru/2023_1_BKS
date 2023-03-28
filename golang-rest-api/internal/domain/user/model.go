package user

import (
	"image"

	"github.com/google/uuid"
)

// Entity

type User struct {
	id       uuid.UUID
	fullname [3]string
	login    string
	email    string
	password []byte
	avatar   image.RGBA
}
