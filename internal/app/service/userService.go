package service

import (
	"github.com/2023_1_BKS/internal/app/model"
	rep "github.com/2023_1_BKS/internal/app/repository"
)

type IUserUseCase interface {
	Register(user *model.User) (string, error)
	Login() (string, error)
	GetUserByToken(string) (*model.User, error)
	Logout(token string) error
}

type UserUseCase struct {
	userRepo rep.IUserRepository
}
