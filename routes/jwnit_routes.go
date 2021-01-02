package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/controllers"
	"github.com/shayd3/jwn-it/data"
)

func addJwnItRoutes(rg *gin.RouterGroup) {
	jwnit := rg.Group("/")

	jwnit.POST("/create", func(c *gin.Context) {
		// urlEntry := URLEntry{}
		// err := c.BindJSON(&urlEntry)
		// urlEntry.Created = time.Now()
		
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H {
		// 		"error": err.Error(),
		// 	})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H {
		// 		"data": urlEntry,
		// 		"error": "",
		// 	})
		// }

		// err = addURLEntry(db, urlEntry)
		// if err != nil {
		// 	return 
		// }
	})

	jwnit.GET("/urls", func(c *gin.Context) {
		fmt.Print(c.Keys)
		urlEntries, err := controllers.GetURLEntries(data.DB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H {
				"data": urlEntries,
				"error": "",
			})
		}
	})
}