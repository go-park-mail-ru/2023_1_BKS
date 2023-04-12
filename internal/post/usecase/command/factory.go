package command

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewCreatePostHandler(
	postRepo domain.CUDRepositoryPost,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreatePostHandler {
	return CreatePostHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewUpdatePostHandler(
	postRepo domain.CUDRepositoryPost,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdatePostHandler {
	return UpdatePostHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewDeletePostHandler(
	postRepo domain.CUDRepositoryPost,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeletePostHandler {
	return DeletePostHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewClosePostHandler(
	postRepo domain.CUDRepositoryPost,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) ClosePostHandler {
	return ClosePostHandler{postRepo: postRepo, validator: validator, loger: loger}
}

func NewAddCartHandler(
	cartRepo domain.CUDRepositoryCart,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) AddCartHandler {
	return AddCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}

func NewRemoveCartHandler(
	cartRepo domain.CUDRepositoryCart,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) RemoveCartHandler {
	return RemoveCartHandler{cartRepo: cartRepo, validator: validator, loger: loger}
}
