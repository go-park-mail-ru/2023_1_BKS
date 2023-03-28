package user

import (
	"image"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	fullName [3]string
	login    string
	password []byte
	avatar   image.RGBA
}
