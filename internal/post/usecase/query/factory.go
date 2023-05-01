package query

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetIdHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetIdHandler {
	return GetIdHandler{postRepo: postRepo, loger: loger}
}

func NewGetSortNewHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetSortNewHandler {
	return GetSortNewHandler{postRepo: postRepo, loger: loger}
}

func NewGetByUserIdHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetByUserIdHandler {
	return GetByUserIdHandler{postRepo: postRepo, loger: loger}
}
