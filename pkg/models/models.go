package models

import (
	"errors"
	"time"
)

// ErrNoRecord ..
var ErrNoRecord = errors.New("models: not matching record found")

// Snippet ..
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
