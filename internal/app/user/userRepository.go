package user

import "image"

type IUserRepository interface {
	CreateUser(u *User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(id int) (*User, error)
	FindByFN(fullname [3]string) (*User, error)
	FindByLogin(login string) (*User, error)
	EditUser(u *User) error
	AddImage(image image.RGBA) error
	DeleteUser(id int) error
	CheckPassword(password string) (bool, error)
	CheckLogin(login string) (bool, error)
}
