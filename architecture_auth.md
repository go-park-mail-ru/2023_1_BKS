## Entity
# Модель пользователя.
```
type User struct {
	ID				int
	FirstName		string
	SecondName		string
	Email			string
	Password		string \\  В зашифрованном виде
	CardNumber		string \\  В зашифрованном виде
	CardCvv			string \\  В зашифрованном виде
	Image			image.RGBA
}
```
# Модель сервера.
```
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}
```
## Repository.
```
type UserRepository interface {
  Create(*entity.User) error
  EditUserData(*entity.User) error
  FindByID(int) (*entity.User, error)
  FindByEmail(string) (*entity.User, error)
}
```
## Методы модели сервера.
```
func (*server) authUser(http.Handler) http.Handler
func (*server) handleUsersCreate() http.HandlerFunc
func (*server) handleSessiaonCreate() http.HandlerFunc
```
## Методы модели пользователя.
```
func (*User) Validate() error
func (*User) ComparePassword(string) bool
func (*User) encryptString(string) (string, error)
```
## Основной принцип работы
В данном файле рассмотрены основные методы и функции для того, чтобы реализовать авторизацию и регистрацию пользователей.
Основной функционал будет реализовываться с помощью интерфейса ```UserRepository```, когда на сервер поступает запрос на создание пользователя, мы будем распаршивать json, который получили в модель пользователя, затем валидация данных и шифрование пароля, с помощью интерфейса ```UserRepository``` отправлять в БД данные.

Когда пользователь будет пытаться авторизоваться, на сервер будет поступать запрос на авторизацию, мы распаршиваем json с полученными данными и пробуем найти пользователя по почте в БД, затем сравниваем пароли и создаем сессию для пользователя.
