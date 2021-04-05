package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/models"
	"github.com/shayd3/jwn-it/services"
)

// Home routes back to the home page!
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "Hey, welcome home!",
	})
}

// GetURLEntries returns a list of URLEntry objects.
func GetURLEntries(c *gin.Context) {
	urlEntries, err := services.GetURLEntries()
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, urlEntries)
	}
}

// GetURLEntry gets a URLEntry on the slug
func GetURLEntry(c *gin.Context) (models.URLEntry, error) {	
	urlEntry, err := services.GetURLEntry(strings.TrimLeft(c.Request.RequestURI, "/"))
	return urlEntry, err
}

// AddURLEntry adds an URLEntry to the db. If TargetURL does not contain 'http://' or 'https://', it will automatically add 'https://'
func AddURLEntry(c *gin.Context) {
	urlEntry := models.URLEntry{}
	err := c.BindJSON(&urlEntry)
	if err != nil {
		panic(err)
	}
	
	urlEntry, err = services.AddURLEntry(urlEntry)
	
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, urlEntry)
	}
	fmt.Println("Added URL Entry")
}

