package domain

import (
	"image"

	"github.com/google/uuid"
)

func CreateUser(
	email,
	login,
	password,
	firstName,
	secondName,
	patronimic string,
	avatar image.RGBA,
) (User, []error) {
	/*var allErr []error

	if err := EmailAndSpecification.IsValid(email); err != nil {
		allErr = append(allErr, err...)
	}

	if err := LoginAndSpecification.IsValid(login); err != nil {
		allErr = append(allErr, err...)
	}

	if err := PassAndSpecification.IsValid(password); err != nil {
		allErr = append(allErr, err...)
	}

	if err := FirstNameAndSpecification.IsValid(firstName); err != nil {
		allErr = append(allErr, err...)
	}

	if err := SecondNameAndSpecification.IsValid(secondName); err != nil {
		allErr = append(allErr, err...)
	}

	if err := PatronimicAndSpecification.IsValid(patronimic); err != nil {
		allErr = append(allErr, err...)
	}

	if err := AvatarAndSpecification.IsValid(avatar); err != nil {
		allErr = append(allErr, err...)
	}*/

	return User{
		Id:       uuid.New(),
		Email:    CreateEmail(email),
		Login:    CreateLogin(login),
		Password: CreatePassword(password),
		FullName: CreateFullName(secondName, firstName, patronimic),
		Avatar:   CreateAvatar(avatar),
	}, nil
}
