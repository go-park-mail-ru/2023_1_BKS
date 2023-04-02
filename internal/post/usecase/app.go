package app

//Корневой агрегат для доступа к бизнес-логике сервиса постов
type Application struct {
	Commands Commands
	Queries  Queries
}

//В данном агрегате перечисленны все команды сервиса постов
type Commands struct {
}

//В данном агрегате перечисленны все запросы сервиса постов
type Queries struct {
}
