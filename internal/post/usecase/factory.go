package usecase

import (
	"context"
	"fmt"
	"post/config"
	"post/domain"
	"post/repository"
	"post/usecase/command"
	"post/usecase/query"

	"github.com/sirupsen/logrus"
)

func NewUsecase(ctx context.Context, cfg config.Config) (Commands, Queries) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)

	validator := domain.CreateSpecificationManager(cfg)

	postRepository := repository.CreatePostgressRepository(dsn)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreatePost: command.NewCreateHandler(&postRepository, validator, logger),
			UpdatePost: command.NewUpdateHandler(&postRepository, validator, logger),
			DeletePost: command.NewDeleteHandler(&postRepository, validator, logger),
			ClosePost:  command.NewCloseHandler(&postRepository, validator, logger),
		},
		Queries{
			GetIdPost:      query.NewGetIdHandler(postRepository, logger),
			GetSortNewPost: query.NewGetSortNewHandler(postRepository, logger),
			GetUserIdPost:  query.NewGetByUserIdHandler(postRepository, logger),
		}
}
