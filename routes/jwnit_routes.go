package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/controllers"
)

func addJwnItRoutes(rg *gin.RouterGroup) {
	jwnit := rg.Group("/")

	jwnit.POST("/create", controllers.AddURLEntry)
	jwnit.GET("/urls", controllers.GetURLEntries)
	jwnit.GET("/", controllers.Home)
}