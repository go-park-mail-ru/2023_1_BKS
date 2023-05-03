package v2

import app "auth/usecase"

func CreateHttpServer(command app.Commands, query app.Queries) HttpServer {
	return HttpServer{
		command: command,
		query:   query,
	}
}
