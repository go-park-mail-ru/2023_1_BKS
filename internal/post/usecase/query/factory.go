package query

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetIdPostHandler(
	postRepo domain.RRepositoryPost,
	loger *logrus.Entry,
) GetIdPostHandler {
	return GetIdPostHandler{postRepo: postRepo, loger: loger}
}

func NewGetSortNewPostHandler(
	postRepo domain.RRepositoryPost,
	loger *logrus.Entry,
) GetSortNewPostHandler {
	return GetSortNewPostHandler{postRepo: postRepo, loger: loger}
}

func NewGetCartHandler(
	cartRepo domain.RRepositoryCart,
	loger *logrus.Entry,
) GetCartHandler {
	return GetCartHandler{cartRepo: cartRepo, loger: loger}
}
