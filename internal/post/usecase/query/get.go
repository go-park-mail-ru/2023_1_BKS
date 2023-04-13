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

type GetSortNewHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetSortNewHandler) Handle(
	ctx context.Context,
	number uint,
) (domain.Post, error) {
	return h.postRepo.GetSortNew(ctx, number)
}
