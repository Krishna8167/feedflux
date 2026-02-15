package model

/*
Package model contains the domain entities of FeedFlux.

Models define the core data structures shared across layers.
They remain independent from:

- HTTP concerns
- Business logic
- Storage implementation
- External dependencies

This ensures clean separation between domain representation
and infrastructure details.
*/

import "time"

/*
Article represents a single RSS entry fetched from a feed source.

It models the persistent unit of content stored in the system.

Field Breakdown:
----------------

ID:
    Unique identifier of the article.
    In later phases, this may be generated as a UUID.

FeedID:
    Foreign reference to the Feed that owns this article.
    Establishes relationship between feeds and articles.

Title:
    The headline/title of the RSS entry.

Link:
    Canonical URL of the article.
    This is typically used for idempotency checks
    (ensuring duplicate articles are not stored).

Description:
    Short summary or content excerpt from the RSS feed.

PublishedAt:
    Timestamp from the RSS feed indicating when
    the article was originally published.
    Pointer allows nullability in case the feed
    does not provide a publish date.

CreatedAt:
    Timestamp representing when the article
    was stored in FeedFlux.

Architectural Note:
-------------------
The model is intentionally passive (no methods attached).
All validation, deduplication logic, and business rules
are implemented in the service layer.

In Phase 2, database-level constraints (e.g., unique index on Link)
will reinforce idempotency at the persistence layer.
*/

type Article struct {
	ID          string
	FeedID      string
	Title       string
	Link        string
	Description string
	PublishedAt *time.Time
	CreatedAt   time.Time
}
