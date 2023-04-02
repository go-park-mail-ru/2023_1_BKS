package object

import "image"

func ValidateTitle(body string) (bool, error) {
	return true, nil
}

func ValidateBody(body string) (bool, error) {
	return true, nil
}

func ValidateImage(image []image.RGBA) (bool, error) {
	return true, nil
}
