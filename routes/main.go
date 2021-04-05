package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start server
func Run() {
	router.Run()
}

// GetRouter will return the global variable router
func GetRouter() *gin.Engine {
	return router
}

// SetupRouter sets up router with all it's routes
func SetupRouter() {
	getRoutes()
}

func getRoutes() {
	v1 := router.Group("/v1")
	addJwnItRoutes(v1)
}