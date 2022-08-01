package service

import (
	"log"

	"news-app/internal/news"
	"news-app/internal/store"
)

type Store interface {
	GetAllSources() ([]store.SourceRecord, error)
}

type Service struct {
	store Store
}

func New(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetAllSources() ([]news.Source, error) {
	dbRecords, err := s.store.GetAllSources()
	if err != nil {
		log.Println("error getting news sources from database: " + err.Error())
		return nil, err
	}

	sources := make([]news.Source, 0, len(dbRecords))
	for _, r := range dbRecords {
		sources = append(sources, news.Source{
			Name: r.SourceName,
			URL:  r.SourceURL,
		})
	}
	return sources, nil
}

//TODO: ALL THE TESTS.

//TODO: inject the news fetcher into the service so that the http handler will only use the service.
