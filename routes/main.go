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
	getRoutes()
	getRouterDefault()
	router.Run()
}

func getRoutes() {
	v1 := router.Group("/v1")
	addJwnItRoutes(v1)
	addPingRoutes(v1)
}

func getRouterDefault() {
	router.NoRoute(controllers.RouteToTargetURL)
}
