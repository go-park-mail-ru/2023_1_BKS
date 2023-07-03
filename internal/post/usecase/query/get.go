package query

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"

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
) (domain.Post, domain.WrapperError) {
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
) ([]domain.Post, domain.WrapperError) {
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
) ([]domain.Post, domain.WrapperError) {
	return h.postRepo.GetByUserIdClose(ctx, idUser, number)
}

type GetSortNewHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetSortNewHandler) Handle(
	ctx context.Context,
	number int,
) ([]domain.Post, domain.WrapperError) {
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
) ([]domain.Post, domain.WrapperError) {
	return h.postRepo.GetByTag(ctx, tag, number)
}

type GetCartHandler struct {
	cartRepo domain.RCartRepository
	loger    *logrus.Entry
}

func (h GetCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
) ([]uuid.UUID, domain.WrapperError) {
	return h.cartRepo.Get(ctx, userId)
}

type GetFavoriteHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetFavoriteHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
) ([]uuid.UUID, domain.WrapperError) {
	return h.postRepo.GetFavorite(ctx, userId)
}

type GetByArrayHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetByArrayHandler) Handle(
	ctx context.Context,
	postId []uuid.UUID,
) ([]domain.Post, domain.WrapperError) {
	return h.postRepo.GetByArray(ctx, postId)
}

type SearchPostHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h SearchPostHandler) Handle(
	ctx context.Context,
	search string,
) ([]uuid.UUID, domain.WrapperError) {
	return h.postRepo.Search(ctx, search)
}
