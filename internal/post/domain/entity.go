package domain

import (
	"github.com/google/uuid"
)

type Post struct {
	Id         uuid.UUID
	UserID     uuid.UUID
	Images     Images
	Desciption Desciption
	Title      Title
	Tags       Tags
}
