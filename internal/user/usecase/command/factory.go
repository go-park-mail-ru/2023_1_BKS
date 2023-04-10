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

func NewUpdateUserHandler(
	userRepo domain.CUDRepository,
	loger *logrus.Entry,
) UpdateUserHandler {
	return UpdateUserHandler{userRepo: userRepo, loger: loger}
}

func NewDeleteUserHandler(
	userRepo domain.CUDRepository,
	loger *logrus.Entry,
) DeleteUserHandler {
	return DeleteUserHandler{userRepo: userRepo, loger: loger}
}
