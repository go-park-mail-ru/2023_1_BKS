package app

import (
	"post/usecase/command"
	"post/usecase/query"
)

//В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	Create command.CreateHandler
	Update command.UpdateHandler
	Delete command.DeleteHandler
	Close  command.CloseHandler
}

//В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetId      query.GetIdHandler
	GetSortNew query.GetSortNewHandler
}
