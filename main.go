package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

func main() {
	// Set up DB
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creates a gin router with default middleware
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"data": gin.H {
				"message": "pong",
			},
			"error": "",
		})
	})

	router.POST("/v1/create", func(c *gin.Context) {
		urlEntry := URLEntry{}
		err := c.BindJSON(&urlEntry)
		urlEntry.Created = time.Now()
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H {
				"data": urlEntry,
				"error": "",
			})
		}

		fmt.Printf("creating short url")
	})

	// Route slug to appropriate url
	router.NoRoute(func(c *gin.Context) {
		q := c.Request.URL.EscapedPath()
		c.JSON(200, gin.H {
			"slug": q,
		})
	})

	// By default, it serves on :8080 unless
	// a PORT environment variable was defined
	router.Run()
	// router.Run(":3000") for a hard coded port
}

func addURLEntry(db *bolt.DB, entry URLEntry) error {
	err := db.Update(func(t *bolt.Tx) error {
		encoded, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("could not marshall URLEntry object: %v", err)
		}
		err = t.Bucket([]byte("JWNIT")).Put([]byte(entry.Slug), encoded)
		if err != nil {
			return fmt.Errorf("could not insert URLEntry: %v", err)
		}
		return nil
	})
	fmt.Println("Added URL Entry")
	return err
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
		_, err := t.CreateBucketIfNotExists([]byte("JWNIT"))
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

// URLEntry is a data object keeping track of
// the target (original) url and the slug for the
// short url. Slug is concidered the key
type URLEntry struct {
	Slug string `json:"slug"`
	Created time.Time 
	TargetURL string `json:"targetURL"`
}