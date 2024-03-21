package handler

import (
	imagePkg "image"
	"net/http"

	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/gin-gonic/gin"

	"github.com/drew138/graphics-api/internal/image"
)

type Image struct {
	service image.Service
}

func NewImage(service image.Service) *Image {
	return &Image{service}
}

func (s *Image) CreateSharpen() gin.HandlerFunc {
	return func(c *gin.Context) {
		image, exists := c.Get("image")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found in request"})
			return
		}

		bytes, err := s.service.TransformImage(image.(imagePkg.Image), kernels.Sharpen)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sharpen image"})
			return
		}

		c.Header("Content-Type", "image/jpeg")
		c.Data(http.StatusOK, "image/jpeg", bytes)
	}
}

func (s *Image) CreateEdgeDetection() gin.HandlerFunc {
	return func(c *gin.Context) {
		image, exists := c.Get("image")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found in request"})
			return
		}

		bytes, err := s.service.TransformImage(image.(imagePkg.Image), kernels.EdgeDetection)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sharpen image"})
			return
		}

		c.Header("Content-Type", "image/jpeg")
		c.Data(http.StatusOK, "image/jpeg", bytes)
	}
}

func (s *Image) CreateGaussianBlur() gin.HandlerFunc {
	return func(c *gin.Context) {
		image, exists := c.Get("image")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found in request"})
			return
		}

		bytes, err := s.service.TransformImage(image.(imagePkg.Image), kernels.GaussianBlur)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sharpen image"})
			return
		}

		c.Header("Content-Type", "image/jpeg")
		c.Data(http.StatusOK, "image/jpeg", bytes)
	}
}

func (s *Image) CreateBoxBlur() gin.HandlerFunc {
	return func(c *gin.Context) {
		image, exists := c.Get("image")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found in request"})
			return
		}

		bytes, err := s.service.TransformImage(image.(imagePkg.Image), kernels.BoxBlur)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sharpen image"})
			return
		}

		c.Header("Content-Type", "image/jpeg")
		c.Data(http.StatusOK, "image/jpeg", bytes)
	}
}
