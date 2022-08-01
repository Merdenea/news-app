package news

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Fetcher struct {
	sourceURLs []string
}

func NewFetcher(urls []string) *Fetcher {
	return &Fetcher{sourceURLs: urls}
}

func (f *Fetcher) FetchFromAllSources() []Source {
	sources := make([]Source, 0, len(f.sourceURLs))
	for _, u := range f.sourceURLs {
		sources = append(sources, f.FetchNewsFromSource(u))
	}
	return sources
}

func (f *Fetcher) FetchNewsFromSource(url string) Source {
	log.Println("Fetching news from source: " + url)
	res, err := http.Get(url)
	if err != nil {
		return Source{}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Source{}
	}

	var rss Rss
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		log.Fatal(err)
		return Source{}
	}
	return toSource(rss)
}

func toSource(rss Rss) Source {
	articles := make([]Article, 0, len(rss.Channel.Items))

	for _, v := range rss.Channel.Items {
		articles = append(articles, Article{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Date:        v.Date,
		})
	}
	return Source{
		Name:     rss.Channel.Title,
		URL:      rss.Channel.Link,
		Articles: articles,
	}
}
