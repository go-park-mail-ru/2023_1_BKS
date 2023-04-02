package error

// Ошибка, которая возникает при пустом заголовке
type TitleEmptyErr struct{}

func (p *TitleEmptyErr) Error() string {
	return "Заголовок не должен быть пустым"
}

// Ошибка, которая возникает при пустом описании
type BodyEmptyErr struct{}

func (p *BodyEmptyErr) Error() string {
	return "Описание не должно быть пустым"
}

// Ошибка, которая возникает при отсутствии изображений товара
type ImageEmptyErr struct{}

func (p *ImageEmptyErr) Error() string {
	return "Вставьте хотя бы 1 изображение товара"
}
