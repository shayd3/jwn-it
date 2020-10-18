package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware
	// logger and recovery (crash-free) middleware
	router := gin.Default()


	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
	
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)

	})


	// By default, it serves on :8080 unless
	// a PORT environment variable was defined
	router.Run()
	// router.Run(":3000") for a hard coded port
}