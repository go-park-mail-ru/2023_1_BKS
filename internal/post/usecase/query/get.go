package query

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetIdPostHandler struct {
	postRepo domain.RRepositoryPost
	loger    *logrus.Entry
}

func (h GetIdPostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.Post, error) {
	return h.postRepo.GetId(ctx, id)
}

type GetSortNewPostHandler struct {
	postRepo domain.RRepositoryPost
	loger    *logrus.Entry
}

func (h GetSortNewPostHandler) Handle(
	ctx context.Context,
	number uint,
) (domain.Post, error) {
	return h.postRepo.GetSortNew(ctx, number)
}

type GetCartHandler struct {
	cartRepo domain.RRepositoryCart
	loger    *logrus.Entry
}

func (h GetCartHandler) Handle(
	ctx context.Context,
) ([]domain.Cart, error) {
	return h.cartRepo.Get(ctx)
}
