package query

import (
	"context"
	"user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type GetIdUserHandler struct {
	userRepo domain.RRepository
	loger    *logrus.Entry
}

func (h GetIdUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) (domain.User, error) {
	return h.userRepo.GetId(ctx, id)
}
