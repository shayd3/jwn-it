package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/shayd3/jwn-it/internal/routes"
	bolt "go.etcd.io/bbolt"
)

func main() {
	// Set up DB
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	routes.Run()
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

func getURLEntries(db *bolt.DB) ([]URLEntry, error) {
	urlEntrys := []URLEntry{}
	err := db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte("JWNIT"))
		bucket.ForEach(func(k, v []byte) error {
			var urlEntry URLEntry
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