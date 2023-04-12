package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdatePostHandler struct {
	postRepo  domain.CUDRepositoryPost
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdatePostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
	postDelivery domain.PostDeliveryInterface,
) error {

	post := domain.Post{
		Id:         id,
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

type ClosePostHandler struct {
	postRepo  domain.CUDRepositoryPost
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *ClosePostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.postRepo.Close(ctx, id)
	return err
}
