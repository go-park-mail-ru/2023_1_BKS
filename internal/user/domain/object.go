package domain

import "image"

type Email string

type Login string

//Используется для хранения пароля в зашифрованном ввиде.
type Password []byte

//Используется для хранения аккаунтов людей.
//Первое поле - Фамилия, второе поле - Имя, третье поле - Отчество.
type FullName [3]string

//Используется для хранения аккаунтов компаний
//Первое поле - Тип компании, второе поле - Название.
type CompanyName [2]string

//Используется для хранения адрессов
type Adress string

//Используется для хранения изображений
type Avatar image.Image
