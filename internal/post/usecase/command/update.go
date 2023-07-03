package command

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdateHandler) Handle(
	ctx context.Context,
	postDelivery domain.Post,
) domain.WrapperError {
	err := h.postRepo.Update(ctx, postDelivery)
	return err
}

type CloseHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CloseHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) domain.WrapperError {
	err := h.postRepo.Close(ctx, id)
	return err
}
