package app

//Корневой агрегат для доступа к бизнес-логике сервиса авторизации
type Application struct {
	Commands Commands
	Queries  Queries
}

//В данном агрегате перечисленны все команды сервиса авторизации
type Commands struct {
}

//В данном агрегате перечисленны все запросы сервиса авторизации
type Queries struct {
}
