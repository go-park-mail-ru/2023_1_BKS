package query

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetIdHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetIdHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.Post, error) {
	return h.postRepo.GetId(ctx, id)
}

type GetByUserIdHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetByUserIdHandler) Handle(
	ctx context.Context,
	idUser uuid.UUID,
	number int,
) (domain.Post, error) {
	return h.postRepo.GetByUserId(ctx, idUser, number)
}

type GetSortNewHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetSortNewHandler) Handle(
	ctx context.Context,
	number int,
) (domain.Post, error) {
	return h.postRepo.GetSortNew(ctx, number)
}
