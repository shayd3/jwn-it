package services

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/models"
	"github.com/shayd3/jwn-it/rand"
	"github.com/shayd3/rand"
	bolt "go.etcd.io/bbolt"
)

const db = "JWNIT"
const randSlugLength = 10

func GetURLEntries() ([]models.URLEntry, error) {
	urlEntries := []models.URLEntry{}
	err := data.DB.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(db))
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

	return urlEntries, err
}

func GetURLEntry(slug string) (models.URLEntry, error) {
	urlEntry := models.URLEntry{}

	err := data.DB.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(db))
		key := []byte(slug)
		err := json.Unmarshal(bucket.Get(key), &urlEntry)
		if err != nil {
			return err
		}
		return nil
	})

	return urlEntry, err
}

func AddURLEntry(urlEntry models.URLEntry) (models.URLEntry, error) {

	// Check if URLEntry already exists for given slug 
	_, err := GetURLEntry(urlEntry.Slug)
	if(err == nil) {
		return urlEntry, fmt.Errorf("entry for given slug '%s' already exists", urlEntry.Slug)
	}

	// Generate random slug if urlEntry.Slug is empty
	if(urlEntry.Slug == "") {
		urlEntry.Slug = generateSlug(randSlugLength)
	}

	urlEntry.Created = time.Now()
	if (!hasHTTPProtocol(urlEntry.TargetURL)) {
		urlEntry.TargetURL = addHTTPSToURL(urlEntry.TargetURL)
	}

	err = data.DB.Update(func(t *bolt.Tx) error {
		encoded, err := json.Marshal(urlEntry)
		if err != nil {
			return fmt.Errorf("could not marshall URLEntry object: %v", err)
		}
		err = t.Bucket([]byte(db)).Put([]byte(urlEntry.Slug), encoded)
		if err != nil {
			return fmt.Errorf("could not insert URLEntry: %v", err)
		}
		return nil
	})

	return urlEntry, err
}

func addHTTPSToURL(url string) string {
	return "https://" + url
}

func hasHTTPProtocol(url string) bool {
	return strings.Contains(url, "http://") || strings.Contains(url, "https://")
}

func generateSlug(length int) string {
	return rand.String(length)
}