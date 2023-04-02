package domain

import "fmt"

// TitleEmptyErr возвращается при пустом заголовке
type TitleEmptyErr struct{}

func (e *TitleEmptyErr) Error() string {
	return "Заголовок не должен быть пустым"
}

// BodyEmptyErr возвращается при пустом описании
type BodyEmptyErr struct{}

func (e *BodyEmptyErr) Error() string {
	return "Описание не должно быть пустым"
}

// ImageEmptyErr возвращается при отсутствии изображений товара
type ImageEmptyErr struct{}

func (e *ImageEmptyErr) Error() string {
	return "Вставьте хотя бы 1 изображение товара"
}

// LongBodyErr возвращается при превышении максимальной длины описания
type LongBodyErr struct {
	MaxLength int
}

func (e *LongBodyErr) Error() string {
	return fmt.Sprintf("Описание должно содержать не более %d символов\n", e.MaxLength)
}

// LongTitleErr возвращается при превышении максимальной длины заголовка
type LongTitleErr struct {
	MaxLength int
}

func (e *LongTitleErr) Error() string {
	return fmt.Sprintf("Заголовок должен содержать не более %d символов\n", e.MaxLength)
}

// ManyImagesErr возвращается при превышении максимального количества изображений
type ManyImagesErr struct {
	MaxCount int
}

func (e *ManyImagesErr) Error() string {
	return fmt.Sprintf("Всего может быть не более %d картинок\n", e.MaxCount)
}

type HeavyImageErr struct{}

func (e *HeavyImageErr) Error() string {
	return "Размер изображения не должен превышать 512 МБ"
}
