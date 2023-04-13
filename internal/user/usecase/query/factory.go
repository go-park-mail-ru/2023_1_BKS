package query

import (
	"user/domain"

	"github.com/sirupsen/logrus"
)

func NewGetIdUserHandler(
	userRepo domain.RRepository,
	loger *logrus.Entry,
) GetIdUserHandler {
	return GetIdUserHandler{userRepo: userRepo, loger: loger}
}

func NewCheckUserHandler(
	userRepo domain.RRepository,
	loger *logrus.Entry,
) CheckUserHandler {
	return CheckUserHandler{userRepo: userRepo, loger: loger}
}
