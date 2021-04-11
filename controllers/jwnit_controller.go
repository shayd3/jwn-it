package controllers

import (
	"net/http"

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
func GetURLEntry(c *gin.Context) {
	urlEntry, err := services.GetURLEntry(c.Param("slug"))

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, urlEntry)
	}
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
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, urlEntry)
	}
}

