package delivery

import app "auth/usecase"

func CreateGrpcServer(command app.Commands, query app.Queries) GrpcServer {
	return GrpcServer{
		command: command,
		query:   query,
	}
}
