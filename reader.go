package reader

import (
	"errors"
	"net/http"

	"github.com/yanchevdimitar/go-rss-reader/dto"
)

var (
	ErrHTTP = errors.New("error on http request")
)

type Parser interface {
	Parse() (rss []dto.RssResponse, err error)
}

type DefaultReader struct {
	urls           []string
	responseReader ResponseReader
}

func NewDafaultReader(urls []string, responseReader ResponseReader) Parser {
	return DefaultReader{
		urls,
		responseReader,
	}
}

func (rr DefaultReader) Parse() (rss []dto.RssResponse, err error) {
	var items []dto.RssItem
	var resultQueue []dto.URLResponse

	for i := 0; i < len(rr.urls); i++ {
		resultQueue = append(resultQueue, <-rr.fetch())
		if resultQueue[i].Error != nil {
			return nil, ErrHTTP
		}

		items, err = rr.responseReader.Read(resultQueue[i].Response)
		if err != nil {
			return nil, err
		}

		rssResponse := dto.RssResponse{
			Items: map[string][]dto.RssItem{resultQueue[i].URL: items},
		}

		rss = append(rss, rssResponse)
	}

	return
}

func (rr DefaultReader) fetch() <-chan dto.URLResponse {
	c := make(chan dto.URLResponse)

	for _, url := range rr.urls {
		go func(url string) {
			req, _ := http.NewRequest("GET", url, nil)
			resp, err := getHTTPClient().Do(req)

			c <- dto.URLResponse{URL: url, Response: resp, Error: err}
		}(url)
	}

	return c
}
