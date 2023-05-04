package v2

import app "github.com/go-park-mail-ru/2023_1_BKS/internal/auth/usecase"

func CreateHttpServer(command app.Commands, query app.Queries) HttpServer {
	return HttpServer{
		command: command,
		query:   query,
	}
}
