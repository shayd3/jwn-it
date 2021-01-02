package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/shayd3/jwn-it/models"
	bolt "go.etcd.io/bbolt"
)

// GetURLEntries returns a list of URLEntry objects.
func GetURLEntries(db *bolt.DB) ([]models.URLEntry, error) {
	urlEntrys := []models.URLEntry{}
	err := db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("JWNIT"))
		bucket.ForEach(func(k, v []byte) error {
			var urlEntry models.URLEntry
			err := json.Unmarshal(v, &urlEntry)
			if err != nil {
				return err
			}
			urlEntrys = append(urlEntrys, urlEntry)
			return nil
		})
		return nil
	})
	return urlEntrys, err
}

// AddURLEntry adds an URLEntry to the db. 
func AddURLEntry(db *bolt.DB, entry models.URLEntry) error {
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

