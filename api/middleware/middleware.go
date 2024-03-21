package middleware

import (
	"bytes"
	"image"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseImage() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.Request.Header.Get("Content-Type")
		if !strings.Contains(contentType, "image/jpeg") {
			c.JSON(400, gin.H{"message": "No image found in request body"})
			c.Abort()
			return
		}

		file, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(400, gin.H{"message": "Error reading file"})
			c.Abort()
			return
		}

		img, _, err := image.Decode(bytes.NewReader(file))

		if err != nil {
			c.JSON(400, gin.H{"message": "Error decoding image"})
			c.Abort()
			return
		}

		c.Set("image", img)

		c.Next()
	}
}
