package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware
	// logger and recovery (crash-free) middleware
	r := gin.Default()


	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	// By default, it serves on :8080 unless
	// a PORT environment variable was defined
	r.Run()
	// router.Run(":3000") for a hard coded port
}