package query

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetIdPostHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetIdPostHandler {
	return GetIdPostHandler{postRepo: postRepo, loger: loger}
}
