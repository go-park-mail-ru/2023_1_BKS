package domain

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
)

const (
	maxImageSize = 1024 * 512
)

func ValidateTitle(title string, maxTitleLength int) error {
	if len(title) == 0 {
		return &TitleEmptyErr{}
	}

	if len(title) > maxTitleLength {
		return &LongTitleErr{MaxLength: maxTitleLength}
	}

	return nil
}

func ValidateBody(body string, maxBodyLength int) error {
	if len(body) == 0 {
		return &BodyEmptyErr{}
	}

	if len(body) > maxBodyLength {
		return &LongBodyErr{MaxLength: maxBodyLength}
	}

	return nil
}

func ValidateImages(images []image.RGBA, maxImageCount int) error {
	if len(images) == 0 {
		return &ImageEmptyErr{}
	}
	if len(images) > maxImageCount {
		return &ManyImagesErr{MaxCount: maxImageCount}
	}

	for _, img := range images {
		var format string
		// Определяем расширение
		if img.Opaque() {
			format = "jpeg"
		} else {
			format = "png"
		}

		// Проверяем размер изображения
		buf := new(bytes.Buffer)

		switch format {
		case "png":
			if err := png.Encode(buf, &img); err != nil {
				return err
			}
		case "jpeg":
			// Создаем промежуточный объект
			bounds := img.Bounds()
			jpegImage := image.NewRGBA(bounds)
			jpegImage.Pix = img.Pix

			if err := jpeg.Encode(buf, jpegImage, nil); err != nil {
				return err
			}
		}

		if buf.Len() > maxImageSize {
			return &HeavyImageErr{}
		}
	}

	return nil
}
