package command

import (
	"context"
	"user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateUserHandler struct {
	userRepo domain.CUDRepository
	loger    *logrus.Entry
}

func (h *UpdateUserHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
	email,
	login,
	phoneNumber,
	secondName,
	firstName,
	patronimic,
	password,
	passwordCheck,
	avatar string,
) error {
	if password != passwordCheck {
		return domain.PassNonComporableErr{}
	}
	user := domain.User{
		Email:       domain.CreateEmail(email),
		Login:       domain.CreateLogin(login),
		PhoneNumber: domain.CreatePhoneNumber(phoneNumber),
		Password:    domain.CreatePassword(phoneNumber),
		FullName:    domain.CreateFullName(secondName, firstName, patronimic),
		Avatar:      domain.CreateAvatar(avatar),
	}
	err := h.userRepo.Update(ctx, user)
	return err
}
