package app

import (
	"post/usecase/command"
	"post/usecase/query"
)

// В данном агрегате перечисленны все команды сервиса объявлений
type Commands struct {
	CreatePost command.CreatePostHandler
	UpdatePost command.UpdatePostHandler
	DeletePost command.DeletePostHandler
}

// В данном агрегате перечисленны все запросы сервиса объявлений
type Queries struct {
	GetPost query.GetIdPostHandler
}
