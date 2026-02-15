package repository

/*
Package repository defines persistence contracts for FeedFlux.

The ArticleRepository interface abstracts article storage operations
from the business logic layer.

Architectural Principle:
-------------------------
The service layer depends on this interface rather than a concrete
storage implementation.

This enables:
- Seamless transition from in-memory storage (Phase 1)
  to PostgreSQL (Phase 2).
- Easier mocking for unit tests.
- Clear separation of persistence concerns.
*/

import "github.com/Krishna8167/feedflux/internal/model"

/*
ArticleRepository defines the contract for storing and retrieving articles.

Any storage implementation must satisfy this interface,
whether in-memory or database-backed.

Method Responsibilities:
------------------------

Create:
    Persists a new Article entity.
    In later phases, this may enforce idempotency
    (e.g., prevent duplicate articles based on Link).

GetAll:
    Retrieves all articles with pagination support.
    - limit  → maximum number of records to return.
    - offset → starting position for pagination.

GetByFeedID:
    Retrieves articles belonging to a specific feed.
    Supports pagination via limit and offset.
    Enables filtering of articles by feed source.

Design Note:
------------
Pagination parameters are included at repository level
to allow efficient implementation in database-backed storage,
where LIMIT and OFFSET can be applied directly in queries.
*/

type ArticleRepository interface {
	Create(article *model.Article) error
	GetAll(limit, offset int) ([]model.Article, error)
	GetByFeedID(feedID string, limit, offset int) ([]model.Article, error)
}
