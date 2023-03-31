package usecase

import "github.com/2023_1_BKS/internal/app/user"

type UserUseCase struct {
	userRepo user.IUserUsecase
}
