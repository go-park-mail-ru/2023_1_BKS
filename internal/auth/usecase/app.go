package usecase

import "github.com/go-park-mail-ru/2023_1_BKS/internal/auth/usecase/commands"

// В данном агрегате перечисленны все команды сервиса авторизации
type Commands struct {
	CreateToken commands.CreateJWSHandle
}

// В данном агрегате перечисленны все запросы сервиса авторизации
type Queries struct {
}
