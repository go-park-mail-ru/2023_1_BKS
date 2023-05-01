package usecase

import (
	"post/usecase/command"
	"post/usecase/query"
)

// В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	CreatePost command.CreateHandler
	UpdatePost command.UpdateHandler
	DeletePost command.DeleteHandler
	ClosePost  command.CloseHandler
}

// В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetIdPost      query.GetIdHandler
	GetUserIdPost  query.GetByUserIdHandler
	GetSortNewPost query.GetSortNewHandler
}
