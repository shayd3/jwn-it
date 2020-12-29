package services

import (
	"encoding/json"

	bolt "go.etcd.io/bbolt"
)

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