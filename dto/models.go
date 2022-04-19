package dto

import "time"

type Item struct {
	Title       string     `xml:"title" json:"title"`
	Link        string     `xml:"link" json:"link"`
	Description string     `xml:"description" json:"description"`
	PubDate     string     `xml:"pubDate"`
	Source      ItemSource `xml:"source"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type ItemSource struct {
	URL   string `xml:"url,attr"`
	Title string `xml:",chardata"`
}

type RssItem struct {
	Title       string     `json:"title"`
	Source      string     `json:"source"`
	SourceURL   string     `json:"sourceUrl"`
	Link        string     `json:"link"`
	PublishDate *time.Time `json:"publishDate"`
	Description string     `json:"description"`
}
