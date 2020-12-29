package routes

import (
	"github.com/gin-gonic/gin"
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
		// urlEntries, err := getURLEntries(db)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H {
		// 		"error": err.Error(),
		// 	})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H {
		// 		"data": urlEntries,
		// 		"error": "",
		// 	})
		// }
	})
}