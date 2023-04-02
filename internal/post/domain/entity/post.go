package entity

import (
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain/object"
	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID
	UserUUID object.UserID
	Title    object.Title
	Body     object.Body
	Image    object.Image
}
