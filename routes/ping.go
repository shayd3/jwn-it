package routes

import (
	"github.com/gin-gonic/gin"
)


func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/")

	ping.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}