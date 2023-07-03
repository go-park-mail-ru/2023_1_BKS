package usecase

import (
	"context"

	config "github.com/go-park-mail-ru/2023_1_BKS/config/post"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/repository"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/usecase/command"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/usecase/query"

	"github.com/sirupsen/logrus"
)

func NewUsecase(ctx context.Context, cfg config.Config) (Commands, Queries) {
	validator := domain.CreateSpecificationManager(cfg)

	cartRepository := repository.CreateRedisRepository(cfg)
	postRepository := repository.CreatePostgressRepository(cfg)

	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreatePost:     command.NewCreateHandler(&postRepository, validator, logger),
			UpdatePost:     command.NewUpdateHandler(&postRepository, validator, logger),
			DeletePost:     command.NewDeleteHandler(&postRepository, validator, logger),
			ClosePost:      command.NewCloseHandler(&postRepository, validator, logger),
			AddFavorite:    command.NewAddFavoriteHandler(&postRepository, validator, logger),
			RemoveFavorite: command.NewRemoveFavoriteHandler(&postRepository, validator, logger),
			AddCart:        command.NewAddCartHandler(&cartRepository, validator, logger),
			RemoveCart:     command.NewRemoveCartHandler(&cartRepository, validator, logger),
		},
		Queries{
			GetIdPost:   query.NewGetIdHandler(postRepository, logger),
			GetMiniPost: query.NewGetMiniPostHandler(postRepository, logger),
			GetCart:     query.NewGetCartHandler(&cartRepository, logger),
			GetFavorite: query.NewGetFavoriteHandler(postRepository, logger),
			GetByArray:  query.NewGetByArrayHandler(postRepository, logger),
			SearhPost:   query.NewSearchPostHandler(postRepository, logger),
		}
}
