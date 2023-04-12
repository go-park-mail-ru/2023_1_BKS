package post

import (
	"context"
	"fmt"
	"os"
	"post/config"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/go-redis/redis"
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
	rdo := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	validator := domain.CreateSpecificationManager(cfg)

	postRepository := repository.CreatePostgressRepository(dsn)
	redisRepository := repository.CartRedisRepository(rdo)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return app.Commands{
			CreatePost: command.NewCreatePostHandler(&postRepository, validator, logger),
			UpdatePost: command.NewUpdatePostHandler(&postRepository, validator, logger),
			DeletePost: command.NewDeletePostHandler(&postRepository, validator, logger),
			ClosePost:  command.NewClosePostHandler(&postRepository, validator, logger),
			AddCart:    command.NewAddCartHandler(&postRepository, validator, logger),
			RemoveCart: command.NewRemoveCartHandler(&postRepository, validator, logger),
		},
		app.Queries{
			GetIdPost:      query.NewGetIdPostHandler(postRepository, logger),
			GetSortNewPost: query.NewGetSortNewPostHandler(postRepository, logger),
			GetCart:        query.NewGetCartHandler(),
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
