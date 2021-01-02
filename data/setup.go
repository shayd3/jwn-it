package data

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

// DB is a global variable for our DB instance
var DB *bolt.DB

// ConnectDatabase sets up the bolt db connection/bucket creation
func ConnectDatabase() {
	// set up bboltdb
	// Open my.db data file in current dir
	db, err := bolt.Open("jwn.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	// Setup buckets
	err = db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte("JWNIT"))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	// Check if there was a problem setting up buckets
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Setup Complete!")

	DB = db
}