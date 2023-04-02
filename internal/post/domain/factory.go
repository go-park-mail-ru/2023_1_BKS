package domain

import (
	"image"

	util "github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain/utils"
)

const (
	maxTitleLength = 50
	maxBodyLength  = 500
	maxImagesCount = 5
	maxSize        = 20 * 1024 * 1024
)

func CreateBody(body string) (Body, error) {
	if err := ValidateBody(body, maxBodyLength); err != nil {
		return Body(""), err
	}

	return Body(body), nil
}

func CreateTitle(title string) (Title, error) {
	if err := ValidateTitle(title, maxTitleLength); err != nil {
		return Title(""), err
	}

	return Title(title), nil
}

func CreateImage(image []image.RGBA) (Image, error) {
	image, err := util.CompressImages(image, maxSize)
	if err != nil {
		return Image(nil), err
	}

	if ok := ValidateImages(image, maxImagesCount); ok != nil {
		if ok != nil {
			return Image(nil), err
		}
	}

	return Image(image), nil
}
