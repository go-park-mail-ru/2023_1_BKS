package command

import (
	"context"
	"post/domain"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreateHandler) Handle(
	ctx context.Context,
	postDelivery domain.Post,
) error {
	//Тут должна быть валидация
	post := domain.Post{
		Id:         uuid.New(),
		UserID:     postDelivery.UserID,
		Title:      postDelivery.Title,
		Desciption: postDelivery.Desciption,
		Price:      postDelivery.Price,
		Tags:       postDelivery.Tags,
		PathImages: postDelivery.PathImages,
		Time:       time.Now(),
	}
	err := h.postRepo.Create(ctx, post)
	return err
}

type AddCartHandler struct {
	cartRepo  domain.CUDCartRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *AddCartHandler) Handle(
	ctx context.Context,
	userId uuid.UUID,
	postId uuid.UUID,
) error {
	err := h.cartRepo.Add(ctx, userId, postId)
	return err
}
