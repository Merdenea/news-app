package news

import "encoding/xml"

type Source struct {
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	Articles []Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Date        string `json:"date"`
}

type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Channel     Channel  `xml:"channel"`
	Description string   `xml:"description"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Date        string   `xml:"pubDate"`
}
