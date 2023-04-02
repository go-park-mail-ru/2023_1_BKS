package object

import "image"

func CreateBody(body string) (Body, []error) {
	var err []error
	/*
		Валидацию сюда
	*/
	return Body(body), err
}

func CreateTitle(title string) (Title, []error) {
	var err []error
	/*
		Валидацию сюда
	*/
	return Title(title), err
}

func CreateImage(image []image.RGBA) (Image, []error) {
	var err []error
	/*
		Валидацию сюда
	*/
	return Image(image), err
}
