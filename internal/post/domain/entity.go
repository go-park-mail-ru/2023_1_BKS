package domain

import (
	"github.com/google/uuid"
)

type Post struct {
	Id         uuid.UUID
	UserID     uuid.UUID
	Close      Close
	Title      Title
	Desciption Desciption
	Price      Price
	Tags       Tags
	Images     Images
	Time       TimeStamp
}
