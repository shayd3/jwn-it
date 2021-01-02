package routes

import (
	"github.com/gin-gonic/gin"
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
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(301,"https://google.com/")
	})
}
