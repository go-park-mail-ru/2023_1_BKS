package domain

func CreateEmail(email string) Email {
	return Email(email)
}

func CreateLogin(login string) Login {
	return Login(login)
}

func CreatePassword(password string) Password {
	return Password(password)
}

func CreatePhoneNumber(phoneNumber string) PhoneNumber {
	return PhoneNumber(phoneNumber)
}

func CreateFullName(secondName, firstName, patronimic string) FullName {
	stringMassive := [3]string{secondName, firstName, patronimic}
	return FullName(stringMassive)
}

func CreateCompanyName(typeCompany, name string) CompanyName {
	stringMassive := [2]string{typeCompany, name}
	return CompanyName(stringMassive)
}

func CreateAvatar(avatar string) Avatar {
	return Avatar(avatar)
}
