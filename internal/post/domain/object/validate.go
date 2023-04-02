package object

import (
	"image"

	err "github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain/error"
)

func ValidateTitle(title string, maxTitleLength int) error {
	if len(title) == 0 {
		return &err.TitleEmptyErr{}
	}

	if len(title) > maxTitleLength {
		return &err.LongTitleErr{MaxLength: maxTitleLength}
	}

	return nil
}

func ValidateBody(body string, maxBodyLength int) error {
	if len(body) == 0 {
		return &err.BodyEmptyErr{}
	}

	if len(body) > maxBodyLength {
		return &err.LongBodyErr{MaxLength: maxBodyLength}
	}

	return nil
}

func ValidateImages(images []*image.RGBA, maxImageCount int) error {
	if len(images) == 0 {
		return &err.ImageEmptyErr{}
	}

	if len(images) > maxImageCount {
		return &err.ManyImagesErr{MaxCount: maxImageCount}
	}

	return nil
}
