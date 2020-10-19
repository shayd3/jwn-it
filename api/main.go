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
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

func setupDB() (*bolt.DB, error) {
	// set up bboltdb
	// Open my.db data file in current dir
	db, err := bolt.Open("jwn.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	// Setup buckets
	err = db.Update(func(t *bolt.Tx) error {
		root, err := t.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("JWNIT"))
		if err != nil {
			return fmt.Errorf("could not create jwnit bucket: %v", err)
		}
		return nil
	})
	// Check if there was a problem setting up buckets
	if err != nil {
		return nil, fmt.Errorf("could not setup buckets: %v", err)
	}
	fmt.Println("DB Setup Complete!")
	return db, nil
}