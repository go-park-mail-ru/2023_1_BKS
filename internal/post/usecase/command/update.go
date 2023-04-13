package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdateHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
	postDelivery domain.Post,
) error {
	//Тут должна быть валидация
	post := domain.Post{
		Id:         id,
		UserID:     postDelivery.UserID,
		Title:      postDelivery.Title,
		Desciption: postDelivery.Desciption,
		Price:      postDelivery.Price,
		Tags:       postDelivery.Tags,
		PathImages: postDelivery.PathImages,
		Time:       postDelivery.Time,
	}
	err := h.postRepo.Update(ctx, post)
	return err
}

type CloseHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CloseHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := h.postRepo.Close(ctx, id)
	return err
}
