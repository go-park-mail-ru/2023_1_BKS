package app

import (
	"user/usecase/command"
	"user/usecase/query"
)

//Корневой агрегат для доступа к бизнес-логике сервиса пользователя
type Application struct {
	Commands Commands
	Queries  Queries
}

//В данном агрегате перечисленны все команды сервиса пользователя
type Commands struct {
	CreateUser command.CreateUserHandler
}

//В данном агрегате перечисленны все запросы сервиса пользователя
type Queries struct {
	GetUser query.GetUserHandler
}
