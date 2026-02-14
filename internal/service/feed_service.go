package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/Krishna8167/feedflux/internal/model"
	"github.com/Krishna8167/feedflux/internal/repository"
)

type FeedService struct {
	repo repository.FeedRepository
}

func NewFeedService(repo repository.FeedRepository) *FeedService {
	return &FeedService{repo: repo}
}

func (s *FeedService) AddFeed(url, name string) (*model.Feed, error) {
	if url == "" || name == "" {
		return nil, errors.New("url and name are required")
	}

	feed := &model.Feed{
		ID:        uuid.NewString(),
		URL:       url,
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func (s *FeedService) ListFeeds() ([]model.Feed, error) {
	return s.repo.GetAll()
}
