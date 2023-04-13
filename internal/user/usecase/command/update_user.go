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
	if err := h.validator.Email.IsValid(userDelivery.Email); err != nil {
		return err
	}
	if err := h.validator.Login.IsValid(userDelivery.Login); err != nil {
		return err
	}
	if err := h.validator.PhoneNumber.IsValid(userDelivery.PhoneNumber); err != nil {
		return err
	}
	if err := h.validator.FirstName.IsValid(userDelivery.FirstName); err != nil {
		return err
	}
	if err := h.validator.SecondName.IsValid(userDelivery.SecondName); err != nil {
		return err
	}
	if err := h.validator.Patronimic.IsValid(userDelivery.Patronimic); err != nil {
		return err
	}
	if err := h.validator.Password.IsValid(userDelivery.Password); err != nil {
		return err
	}
	if err := h.validator.Avatar.IsValid(userDelivery.PathToAvatar); err != nil {
		return err
	}

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
