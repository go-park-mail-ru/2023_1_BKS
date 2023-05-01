package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id         uuid.UUID
	UserID     uuid.UUID
	Close      bool
	Title      string
	Desciption string
	Price      string
	Tags       string
	PathImages []string
	Time       time.Time
	Views      int
}
