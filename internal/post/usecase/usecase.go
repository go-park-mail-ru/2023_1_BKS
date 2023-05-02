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
	AddCart    command.AddCartHandler
	RemoveCart command.RemoveCartHandler
}

// В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetIdPost          query.GetIdHandler
	GetUserIdOpenPost  query.GetByUserIdOpenHandler
	GetUserIdClosePost query.GetByUserIdCloseHandler
	GetSortNewPost     query.GetSortNewHandler
	GetTagPost         query.GetTagHandler
	GetCart            query.GetCartHandler
}
