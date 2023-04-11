package domain

type Email string

func (l Email) String() string {
	return string(l)
}

type Login string

func (l Login) String() string {
	return string(l)
}

//Используется для хранения пароля в зашифрованном ввиде.
type Password []byte

func (l Password) String() string {
	return string(l)
}

type PhoneNumber string

func (l PhoneNumber) String() string {
	return string(l)
}

//Используется для хранения аккаунтов людей.
//Первое поле - Фамилия, второе поле - Имя, третье поле - Отчество.
type FullName [3]string

func (l FullName) SecondName() string {
	return l[0]
}

func (l FullName) FirstName() string {
	return l[1]
}

func (l FullName) Patronimic() string {
	return l[2]
}

//Используется для хранения аккаунтов компаний
//Первое поле - Тип компании, второе поле - Название.
type CompanyName [2]string

func (l CompanyName) Type() string {
	return l[0]
}

func (l CompanyName) Name() string {
	return l[1]
}

//Используется для хранения адрессов
type Adress string

func (l Adress) String() string {
	return string(l)
}

//Используется для хранения изображений
type Avatar string

func (l Avatar) String() string {
	return string(l)
}
