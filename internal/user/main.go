package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	v2 "user/delivery/http/v2"
	"user/repository"
	app "user/usecase"
	"user/usecase/command"
	"user/usecase/query"
)

func NewApplication(ctx context.Context) (app.Commands, app.Queries) {
	/*postgresClient, err :=
	if err != nil {
		panic(err)
	}
	*/
	var userRepository repository.UserCacheRepository
	userRepository.New()
	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Commands{
			CreateUser: command.NewCreateUserHandler(&userRepository, logger),
			/*UpdateUser: command.NewUpdateUserHandler(userRepository, logger),
			DeleteUser: command.NewDeleteUserHandler(userRepository, logger),*/
		},
		app.Queries{
			GetUser: query.NewGetIdUserHandler(userRepository, logger),
		}
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	ctx := context.Background()
	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	command, query := NewApplication(ctx)

	serverHandler := v2.CreateHttpServer(command, query)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	v2.RegisterHandlers(e, &serverHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))

}
