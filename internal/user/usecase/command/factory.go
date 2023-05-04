package command

import (
	"github.com/go-park-mail-ru/2023_1_BKS/internal/user/domain"

	"github.com/sirupsen/logrus"
)

func NewCreateUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreateUserHandler {
	return CreateUserHandler{userRepo: userRepo,
		validator: validator, loger: loger}
}

func NewUpdateUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdateUserHandler {
	return UpdateUserHandler{userRepo: userRepo,
		validator: validator, loger: loger}
}

func NewDeleteUserHandler(
	userRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeleteUserHandler {
	return DeleteUserHandler{userRepo: userRepo, validator: validator, loger: loger}
}
