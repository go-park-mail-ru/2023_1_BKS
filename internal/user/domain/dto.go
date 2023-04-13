package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	Email       string
	PhoneNumber string

	Login         string
	Password      string
	CheckPassword string

	SecondName string
	FirstName  string
	Patronimic string

	PathToAvatar string
}
