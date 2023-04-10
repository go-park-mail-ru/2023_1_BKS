package user

import (
	"config"
	"context"
	"fmt"
	"os"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	v2 "user/delivery/http/v2"
	"user/domain"
	"user/repository"
	app "user/usecase"
	"user/usecase/command"
	"user/usecase/query"
)

func NewApplication(ctx context.Context, cfg config.Config) (app.Commands, app.Queries) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)
	validator := domain.CreateSpecificationManager(cfg)
	userRepository := repository.CreatePostgressRepository(dsn)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Commands{
			CreateUser: command.NewCreateUserHandler(&userRepository, validator, logger),
			UpdateUser: command.NewUpdateUserHandler(&userRepository, validator, logger),
			DeleteUser: command.NewDeleteUserHandler(&userRepository, validator, logger),
		},
		app.Queries{
			GetUser: query.NewGetIdUserHandler(userRepository, logger),
		}
}

func Run(cfg config.Config) {

	ctx := context.Background()
	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	command, query := NewApplication(ctx, cfg)

	serverHandler := v2.CreateHttpServer(command, query)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	v2.RegisterHandlers(e, &serverHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)))
}
