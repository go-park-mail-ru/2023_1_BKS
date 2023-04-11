package command

import (
	"context"
	"post/domain"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdatePostHandler struct {
	postRepo  domain.CUDRepository
	validator domain.SpecificationManager
	loger     *logrus.Entry
}

func (h *UpdatePostHandler) Handle(
	ctx context.Context,
	id uuid.UUID,
	title,
	description string,
	image []string,
	tags []string,
) error {
	if err := h.validator.Title.IsValid(title); err != nil {
		return err
	}

	post := domain.Post{
		Title:      domain.CreateTitle(title),
		Desciption: domain.CreateDescription(description),
		Images:     domain.CreateImages(image),
		Tags:       domain.CreateTags(tags),
	}
	err := h.postRepo.Update(ctx, post)
	return err
}
