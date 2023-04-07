package run

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"user/repository"
	app "user/usecase"
	"user/usecase/command"
	"user/usecase/query"
)

func NewApplication(ctx context.Context) app.Application {
	/*postgresClient, err :=
	if err != nil {
		panic(err)
	}
	*/

	userRepository := repository.NewPostgresUserRepository(postgresClient, hourFactory)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(userRepository, logger),
			/*UpdateUser: command.NewUpdateUserHandler(userRepository, logger),
			DeleteUser: command.NewDeleteUserHandler(userRepository, logger),*/
		},
		Queries: app.Queries{
			GetUser: query.NewGetUserHandler(userRepository, logger),
		},
	}
}

func main() {
	//logs.Init()

	ctx := context.Background()

	application := NewApplication(ctx)

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		//go loadFixtures(application)

		server.RunHTTPServer(func(router echo.Router) http.Handler {
			return ports.HandlerFromMux(
				ports.NewHttpServer(application),
				router,
			)
		})
	/*case "grpc":
	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		trainer.RegisterTrainerServiceServer(server, svc)
	})*/
	default:
		panic(fmt.Sprintf("Тип сервера '%s' не поддерживается", serverType))
	}
}
