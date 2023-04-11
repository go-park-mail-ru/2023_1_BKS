package command

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewCreateUserHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreateUserHandler {
	return CreateUserHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewUpdateUserHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdateUserHandler {
	return UpdateUserHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewDeleteUserHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeleteUserHandler {
	return DeleteUserHandler{postRepo: postRepo, validator: validator, loger: loger}
}
