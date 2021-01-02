package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/models"
	bolt "go.etcd.io/bbolt"
)

// GetURLEntries returns a list of URLEntry objects.
func GetURLEntries(c *gin.Context) {
	urlEntries := []models.URLEntry{}
	err := data.DB.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("JWNIT"))
		bucket.ForEach(func(k, v []byte) error {
			var urlEntry models.URLEntry
			err := json.Unmarshal(v, &urlEntry)
			if err != nil {
				return err
			}
			urlEntries = append(urlEntries, urlEntry)
			return nil
		})
		return nil
	})
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
}

// AddURLEntry adds an URLEntry to the db. 
func AddURLEntry(c *gin.Context) {
	urlEntry := models.URLEntry{}
	err := c.BindJSON(&urlEntry)
	urlEntry.Created = time.Now()

	err = data.DB.Update(func(t *bolt.Tx) error {
		encoded, err := json.Marshal(urlEntry)
		if err != nil {
			return fmt.Errorf("could not marshall URLEntry object: %v", err)
		}
		err = t.Bucket([]byte("JWNIT")).Put([]byte(urlEntry.Slug), encoded)
		if err != nil {
			return fmt.Errorf("could not insert URLEntry: %v", err)
		}
		return nil
	})

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
	fmt.Println("Added URL Entry")
}

