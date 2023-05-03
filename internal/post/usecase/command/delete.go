package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DeleteHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *DeleteHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.postRepo.Delete(ctx, id)
	return err
}

type RemoveCartHandler struct {
	cartRepo  domain.CUDCartRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *RemoveCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) error {
	err := h.cartRepo.Remove(ctx, userId, postId)
	return err
}

type RemoveFavoriteHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *RemoveFavoriteHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) error {
	err := h.postRepo.RemoveFavorite(ctx, userId, postId)
	return err
}
