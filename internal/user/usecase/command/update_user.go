package command

import (
	"context"
	"user/domain"

	"github.com/sirupsen/logrus"
)

type UpdateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdateUserHandler) Handle(
	ctx context.Context,
	userDelivery domain.User,
) error {

	user := domain.User{
		Id: userDelivery.Id,

		Email:       userDelivery.Email,
		PhoneNumber: userDelivery.PhoneNumber,

		Login:    userDelivery.Login,
		Password: userDelivery.Password,

		SecondName: userDelivery.SecondName,
		FirstName:  userDelivery.FirstName,
		Patronimic: userDelivery.Patronimic,

		PathToAvatar: userDelivery.PathToAvatar,
	}
	err := h.userRepo.Update(ctx, user)
	return err
}
