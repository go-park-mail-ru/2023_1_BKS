package command

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewCreateHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreateHandler {
	return CreateHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewUpdateHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdateHandler {
	return UpdateHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewDeleteHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeleteHandler {
	return DeleteHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewCloseHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CloseHandler {
	return CloseHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewAddCartHandler(
	cartRepo domain.CUDCartRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) AddCartHandler {
	return AddCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}

func NewRemoveCartHandler(
	cartRepo domain.CUDCartRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) RemoveCartHandler {
	return RemoveCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}
