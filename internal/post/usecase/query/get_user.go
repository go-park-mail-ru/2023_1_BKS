package query

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetIdUserHandler struct {
	postRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetIdUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.User, error) {
	return h.postRepo.Get(ctx, id)
}
