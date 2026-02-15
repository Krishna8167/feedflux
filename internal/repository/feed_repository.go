package repository

/*
Package repository defines persistence abstractions for FeedFlux.

The repository layer introduces interfaces that abstract
data storage operations from business logic.

Core Principle:
---------------
The service layer depends on interfaces, not concrete implementations.

This allows:
- Swapping in-memory storage with PostgreSQL without modifying services.
- Easier unit testing through mock implementations.
- Clean separation between business rules and persistence logic.
*/

import "github.com/Krishna8167/feedflux/internal/model"

/*
FeedRepository defines the contract for feed persistence operations.

Any storage implementation (in-memory, PostgreSQL, etc.)
must satisfy this interface.

This ensures the service layer remains storage-agnostic.

Method Responsibilities:
------------------------

Create:
    Persists a new Feed entity.
    Must return an error if persistence fails
    or if business constraints are violated
    (e.g., duplicate URL in future implementations).

GetAll:
    Retrieves all stored feeds.
    Pagination may be implemented at service level.

GetByID:
    Fetches a single feed by its unique identifier.
    Returns an error if the feed does not exist.
*/

type FeedRepository interface {
	Create(feed *model.Feed) error
	GetAll() ([]model.Feed, error)
	GetByID(id string) (*model.Feed, error)
}
