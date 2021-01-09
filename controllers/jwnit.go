package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/models"
	bolt "go.etcd.io/bbolt"
)

// Home routes back to the home page!
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"data": gin.H {
			"message": "Hey, welcome home!",
		},
		"error" : "",
	})
}

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

// GetURLEntry gets a URLEntry on the slug
func GetURLEntry(c *gin.Context) (models.URLEntry, error) {
	slug := strings.TrimLeft(c.Request.RequestURI, "/")
	urlEntry := models.URLEntry{}

	err := data.DB.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("JWNIT"))
		key := []byte(slug)
		err := json.Unmarshal(bucket.Get(key), &urlEntry)
		if err != nil {
			return err
		}
		return nil
	})

	return urlEntry, err
}

// AddURLEntry adds an URLEntry to the db. If TargetURL does not contain 'http://' or 'https://', it will automatically add 'https://'
func AddURLEntry(c *gin.Context) {
	urlEntry := models.URLEntry{}
	err := c.BindJSON(&urlEntry)

	urlEntry.Created = time.Now()
	if (!hasHTTPProtocol(urlEntry.TargetURL)) {
		urlEntry.TargetURL = addHTTPSToURL(urlEntry.TargetURL)
	}

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

// RouteToTargetURL will map a slug to a targetURL and redirect to that targetURL
// if URLEntry with given slug does not exist, this will re-route to /
func RouteToTargetURL(c *gin.Context) {
	urlEntry, err := GetURLEntry(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H {
			"error": err.Error(),
		})
	} else {
		c.Redirect(301,urlEntry.TargetURL)
	}
}

func addHTTPSToURL(url string) string {
	return "https://" + url
}

func hasHTTPProtocol(url string) bool {
	return strings.Contains(url, "http://") || strings.Contains(url, "https://")
}
