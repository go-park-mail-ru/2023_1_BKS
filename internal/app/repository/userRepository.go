package repository

import (
	"database/sql"
	"image"

	"github.com/2023_1_BKS/internal/app/model"
)

type UserRepository struct {
	db *sql.DB
}

type IUserRepository interface {
	CreateUser(u *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(id int) (*model.User, error)
	FindByFN(fullname [3]string) (*model.User, error)
	FindByLogin(login string) (*model.User, error)
	EditUser(u *model.User) error
	AddImage(image image.RGBA) error
	DeleteUser(id int) error
	CheckPassword(password string) (bool, error)
	CheckLogin(login string) (bool, error)
}
