package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DeletePostHandler struct {
	postRepo  domain.CUDRepository
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
