package cache

import (
	"log"
	"sync"
	"time"

	"news-app/internal/news"
)

type Fetcher interface {
	FetchNewsFromSource(url string) news.Source
}

type CachedNewsFetcher struct {
	// This should be a Redis cache or equivalent in a production app.
	cache sync.Map

	sourceURLs []string
	fetcher    Fetcher
	ttl        time.Duration
}

func NewCachedNewsFetcher(sourceURLs []string, fetcher Fetcher, ttl time.Duration) *CachedNewsFetcher {
	return &CachedNewsFetcher{
		cache: sync.Map{},
		// The sources are injected from the config file. It should use the ones store in the DB #TODO
		sourceURLs: sourceURLs,
		fetcher:    fetcher,
		ttl:        ttl,
	}
}

type CachedSource struct {
	expiry time.Time
	source news.Source
}

func (f *CachedNewsFetcher) FetchFromAllSources() ([]news.Source, error) {
	sources := make([]news.Source, 0, len(f.sourceURLs))
	for _, sourceURL := range f.sourceURLs {
		s, err := f.FetchFromOneSource(sourceURL)
		if err != nil {
			log.Println("error getting news from source: " + sourceURL + err.Error())
		}
		sources = append(sources, s)
	}
	return sources, nil
}

func (f *CachedNewsFetcher) FetchFromOneSource(sourceURL string) (news.Source, error) {
	if cacheValue, ok := f.cache.Load(sourceURL); ok {
		log.Println("Cache hit for " + sourceURL)
		cachedSource := cacheValue.(CachedSource)
		if cachedSource.expiry.Before(time.Now().UTC()) {
			log.Println("stale cache entry for " + sourceURL)
			return f.getSourceFromRemote(sourceURL)
		} else {
			return cachedSource.source, nil
		}
	} else {
		return f.getSourceFromRemote(sourceURL)
	}
}

func (f *CachedNewsFetcher) getSourceFromRemote(sourceURL string) (news.Source, error) {
	source := f.fetcher.FetchNewsFromSource(sourceURL)
	f.cache.Store(sourceURL, CachedSource{
		expiry: time.Now().UTC().Add(f.ttl),
		source: source,
	})
	return source, nil
}

//TODO: Add a clear cache function to run in a separate goroutine and remove expired elements
//TODO: ALL THE UNIT TESTS.
