package domain

import (
	"regexp"
	"unicode"
)

type typeSpecification interface {
	string | rune
}

type Specification[T typeSpecification] interface {
	IsValid(verifiable T) error
}

type PassSpecification struct {
	specifications map[string]Specification[string]
}

func (e PassSpecification) IsValid(pass string) error {
	if err := e.specifications["AND"].IsValid(pass); err != nil {
		return err
	}
	if err := e.specifications["OR"].IsValid(pass); err != nil {
		return err
	}
	return nil
}

type PassAndSpecification struct {
	specifications map[string]Specification[string]
}

func (e PassAndSpecification) IsValid(pass string) error {
	for _, specification := range e.specifications {
		if err := specification.IsValid(pass); err != nil {
			return err
		}
	}
	return nil
}

type PassLengthValidation struct {
	maxLength uint
	minLength uint
}

func (e PassLengthValidation) IsValid(pass string) error {
	if e.minLength < uint(len(pass)) && e.maxLength < uint(len(pass)) {
		return PassMaxLengthErr{e.maxLength}
	}
	if e.minLength > uint(len(pass)) && e.maxLength > uint(len(pass)) {
		return PassMinLengthErr{e.minLength}
	}
	return nil
}

type PassForRuneValidation struct {
	specifications map[string]Specification[rune]
}

func (e PassForRuneValidation) IsValid(pass string) error {
	var errSC error
	var errUC error
	var errLC error
	for _, char := range pass {
		errSC = e.specifications["PassSpecialCharacters"].IsValid(char)
		errUC = e.specifications["PassUpperCase"].IsValid(char)
		errLC = e.specifications["PassLowerCase"].IsValid(char)
	}
	if errSC != nil {
		return errSC
	}
	if errUC != nil {
		return errUC
	}
	if errLC != nil {
		return errLC
	}
	return nil

}

type passSpecialCharactersValidation struct {
	specialChar map[rune]bool
}

func (e passSpecialCharactersValidation) IsValid(char rune) error {
	if e.specialChar[char] {
		return nil
	}
	return PassSpecialCharactersErr{e.specialChar}
}

type passUpperCaseValidation struct{}

func (e passUpperCaseValidation) IsValid(char rune) error {
	if unicode.IsUpper(char) {
		return nil
	}
	return PassUpperCaseErr{}
}

type passLowerCaseValidation struct{}

func (e passLowerCaseValidation) IsValid(char rune) error {
	if unicode.IsLower(char) {
		return nil
	}
	return PassLowerCaseErr{}
}

type EmailAndSpecification struct {
	specifications map[string]Specification[string]
}

type EmailRequiredValueValidation struct{}

func (e EmailRequiredValueValidation) IsValid(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return EmailRequiredValueErr{}
	} else {
		return nil
	}
}

type EmailLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e EmailLengthValidation) IsValid(email string) error {
	if e.minLength < uint(len(email)) && e.maxLength < uint(len(email)) {
		return EmailMaxLengthErr{e.maxLength}
	}
	if e.minLength > uint(len(email)) && e.maxLength > uint(len(email)) {
		return EmailMinLengthErr{e.minLength}
	}
	return nil
}

type PhoneNumberSpecification interface {
	IsValid(phone string) error
}

type PhoneNumberOrSpecification struct {
	specifications map[string]Specification[string]
}

type PhoneNumberAndSpecification struct {
	specifications map[string]Specification[string]
}

type PhoneNumberRussianRequiredValueValidation struct{}

func (e PhoneNumberRussianRequiredValueValidation) IsValid(phone string) error {
	emailRegex := regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`)
	if !emailRegex.MatchString(phone) {
		return EmailRequiredValueErr{}
	} else {
		return nil
	}
}

type PhoneNumberLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e PhoneNumberLengthValidation) IsValid(phone string) error {
	if e.minLength < uint(len(phone)) && e.maxLength < uint(len(phone)) {
		return EmailMaxLengthErr{e.maxLength}
	}
	if e.minLength > uint(len(phone)) && e.maxLength > uint(len(phone)) {
		return EmailMinLengthErr{e.minLength}
	}
	return nil
}

type LoginAndSpecification struct {
	specifications map[string]Specification[string]
}

type LoginAcceptableValuesValidation struct {
	NonAcceptableValues map[rune]bool
}

func (e LoginAcceptableValuesValidation) IsValid(email string) error {
	for char := range e.NonAcceptableValues {
		if e.NonAcceptableValues[char] {
			return nil
		}
	}
	return LoginAcceptableValuesErr{e.NonAcceptableValues}
}

type LoginLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e LoginLengthValidation) IsValid(login string) error {
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return LoginMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return LoginMinLengthErr{e.minLength}
	}
	return nil
}

type FirstNameAndSpecification struct {
	specifications map[string]Specification[string]
}

type FirstNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e FirstNameLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return FirstNameMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return FirstNameMinLengthErr{e.minLength}
	}
	return nil
}

type SecondNameAndSpecification struct {
	specifications map[string]Specification[string]
}

type SecondNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e SecondNameLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return SecondNameMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return SecondNameMinLengthErr{e.minLength}
	}
	return nil
}

type PatronimicAndSpecification struct {
	specifications map[string]Specification[string]
}

type PatronimicLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e PatronimicLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return PatronimicMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return PatronimicMinLengthErr{e.minLength}
	}
	return nil
}

type CompanyTypeAndSpecification struct {
	specifications map[string]Specification[string]
}

type CompanyTypeValidation struct {
	companyType map[string]bool
}

func (e CompanyTypeValidation) IsValid(companyType string) error {
	if e.companyType[companyType] {
		return nil
	}
	return CompanyTypeErr{e.companyType}
}

type CompanyNameAndSpecification struct {
	specifications map[string]Specification[string]
}

type CompanyNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e CompanyNameLengthValidation) IsValid(name string) error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return CompanyNameMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return CompanyNameMinLengthErr{e.minLength}
	}
	return nil
}

type AdressAndSpecification struct {
	specifications map[string]Specification[string]
}

type AdressLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e AdressLengthValidation) IsValid(adress string) error {
	if e.minLength < uint(len(adress)) && e.maxLength < uint(len(adress)) {
		return AdressMaxLengthErr{e.maxLength}
	}
	if e.minLength < uint(len(adress)) && e.maxLength < uint(len(adress)) {
		return AdressMinLengthErr{e.minLength}
	}
	return nil
}

type AvatarAndSpecification struct {
	specifications map[string]Specification[string]
}

type AvatarWightValidation struct {
	maxImageSize uint
}

func (e AvatarWightValidation) IsValid(avatar string) error {
	return nil
}
