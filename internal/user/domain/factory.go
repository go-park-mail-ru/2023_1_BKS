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

func CreateFullName(second_name, first_name, patronimic string) FullName {
	string_massive := [3]string{second_name, first_name, patronimic}
	return FullName(string_massive)
}

func CreateCompanyName(type_company, name string) CompanyName {
	string_massive := [2]string{type_company, name}
	return CompanyName(string_massive)
}
