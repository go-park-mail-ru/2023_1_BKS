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

func NewGetByUserIdOpenHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetByUserIdOpenHandler {
	return GetByUserIdOpenHandler{postRepo: postRepo, loger: loger}
}

func NewGetByUserIdCloseHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetByUserIdCloseHandler {
	return GetByUserIdCloseHandler{postRepo: postRepo, loger: loger}
}

func NewGetByTagHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetTagHandler {
	return GetTagHandler{postRepo: postRepo, loger: loger}
}

func NewGetCartHandler(
	cartRepo domain.RCartRepository,
	loger *logrus.Entry,
) GetCartHandler {
	return GetCartHandler{cartRepo: cartRepo, loger: loger}
}

func NewGetFavoriteHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetFavoriteHandler {
	return GetFavoriteHandler{postRepo: postRepo, loger: loger}
}

func NewGetByArrayHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetByArrayHandler {
	return GetByArrayHandler{postRepo: postRepo, loger: loger}
}

func NewSearchPostHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) SearchPostHandler {
	return SearchPostHandler{postRepo: postRepo, loger: loger}
}
