package repository

import "github.com/Krishna8167/feedflux/internal/model"

type FeedRepository interface {
	Create(feed *model.Feed) error
	GetAll() ([]model.Feed, error)
	GetByID(id string) (*model.Feed, error)
}
