package image

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/drew138/go-graphics/filters"
	"github.com/drew138/go-graphics/filters/kernels"
)

type Service interface {
	TransformImage(image image.Image, kernel kernels.Kernel) ([]byte, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (sv *service) TransformImage(image image.Image, kernel kernels.Kernel) ([]byte, error) {
	img := filters.ApplyFilter(image, kernel)

	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
