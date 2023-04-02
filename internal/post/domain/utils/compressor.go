package imageutil

import (
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
)

// CompressImages сжимает каждое изображение массива images до 20 Мб,
// пропорционально изменяя его размер. Функция возвращает массив сжатых
// изображений.
func CompressImages(images []image.RGBA, maxSize int64) ([]image.RGBA, error) {
	compressedImages := make([]image.RGBA, 0, len(images))

	for _, img := range images {
		// Создаем новое изображение типа RGBA с такими же размерами
		// как оригинальное изображение
		newImg := image.NewRGBA(img.Bounds())

		// Копируем пиксели из оригинального изображения в новое изображение
		draw.Draw(newImg, newImg.Bounds(), image.Image(&img), image.Point{}, draw.Src)

		// Вычисляем коэффициент масштабирования для изменения размера
		// изображения до нужного размера
		scale := float64(maxSize) / float64(newImg.Bounds().Dx()*newImg.Bounds().Dy()*4)

		// Изменяем размер изображения с помощью библиотеки disintegration/imaging
		newWidth := int(float64(newImg.Bounds().Dx()) * scale)
		newHeight := int(float64(newImg.Bounds().Dy()) * scale)
		compressedImg := imaging.Resize(newImg, newWidth, newHeight, imaging.Lanczos)

		// Создаем новое изображение типа RGBA и копируем пиксели из сжатого
		// изображения в новое изображение
		newCompressedImg := image.NewRGBA(compressedImg.Bounds())
		draw.Draw(newCompressedImg, newCompressedImg.Bounds(), compressedImg, image.Point{}, draw.Src)

		// Добавляем сжатое изображение в результирующий массив
		compressedImages = append(compressedImages, *newCompressedImg)
	}

	return compressedImages, nil
}
