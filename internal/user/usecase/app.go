package app

import (
	"user/usecase/command"
	"user/usecase/query"
)

//В данном агрегате перечисленны все команды сервиса пользователя
type Commands struct {
	CreateUser command.CreateUserHandler
	UpdateUser command.UpdateUserHandler
	DeleteUser command.DeleteUserHandler
}

//В данном агрегате перечисленны все запросы сервиса пользователя
type Queries struct {
	GetUser query.GetIdUserHandler
}
