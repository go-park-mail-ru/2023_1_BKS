package app

import (
	"post/usecase/command"
	"post/usecase/query"
)

//В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	CreateUser command.CreateUserHandler
	UpdateUser command.UpdateUserHandler
	DeleteUser command.DeleteUserHandler
}

//В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetUser query.GetIdUserHandler
}
