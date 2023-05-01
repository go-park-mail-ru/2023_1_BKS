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

type GetByUserIdOpenHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetByUserIdOpenHandler) Handle(
	ctx context.Context,
	idUser uuid.UUID,
	number int,
) ([]domain.Post, error) {
	return h.postRepo.GetByUserIdOpen(ctx, idUser, number)
}

type GetByUserIdCloseHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetByUserIdCloseHandler) Handle(
	ctx context.Context,
	idUser uuid.UUID,
	number int,
) ([]domain.Post, error) {
	return h.postRepo.GetByUserIdClose(ctx, idUser, number)
}

type GetSortNewHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetSortNewHandler) Handle(
	ctx context.Context,
	number int,
) ([]domain.Post, error) {
	return h.postRepo.GetSortNew(ctx, number)
}

type GetTagHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetTagHandler) Handle(
	ctx context.Context,
	tag string,
	number int,
) ([]domain.Post, error) {
	return h.postRepo.GetByTag(ctx, tag, number)
}
