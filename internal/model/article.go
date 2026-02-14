package model

import "time"

type Article struct {
	ID          string
	FeedID      string
	Title       string
	Link        string
	Description string
	PublishedAt *time.Time
	CreatedAt   time.Time
}
