package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DeletePostHandler struct {
	postRepo  domain.CUDRepositoryPost
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *DeletePostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.postRepo.Delete(ctx, id)
	return err
}

type RemoveCartHandler struct {
	cartRepo  domain.CUDRepositoryCart
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *RemoveCartHandler) Handle(
	ctx context.Context,
	cart domain.Cart,
) error {
	err := h.cartRepo.Remove(ctx, cart)
	return err
}
