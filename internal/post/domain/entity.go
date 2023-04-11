package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Email       Email
	Login       Login
	PhoneNumber PhoneNumber
	Password    Password
	FullName    FullName
	Avatar      Avatar
}
