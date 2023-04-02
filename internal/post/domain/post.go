package domain

import (
	"github.com/google/uuid"
)

type Post struct {
	ID     uuid.UUID
	UserID int
	Title  Title
	Body   Body
	Image  Image
}
