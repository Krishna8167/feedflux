package model

import "time"

type Feed struct {
	ID            string
	URL           string
	Name          string
	LastFetchedAt *time.Time
	CreatedAt     time.Time
}
