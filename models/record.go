package models

import (
	"time"
)

type Record struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Runner   string
	Time     time.Time `json:"time"`
	Url      string    `json:"url"`
}
