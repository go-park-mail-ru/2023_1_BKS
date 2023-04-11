package query

import (
	"post/domain"

	"github.com/sirupsen/logrus"
)

func NewGetIdUserHandler(
	postRepo domain.RRepository,
	loger *logrus.Entry,
) GetIdUserHandler {
	return GetIdUserHandler{postRepo: postRepo, loger: loger}
}
