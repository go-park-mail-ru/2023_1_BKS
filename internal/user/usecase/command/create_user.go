package command

import (
	"context"
	"user/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateUserHandler struct {
	userRepo domain.CUDRepository
	loger    *logrus.Entry
}

func (h *CreateUserHandler) Handle(
	ctx context.Context,
	email string,
	login string,
	phoneNumber string,
	secondName string,
	firstName string,
	patronimic string,
	password string,
	passwordCheck string,
	avatar string,
) error {
	user := domain.User{
		Id:          uuid.New(),
		Email:       domain.CreateEmail(email),
		Login:       domain.CreateLogin(login),
		PhoneNumber: domain.CreatePhoneNumber(phoneNumber),
		Password:    domain.CreatePassword(phoneNumber),
		FullName:    domain.CreateFullName(secondName, firstName, patronimic),
		Avatar:      domain.Avatar(avatar),
	}
	err := h.userRepo.Create(ctx, user)
	return err
}
