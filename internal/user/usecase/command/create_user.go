package command

import (
	"context"
	"user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateUserHandler struct {
	userRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreateUserHandler) Handle(
	ctx context.Context,
	passwordCheck string,
	userDelivery domain.User,
) error {
	if userDelivery.Password != passwordCheck {
		return domain.PassNonComporableErr{}
	}

	user := domain.User{
		Id: uuid.New(),

		Email:       userDelivery.Email,
		PhoneNumber: userDelivery.PhoneNumber,

		Login:    userDelivery.Login,
		Password: userDelivery.Password,

		SecondName: userDelivery.SecondName,
		FirstName:  userDelivery.FirstName,
		Patronimic: userDelivery.Patronimic,

		PathToAvatar: userDelivery.PathToAvatar,
	}
	err := h.userRepo.Create(ctx, user)
	return err
}
