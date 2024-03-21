package router

import (
	"github.com/gin-gonic/gin"

	"github.com/drew138/graphics-api/api/handler"
	"github.com/drew138/graphics-api/api/middleware"
	"github.com/drew138/graphics-api/internal/image"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
}

func NewRouter(eng *gin.Engine) Router {
	return &router{eng: eng}
}

func (r *router) setGlobalMiddleware() {
	r.eng.Use(middleware.ParseImage())
}

func (r *router) MapRoutes() {
	r.setGlobalMiddleware()
	r.buildImageRoutes()
}

func (r *router) buildImageRoutes() {
	service := image.NewService()
	handler := handler.NewImage(service)

	r.eng.POST("/sharpen", handler.CreateSharpen())
	r.eng.POST("/edgedetection", handler.CreateEdgeDetection())
	r.eng.POST("/gaussianblur", handler.CreateGaussianBlur())
	r.eng.POST("/boxblur", handler.CreateBoxBlur())
}
