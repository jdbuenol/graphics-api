package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseImageMiddleware(t *testing.T) {
	// Create a new Gin router and add the middleware to it
	router := gin.New()
	router.Use(ParseImage())

	t.Run("NoMultipartFormError", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/upload", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "no multipart form")
	})

	t.Run("ErrorParsingFormData", func(t *testing.T) {
		body := new(bytes.Buffer)
		body.WriteString("invalid")

		req, err := http.NewRequest("POST", "/upload", body)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "multipart/form-data")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Error parsing form data")
	})

	t.Run("FileNotFoundError", func(t *testing.T) {
		body := new(bytes.Buffer)
		req, err := http.NewRequest("POST", "/upload", body)
		assert.NoError(t, err)

		form := new(bytes.Buffer)
		form.WriteString(fmt.Sprintf("--%s\r\n", "boundary"))
		form.WriteString(
			fmt.Sprintf(
				"Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"\r\n",
				"file",
				"example.jpg",
			),
		)
		form.WriteString("Content-Type: application/octet-stream\r\n\r\n")
		form.WriteString("file content here\r\n")
		form.WriteString(fmt.Sprintf("--%s--\r\n", "boundary"))

		req.Body = http.MaxBytesReader(nil, io.NopCloser(form), 32<<20)
		req.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "File not found in form data")
	})

	t.Run("Success", func(t *testing.T) {
		body := new(bytes.Buffer)
		req, err := http.NewRequest("POST", "/upload", body)
		assert.NoError(t, err)

		form := new(bytes.Buffer)
		form.WriteString(fmt.Sprintf("--%s\r\n", "boundary"))
		form.WriteString(
			fmt.Sprintf(
				"Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"\r\n",
				"file",
				"example.jpg",
			),
		)
		form.WriteString("Content-Type: application/octet-stream\r\n\r\n")
		form.WriteString("file content here\r\n")
		form.WriteString(fmt.Sprintf("--%s--\r\n", "boundary"))

		req.Body = http.MaxBytesReader(nil, io.NopCloser(form), 32<<20)
		req.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
