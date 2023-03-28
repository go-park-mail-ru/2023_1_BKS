# Domain

```
type Email string
```
```
func сreateEmail(email string) Email
```
```
type Login string
```
```
func сreateLogin(login string) Login
```

```
type Password []byte
```
```
func createPassword(password string) Password
```
```
type FullName [3]string
```
```
func createFullName(first_name, second_name, patronimic string) FullName
```

```
type User struct {
    id          UU
    fullname    FullName
    email       Email
    login       Login
    password    Password
    avatar      image.RGBA
}
```
Где доступен: Данный класс доступен на всех слоях.

Обозначения: id - уникальный id пользователя, fullname - ФИО, email - почта, login - логин, password - пароль в зашифрованом виде, avatar - аватар пользователя.

Требования: Для изменения состояний используем методы, для получения значения дергаем нужное поле(сеттеры есть, геттеров нет).

```
type DomainFactory interface {
	CreateEmail(email string) Email
	CreateLogin(login string) Login
    CreatePassword(password string) Password
    CreateFullName(first_name, second_name, patronimic string) FullName
}
```







# UseCase

```
func CompPassword(l_pass, r_pass string) error
```

Где доступен: Данный класс доступен на всех слоях ниже usecase.

Описание: Данная функция сравнивет пароли, при совпадении возвращает nil, в остальных случаях ошибку.

Обозначения: l_pass - первый пароль, r_pass - второй пароль.

Требования: Проверка осуществляется перед шифрованием пароля.

```
func ValidatePassword(pass string) []error
```

Где доступен: Данный класс доступен на всех слоях ниже usecase.

Описание: Данная функция проверяет требования в написании необходимых символов пароля(Спец символы, большие буквы и.т.д.) и количество символов.

Обозначения: pass - пароль.

Требования: Проверка осуществляется перед шифрованием пароля.

```
func ValidateLogin(login string) []error
```

Где доступен: Данный класс доступен на всех слоях ниже usecase.

Описание: Данная функция проверяет требования в написании необходимых символов логина(Спец символы, большие буквы и.т.д.) и количество символов.

Обозначения: login - логин.


# Delivery

https://rietta.com/blog/bcrypt-not-sha-for-passwords/

Шифрование пароля будет осуществляться модулем bcrypt.


# Repository


```
type UserRepository struct {
    database *sql.DB
}

```

```
func (rep *UserRpository) Open(conf Config)  error
```

```
func (rep *UserRpository) Close() error
```

```
func (rep *UserRpository) GetUserById(id int) (User, error)
```

```
func (rep *UserRpository) GetUserByFIO(first_name, last_name, patronym string) (User, error)
```

```
func (rep *UserRpository) GetUserByLogin(login string) (User, error)
```

```
func (rep *UserRpository) GetUserByEmail(email string) (User, error)
```

```
func (rep *UserRpository) InsertUser(user User) error
```

```
func (rep *UserRpository) EditUser(user User) error
```

```
func (rep *UserRpository) DeleteUser(id int) error
```