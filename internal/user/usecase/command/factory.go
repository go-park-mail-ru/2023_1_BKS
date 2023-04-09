package command

import (
	"user/domain"

	"github.com/sirupsen/logrus"
)

func NewCreateUserHandler(
	userRepo domain.CUDRepository,
	loger *logrus.Entry,
) CreateUserHandler {
	return CreateUserHandler{userRepo: userRepo, loger: loger}
}
