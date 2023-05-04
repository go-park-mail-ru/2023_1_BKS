package delivery

import app "github.com/go-park-mail-ru/2023_1_BKS/internal/user/usecase"

func CreateGrpcServer(command app.Commands, query app.Queries) GrpcServer {
	return GrpcServer{
		command: command,
		query:   query,
	}
}
