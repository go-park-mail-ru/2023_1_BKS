package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreatePostHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *CreatePostHandler) Handle(
	ctx context.Context,
	title,
	description string,
	images []string,
	tags []string,
) error {
	if err := h.validator.Title.IsValid(title); err != nil {
		return err
	}
	if err := h.validator.Desciption.IsValid(description); err != nil {
		return err
	}
	for i := 0; i < len(images); i++ {
		if err := h.validator.Images.IsValid(images[i]); err != nil {
			return err
		}
	}
	for i := 0; i < len(images); i++ {
		if err := h.validator.Tags.IsValid(tags[i]); err != nil {
			return err
		}
	}

	post := domain.Post{
		Id:         uuid.New(),
		Title:      domain.CreateTitle(title),
		Desciption: domain.CreateDescription(description),
		Images:     domain.CreateImages(images),
		Tags:       domain.CreateTags(tags),
	}
	err := h.postRepo.Create(ctx, post)
	return err
}
