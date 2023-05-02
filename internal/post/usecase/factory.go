package usecase

import (
	"context"
	"post/config"
	"post/domain"
	"post/repository"
	"post/usecase/command"
	"post/usecase/query"

	"github.com/sirupsen/logrus"
)

func NewUsecase(ctx context.Context, cfg config.Config) (Commands, Queries) {
	validator := domain.CreateSpecificationManager(cfg)

	cartRepository := repository.CreateRedisRepository(cfg)
	postRepository := repository.CreatePostgressRepository(cfg)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreatePost: command.NewCreateHandler(&postRepository, validator, logger),
			UpdatePost: command.NewUpdateHandler(&postRepository, validator, logger),
			DeletePost: command.NewDeleteHandler(&postRepository, validator, logger),
			ClosePost:  command.NewCloseHandler(&postRepository, validator, logger),
			AddCart:    command.NewAddCartHandler(&cartRepository, validator, logger),
			RemoveCart: command.NewRemoveCartHandler(&cartRepository, validator, logger),
		},
		Queries{
			GetIdPost:          query.NewGetIdHandler(postRepository, logger),
			GetSortNewPost:     query.NewGetSortNewHandler(postRepository, logger),
			GetUserIdOpenPost:  query.NewGetByUserIdOpenHandler(postRepository, logger),
			GetUserIdClosePost: query.NewGetByUserIdCloseHandler(postRepository, logger),
			GetTagPost:         query.NewGetByTagHandler(postRepository, logger),
			GetCart:            query.NewGetCartHandler(&cartRepository, logger),
		}
}
