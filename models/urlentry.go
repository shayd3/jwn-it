package models

import "time"

// URLEntry is a data object keeping track of
// the target (original) url and the slug for the
// short url. Slug is concidered the key
type URLEntry struct {
	Slug string `json:"slug"`
	Created time.Time 
	TargetURL string `json:"targetURL"`
}