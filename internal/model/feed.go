package model

/*
Package model defines the core domain entities used across FeedFlux.

Models represent the fundamental data structures of the system.
They are intentionally kept free of business logic and infrastructure concerns.

Design Principles:
------------------
- Models define the shape of domain data.
- They are shared across handler, service, and repository layers.
- They must remain storage-agnostic (not tied to memory or database implementation).
*/

import "time"

/*

Feed represents a registered RSS feed source in the system.

It models the essential attributes required to:
- Identify a feed
- Store its source URL
- Track metadata related to ingestion

Field Breakdown:
----------------
ID:
    Unique identifier of the feed.
    In Phase 1 (in-memory), this may be generated manually.
    In Phase 2 (PostgreSQL), this may become a UUID.

URL:
    The RSS endpoint from which articles are fetched.
    Expected to be unique across the system.

Name:
    Human-readable name of the feed.

LastFetchedAt:
    Timestamp indicating when the feed was last processed
    by the background worker.
    Pointer is used to allow nullability (feed may not
    have been fetched yet).

CreatedAt:
    Timestamp representing when the feed was registered
    in the system.

Architectural Note:
-------------------
This struct intentionally contains no validation or behavior.
All business rules are handled in the service layer.

*/

type Feed struct {
	ID            string
	URL           string
	Name          string
	LastFetchedAt *time.Time
	CreatedAt     time.Time
}
