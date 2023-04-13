package delivery

import app "user/usecase"

func CreateGrpcServer(command app.Commands, query app.Queries) GrpcServer {
	return GrpcServer{
		command: command,
		query:   query,
	}
}
