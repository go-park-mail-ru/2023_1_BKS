package user

type UserUseCase interface {
	CompPassword(l_pass, r_pass string) error
	ValidatePassword(pass string) []error
	ValidateLogin(login string) []error
}
