package usecase

import (
	"config"
	"context"
	"fmt"
	"user/domain"
	"user/repository"
	"user/usecase/command"
	"user/usecase/query"

	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context, cfg config.Config) (Commands, Queries) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Db.User, cfg.Db.DataBaseName, cfg.Db.Password, cfg.Db.Host,
		cfg.Db.Port, cfg.Db.Sslmode)
	validator := domain.CreateSpecificationManager(cfg)
	Repository := repository.CreatePostgressRepository(dsn)
	logger := logrus.NewEntry(logrus.StandardLogger())

	return Commands{
			CreateUser: command.NewCreateUserHandler(&Repository, validator, logger),
			UpdateUser: command.NewUpdateUserHandler(&Repository, validator, logger),
			DeleteUser: command.NewDeleteUserHandler(&Repository, validator, logger),
		},
		Queries{
			GetUser: query.NewGetIdUserHandler(Repository, logger),
		}
}
