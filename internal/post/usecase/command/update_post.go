package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateUserHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
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

	if err := h.validator.Email.IsValid(email); err != nil {
		return err
	}
	if err := h.validator.Login.IsValid(login); err != nil {
		return err
	}
	if err := h.validator.PhoneNumber.IsValid(phoneNumber); err != nil {
		return err
	}
	if err := h.validator.FirstName.IsValid(firstName); err != nil {
		return err
	}
	if err := h.validator.SecondName.IsValid(secondName); err != nil {
		return err
	}
	if err := h.validator.Patronimic.IsValid(patronimic); err != nil {
		return err
	}
	if err := h.validator.Password.IsValid(password); err != nil {
		return err
	}
	if err := h.validator.Avatar.IsValid(avatar); err != nil {
		return err
	}

	post := domain.User{
		Email:       domain.CreateEmail(email),
		Login:       domain.CreateLogin(login),
		PhoneNumber: domain.CreatePhoneNumber(phoneNumber),
		Password:    domain.CreatePassword(phoneNumber),
		FullName:    domain.CreateFullName(secondName, firstName, patronimic),
		Avatar:      domain.CreateAvatar(avatar),
	}
	err := h.postRepo.Update(ctx, post)
	return err
}
