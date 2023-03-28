package user

type UserRepository interface {
	GetUserById(id int) (User, error)
	GetUserByLogin(login string) (User, error)
	GetUserByEmail(email string) (User, error)
	InsertUser(user User) error
	Delete(id int) error
}
