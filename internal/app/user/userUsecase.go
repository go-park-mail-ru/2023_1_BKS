package user

type IUserUsecase interface {
	Register(user *User) (string, error)
	Login() (string, error)
	GetUserByToken(string) (*User, error)
	Logout(token string) error
}
