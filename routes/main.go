package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/controllers"
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
	getRouterDefault()
}

func getRoutes() {
	v1 := router.Group("/v1")
	addJwnItRoutes(v1)
}

func getRouterDefault() {
	router.NoRoute(controllers.RouteToTargetURL)
}
