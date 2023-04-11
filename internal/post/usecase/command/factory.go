package command

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewCreatePostHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) CreatePostHandler {
	return CreatePostHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewUpdatePostHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) UpdatePostHandler {
	return UpdatePostHandler{postRepo: postRepo,
		validator: validator, loger: loger}
}

func NewDeletePostHandler(
	postRepo domain.CUDRepository,
	validator domain.SpecificationManager,
	loger *logrus.Entry,
) DeletePostHandler {
	return DeletePostHandler{postRepo: postRepo, validator: validator, loger: loger}
}
