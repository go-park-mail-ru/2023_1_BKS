package query

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetIdPostHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetIdPostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.Post, error) {
	return h.postRepo.Get(ctx, id)
}
