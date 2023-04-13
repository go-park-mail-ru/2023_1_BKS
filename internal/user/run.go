package user

import (
	"config"
	"context"
	"fmt"
	"os"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v2 "user/delivery/http/v2"
	"user/usecase"
)

func Run(cfg config.Config) {

	ctx := context.Background()
	swagger, err := v2.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка загрузки спецификации swagger\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	command, query := usecase.NewApplication(ctx, cfg)

	serverHandler := v2.CreateHttpServer(command, query)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(oapimiddleware.OapiRequestValidator(swagger))
	v2.RegisterHandlers(e, &serverHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)))
}
