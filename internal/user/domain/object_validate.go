package domain

import (
	"image"
	"regexp"
	"unicode"
)

type PassSpecification interface {
	IsValid(password string) []error
}

type PassAndSpecification struct {
	specifications []PassSpecification
}

func (e PassAndSpecification) IsValid(pass string) []error {
	var errors []error
	for _, specification := range e.specifications {
		if err := specification.IsValid(pass); err != nil {
			errors = append(errors, err...)
		}
	}
	return errors
}

type PassLengthValidation struct {
	maxLength uint
	minLength uint
}

func (e PassLengthValidation) IsValid(pass string) []error {
	if e.minLength < uint(len(pass)) && e.maxLength < uint(len(pass)) {
		return append([]error{}, PassMaxLengthErr{e.maxLength})
	}
	if e.minLength > uint(len(pass)) && e.maxLength > uint(len(pass)) {
		return append([]error{}, PassMinLengthErr{e.minLength})
	}
	return nil
}

type PassSpecialCharactersValidation struct {
	specialChar map[rune]bool
}

func (e PassSpecialCharactersValidation) IsValid(pass string) []error {
	for _, char := range pass {
		if e.specialChar[char] {
			return nil
		}
	}
	return append([]error{}, PassSpecialCharactersErr{e.specialChar})
}

type PassUpperCaseValidation struct{}

func (e PassUpperCaseValidation) IsValid(pass string) []error {
	for _, char := range pass {
		if !unicode.IsLower(char) {
			return nil
		}
	}
	return append([]error{}, PassUpperCaseErr{})
}

type PassLowerCaseValidation struct{}

func (e PassLowerCaseValidation) IsValid(pass string) []error {
	for _, char := range pass {
		if unicode.IsLower(char) {
			return nil
		}
	}
	return append([]error{}, PassLowerCaseErr{})
}

type EmailSpecification interface {
	IsValid(email string) []error
}

type EmailAndSpecification struct {
	specifications []EmailSpecification
}

type EmailRequiredValueValidation struct{}

func (e EmailRequiredValueValidation) IsValid(email string) []error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return append([]error{}, EmailRequiredValueErr{})
	} else {
		return nil
	}
}

type EmailLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e EmailLengthValidation) IsValid(email string) []error {
	if e.minLength < uint(len(email)) && e.maxLength < uint(len(email)) {
		return append([]error{}, EmailMaxLengthErr{e.maxLength})
	}
	if e.minLength > uint(len(email)) && e.maxLength > uint(len(email)) {
		return append([]error{}, EmailMinLengthErr{e.minLength})
	}
	return nil
}

type LoginSpecification interface {
	IsValid(login string) []error
}

type LoginAndSpecification struct {
	specifications []LoginSpecification
}

type LoginAcceptableValuesValidation struct {
	NonAcceptableValues map[rune]bool
}

func (e LoginAcceptableValuesValidation) IsValid(email string) []error {
	for char := range e.NonAcceptableValues {
		if e.NonAcceptableValues[char] {
			return nil
		}
	}
	return append([]error{}, LoginAcceptableValuesErr{e.NonAcceptableValues})
}

type LoginLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e LoginLengthValidation) IsValid(login string) []error {
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return append([]error{}, LoginMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(login)) && e.maxLength < uint(len(login)) {
		return append([]error{}, LoginMinLengthErr{e.minLength})
	}
	return nil
}

type FirstNameSpecification interface {
	IsValid(firstName string) []error
}

type FirstNameAndSpecification struct {
	specifications []FirstNameSpecification
}

type FirstNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e FirstNameLengthValidation) IsValid(name string) []error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, FirstNameMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, FirstNameMinLengthErr{e.minLength})
	}
	return nil
}

type SecondNameSpecification interface {
	IsValid(secondName string) []error
}

type SecondNameAndSpecification struct {
	specifications []SecondNameSpecification
}

type SecondNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e SecondNameLengthValidation) IsValid(name string) []error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, SecondNameMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, SecondNameMinLengthErr{e.minLength})
	}
	return nil
}

type PatronimicSpecification interface {
	IsValid(patronimic string) []error
}

type PatronimicAndSpecification struct {
	specifications []PatronimicSpecification
}

type PatronimicLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e PatronimicLengthValidation) IsValid(name string) []error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, PatronimicMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, PatronimicMinLengthErr{e.minLength})
	}
	return nil
}

type CompanyTypeSpecification interface {
	IsValid(companyType string) []error
}

type CompanyTypeAndSpecification struct {
	specifications []CompanyTypeSpecification
}

type CompanyTypeValidation struct {
	companyType map[string]bool
}

func (e CompanyTypeValidation) IsValid(companyType string) []error {
	if e.companyType[companyType] {
		return nil
	}
	return append([]error{}, CompanyTypeErr{e.companyType})
}

type CompanyNameSpecification interface {
	IsValid(companyName string) []error
}

type CompanyNameAndSpecification struct {
	specifications []CompanyNameSpecification
}

type CompanyNameLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e CompanyNameLengthValidation) IsValid(name string) []error {
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, CompanyNameMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(name)) && e.maxLength < uint(len(name)) {
		return append([]error{}, CompanyNameMinLengthErr{e.minLength})
	}
	return nil
}

type AdressSpecification interface {
	IsValid(adress string) []error
}

type AdressAndSpecification struct {
	specifications []AdressSpecification
}

type AdressLengthValidation struct {
	minLength uint
	maxLength uint
}

func (e AdressLengthValidation) IsValid(adress string) []error {
	if e.minLength < uint(len(adress)) && e.maxLength < uint(len(adress)) {
		return append([]error{}, AdressMaxLengthErr{e.maxLength})
	}
	if e.minLength < uint(len(adress)) && e.maxLength < uint(len(adress)) {
		return append([]error{}, AdressMinLengthErr{e.minLength})
	}
	return nil
}

type AvatarSpecification interface {
	IsValid(avatar image.RGBA) []error
}

type AvatarAndSpecification struct {
	specifications []AvatarSpecification
}

type AvatarWightValidation struct {
	maxImageSize uint
}

func (e AvatarWightValidation) IsValid(avatar image.Image) []error {

}
