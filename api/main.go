package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

func main() {
	// set up bboltdb
	// Open my.db data file in current dir
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create bucket
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		log.Printf("Bucket Depth: %v", bucket.Stats().Depth) 
		return nil
	})


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