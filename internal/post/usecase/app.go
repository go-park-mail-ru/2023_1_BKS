package app

import (
	"post/usecase/command"
	"post/usecase/query"
)

//В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	CreatePost command.CreatePostHandler
	UpdatePost command.UpdatePostHandler
	DeletePost command.DeletePostHandler
	ClosePost  command.ClosePostHandler
	AddCart    command.AddCartHandler
	RemoveCart command.RemoveCartHandler
}

//В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetIdPost      query.GetIdPostHandler
	GetSortNewPost query.GetSortNewPostHandler
	GetCart        query.GetCartHandler
}
