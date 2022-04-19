package reader

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yanchevdimitar/go-rss-reader/dto"
)

var (
	ErrXML = errors.New("wrong XML")
)

type ResponseReader interface {
	Read(resp *http.Response) (items []dto.RssItem, err error)
}

type DefaultParser struct {
}

func NewDefaultParser() ResponseReader {
	return DefaultParser{}
}

func (pr DefaultParser) Read(resp *http.Response) (items []dto.RssItem, err error) {
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	data := dto.Rss{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return nil, ErrXML
	}

	return pr.buildRssItems(&data.Channel)
}

func (pr DefaultParser) buildRssItems(channel *dto.Channel) (items []dto.RssItem, err error) {
	for _, item := range channel.Items {
		rssItem := dto.RssItem{
			Title:       item.Title,
			Source:      item.Source.Title,
			SourceURL:   item.Source.URL,
			Link:        item.Link,
			Description: item.Description,
		}

		var t time.Time
		t, err = time.Parse(time.RFC3339Nano, item.PubDate)
		if err == nil {
			v := t.UTC()
			rssItem.PublishDate = &v
		}

		items = append(items, rssItem)
	}

	return items, nil
}
