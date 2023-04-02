package object

func CreateEmail(email string) (Email, []error) {
	var err []error
	/*
		Тут должна быть валидация
	*/
	return Email(email), err
}

func CreateLogin(login string) (Login, []error)

func CreatePassword(password string) (Password, []error)

func CreateFullName(second_name, first_name, patronimic string) (FullName, []error)

func CreateCompanyName(type_company, name string) (CompanyName, []error)
