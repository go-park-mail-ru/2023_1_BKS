package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreatePostHandler struct {
	postRepo  domain.CUDRepositoryPost
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreatePostHandler) Handle(
	ctx context.Context,
	postDelivery domain.PostBody,
) error {

	post := domain.Post{
		Id:         uuid.New(),
		UserID:     postDelivery.UserID,
		Title:      domain.CreateTitle(postDelivery.Title),
		Desciption: domain.CreateDescription(postDelivery.Desciption),
		Price:      domain.CreatePrice(postDelivery.Price),
		Tags:       domain.CreateTags(postDelivery.Tags),
		Images:     domain.CreateImages(postDelivery.Images),
		Time:       domain.CreateTimeStamp(postDelivery.Time),
	}
	err := h.postRepo.Create(ctx, post)
	return err
}
