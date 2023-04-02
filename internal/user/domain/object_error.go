package domain

import (
	"fmt"
	"strconv"
)

type PassLengthErr struct {
	length int
}

func (e *PassLengthErr) Error() string {
	return fmt.Sprintf("Минимальная длина пароля: %d\n", e.length)
}

type PassSpecialCharactersErr struct {
	specialChar []rune
}

func (e *PassSpecialCharactersErr) Error() string {
	var str string
	for _, char := range e.specialChar {
		str = fmt.Sprintf(str, strconv.QuoteRune(char), ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Пароль не содержит специальных символов из данного списка: %s\n", str)
}

type PassUpperCaseErr struct{}

func (e *PassUpperCaseErr) Error() string {
	return "Пароль не содержит заглавные буквы"
}

type PassLowerCaseErr struct{}

func (e *PassLowerCaseErr) Error() string {
	return "Пароль не содержит строчные буквы"
}

type EmailRequiredValueErr struct{}

func (e *EmailRequiredValueErr) Error() string {
	return "Введите валидный email"
}

type LoginAcceptableValuesErr struct {
	NonAcceptableValues []rune
}

func (e *LoginAcceptableValuesErr) Error() string {
	var str string
	for _, char := range e.NonAcceptableValues {
		str = fmt.Sprintf(str, strconv.QuoteRune(char), ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Логин содержит недопустимые символы: %s\n", str)
}

type FirstNameLengthErr struct{}

func (e *FirstNameLengthErr) Error() string {
	return "Имя не было введено"
}

type SecondNameLengthErr struct{}

func (e *SecondNameLengthErr) Error() string {
	return "Фамилия не была введена"
}

type CompanyTypeErr struct {
	typeComp []string
}

func (e *CompanyTypeErr) Error() string {
	var str string
	for _, typeComp := range e.typeComp {
		str = fmt.Sprintf(str, typeComp, ", ")
	}
	str = str[0 : len(str)-3]
	return fmt.Sprintf("Неправильные тип компании. Введите допустимый из данного списка: %s\n", str)
}

type CompanyNameLengthErr struct{}

func (e *CompanyNameLengthErr) Error() string {
	return "Название компании не введено"
}

type AdressRequiredValueErr struct{}

func (e *AdressRequiredValueErr) Error() string {
	return "Введён не валидный адресс"
}
