package post

import (
	"context"
	"fmt"
	"os"
	"post/config"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	v2 "post/delivery/http/v2"
	"post/domain"
	"post/repository"
	app "post/usecase"
	"post/usecase/command"
	"post/usecase/query"
)

func NewApplication(ctx context.Context, cfg config.Config) (app.Commands, app.Queries) {
	dsn := fmt.Sprintf("post=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)
	validator := domain.CreateSpecificationManager(cfg)
	postRepository := repository.CreatePostgressRepository(dsn)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Commands{
			CreateUser: command.NewCreateUserHandler(&postRepository, validator, logger),
			UpdateUser: command.NewUpdateUserHandler(&postRepository, validator, logger),
			DeleteUser: command.NewDeleteUserHandler(&postRepository, validator, logger),
		},
		app.Queries{
			GetUser: query.NewGetIdUserHandler(postRepository, logger),
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
