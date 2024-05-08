package store

import "time"

// Item is an item in the store.
type Item struct {
	SKU   string
	Name  string
	Price int // In ¢
	Added time.Time
	Image []byte
}
