package handler

import (
	"bytes"
	"errors"
	"fmt"
	imagePkg "image"
	"image/jpeg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/drew138/graphics-api/api/middleware"
	"github.com/drew138/graphics-api/mocks"
)

func TestCreateSharpenHandler(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateSharpen()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/sharpen", sharpenHandler)
	req, _ := http.NewRequest("GET", "/sharpen", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior
	mockService.On("TransformImage", mock.Anything, mock.Anything).Return(buf.Bytes(), nil).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateSharpenHandler_ImageNotFound(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateSharpen()

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/sharpen", sharpenHandler)
	req, _ := http.NewRequest("GET", "/sharpen", nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateSharpenHandler_FailedToSharpen(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateSharpen()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/sharpen", sharpenHandler)
	req, _ := http.NewRequest("GET", "/sharpen", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior to simulate error
	mockService.On("TransformImage", mock.Anything, mock.Anything).
		Return(nil, errors.New("failed to sharpen")).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateGaussianBlurHandler(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateGaussianBlur()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/gaussian-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/gaussian-blur", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior
	mockService.On("TransformImage", mock.Anything, mock.Anything).Return(buf.Bytes(), nil).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateGaussianBlur_ImageNotFound(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateGaussianBlur()

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/gaussian-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/gaussian-blur", nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateGaussianBlur_FailedToSharpen(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateGaussianBlur()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/gaussian-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/gaussian-blur", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior to simulate error
	mockService.On("TransformImage", mock.Anything, mock.Anything).
		Return(nil, errors.New("failed to sharpen")).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateEdgeDetectionHandler(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateEdgeDetection()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/edge-detection", sharpenHandler)
	req, _ := http.NewRequest("GET", "/edge-detection", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior
	mockService.On("TransformImage", mock.Anything, mock.Anything).Return(buf.Bytes(), nil).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateEdgeDetection_ImageNotFound(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateEdgeDetection()

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/edge-detection", sharpenHandler)
	req, _ := http.NewRequest("GET", "/edge-detection", nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateEdgeDetection_FailedToSharpen(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateEdgeDetection()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/edge-detection", sharpenHandler)
	req, _ := http.NewRequest("GET", "/edge-detection", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior to simulate error
	mockService.On("TransformImage", mock.Anything, mock.Anything).
		Return(nil, errors.New("failed to sharpen")).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateBoxBlurHandler(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateBoxBlur()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/box-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/box-blur", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior
	mockService.On("TransformImage", mock.Anything, mock.Anything).Return(buf.Bytes(), nil).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateBoxBlur_ImageNotFound(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateBoxBlur()

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/box-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/box-blur", nil)

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}

func TestCreateBoxBlur_FailedToSharpen(t *testing.T) {
	mockService := mocks.NewService(t)
	sharpenHandler := NewImage(mockService).CreateBoxBlur()

	// Prepare a sample image
	img := imagePkg.NewRGBA(imagePkg.Rect(0, 0, 100, 100))
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, nil)

	// Set up Gin context
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.ParseImage())
	r.GET("/box-blur", sharpenHandler)
	req, _ := http.NewRequest("GET", "/box-blur", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Content-Length", fmt.Sprint(buf.Len()))

	// Mock service behavior to simulate error
	mockService.On("TransformImage", mock.Anything, mock.Anything).
		Return(nil, errors.New("failed to sharpen")).Once()

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotNil(t, w.Body)
	// You may add more assertions based on your handler's behavior
}
