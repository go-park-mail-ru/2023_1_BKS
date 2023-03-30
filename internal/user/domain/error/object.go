package error

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

type FirstNameErr struct{}

func (e *FirstNameErr) Error() string {
	return "Имя не было введено"
}

type SecondNameErr struct{}

func (e *SecondNameErr) Error() string {
	return "Фамилия не была введена"
}
