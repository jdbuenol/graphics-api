package main

import (
	"github.com/gin-gonic/gin"

	router "github.com/drew138/graphics-api/api/routes"
)

func main() {
	eng := gin.Default()

	router := router.NewRouter(eng)
	router.MapRoutes()

	if err := eng.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}

}
